package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/shopspring/decimal"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

const (
	ConfirmHeight                        = 0
	ScanSpan                             = 1000
	SyncIntervalBelow1P0                 = 60 //in secs
	SyncIntervalBelow1P1                 = 60
	SyncIntervalBelow1P5                 = 360
	SyncIntervalBelow2P0                 = 720
	SyncIntervalAbove2P0                 = 1800
	SyncIntervalBackGround               = 600
	SyncIntervalForMarkets               = 6
	MonitorLiquidationInterval           = 120
	ForbiddenPeriodForBadLiquidation     = 200 //200 block
	ForbiddenPeriodForPendingLiquidation = 200
	FeedPriceUpdateInterval              = 2 //
	ProfitSymbol                         = "vBUSD"
	GasSymbol                            = "vBNB"
)

type TokenInfo struct {
	Address               common.Address
	UnderlyingAddress     common.Address
	UnderlyingDecimals    uint8
	CollateralFactor      decimal.Decimal
	Price                 decimal.Decimal
	PriceUpdateHeight     uint64
	FeededPrice           decimal.Decimal
	FeedPriceUpdateHeihgt uint64
	Oracle                common.Address
}

type Asset struct {
	Symbol           string
	Balance          decimal.Decimal
	Loan             decimal.Decimal
	CollateralFactor decimal.Decimal
	BalanceValue     decimal.Decimal
	CollateralValue  decimal.Decimal
	LoanValue        decimal.Decimal
	Price            decimal.Decimal
	ExchangeRate     decimal.Decimal
}

type AccountInfo struct {
	Account       common.Address
	HealthFactor  decimal.Decimal
	MaxLoanSymbol string
	MaxLoanValue  decimal.Decimal
	MintedVAIS    decimal.Decimal
	Height        uint64
	Assets        []Asset
}

type FeededPrice struct {
	Symbol   string
	Address  common.Address
	Price    decimal.Decimal
	GasPrice decimal.Decimal
	Hash     common.Hash
	Height   uint64
}

type AccountsWithFeededPrice struct {
	Addresses   []common.Address
	FeededPrice *FeededPrice
}

type Liquidation struct {
	AccInfo     AccountInfo
	FeededPrice *FeededPrice
}

type ConcernedAccountInfo struct {
	AccInfo     AccountInfo
	FeededPrice FeededPrice
}

type semaphore chan struct{}

var (
	EXPSACLE              = decimal.New(1, 18)
	ExpScaleFloat, _      = big.NewFloat(0).SetString("1000000000000000000")
	BigZero               = big.NewInt(0)
	DecimalMax            = decimal.New(1, 128) // solidity's 2^256 = 4^128 < decimal's 10^128
	Decimal1P0, _         = decimal.NewFromString("1.0")
	Decimal1P1, _         = decimal.NewFromString("1.1")
	Decimal1P5, _         = decimal.NewFromString("1.5")
	Decimal2P0, _         = decimal.NewFromString("2.0")
	Decimal3P0, _         = decimal.NewFromString("3.0")
	DecimalNonProfit, _   = decimal.NewFromString("255") //magicnumber for nonprofit
	vBNBAddress           = common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36")
	wBNBAddress           = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	VAIAddress            = common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7")
	VAIControllerAddress  = common.HexToAddress("0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE")
	ProfitThreshold       = decimal.New(5, 18)   //5 USDT
	MaxLoanValueThreshold = decimal.New(100, 18) //100 USDT
)

type Syncer struct {
	c  *ethclient.Client
	db *leveldb.DB

	oracle                    *venus.PriceOracle
	comptroller               *venus.Comptroller
	pancakeRouter             *venus.IPancakeRouter02
	pancakeFactory            *venus.IPancakeFactory
	closeFactor               decimal.Decimal
	symbols                   map[common.Address]string
	tokens                    map[string]*TokenInfo
	flashLoanPools            map[string][]common.Address
	paths                     map[string][][]common.Address //"srcSymbol:dstSymbol"
	vbep20s                   map[common.Address]*venus.Vbep20
	liquidator                *venus.IQingsuan
	PrivateKey                *ecdsa.PrivateKey
	m                         sync.Mutex
	wg                        sync.WaitGroup
	quitCh                    chan struct{}
	feededPriceCh             chan *FeededPrice
	liquidationCh             chan *Liquidation
	priortyLiquidationCh      chan *Liquidation
	concernedAccountInfoCh    chan *ConcernedAccountInfo
	backgroundAccountSyncCh   chan common.Address
	lowPriorityAccountSyncCh  chan *AccountsWithFeededPrice
	highPriorityAccountSyncCh chan *AccountsWithFeededPrice
}

func init() {
}

func (s semaphore) Acquire() {
	s <- struct{}{}
}

func (s semaphore) Release() {
	<-s
}

func NewSyncer(
	c *ethclient.Client,
	db *leveldb.DB,
	comptrollerAddress string,
	oracleAddress string,
	pancakeRouterAddress string,
	liquidatorAddress string,
	privatekey string) *Syncer {

	exist, err := db.Has(dbm.BorrowerNumberKey(), nil)
	if !exist {
		db.Put(dbm.BorrowerNumberKey(), big.NewInt(0).Bytes(), nil)
	}

	comptroller, err := venus.NewComptroller(common.HexToAddress(comptrollerAddress), c)
	if err != nil {
		panic(err)
	}

	bigCloseFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		panic(err)
	}
	closeFactor := decimal.NewFromBigInt(bigCloseFactor, 0)

	oracle, err := venus.NewPriceOracle(common.HexToAddress(oracleAddress), c)
	if err != nil {
		panic(err)
	}

	pancakeRouter, err := venus.NewIPancakeRouter02(common.HexToAddress(pancakeRouterAddress), c)
	if err != nil {
		panic(err)
	}

	factoryAddress, err := pancakeRouter.Factory(nil)
	if err != nil {
		panic(err)
	}

	pancakeFactory, err := venus.NewIPancakeFactory(factoryAddress, c)
	if err != nil {
		panic(err)
	}

	liquidator, err := venus.NewIQingsuan(common.HexToAddress(liquidatorAddress), c)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		panic(err)
	}

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		panic(err)
	}

	symbols := make(map[common.Address]string)
	tokens := make(map[string]*TokenInfo)
	vbep20s := make(map[common.Address]*venus.Vbep20)

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(markets))

	sem := make(semaphore, runtime.NumCPU())
	for _, market_ := range markets {
		market := market_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()

			vbep20, err := venus.NewVbep20(market, c)
			if err != nil {
				panic(err)
			}

			symbol, err := vbep20.Symbol(nil)
			if err != nil {
				panic(err)
			}

			marketDetail, err := comptroller.Markets(nil, market)
			if err != nil {
				panic(err)
			}

			bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				bigPrice = big.NewInt(0)
			}
			var underlyingAddress common.Address
			if market == vBNBAddress {
				underlyingAddress = wBNBAddress
			} else {
				underlyingAddress, err = vbep20.Underlying(nil)
				if err != nil {
					panic(err)
				}
			}

			bep20, err := venus.NewVbep20(underlyingAddress, c)
			underlyingDecimals, err := bep20.Decimals(nil)
			if err != nil {
				panic(err)
			}

			collaterFactor := decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
			price := decimal.NewFromBigInt(bigPrice, 0)

			feedSymbol := strings.TrimPrefix(symbol, "v")
			if feedSymbol == "BTC" {
				feedSymbol = "BTCB"
			}

			var finalOracle common.Address
			if symbol != "vCAN" {
				feedAddr, err := oracle.GetFeed(nil, feedSymbol)
				if err != nil {
					panic(err)
				}

				priceFeed, err := venus.NewPriceFeed(feedAddr, c)
				if err != nil {
					panic(err)
				}

				//feedDecimal, err := priceFeed.Decimals(nil)
				//if err != nil {
				//	panic(err)
				//}

				//logger.Printf("%v feeddecimal:%v\n", symbol, feedDecimal)

				finalOracle, err = priceFeed.Aggregator(nil)
				//logger.Printf("symbol:%v, priceFeed:%v, oracle:%v\n", symbol, feedAddr, finalOracle)
				if err != nil {
					panic(err)
				}
			}

			height, err := c.BlockNumber(context.Background())
			if err != nil {
				panic(err)
			}

			token := &TokenInfo{
				Address:            market,
				UnderlyingAddress:  underlyingAddress,
				UnderlyingDecimals: underlyingDecimals,
				CollateralFactor:   collaterFactor,
				Price:              price,
				PriceUpdateHeight:  height,
				Oracle:             finalOracle,
			}

			m.Lock()
			//logger.Printf("market:%v, symbol:%v, token:%+v\n", market, symbol, token)
			tokens[symbol] = token
			symbols[market] = symbol
			vbep20s[market] = vbep20
			m.Unlock()
		}()
	}
	wg.Wait()

	tokens["VAI"] = &TokenInfo{
		Address:            VAIControllerAddress,
		UnderlyingAddress:  VAIAddress,
		UnderlyingDecimals: 18,
		CollateralFactor:   decimal.Zero,
		Price:              decimal.New(1, 18),
		Oracle:             common.Address{},
	}

	flashLoanPool := buildFlashLoanPool(pancakeFactory, c, tokens)
	paths := buildPaths(pancakeRouter, tokens)

	return &Syncer{
		c:                         c,
		db:                        db,
		oracle:                    oracle,
		comptroller:               comptroller,
		pancakeRouter:             pancakeRouter,
		pancakeFactory:            pancakeFactory,
		closeFactor:               closeFactor,
		tokens:                    tokens,
		flashLoanPools:            flashLoanPool,
		paths:                     paths,
		symbols:                   symbols,
		vbep20s:                   vbep20s,
		liquidator:                liquidator,
		PrivateKey:                privateKey,
		m:                         m,
		quitCh:                    make(chan struct{}),
		feededPriceCh:             make(chan *FeededPrice, 64),
		liquidationCh:             make(chan *Liquidation, 1024),
		priortyLiquidationCh:      make(chan *Liquidation, 1024),
		concernedAccountInfoCh:    make(chan *ConcernedAccountInfo, 4096),
		backgroundAccountSyncCh:   make(chan common.Address, 8192),
		lowPriorityAccountSyncCh:  make(chan *AccountsWithFeededPrice, 1024),
		highPriorityAccountSyncCh: make(chan *AccountsWithFeededPrice, 248),
	}
}

func (s *Syncer) Start() {
	log.Info("server start")
	fmt.Println("server start")

	s.wg.Add(9)
	go s.SyncMarketsAndPricesLoop()
	go s.ProcessFeededPriceLoop()
	go s.SearchNewBorrowerLoop()
	go s.BackgroundSyncLoop()
	go s.syncAccountLoop()
	go s.MonitorLiquidationEventLoop()
	go s.PrintConcernedAccountInfoLoop()
	go s.MonitorTxPoolLoop()
	go s.ProcessLiquidationLoop()
}

func (s *Syncer) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Syncer) SyncMarketsAndPricesLoop() {
	defer s.wg.Done()
	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			//logger.Printf("%v th sync markers and prices @ %v\n", count, time.Now())
			count++
			s.doSyncMarketsAndPrices()
			t.Reset(time.Second * SyncIntervalForMarkets)
		}
	}
}

func (s *Syncer) doSyncMarketsAndPrices() {
	comptroller := s.comptroller
	oracle := s.oracle
	c := s.c

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		return
	}

	bigCloseFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		return
	}
	s.closeFactor = decimal.NewFromBigInt(bigCloseFactor, 0)
	for _, market_ := range markets {
		market := market_
		vbep20, err := venus.NewVbep20(market, c)
		if err != nil {
			continue
		}
		symbol, err := vbep20.Symbol(nil)
		if err != nil {
			continue
		}

		marketDetail, err := comptroller.Markets(nil, market)
		if err != nil {
			continue
		}

		bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
		if err != nil {
			bigPrice = big.NewInt(0)
		}

		height, err := s.c.BlockNumber(context.Background())
		if err != nil {
			continue
		}
		s.m.Lock()
		if s.tokens[symbol] != nil {
			s.tokens[symbol].CollateralFactor = decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
			s.tokens[symbol].Price = decimal.NewFromBigInt(bigPrice, 0)
			s.tokens[symbol].PriceUpdateHeight = height
		} else {
			logger.Printf("a new symbol:%v added, please restart venusd\n", symbol)
		}
		s.m.Unlock()
	}
}

func (s *Syncer) ProcessFeededPriceLoop() {
	defer s.wg.Done()
	for {
		select {
		case <-s.quitCh:
			return
		case feededPrice := <-s.feededPriceCh:
			s.processFeededPrice(feededPrice)
		}
	}
}

func (s *Syncer) processFeededPrice(feededPrice *FeededPrice) {
	db := s.db
	var accountsSymbol []common.Address
	var accountsSymobl2 []common.Address
	wg := s.wg

	s.m.Lock()
	tokens := copyTokens(s.tokens)
	s.m.Unlock()

	symbol := feededPrice.Symbol
	price := tokens[symbol].Price
	priceDeltaRatio := price.Sub(feededPrice.Price).Abs().Div(price)
	if priceDeltaRatio.Cmp(decimal.New(10, -2)) == 1 {
		logger.Printf("processFeededPrices, sybmol:%v，feedPrices %v, originalPrices:%v, vibration %v exceeds 5 percent, height%v\n", symbol, feededPrice.Price, price, priceDeltaRatio, feededPrice.Height)
		return
	}

	symbol2 := ""
	if symbol == "vBETH" {
		symbol2 = "vETH"
	}
	if symbol == "vETH" {
		symbol2 = "vBETH"
	}

	s.m.Lock()
	s.tokens[symbol].FeededPrice = feededPrice.Price
	s.tokens[symbol].FeedPriceUpdateHeihgt = feededPrice.Height

	if symbol2 != "" {
		s.tokens[symbol2].FeededPrice = feededPrice.Price
		s.tokens[symbol2].FeedPriceUpdateHeihgt = feededPrice.Height
	}
	s.m.Unlock()

	wg.Add(2)
	go func() {
		defer wg.Done()
		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		iter := db.NewIterator(util.BytesPrefix(prefix), nil)
		defer iter.Release()
		for iter.Next() {
			account := common.BytesToAddress(iter.Value())
			accountsSymbol = append(accountsSymbol, account)
		}
	}()

	go func() {
		defer wg.Done()
		if symbol2 != "" {
			prefix2 := append(dbm.MarketPrefix, []byte(symbol2)...)
			iter := db.NewIterator(util.BytesPrefix(prefix2), nil)
			defer iter.Release()
			for iter.Next() {
				account := common.BytesToAddress(iter.Value())
				accountsSymobl2 = append(accountsSymobl2, account)
			}
		}
	}()
	wg.Wait()
	accountsSymbol = append(accountsSymbol, accountsSymobl2...)

	priorityAccountMap := make(map[common.Address]bool)
	iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P1Prefix), nil)
	for iter.Next() {
		account := common.BytesToAddress(iter.Value())
		priorityAccountMap[account] = true
	}

	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P0Prefix), nil)
	for iter.Next() {
		account := common.BytesToAddress(iter.Value())
		priorityAccountMap[account] = true
	}
	iter.Release()

	var highPriorityAccounts []common.Address
	var lowPriorityAccounts []common.Address
	for _, account := range accountsSymbol {
		if priorityAccountMap[account] {
			highPriorityAccounts = append(highPriorityAccounts, account)
		} else {
			lowPriorityAccounts = append(lowPriorityAccounts, account)
		}
	}

	//sort accounts according ascending healthfactor
	healthFactors := make([]decimal.Decimal, 0)
	finalHighPriorityAccounts := make([]common.Address, 0)
	errorHighPriorityAccounts := make([]common.Address, 0)
	for _, account := range highPriorityAccounts {
		bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
		if err != nil {
			errorHighPriorityAccounts = append(errorHighPriorityAccounts, account)
		}
		var info AccountInfo
		err = json.Unmarshal(bz, &info)
		if err != nil {
			errorHighPriorityAccounts = append(errorHighPriorityAccounts, account)
		} else {
			healthFactors = append(healthFactors, info.HealthFactor)
			finalHighPriorityAccounts = append(finalHighPriorityAccounts, account)
		}
	}

	for i := 0; i < len(finalHighPriorityAccounts)-1; i++ {
		for j := i + 1; j < len(finalHighPriorityAccounts); j++ {
			if healthFactors[i].Cmp(healthFactors[j]) == 1 {
				healthFactors[i], healthFactors[j] = healthFactors[j], healthFactors[j]
				finalHighPriorityAccounts[i], finalHighPriorityAccounts[j] = finalHighPriorityAccounts[j], finalHighPriorityAccounts[i]
			}
		}
	}

	//send highPriorityAccount as a batch
	if len(finalHighPriorityAccounts) != 0 {
		s.highPriorityAccountSyncCh <- &AccountsWithFeededPrice{
			Addresses:   finalHighPriorityAccounts,
			FeededPrice: feededPrice,
		}

		s.highPriorityAccountSyncCh <- &AccountsWithFeededPrice{
			Addresses:   errorHighPriorityAccounts,
			FeededPrice: feededPrice,
		}
	}
}

func (s *Syncer) syncAccountLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case req := <-s.highPriorityAccountSyncCh:
			s.processHighPriorityAccountSync(req)

		case req := <-s.lowPriorityAccountSyncCh:
			if len(s.highPriorityAccountSyncCh) != 0 {
				continue
			}
			accounts := req.Addresses
			feededPrices := req.FeededPrice
			s.syncOneAccountWithFeededPrice(accounts[0], feededPrices)

		case account := <-s.backgroundAccountSyncCh:
			if len(s.highPriorityAccountSyncCh) != 0 {
				continue
			}
			s.syncOneAccount(account)
		}
	}
}

func (s *Syncer) processHighPriorityAccountSync(req *AccountsWithFeededPrice) {
	accounts := req.Addresses
	feededPrice := req.FeededPrice

	var wg sync.WaitGroup
	wg.Add(len(accounts))
	sem := make(semaphore, runtime.NumCPU())
	for _, account_ := range accounts {
		account := account_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			s.syncOneAccountWithFeededPrice(account, feededPrice)
		}()
	}
	wg.Wait()
}

func (s *Syncer) SearchNewBorrowerLoop() {
	defer s.wg.Done()
	db := s.db
	c := s.c
	ctx := context.Background()
	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	s.m.Unlock()

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	vbep20Abi, _ := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))

	var logs []types.Log
	var addresses []common.Address

	for _, token := range tokens {
		addresses = append(addresses, token.Address)
	}

	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-s.quitCh:
			return

		case <-t.C:
			currentHeight, err := c.BlockNumber(ctx)
			if err != nil {
				t.Reset(time.Second * 3)
				continue
			}

			bz, err := db.Get(dbm.LatestHandledHeightStoreKey(), nil)
			if err != nil {
				t.Reset(time.Millisecond * 20)
				continue
			}

			lastHandledHeight := big.NewInt(0).SetBytes(bz).Uint64()
			startHeight := lastHandledHeight + 1
			endHeight := currentHeight
			if startHeight+ConfirmHeight >= currentHeight {
				t.Reset(time.Second * 3)
				continue
			}

			if currentHeight-lastHandledHeight >= ScanSpan {
				endHeight = startHeight + ScanSpan - 1
			}

			//logger.Printf("startHeight:%v, endHeight:%v\n", startHeight, endHeight)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startHeight)),
				ToBlock:   big.NewInt(int64(endHeight)),
				Addresses: addresses,
				Topics:    [][]common.Hash{{topicBorrow}},
			}

			logs, err = c.FilterLogs(context.Background(), query)
			if err != nil {
				logger.Printf("SearchNewBorrowerLoop, fail to filter logs, err:%v\n", err)
				goto EndWithoutUpdateHeight
			}

			for i, log := range logs {
				var borrowEvent venus.Vbep20Borrow
				err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
				logger.Printf("%v height:%v, name:%v account:%v\n", (i + 1), log.BlockNumber, symbols[log.Address], borrowEvent.Borrower)

				account := borrowEvent.Borrower
				err := s.syncOneAccountWithIncreaseAccountNumber(account)
				if err != nil {
					goto EndWithoutUpdateHeight
				}
			}

			err = db.Put(dbm.LatestHandledHeightStoreKey(), big.NewInt(int64(endHeight)).Bytes(), nil)
			if err != nil {
				goto EndWithoutUpdateHeight
			}

		EndWithoutUpdateHeight:
			t.Reset(time.Millisecond * 20)
		}
	}
}

func (s *Syncer) BackgroundSyncLoop() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			logger.Printf("%vth background sync t @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			for _, account := range accounts {
				s.backgroundAccountSyncCh <- account
			}

			t.Reset(time.Second * SyncIntervalBackGround)
		}
	}
}

func (s *Syncer) MonitorLiquidationEventLoop() {
	db := s.db
	defer s.wg.Done()

	t := time.NewTimer(0)
	defer t.Stop()

	s.m.Lock()
	tokens := s.tokens
	s.m.Unlock()

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	if err != nil {
		panic(err)
	}

	bz, err := db.Get(dbm.LatestHandledHeightStoreKey(), nil)
	if err != nil {
		panic(err)
	}

	monitorStartHeight := big.NewInt(0).SetBytes(bz).Uint64()

	topicLiquidateBorrow := common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")

	var vTokenAddresses []common.Address
	for _, token := range tokens {
		vTokenAddresses = append(vTokenAddresses, token.Address)
	}

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			monitorEndHeight, err := s.c.BlockNumber(context.Background())
			if err != nil {
				monitorEndHeight = monitorStartHeight
			}
			logger.Printf("%vth sync monitor LiquidationBorrow event, startHeight:%v, endHeight:%v \n", count, monitorStartHeight, monitorEndHeight)

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(monitorStartHeight)),
				ToBlock:   big.NewInt(int64(monitorEndHeight)),
				Addresses: vTokenAddresses,
				Topics:    [][]common.Hash{{topicLiquidateBorrow}},
			}

			logs, err := s.c.FilterLogs(context.Background(), query)
			if err != nil {
				logger.Printf("%vth sync monitor LiquidationBorrow event, startHeight:%v, endHeight:%v, err:%v \n", count, monitorStartHeight, monitorEndHeight, err)
			} else {
				for _, log := range logs {
					var eve venus.Vbep20LiquidateBorrow
					err = vbep20Abi.UnpackIntoInterface(&eve, "LiquidateBorrow", log.Data)
					if err == nil {
						logger.Printf("LiquidateBorrow event happen @ height:%v, txhash:%v, liquidator:%v borrower:%v, repayAmount:%v, collateral:%v, seizedAmount:%v\n", log.BlockNumber, log.TxHash, eve.Liquidator, eve.Borrower, eve.RepayAmount, eve.VTokenCollateral, eve.SeizeTokens)
					}
				}
				monitorStartHeight = monitorEndHeight
				t.Reset(time.Second * MonitorLiquidationInterval)
			}
			count++
		}
	}
}

func (s *Syncer) MonitorTxPoolLoop() {
	defer s.wg.Done()

	v := reflect.ValueOf(s.c).Elem()
	f := v.FieldByName("c")
	rf := reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
	concrete_client, _ := rf.Interface().(*rpc.Client)

	txPoolTXs := make(chan common.Hash, 2048)
	concrete_client.EthSubscribe(
		context.Background(), txPoolTXs, "newPendingTransactions",
	)
	aggregatorABI, _ := venus.AggregatorMetaData.GetAbi()
	OracleToVTokenMap := make(map[common.Address]common.Address)

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	s.m.Unlock()

	for symbol, token := range tokens {
		if symbol != "vCAN" {
			OracleToVTokenMap[token.Oracle] = token.Address
		}
	}

	for {
		select {
		case <-s.quitCh:
			return

		case txHash := <-txPoolTXs:
			height, _ := s.c.BlockNumber(context.Background())
			txn, is_pending, err := s.c.TransactionByHash(context.Background(), txHash)

			if err == nil && txn != nil && txn.To() != nil && is_pending == true {
				vTokenAddress, ok := OracleToVTokenMap[*txn.To()]
				if !ok || len(txn.Data()) < 5 {
					continue
				}

				method, err := aggregatorABI.MethodById(txn.Data()[0:4])
				if err != nil {
					continue
				}

				if method.Name == "transmit" {
					inputData := make(map[string]interface{})
					err = method.Inputs.UnpackIntoMap(inputData, txn.Data()[4:])
					if err != nil {
						continue
					}

					data := inputData["_report"].([]byte)
					numbering := data[32+32+32+32:]
					midpos := len(numbering) / 32 / 2
					numberingmid := numbering[midpos*32 : midpos*32+32]
					bigFeededPrice := big.NewInt(0).SetBytes(numberingmid)

					symobl := symbols[vTokenAddress]
					decimalDelta := 10 + 18 - tokens[symobl].UnderlyingDecimals //VenusChainlinkOracle.sol line46

					feededPrice := &FeededPrice{
						Symbol:   symobl,
						Address:  vTokenAddress,
						Price:    decimal.NewFromBigInt(bigFeededPrice, int32(decimalDelta)),
						Hash:     txHash,
						GasPrice: decimal.NewFromBigInt(txn.GasPrice(), 0),
						Height:   height,
					}

					logger.Printf("catch feedPrice, height:%v, symbol:%v, price:%+v\n", height, symobl, feededPrice)
					s.feededPriceCh <- feededPrice
				}
			}
		}
	}
}

func (s *Syncer) PrintConcernedAccountInfoLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return
		case info := <-s.concernedAccountInfoCh:
			logger.Printf("ConernedAccountInfo:%+v\n", info)
		}
	}
}

// only for tests
func (s *Syncer) syncAccounts(accounts []common.Address) {
	var wg sync.WaitGroup
	wg.Add(len(accounts))
	sem := make(semaphore, runtime.NumCPU())

	for _, account_ := range accounts {
		account := account_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			s.syncOneAccount(account)
		}()
	}
	wg.Wait()
}

func (s *Syncer) syncOneAccount(account common.Address) error {
	//ctx := context.Background()
	comptroller := s.comptroller

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	s.m.Unlock()

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	var assets []Asset
	bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		logger.Printf("syncOneAccount, fail to get MintedVAIs, err:%v\n", err)
		return err
	}

	mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil || len(markets) == 0 {
		logger.Printf("syncOneAccount, fail to get GetAssetsIn or account is not in any markets, err:%v\n", err)
		return err
	}

	maxLoanValue := decimal.NewFromInt(0)
	maxLoanSymbol := ""
	var fakedFeedPrice *FeededPrice
	for _, market := range markets {
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil || errCode.Cmp(BigZero) != 0 {
			logger.Printf("syncOneAccount, fail to get GetAccountSnapshot, err:%v, errCode:%v\n", err, errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			continue
		}

		symbol := symbols[market]
		collateralFactor := tokens[symbol].CollateralFactor
		price := tokens[symbol].Price
		if tokens[symbol].FeedPriceUpdateHeihgt+FeedPriceUpdateInterval > tokens[symbol].PriceUpdateHeight {
			price = tokens[symbol].FeededPrice
			fakedFeedPrice = &FeededPrice{
				Symbol: symbol,
				Price:  price,
				Height: tokens[symbol].FeedPriceUpdateHeihgt,
			}
		}

		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)
		borrow := decimal.NewFromBigInt(bigBorrow, 0)

		multiplier := price.Mul(exchangeRate).Div(EXPSACLE).Div(EXPSACLE)
		balanceValue := balance.Mul(multiplier)
		collateral := balanceValue.Mul(collateralFactor).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateral.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))

		asset := Asset{
			Symbol:           symbol,
			Balance:          balance,
			Loan:             borrow,
			CollateralFactor: collateralFactor,
			BalanceValue:     balanceValue,
			CollateralValue:  collateral,
			LoanValue:        loan,
			Price:            price,
			ExchangeRate:     exchangeRate,
		}

		//logger.Printf("syncOneAccount, symbol:%v, price:%v, exchangeRate:%v, asset:%+v\n", symbol, price, bigExchangeRate, asset)
		assets = append(assets, asset)
		if loan.Cmp(maxLoanValue) == 1 {
			maxLoanValue = loan
			maxLoanSymbol = symbol
		}
	}

	totalLoan = totalLoan.Add(mintedVAIS)
	healthFactor := decimal.New(100, 0)
	if totalLoan.Cmp(decimal.Zero) == 1 {
		healthFactor = totalCollateral.Div(totalLoan)
	}

	if mintedVAIS.Cmp(maxLoanValue) == 1 {
		maxLoanValue = mintedVAIS
		maxLoanSymbol = "VAI"
	}
	currentHeight, _ := s.c.BlockNumber(context.Background())
	info := AccountInfo{
		Account:       account,
		HealthFactor:  healthFactor,
		MaxLoanValue:  maxLoanValue,
		MaxLoanSymbol: maxLoanSymbol,
		MintedVAIS:    mintedVAIS,
		Height:        currentHeight,
		Assets:        assets,
	}
	s.updateDB(account, info)
	if maxLoanValue.Cmp(MaxLoanValueThreshold) == 1 {
		errCode, _, shortfall, err := comptroller.GetAccountLiquidity(nil, account)
		if err == nil && errCode.Cmp(BigZero) == 0 && shortfall.Cmp(BigZero) == 1 {
			liquidation := &Liquidation{
				FeededPrice: fakedFeedPrice,
				AccInfo:     info,
			}
			s.liquidationCh <- liquidation
		}

		//logger.Printf("syncOneAccount,account:%v, height:%v,totalCollateral:%v, totalLoan:%v,info:%+v\n", account, currentHeight, totalCollateral, totalLoan, info.toReadable())
		if healthFactor.Cmp(Decimal1P1) == -1 {
			cinfo := &ConcernedAccountInfo{
				AccInfo: info,
			}
			s.concernedAccountInfoCh <- cinfo
		}
	}

	return nil
}

func (s *Syncer) syncOneAccountWithFeededPrice(account common.Address, feededPrice *FeededPrice) error {
	comptroller := s.comptroller

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	s.m.Unlock()
	vbep20s := s.vbep20s

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	var assets []Asset
	bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		return err
	}

	mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)

	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return err
	}

	maxLoanValue := decimal.NewFromInt(0)
	maxLoanSymbol := ""
	for _, market := range markets {
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			logger.Printf("syncOneAccountWithFeededPrices, fail to get GetAccountSnapshot, err:%v\n", err)
			return err
		}

		if errCode.Cmp(BigZero) != 0 {
			logger.Printf("syncOneAccountWithFeededPrices, fail to get GetAccountSnapshot, errCode:%v\n", errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			continue
		}

		symbol := symbols[market]
		collateralFactor := tokens[symbol].CollateralFactor

		//use latest price even it is pending
		price := tokens[symbol].Price
		if tokens[symbol].FeedPriceUpdateHeihgt+FeedPriceUpdateInterval > tokens[symbol].PriceUpdateHeight {
			price = tokens[symbol].FeededPrice
		}

		//apply feeded prices if exist
		if symbols[feededPrice.Address] == symbol {
			price = feededPrice.Price
		} else if strings.Contains(symbols[feededPrice.Address], "ETH") && strings.Contains(symbol, "ETH") {
			price = feededPrice.Price
		}

		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)
		borrow := decimal.NewFromBigInt(bigBorrow, 0)

		multiplier := price.Mul(exchangeRate).Div(EXPSACLE).Div(EXPSACLE)
		balanceValue := balance.Mul(multiplier)
		collateral := balanceValue.Mul(collateralFactor).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateral.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))

		//build account table
		asset := Asset{
			Symbol:           symbol,
			Balance:          balance,
			Loan:             borrow,
			CollateralFactor: collateralFactor,
			BalanceValue:     balanceValue,
			CollateralValue:  collateral,
			LoanValue:        loan,
			Price:            price,
			ExchangeRate:     exchangeRate,
		}

		//logger.Printf("syncOneAccountWithFeededPrices, symbol:%v, exchangeRate:%v,price:%v, asset:%+v\n", symbol, exchangeRate, price, asset)
		assets = append(assets, asset)

		if loan.Cmp(maxLoanValue) == 1 {
			maxLoanValue = loan
			maxLoanSymbol = symbol
		}
	}

	totalLoan = totalLoan.Add(mintedVAIS)
	healthFactor := decimal.New(100, 0)
	if totalLoan.Cmp(decimal.Zero) == 1 {
		healthFactor = totalCollateral.Div(totalLoan)
	}

	if mintedVAIS.Cmp(maxLoanValue) == 1 {
		maxLoanValue = mintedVAIS
		maxLoanSymbol = "VAI"
	}

	currentHeight, _ := s.c.BlockNumber(context.Background())
	info := AccountInfo{
		Account:       account,
		HealthFactor:  healthFactor,
		MaxLoanValue:  maxLoanValue,
		MaxLoanSymbol: maxLoanSymbol,
		MintedVAIS:    mintedVAIS,
		Height:        currentHeight,
		Assets:        assets,
	}
	s.updateDB(account, info)

	//logger.Printf("syncOneAccountWithFeededPrices,account:%v, height:%v,  totalCollateral:%v, totalLoan:%v, info:%+v\n", account, currentHeight, totalCollateral, totalLoan, info.toReadable())
	if healthFactor.Cmp(decimal.NewFromInt(1)) == -1 {
		liquidation := &Liquidation{
			AccInfo:     info,
			FeededPrice: feededPrice,
		}
		s.priortyLiquidationCh <- liquidation
	}

	if healthFactor.Cmp(Decimal1P1) == -1 {
		cinfo := &ConcernedAccountInfo{
			FeededPrice: *feededPrice,
			AccInfo:     info,
		}
		s.concernedAccountInfoCh <- cinfo
	}
	return nil
}

func (s *Syncer) syncOneAccountWithIncreaseAccountNumber(account common.Address) error {
	ctx := context.Background()
	c := s.c
	db := s.db

	accountBytes := account.Bytes()
	exist, err := db.Has(dbm.BorrowersStoreKey(accountBytes), nil)
	if err != nil {
		return err
	}

	byteCode, err := c.CodeAt(ctx, account, nil)
	if len(byteCode) > 0 {
		//ignore smart contract
		return nil
	}

	err = s.syncOneAccount(account)
	if err != nil {
		return err
	}

	if !exist {
		//if account not exist in borrowers table, record it into borrowers table and increase borrowers number
		err = db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
		if err != nil {
			return err
		}

		bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
		num := big.NewInt(0).SetBytes(bz).Int64()
		if err != nil {
			return err
		}

		num += 1
		err = db.Put(dbm.BorrowerNumberKey(), big.NewInt(num).Bytes(), nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Syncer) ProcessLiquidationLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case pending := <-s.priortyLiquidationCh:
			logger.Printf("receive priority liquidation:%v\n", pending)
			s.processLiquidationReq(pending)

		case pending := <-s.liquidationCh:
			logger.Printf("recive liquidation:%v\n", pending)
			s.processLiquidationReq(pending)
		}
	}
}

func (s *Syncer) processLiquidationReq(liquidation *Liquidation) error {
	comptroller := s.comptroller
	//pancakeRouter := s.pancakeRouter
	//pancakeFactory := s.pancakeFactory
	accInfo := liquidation.AccInfo
	account := accInfo.Account
	db := s.db
	feededPrice := liquidation.FeededPrice
	withFeedPrice := false
	if feededPrice != nil {
		withFeedPrice = true
	}

	s.m.Lock()
	//symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	s.m.Unlock()
	vbep20s := s.vbep20s
	closeFactor := s.closeFactor

	//current height
	currentHeight, err := s.c.BlockNumber(context.Background())
	if err != nil {
		logger.Printf("processLiquidationReq, fail to get blockNumber,err:%v\n", err)
		return err
	}
	callOptions := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(currentHeight)),
	}
	//check BadLiquidationTx
	err = s.checkBadLiquidation(account, currentHeight)
	if err != nil {
		return err
	}

	//check PendingLiquidationTx
	err = s.checkPendingLiquidation(account, currentHeight)
	if err != nil {
		return err
	}

	//totalCollateral := decimal.NewFromInt(0)
	//totalLoan := decimal.NewFromInt(0)
	//
	//bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	//if err != nil {
	//	logger.Printf("processLiquidationReq, fail to get MintedVAIs, account:%v, err:%v\n", account, err)
	//	return err
	//}
	//mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)
	//
	//var assets []Asset
	//markets, err := comptroller.GetAssetsIn(nil, account)
	//if err != nil {
	//	logger.Printf("processLiquidationReq, fail to get GetAssetsIn, account:%v, err:%v\n", account, err)
	//	return err
	//}
	//
	//for _, market := range markets {
	//	errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
	//	//logger.Printf("bigBalance:%v, bigBorrow:%v, bigExchangeRate:%v\n", bigBalance, bigBorrow, bigExchangeRate)
	//	if err != nil {
	//		logger.Printf("processLiquidationReq, fail to get GetAccountSnapshot, account:%v, err:%v\n", account, err)
	//		return err
	//	}
	//
	//	if errCode.Cmp(BigZero) != 0 {
	//		logger.Printf("processLiquidationReq, fail to get GetAccountSnapshot, account:%v, errCode:%v\n", account, errCode)
	//		return err
	//	}
	//
	//	if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
	//		continue
	//	}
	//
	//	symbol := symbols[market]
	//	collateralFactor := tokens[symbol].CollateralFactor
	//	price := tokens[symbol].Price
	//
	//	if tokens[symbol].FeedPriceUpdateHeihgt > tokens[symbol].PriceUpdateHeight {
	//		withFeedPrice = true
	//		price = tokens[symbol].FeededPrice
	//	}
	//
	//	//apply feeded prices if exist
	//	if feededPrice != nil {
	//		if symbols[feededPrice.Address] == symbol {
	//			price = feededPrice.Price
	//		} else if strings.Contains(symbols[feededPrice.Address], "ETH") && strings.Contains(symbol, "ETH") {
	//			price = feededPrice.Price
	//		}
	//	}
	//
	//	exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
	//	balance := decimal.NewFromBigInt(bigBalance, 0) //vToken Amount
	//	borrow := decimal.NewFromBigInt(bigBorrow, 0)
	//
	//	multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
	//	multiplier = multiplier.Mul(price).Div(EXPSACLE)
	//	collateralD := balance.Mul(multiplier).Div(EXPSACLE)
	//	totalCollateral = totalCollateral.Add(collateralD.Truncate(0))
	//
	//	loan := borrow.Mul(price).Div(EXPSACLE)
	//	totalLoan = totalLoan.Add(loan.Truncate(0))
	//
	//	asset := Asset{
	//		Symbol:           symbol,
	//		Balance:          decimal.NewFromBigInt(bigBalance, 0),
	//		Loan:             decimal.NewFromBigInt(bigBorrow, 0),
	//		CollateralFactor: collateralFactor.Div(EXPSACLE),
	//		BalanceValue:     balance.Mul(exchangeRate).Mul(price).Div(EXPSACLE).Div(EXPSACLE), //collateral value in 10^-18 USDT
	//		CollateralValue:  collateralD,                                                      //collateral value considering collateralFactor, = Balance * collateralFactor
	//		LoanValue:        loan,                                                             //loan value in 10^-18 USDT
	//		Price:            price,
	//		ExchangeRate:     exchangeRate,
	//	}
	//	//logger.Printf("asset:%+v, address:%v\n", asset, tokens[asset.Symbol].Address)
	//	assets = append(assets, asset)
	//}
	//
	//totalLoan = totalLoan.Add(mintedVAIS)
	//logger.Printf("account:%v, totalCollateralValue:%v, mintedVAISValue:%v, totalLoanValue:%v, calculatedshortfall:%v\n", account, totalCollateral.Div(EXPSACLE), mintedVAIS.Div(EXPSACLE), totalLoan.Div(EXPSACLE), (totalLoan.Sub(totalCollateral)))
	//logger.Printf("account:%v, assets:%v\n", account, assets)

	if !withFeedPrice {
		errCode, _, shortfall, err := comptroller.GetAccountLiquidity(callOptions, account)
		if err != nil {
			logger.Printf("processLiquidationReq, fail to get GetAccountLiquidity, account:%v, err:%v \n", account, err)
			return err
		}
		if errCode.Cmp(BigZero) != 0 || shortfall.Cmp(BigZero) != 1 {
			err := fmt.Errorf("processLiquidationReq, fail to get GetAccountLiquidity, account:%v, errCode:%v, shortfall:%v \n", account, errCode, shortfall)
			logger.Printf("processLiquidationReq, fail to get GetAccountLiquidity, account:%v, errCode:%v, shortfall:%v \n", account, errCode, shortfall)
			return err
		}
	}

	//select the repayed token and seized collateral token
	maxLoanValue := accInfo.MaxLoanValue
	maxLoanSymbol := accInfo.MaxLoanSymbol

	//repayIndex := math.MaxInt32
	//for _, asset := range  assets {
	//	if asset.LoanValue.Cmp(maxLoanValue) == 1 {
	//		maxLoanValue = asset.LoanValue
	//		maxLoanSymbol = asset.Symbol
	//		//repayIndex = i
	//	}
	//}
	//
	//if mintedVAIS.Cmp(maxLoanValue) == 1 {
	//	maxLoanValue = mintedVAIS
	//	maxLoanSymbol = "VAI"
	//}

	//the closeFactor applies to only single borrowed asset
	maxRepayValue := maxLoanValue.Mul(closeFactor).Div(EXPSACLE) //closeFactor = 0.8*10^18
	repaySymbol := maxLoanSymbol
	assets := accInfo.Assets

	repayValue := decimal.NewFromInt(0)
	seizedSymbol := ""
	seizedIndex := 0
	for k, asset := range assets {
		if asset.BalanceValue.Cmp(maxRepayValue) != -1 {
			repayValue = maxRepayValue
			seizedSymbol = asset.Symbol
			seizedIndex = k
			break
		} else {
			if asset.BalanceValue.Cmp(repayValue) == 1 {
				repayValue = asset.BalanceValue
				seizedSymbol = asset.Symbol
				seizedIndex = k
			}
		}
	}

	bigSeizedCTokenAmount := big.NewInt(0)
	errCode := big.NewInt(0)

	var repayAmount decimal.Decimal
	if repaySymbol == "VAI" {
		repayAmount = repayValue.Truncate(0).Mul(decimal.New(999999, -4)) //to avoid TOO_MUCH_REPAY error
		errCode, bigSeizedCTokenAmount, err = comptroller.LiquidateVAICalculateSeizeTokens(callOptions, tokens[seizedSymbol].Address, repayAmount.BigInt())
		if err != nil || errCode.Cmp(BigZero) != 0 {
			logger.Printf("processLiquidationReq, fail to get LiquidateVAICalculateSeizeTokens, account:%v, err:%v, errCode:%v\n", account, err, errCode)
			return err
		}
	} else {
		bigBorrowBalanceStored, err := vbep20s[tokens[repaySymbol].Address].BorrowBalanceStored(callOptions, account)
		//logger.Printf("repaySymbol:%v, bigBorrowBalancedStored:%v\n", repaySymbol, bigBorrowBalanceStored)
		if err != nil {
			logger.Printf("processLiquidationReq, fail to get BorrowBalanceStored, account:%v, err:%v\n", account, err)
			return err
		}
		repayAmount = decimal.NewFromBigInt(bigBorrowBalanceStored, 0).Mul(closeFactor).Div(EXPSACLE)
		repayAmount = repayAmount.Truncate(0).Mul(decimal.New(999999, -4)) //to avoid TOO_MUCH_REPAY error

		errCode, bigSeizedCTokenAmount, err = comptroller.LiquidateCalculateSeizeTokens(callOptions, tokens[repaySymbol].Address, tokens[seizedSymbol].Address, repayAmount.BigInt())
		if err != nil || errCode.Cmp(BigZero) != 0 {
			logger.Printf("processLiquidationReq, fail to get LiquidateCalculateSeizeTokens, account:%v, err:%v, errCode%v\n", account, err, errCode)
			return err
		}
	}
	seizedVTokenAmount := decimal.NewFromBigInt(bigSeizedCTokenAmount, 0)
	seizedUnderlyingTokenAmount := seizedVTokenAmount.Mul(assets[seizedIndex].ExchangeRate).Div(EXPSACLE)
	seizedUnderlyingTokenValue := seizedUnderlyingTokenAmount.Mul(assets[seizedIndex].Price).Div(EXPSACLE)

	ratio := seizedUnderlyingTokenValue.Div(repayValue)
	if ratio.Cmp(decimal.NewFromFloat32(1.11)) == 1 || ratio.Cmp(decimal.NewFromFloat32(1.09)) == -1 {
		logger.Printf("calculated seizedUnerlyingTokenValue != 1.1, calculateRatio:%v, seizedUnderylingTokenValue:%v, repayValue:%v\n", ratio, seizedUnderlyingTokenValue, repayValue)
		//fmt.Errorf("calculated seizedUnerlyingTokenValue != 1.1, calculateRatio:%v, seizedUnderylingTokenValue:%v, repayValue:%v\n", ratio, seizedUnderlyingTokenValue, repayValue)
		//return err
	}

	massProfit := seizedUnderlyingTokenValue.Sub(repayValue)
	if massProfit.Cmp(ProfitThreshold) == -1 {
		logger.Printf("processLiquidationReq, profit:%v < 5 USDT, omit\n", massProfit.Div(EXPSACLE))
		return nil
	}

	flashLoanFeeRatio := decimal.NewFromInt(25).Div(decimal.NewFromInt(9975))
	flashLoanFeeAmount := repayAmount.Mul(flashLoanFeeRatio).Add(decimal.NewFromInt(1))
	flashLoanReturnAmount := repayAmount.Add(flashLoanFeeAmount.Truncate(0))

	bigGasPrice, err := s.c.SuggestGasPrice(context.Background())
	if err != nil {
		bigGasPrice = big.NewInt(5000000000)
	}

	gasPrice := decimal.NewFromBigInt(bigGasPrice, 0) //x gasPrice for PGA
	if withFeedPrice {
		if feededPrice.GasPrice.Cmp(decimal.Zero) == 1 {
			gasPrice = feededPrice.GasPrice //follow feedPrice Tx's gasPrice
		}
	} else {
		gasPrice = decimal.NewFromBigInt(bigGasPrice, 0).Mul(decimal.NewFromFloat32(1.02)) //x1.02 gasPrice for PGA
	}

	gasLimit := decimal.NewFromInt(int64(3000000))
	ethPrice := tokens[GasSymbol].Price

	if repaySymbol == "VAI" {
		addresses := []common.Address{
			VAIControllerAddress,
			tokens[seizedSymbol].Address,
			tokens[seizedSymbol].UnderlyingAddress,
			VAIAddress,
			account,
		}

		if isStalbeCoin(seizedSymbol) {
			//case6, repay VAI and get stablecoin
			//repay VAI, capture vUSDT, vBUSD, vDAI, swap part vUSDT/vBUSD/vDAI to VAI.
			path1, amountIn, err := s.selectPathWithMinAmountIn(seizedSymbol, repaySymbol, flashLoanReturnAmount)
			if err != nil {
				logger.Printf("processLiquidationReq case6, fail to get selectPathWithMinAmountIn, account:%v, seizedSymbol:%v, repaySymbol:%v, flashLoanReturnAmount:%v, err:%v\n", account, seizedSymbol, repaySymbol, flashLoanReturnAmount, err)
				return err
			}

			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, VAIAddress, flashLoanReturnAmount, path1, nil)
			if err != nil {
				logger.Printf("processLiquidationReq case6, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			remain := seizedUnderlyingTokenAmount.Sub(amountIn)
			if remain.Cmp(decimal.Zero) == -1 {
				err := fmt.Errorf("processLiquidationReq case6, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				logger.Printf("processLiquidationReq case6, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				return err
			}

			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			profit := remain.Mul(tokens[seizedSymbol].Price).Div(EXPSACLE).Sub(gasFee)
			logger.Printf("processLiquidationReq case6: repaySymbol is VAI and seizedSymbol is stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, path1, nil, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case6, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(6), flashLoanFrom, path1, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}
			}
			return nil
		} else {
			//case7, repay VAI and seizedSymbol is not stable coin. sell partly seizedSymbol to repay symbol, sell remain to usdt
			//case7.1 repay VAI, capture wBNB,  swap wBNB to VAI, swap wBNB to USDT, flashLoan from VAI-USDT pair
			//case7.2 repay VAI, capture wETH, swap wETH to VAI, swap wETH to USDT, flashLoan from VAI-wBNB pair
			path1, amountIn, err := s.selectPathWithMinAmountIn(seizedSymbol, repaySymbol, flashLoanReturnAmount)
			if err != nil {
				logger.Printf("processLiquidationReq case7, fail to get selectPathWithMinAmountIn, account:%v, seizedSymbol:%v, repaySymbol:%v, flashLoanReturnAmount:%v, err:%v\n", account, seizedSymbol, repaySymbol, flashLoanReturnAmount, err)
				return err
			}

			remain := seizedUnderlyingTokenAmount.Truncate(0).Sub(amountIn)
			if remain.Cmp(decimal.Zero) == -1 {
				err := fmt.Errorf("processLiquidationReq case7, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				logger.Printf("processLiquidationReq case7, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				return err
			}

			path2, amountOut, err := s.selectPathWithMaxAmountOut(seizedSymbol, ProfitSymbol, remain)
			if err != nil {
				logger.Printf("processLiquidationReq case7, fail to get selectPathWithMaxAmountOut, account:%v seizedSymbol:%v,ProfitSymbol:%v, err:%v\n", account, seizedSymbol, ProfitSymbol, err)
				return err
			}

			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, VAIAddress, flashLoanReturnAmount, path1, path2)
			if err != nil {
				logger.Printf("processLiquidationReq case7, fail to get %v selectFlashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			usdtAmount := amountOut
			profit := usdtAmount.Mul(tokens[ProfitSymbol].Price).Div(EXPSACLE).Sub(gasFee)

			logger.Printf("processLiquidationReq case7: repaySymbol is VAI and seizedSymbol is not stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, path1, path2, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case7: profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(7), flashLoanFrom, path1, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}
			}
		}
		return nil
	}

	addresses := []common.Address{
		tokens[repaySymbol].Address,
		tokens[seizedSymbol].Address,
		tokens[seizedSymbol].UnderlyingAddress,
		tokens[repaySymbol].UnderlyingAddress,
		account,
	}

	if seizedSymbol == repaySymbol {
		if isStalbeCoin(seizedSymbol) {
			//case1, seizedSymbol == repaySymbol and symbol is a stable coin
			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			remain := seizedUnderlyingTokenAmount.Sub(flashLoanReturnAmount)
			profit := (remain.Mul(tokens[seizedSymbol].Price).Div(EXPSACLE)).Sub(gasFee)
			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, tokens[repaySymbol].UnderlyingAddress, repayAmount, nil, nil)
			if err != nil {
				logger.Printf("processLiquidationReq case1, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			logger.Printf("processLiquidationReq case1: seizedSymbol == repaySymbol and symbol is a stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, nil, nil, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case1, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(1), flashLoanFrom, nil, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}

			}
		} else {
			//case2, seizedSymbol == repaySymbol and symbol is not a stable coin, after return flashloan, sell remain to usdt
			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			remain := seizedUnderlyingTokenAmount.Sub(flashLoanReturnAmount)

			path2, amountOut, err := s.selectPathWithMaxAmountOut(seizedSymbol, ProfitSymbol, remain)
			if err != nil {
				logger.Printf("processLiquidationReq case2:, fail to get selectPathWithMaxAmountOut, account:%v, seizedSymbol:%v,profitSymbol:%v, remain:%v, err:%v\n", account, seizedSymbol, ProfitSymbol, remain.Truncate(0), err)
				return err
			}

			usdtAmount := amountOut
			profit := (usdtAmount.Mul(tokens[ProfitSymbol].Price).Div(EXPSACLE)).Sub(gasFee)
			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, tokens[repaySymbol].UnderlyingAddress, repayAmount, nil, path2)
			if err != nil {
				logger.Printf("processLiquidationReq case2, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			logger.Printf("processLiquidationReq case2: seizedSymbol == repaySymbol and symbol is not a stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, nil, path2, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case2, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(2), flashLoanFrom, nil, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}
			}
		}
		return nil
	}

	if isStalbeCoin(seizedSymbol) {
		//case3, collateral(i.e. seizedSymbol) is stable coin, repaySymbol may or not be a stable coin, sell part of seized symbol to repaySymbol
		path1, amountIn, err := s.selectPathWithMinAmountIn(seizedSymbol, repaySymbol, flashLoanReturnAmount)
		if err != nil {
			logger.Printf("processLiquidationReq case3, fail to get selectPathWithMinAmountIn, account:%v, seizedSymbol:%v, repaySymbol:%v, flashLoanReturnAmount:%v, err:%v\n", account, seizedSymbol, repaySymbol, flashLoanReturnAmount, err)
			return err
		}

		gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
		remain := seizedUnderlyingTokenAmount.Sub(amountIn)
		if remain.Cmp(decimal.Zero) == -1 {
			err := fmt.Errorf("processLiquidationReq case3, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
			logger.Printf("processLiquidationReq case3, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
			return err
		}
		profit := remain.Mul(tokens[seizedSymbol].Price).Div(EXPSACLE).Sub(gasFee)
		flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, tokens[repaySymbol].UnderlyingAddress, repayAmount, path1, nil)
		if err != nil {
			logger.Printf("processLiquidationReq case3, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
			return err
		}

		logger.Printf("processLiquidationReq case3: seizedSymbol != repaySymbol and seizedSymbol stable coin\n")
		logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
		logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, path1, nil, addresses)

		if profit.Cmp(ProfitThreshold) == 1 {
			logger.Printf("case3, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
			tx, err := s.doLiquidation(big.NewInt(3), flashLoanFrom, path1, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
			if err != nil {
				logger.Printf("doLiquidation error:%v\n", err)
				db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				return err
			}
			if tx != nil {
				logger.Printf("tx success, hash:%v\n", tx.Hash())
				db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
			}
		}
	} else {
		if isStalbeCoin(repaySymbol) {
			//case4, collateral(i.e. seizedSymbol) is not stable coin, repaySymbol is a stable coin, sell all seizedSymbol to repaySymbol
			path1, amountOut, err := s.selectPathWithMaxAmountOut(seizedSymbol, repaySymbol, seizedUnderlyingTokenAmount)
			if err != nil {
				logger.Printf("processLiquidationReq case4, fail to get selectPathWithMaxAmountOut, account:%V, seizedSymbol:%v,repaySymbol:%v, flashLoanReturnAmount:%v, err:%v\n", account, seizedSymbol, repaySymbol, flashLoanReturnAmount, err)
				return err
			}

			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			remain := amountOut.Sub(flashLoanReturnAmount)
			if remain.Cmp(decimal.Zero) == -1 {
				err := fmt.Errorf("processLiquidationReq case4, unexpected paths selected, negative remain, seizedUnderlyingTokenAmount:%v, amountOut:%v, flashLoanReturnAmount:%v, remain:%v\n", seizedUnderlyingTokenAmount, amountOut, flashLoanReturnAmount, remain)
				logger.Printf("processLiquidationReq case4, unexpected paths selected, negative remain, seizedUnderlyingTokenAmount:%v, amountOut:%v, flashLoanReturnAmount:%v, remain:%v\n", seizedUnderlyingTokenAmount, amountOut, flashLoanReturnAmount, remain)
				return err
			}
			profit := remain.Mul(tokens[repaySymbol].Price).Div(EXPSACLE).Sub(gasFee)
			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, tokens[repaySymbol].UnderlyingAddress, repayAmount, path1, nil)
			if err != nil {
				logger.Printf("processLiquidationReq case4, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			logger.Printf("processLiquidationReq case4: seizedSymbol is not stable coin, repaySymbol is stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, path1, nil, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case4, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(4), flashLoanFrom, path1, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}
			}
		} else {
			//case5,  collateral(i.e. seizedSymbol) and repaySymbol are not stable coin. sell partly seizedSymbol to repay symbol, sell remain to usdt
			path1, amountIn, err := s.selectPathWithMinAmountIn(seizedSymbol, repaySymbol, flashLoanReturnAmount)
			if err != nil {
				logger.Printf("processLiquidationReq case5, fail to get selectPathWithMinAmountIn, account:%v, seizedSymbol:%v, repaySymbol:%v, flashLoanReturnAmount:%v, err:%v\n", account, seizedSymbol, repaySymbol, flashLoanReturnAmount, err)
				return err
			}
			remain := seizedUnderlyingTokenAmount.Truncate(0).Sub(amountIn)
			if remain.Cmp(decimal.Zero) == -1 {
				err := fmt.Errorf("processLiquidationReq case5, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				logger.Printf("processLiquidationReq case5, unexpected paths selected, negative remain,flashLoanReturnAmount:%v, amountIn:%v, seizedUnderlyingTokenAmount:%v, remain:%v\n", flashLoanReturnAmount, amountIn, seizedUnderlyingTokenAmount, remain)
				return err
			}
			// swap the remains to usdt
			path2, amountOut, err := s.selectPathWithMaxAmountOut(seizedSymbol, ProfitSymbol, remain)
			if err != nil {
				logger.Printf("processLiquidationReq case5, fail to get selectPathWithMaxAmountOut, account:%v seizedSymbol:%v, ProfitSymbol:%v:, err:%v\n", account, seizedSymbol, ProfitSymbol, err)
				return err
			}
			gasFee := gasPrice.Mul(gasLimit).Mul(ethPrice).Div(EXPSACLE)
			usdtAmount := amountOut
			profit := usdtAmount.Mul(tokens[ProfitSymbol].Price).Div(EXPSACLE).Sub(gasFee)
			flashLoanFrom, err := s.selectFlashLoanFrom(repaySymbol, tokens[repaySymbol].UnderlyingAddress, repayAmount, path1, path2)
			if err != nil {
				logger.Printf("processLiquidationReq case2, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
				return err
			}

			logger.Printf("processLiquidationReq case5: seizedSymbol and repaySymbol are not stable coin\n")
			logger.Printf("height:%v, account:%v, repaySymbol:%v, repayUnderlyingAmount:%v, seizedSymbol:%v, seizedVTokenAmount:%v, seizedUnderlyingAmount:%v, seizedValue:%v, flashLoanReturnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", currentHeight, account, repaySymbol, repayAmount, seizedSymbol, seizedVTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue, flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))
			logger.Printf("flashLoanFrom:%v, path1:%+v, path2:%+v, addresses:%+v\n", flashLoanFrom, path1, path2, addresses)

			if profit.Cmp(ProfitThreshold) == 1 {
				logger.Printf("case5: profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				tx, err := s.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt(), gasLimit.BigInt().Uint64())
				if err != nil {
					logger.Printf("doLiquidation error:%v\n", err)
					db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
					return err
				}
				if tx != nil {
					logger.Printf("tx success, hash:%v\n", tx.Hash())
					db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
				}
			}
		}
	}
	return nil
}

func (s *Syncer) updateDB(account common.Address, info AccountInfo) {
	s.deleteAccount(account)
	s.storeAccount(account, info)
}

func (s *Syncer) deleteAccount(account common.Address) {
	db := s.db
	accountBytes := account.Bytes()

	had, _ := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	if had {
		bz, _ := db.Get(dbm.AccountStoreKey(accountBytes), nil)
		var info AccountInfo
		err := json.Unmarshal(bz, &info)
		if err != nil {
			panic(err)
		}

		healthFactor := info.HealthFactor
		assets := info.Assets
		maxLoanValue := info.MaxLoanValue

		for _, asset := range assets {
			db.Delete(dbm.MarketMemberStoreKey([]byte(asset.Symbol), accountBytes), nil)
		}

		if maxLoanValue.Cmp(MaxLoanValueThreshold) == -1 {
			db.Delete(dbm.LiquidationNonProfitStoreKey(accountBytes), nil)
		} else {
			if healthFactor.Cmp(Decimal1P0) == -1 {
				db.Delete(dbm.LiquidationBelow1P0StoreKey(accountBytes), nil)
			} else if healthFactor.Cmp(Decimal1P1) == -1 {
				db.Delete(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
			} else if healthFactor.Cmp(Decimal1P5) == -1 {
				db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
			} else if healthFactor.Cmp(Decimal2P0) == -1 {
				db.Delete(dbm.LiquidationBelow2P0StoreKey(accountBytes), nil)
			} else {
				db.Delete(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
			}
		}

		db.Delete(dbm.AccountStoreKey(accountBytes), nil)
	}
}

func (s *Syncer) storeAccount(account common.Address, info AccountInfo) {
	db := s.db
	accountBytes := account.Bytes()
	healthFactor := info.HealthFactor
	maxLoanValue := info.MaxLoanValue

	for _, asset := range info.Assets {
		db.Put(dbm.MarketMemberStoreKey([]byte(asset.Symbol), accountBytes), accountBytes, nil)
	}

	if maxLoanValue.Cmp(MaxLoanValueThreshold) == -1 {
		db.Put(dbm.LiquidationNonProfitStoreKey(accountBytes), accountBytes, nil)
	} else {
		if healthFactor.Cmp(Decimal1P0) == -1 {
			db.Put(dbm.LiquidationBelow1P0StoreKey(accountBytes), accountBytes, nil)
		} else if healthFactor.Cmp(Decimal1P1) == -1 {
			db.Put(dbm.LiquidationBelow1P1StoreKey(accountBytes), accountBytes, nil)
		} else if healthFactor.Cmp(Decimal1P5) == -1 {
			db.Put(dbm.LiquidationBelow1P5StoreKey(accountBytes), accountBytes, nil)
		} else if healthFactor.Cmp(Decimal2P0) == -1 {
			db.Put(dbm.LiquidationBelow2P0StoreKey(accountBytes), accountBytes, nil)
		} else {
			db.Put(dbm.LiquidationAbove2P0StoreKey(accountBytes), accountBytes, nil)
		}
	}
	bz, _ := json.Marshal(info)
	db.Put(dbm.AccountStoreKey(accountBytes), bz, nil)
	db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
}

func copySymbols(src map[common.Address]string) map[common.Address]string {
	dst := make(map[common.Address]string)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyTokens(src map[string]*TokenInfo) map[string]*TokenInfo {
	dst := make(map[string]*TokenInfo)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func isStalbeCoin(symbol string) bool {
	return (symbol == "vUSDT" || symbol == "vDAI" || symbol == "vBUSD")
}

// path1: [seizedSymbol, intermediateSymbol, repaySymbol]
// path2: [seizedSymbol, "BUSD"]
func (s *Syncer) selectFlashLoanFrom(repaySymbol string, repayUnderlyingAddress common.Address, repayAmount decimal.Decimal, path1 []common.Address, path2 []common.Address) (common.Address, error) {
	var pair1, pair2 common.Address
	var err error
	len1 := len(path1)
	if len1 >= 2 {
		pair1, err = s.pancakeFactory.GetPair(nil, path1[len1-2], path1[len1-1])
		if err != nil {
			return common.Address{}, err
		}
	}

	len2 := len(path2)
	if len2 >= 2 {
		pair2, err = s.pancakeFactory.GetPair(nil, path2[len2-2], path2[len2-1])
		if err != nil {
			return common.Address{}, err
		}
	}

	bep20, err := venus.NewBep20(repayUnderlyingAddress, s.c)
	if err != nil {
		return common.Address{}, err
	}

	for _, pair := range s.flashLoanPools[repaySymbol] {
		if pair != pair1 && pair != pair2 && pair != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			balance, err := bep20.BalanceOf(nil, pair)
			if err != nil {
				continue
			}

			if decimal.NewFromBigInt(balance, 0).Cmp(repayAmount) == 1 {
				return pair, nil
			}
		}
	}

	return common.Address{}, fmt.Errorf("no flashLoanFrom")
}

func (s *Syncer) checkBadLiquidation(account common.Address, currentHeight uint64) error {
	db := s.db
	had, err := db.Has(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
	if err != nil {
		logger.Printf("checkBadLiquidation, fail to get BadLiquidationTx,err:%v\n", err)
		return err
	}
	if had {
		bz, err := db.Get(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
		if err != nil {
			logger.Printf("checkBadLiquidation, fail to get BadLiquidationTx,err:%v\n", err)
			return err
		}

		recordHeight := big.NewInt(0).SetBytes(bz).Uint64()
		if currentHeight-recordHeight <= ForbiddenPeriodForBadLiquidation {
			err = fmt.Errorf("checkBadLiquidation, forbidden bad %v liquidation temporay, currentHeight:%v, recordHeight:%v\n", account, currentHeight, recordHeight)
			logger.Printf("checkBadLiquidation, forbidden bad %v liquidation temporay, currentHeight:%v, recordHeight:%v\n", account, currentHeight, recordHeight)
			return err
		}
		db.Delete(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
	}
	return nil
}

func (s *Syncer) checkPendingLiquidation(account common.Address, currentHeight uint64) error {
	db := s.db
	//PendingLiquidationTx check
	had, err := db.Has(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
	if err != nil {
		logger.Printf("checkPendingLiquidation, fail to get PendingLiquidationTx,err:%v\n", err)
		return err
	}
	if had {
		bz, err := db.Get(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
		if err != nil {
			logger.Printf("checkPendingLiquidation, fail to get PendingLiquidationTx,err:%v\n", err)
			return err
		}

		recordHeight := big.NewInt(0).SetBytes(bz).Uint64()
		if currentHeight-recordHeight <= ForbiddenPeriodForPendingLiquidation {
			err = fmt.Errorf("checkPendingLiquidation, forbidden pending %v liquidation temporay, currentHeight:%v, recordHeight:%v\n", account, currentHeight, recordHeight)
			logger.Printf("checkPendingLiquidation, forbidden pending %v liquidation temporay, currentHeight:%v, recordHeight:%v\n", account, currentHeight, recordHeight)
			return err
		}
		db.Delete(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
	}
	return nil
}

func buildFlashLoanPool(factory *venus.IPancakeFactory, c *ethclient.Client, tokens map[string]*TokenInfo) map[string][]common.Address {
	flashLoanPool := make(map[string][]common.Address)
	peerSymbols := []string{"vBNB", "vBUSD", "vUSDT", "vETH", "vCAKE"}

	for symbol1, token1 := range tokens {
		if symbol1 == "vCAN" {
			continue
		}

		pairs := make([]common.Address, 0)
		balances := make([]decimal.Decimal, 0)
		bep20, _ := venus.NewBep20(token1.UnderlyingAddress, c)
		for _, symbol2 := range peerSymbols {
			pair, err := factory.GetPair(nil, token1.UnderlyingAddress, tokens[symbol2].UnderlyingAddress)
			if err != nil || pair == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				continue
			}

			balance, err := bep20.BalanceOf(nil, pair)
			if err != nil {
				continue
			}

			pairs = append(pairs, pair)
			balances = append(balances, decimal.NewFromBigInt(balance, 0))
		}

		length := len(pairs)
		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				if balances[i].Cmp(balances[j]) == -1 {
					balances[i], balances[j] = balances[j], balances[i]
					pairs[i], pairs[j] = pairs[j], pairs[i]
				}
			}
		}
		flashLoanPool[symbol1] = pairs
	}

	//{
	//	symbol1 := "VAI"
	//	pairs := make([]common.Address, 0)
	//	balances := make([]decimal.Decimal, 0)
	//	bep20, _ := venus.NewBep20(VAIAddress, c)
	//	for _, symbol2 := range peerSymbols {
	//		pair, err := factory.GetPair(nil, VAIAddress, tokens[symbol2].UnderlyingAddress)
	//		if err != nil || pair == common.HexToAddress("0x0000000000000000000000000000000000000000") {
	//			continue
	//		}
	//
	//		balance, err := bep20.BalanceOf(nil, pair)
	//		if err != nil {
	//			continue
	//		}
	//
	//		pairs = append(pairs, pair)
	//		balances = append(balances, decimal.NewFromBigInt(balance, 0))
	//	}
	//
	//	length := len(pairs)
	//	for i := 0; i < length-1; i++ {
	//		for j := i + 1; j < length; j++ {
	//			if balances[i].Cmp(balances[j]) == -1 {
	//				balances[i], balances[j] = balances[j], balances[i]
	//				pairs[i], pairs[j] = pairs[j], pairs[i]
	//			}
	//		}
	//	}
	//	flashLoanPool[symbol1] = pairs
	//}
	return flashLoanPool
}

func buildPaths(router *venus.IPancakeRouter02, tokens map[string]*TokenInfo) map[string][][]common.Address {
	candidatePaths := make([][]common.Address, 4)
	paths := make(map[string][][]common.Address, 0)

	amountIn := decimal.New(1000, 18).BigInt()

	for srcSymbol, srcToken := range tokens {
		for dstSymbol, dstToken := range tokens {
			if srcSymbol == dstSymbol || srcSymbol == "vCAN" || dstSymbol == "vCAN" {
				continue
			}

			amountOuts := make([][]*big.Int, 0)
			pathSet := make([][]common.Address, 0)
			candidatePaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
			candidatePaths[1] = []common.Address{srcToken.UnderlyingAddress, tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
			candidatePaths[2] = []common.Address{srcToken.UnderlyingAddress, tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
			candidatePaths[3] = []common.Address{srcToken.UnderlyingAddress, tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

			for i := 0; i < 4; i++ {
				candidateAmountsOut, err := router.GetAmountsOut(nil, amountIn, candidatePaths[i])
				if err == nil && candidateAmountsOut[len(candidatePaths[i])-1].Cmp(BigZero) == 1 {
					pathSet = append(pathSet, candidatePaths[i])
					amountOuts = append(amountOuts, candidateAmountsOut)
				}
			}

			for i := 0; i < len(pathSet)-1; i++ {
				for j := i + 1; j < len(pathSet); j++ {
					ilen := len(amountOuts[i])
					jLen := len(amountOuts[j])

					if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
						amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
						pathSet[i], pathSet[j] = pathSet[j], pathSet[i]
					}
				}
			}
			paths[srcSymbol+":"+dstSymbol] = pathSet
		}
	}

	//{
	//	srcSymbol := "VAI"
	//	for dstSymbol, dstToken := range tokens {
	//		if srcSymbol == dstSymbol || srcSymbol == "vCAN" || dstSymbol == "vCAN" {
	//			continue
	//		}
	//
	//		amountOuts := make([][]*big.Int, 0)
	//		pathSet := make([][]common.Address, 0)
	//		candidatePaths[0] = []common.Address{VAIAddress, dstToken.UnderlyingAddress}
	//		candidatePaths[1] = []common.Address{VAIAddress, tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
	//		candidatePaths[2] = []common.Address{VAIAddress, tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
	//		candidatePaths[3] = []common.Address{VAIAddress, tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}
	//
	//		for i := 0; i < 4; i++ {
	//			candidateAmountsOut, err := router.GetAmountsOut(nil, amountIn, candidatePaths[i])
	//			if err == nil && candidateAmountsOut[len(candidatePaths[i])-1].Cmp(BigZero) == 1 {
	//				pathSet = append(pathSet, candidatePaths[i])
	//				amountOuts = append(amountOuts, candidateAmountsOut)
	//			}
	//		}
	//
	//		for i := 0; i < len(pathSet)-1; i++ {
	//			for j := i + 1; j < len(pathSet); j++ {
	//				ilen := len(amountOuts[i])
	//				jLen := len(amountOuts[j])
	//
	//				if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
	//					amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
	//					pathSet[i], pathSet[j] = pathSet[j], pathSet[i]
	//				}
	//			}
	//		}
	//		paths[srcSymbol+":"+dstSymbol] = pathSet
	//	}
	//}
	//{
	//	dstSymbol := "VAI"
	//	for srcSymbol, srcToken := range tokens {
	//		if srcSymbol == dstSymbol || srcSymbol == "vCAN" || dstSymbol == "vCAN" {
	//			continue
	//		}
	//
	//		amountOuts := make([][]*big.Int, 0)
	//		pathSet := make([][]common.Address, 0)
	//		candidatePaths[0] = []common.Address{srcToken.UnderlyingAddress, VAIAddress}
	//		candidatePaths[1] = []common.Address{srcToken.UnderlyingAddress, tokens["vBNB"].UnderlyingAddress, VAIAddress}
	//		candidatePaths[2] = []common.Address{srcToken.UnderlyingAddress, tokens["vBUSD"].UnderlyingAddress, VAIAddress}
	//		candidatePaths[3] = []common.Address{srcToken.UnderlyingAddress, tokens["vUSDT"].UnderlyingAddress, VAIAddress}
	//
	//		for i := 0; i < 4; i++ {
	//			candidateAmountsOut, err := router.GetAmountsOut(nil, amountIn, candidatePaths[i])
	//			if err == nil && candidateAmountsOut[len(candidatePaths[i])-1].Cmp(BigZero) == 1 {
	//				pathSet = append(pathSet, candidatePaths[i])
	//				amountOuts = append(amountOuts, candidateAmountsOut)
	//			}
	//		}
	//
	//		for i := 0; i < len(pathSet)-1; i++ {
	//			for j := i + 1; j < len(pathSet); j++ {
	//				ilen := len(amountOuts[i])
	//				jLen := len(amountOuts[j])
	//
	//				if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
	//					amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
	//					pathSet[i], pathSet[j] = pathSet[j], pathSet[i]
	//				}
	//			}
	//		}
	//		paths[srcSymbol+":"+dstSymbol] = pathSet
	//	}
	//}

	return paths
}

func (s *Syncer) selectPathWithMaxAmountOut(srcSymbol, dstSymbol string, amountIn decimal.Decimal) ([]common.Address, decimal.Decimal, error) {
	paths := s.paths[srcSymbol+":"+dstSymbol]
	maxAmountOut := decimal.Zero
	var selected []common.Address

	for _, path := range paths {
		bigAmountsOut, err := s.pancakeRouter.GetAmountsOut(nil, amountIn.Truncate(0).BigInt(), path)
		if err != nil {
			continue
		}

		if decimal.NewFromBigInt(bigAmountsOut[len(bigAmountsOut)-1], 0).Cmp(maxAmountOut) == 1 {
			maxAmountOut = decimal.NewFromBigInt(bigAmountsOut[len(bigAmountsOut)-1], 0)
			selected = path
		}
	}
	if maxAmountOut.Cmp(decimal.Zero) == 0 {
		err := fmt.Errorf("%v:%v has no selectPathWithMaxAmountOut with amountIn:%v path", srcSymbol, dstSymbol, amountIn)
		return []common.Address{}, decimal.Zero, err
	}
	return selected, maxAmountOut, nil
}

func (s *Syncer) selectPathWithMinAmountIn(srcSymbol, dstSymbol string, amountOut decimal.Decimal) ([]common.Address, decimal.Decimal, error) {
	paths := s.paths[srcSymbol+":"+dstSymbol]

	minAmountIn := DecimalMax
	var selected []common.Address

	for _, path := range paths {
		bigAmountsIn, err := s.pancakeRouter.GetAmountsIn(nil, amountOut.Truncate(0).BigInt(), path)
		if err != nil {
			continue
		}

		if decimal.NewFromBigInt(bigAmountsIn[0], 0).Cmp(minAmountIn) == -1 {
			minAmountIn = decimal.NewFromBigInt(bigAmountsIn[0], 0)
			selected = path
		}
	}
	if minAmountIn.Cmp(DecimalMax) == 0 {
		err := fmt.Errorf("%v:%v has no selectPathWithMinAmountIn with amountOut:%v path", srcSymbol, dstSymbol, amountOut)
		return []common.Address{}, DecimalMax, err
	}

	return selected, minAmountIn, nil
}

//situcation： 情况 1-7
//ch： 借钱用的pair地址
//path1： 卖的时候的path, seizedSymbol => repaySymbol的path
//path2:  将seizedSymbol => USDT
//tokens：
// Tokens array
// [0] - _flashLoanVToken 要去借的钱（要还给venus的）
// [1] - _seizedVToken 可以赎回来的钱
// [2] - _seizedTokenUnderlying 赎回来的钱的underlying
// [3] - _flashloanTokenUnderlying 借的钱的underlying
// [4] - target 目标账号
//_flashLoanAmount ： 借多少？ 还多少？

func (s *Syncer) doLiquidation(scenarioNo *big.Int, flashLoanFrom common.Address, path1 []common.Address, path2 []common.Address, tokens []common.Address, flashLoanAmount *big.Int, gasPrice *big.Int, gasLimit uint64) (*types.Transaction, error) {
	publicKey := s.PrivateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, big.NewInt(56))
	auth.Value = big.NewInt(0)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = gasLimit

	tx, err := s.liquidator.Qingsuan(auth, scenarioNo, flashLoanFrom, path1, path2, tokens, flashLoanAmount)
	if err != nil {
		return nil, err
	}

	if tx == nil {
		return nil, fmt.Errorf("empty tx")
	}

	return tx, nil
}

func (info *AccountInfo) toReadable() AccountInfo {
	readableInfo := AccountInfo{}
	readableInfo.HealthFactor = info.HealthFactor
	readableInfo.MaxLoanValue = info.MaxLoanValue.Div(EXPSACLE)

	var readableAssets []Asset
	for _, asset := range info.Assets {
		var readableAsset Asset
		readableAsset.Symbol = asset.Symbol
		readableAsset.Balance = asset.Balance
		readableAsset.Loan = asset.Loan
		readableAsset.BalanceValue = asset.BalanceValue.Div(EXPSACLE)
		readableAsset.LoanValue = asset.LoanValue.Div(EXPSACLE)
		readableAssets = append(readableAssets, readableAsset)
	}
	readableInfo.Assets = readableAssets
	return readableInfo
}

func (info *ConcernedAccountInfo) toReadable() ConcernedAccountInfo {
	readableInfo := ConcernedAccountInfo{}
	readableInfo.FeededPrice = info.FeededPrice
	readableInfo.AccInfo = info.AccInfo.toReadable()
	return readableInfo
}
