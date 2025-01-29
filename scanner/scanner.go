package scanner

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/db"
	"github.com/readygo586/LiquidationBot/venus"
	"github.com/shopspring/decimal"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"math/big"
	"os"
	"runtime"
	"sync"
	"time"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

type semaphore chan struct{}

type TokenInfo struct {
	Symbol             string
	Market             common.Address
	UnderlyingAddress  common.Address
	UnderlyingDecimals uint8
	CollateralFactor   decimal.Decimal
	UpdatedHeight      uint64
}

type TokenPrice struct {
	Price         decimal.Decimal
	UpdatedHeight uint64
}

type CloseFactor struct {
	Factor        decimal.Decimal
	UpdatedHeight uint64
}

type PriceChanged struct {
	Market        common.Address
	Price         decimal.Decimal
	UpdatedHeight uint64
}

type NewMarket struct {
	Market        common.Address
	UpdatedHeight uint64
}

type CloseFactorChanged struct {
	CloseFactor   decimal.Decimal
	UpdatedHeight uint64
}

type CollateralFactorChanged struct {
	Market           common.Address
	CollateralFactor decimal.Decimal
	UpdatedHeight    uint64
}

type EnterMarket struct {
	Market        common.Address
	Account       common.Address
	UpdatedHeight uint64
}

type ExitMarket struct {
	Market        common.Address
	Account       common.Address
	UpdatedHeight uint64
}

type RepayVaiAmountChanged struct {
	Account       common.Address
	Amount        decimal.Decimal
	UpdatedHeight uint64
}

type VTokenAmountChanged struct {
	Market        common.Address
	From          common.Address
	To            common.Address
	Amount        decimal.Decimal
	UpdatedHeight uint64
}

type Scanner struct {
	c  *ethclient.Client
	db *leveldb.DB

	//global setting
	comptrollerAddr   common.Address
	vaiControllerAddr common.Address
	oracleAddr        common.Address
	comptroller       *venus.Comptroller
	vaiController     *venus.VaiController
	vai               *venus.Vai
	oracle            *venus.Oracle
	feeders           []common.Address
	closeFactor       *CloseFactor

	//token info
	markets   []common.Address
	tokens    map[common.Address]*TokenInfo
	prices    map[common.Address]*TokenPrice
	vbep20s   map[common.Address]*venus.Vbep20
	feederMap map[common.Address]common.Address //feeder ->market map,

	//self privateKey and address
	PrivateKey *ecdsa.PrivateKey
	Account    common.Address
	vaiBalance decimal.Decimal

	//mutex, wg and channel
	m      sync.Mutex
	wg     sync.WaitGroup
	quitCh chan struct{}

	//event channel
	newMarketCh               chan *NewMarket
	closeFactorChangedCh      chan *CloseFactorChanged
	collateralFactorChangedCh chan *CollateralFactorChanged
	enterMarketCh             chan *EnterMarket
	exitMarketCh              chan *ExitMarket
	repayVaiAmountChangedCh   chan *RepayVaiAmountChanged
	vTokenAmountChangedCh     chan *VTokenAmountChanged //collateralAmount change, including mint, redeem, transfer
	priceChangedCh            chan *PriceChanged

	//four level account sync channels
	topAccountSyncCh    chan []common.Address
	highAccountSyncCh   chan []common.Address
	middleAccountSyncCh chan []common.Address
	lowAccountSyncCh    chan []common.Address

	//liquidation channel
	liquidationCh chan *AccountInfo
}

func init() {
}

func (s semaphore) Acquire() {
	s <- struct{}{}
}

func (s semaphore) Release() {
	<-s
}

func NewScanner(
	c *ethclient.Client,
	db *leveldb.DB,
	_comptroller string,
	_vaiController string,
	_vai string,
	_oracle string,
	_privateKey string,
) *Scanner {
	exist, _ := db.Has(dbm.BorrowerNumberKey(), nil)
	if !exist {
		db.Put(dbm.BorrowerNumberKey(), big.NewInt(0).Bytes(), nil)
	}

	comptroller, err := venus.NewComptroller(common.HexToAddress(_comptroller), c)
	if err != nil {
		panic(err)
	}

	vaiController, err := venus.NewVaiController(common.HexToAddress(_vaiController), c)
	if err != nil {
		panic(err)
	}

	vai, err := venus.NewVai(common.HexToAddress(_vai), c)
	if err != nil {
		panic(err)
	}

	oracle, err := venus.NewOracle(common.HexToAddress(_oracle), c)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		panic(err)
	}
	_publicKey := privateKey.Public()
	publicKey, ok := _publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("failed to cast public key to ECDSA")
	}
	account := crypto.PubkeyToAddress(*publicKey)

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		panic(err)
	}

	var feeders []common.Address
	feedersMap := make(map[common.Address]common.Address)
	for _, market := range markets {
		feeder, err := oracle.Feeder(nil, market)
		if err != nil {
			panic(err)
		}
		feeders = append(feeders, feeder)
		feedersMap[feeder] = market
	}

	//collect all markets information in parallel
	tokens := make(map[common.Address]*TokenInfo)
	prices := make(map[common.Address]*TokenPrice)
	vbep20s := make(map[common.Address]*venus.Vbep20)

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(markets))

	sem := make(semaphore, runtime.NumCPU())
	for _, market := range markets {
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			token, price, err := newMarket(c, comptroller, oracle, market)
			if err != nil {
				panic(err)
			}
			vbep20, err := venus.NewVbep20(market, c) //vBep20
			if err != nil {
				panic(err)
			}

			logger.Printf("symbol:%v, market:%v, underlying:%v ", token.Symbol, market, token.UnderlyingAddress)
			m.Lock()
			tokens[market] = token
			prices[market] = price
			vbep20s[market] = vbep20
			m.Unlock()
		}()
	}
	wg.Wait()

	height, err := c.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	//special processing for vai
	{
		vaiSymbol, err := vai.Symbol(nil)
		if err != nil {
			panic(err)
		}

		vaiDecimals, err := vai.Decimals(nil)
		if err != nil {
			panic(err)
		}

		tokens[common.HexToAddress(_vaiController)] = &TokenInfo{
			Symbol:             vaiSymbol,
			Market:             common.HexToAddress(_vaiController),
			UnderlyingAddress:  common.HexToAddress(_vai),
			UnderlyingDecimals: vaiDecimals,
			CollateralFactor:   decimal.Zero,
			UpdatedHeight:      height,
		}

		prices[common.HexToAddress(_vai)] = &TokenPrice{
			Price: decimal.New(1, 18),
		}
	}

	_closeFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		panic(err)
	}

	closeFactor := &CloseFactor{
		Factor:        decimal.NewFromBigInt(_closeFactor, 0),
		UpdatedHeight: height,
	}

	_vaiBalance, err := vai.BalanceOf(nil, account)
	if err != nil {
		panic(err)
	}
	vaiBalance := decimal.NewFromBigInt(_vaiBalance, 0)
	//_vaiLoan, err := vaiController.GetVAIRepayAmount(nil, account)
	//if err != nil {
	//	panic(err)
	//}
	//logger.Printf("vaiBalance:%v, vaiLoan:%v\n", _vaiBalance, _vaiLoan)
	return &Scanner{
		c:  c,
		db: db,

		comptrollerAddr:   common.HexToAddress(_comptroller),
		vaiControllerAddr: common.HexToAddress(_vaiController),
		oracleAddr:        common.HexToAddress(_oracle),
		feeders:           feeders,
		comptroller:       comptroller,
		vaiController:     vaiController,
		vai:               vai,
		oracle:            oracle,
		closeFactor:       closeFactor,

		markets:   markets,
		tokens:    tokens,
		prices:    prices,
		vbep20s:   vbep20s,
		feederMap: feedersMap,

		PrivateKey: privateKey,
		Account:    account,
		vaiBalance: vaiBalance,

		m:      m,
		quitCh: make(chan struct{}),

		newMarketCh:               make(chan *NewMarket, 8),
		closeFactorChangedCh:      make(chan *CloseFactorChanged, 8),
		collateralFactorChangedCh: make(chan *CollateralFactorChanged, 8),
		enterMarketCh:             make(chan *EnterMarket, 512),
		exitMarketCh:              make(chan *ExitMarket, 512),
		repayVaiAmountChangedCh:   make(chan *RepayVaiAmountChanged, 1024),
		vTokenAmountChangedCh:     make(chan *VTokenAmountChanged, 1024),
		priceChangedCh:            make(chan *PriceChanged, 1024),

		//account sync channel
		topAccountSyncCh:    make(chan []common.Address, 512),
		highAccountSyncCh:   make(chan []common.Address, 512),
		middleAccountSyncCh: make(chan []common.Address, 512),
		lowAccountSyncCh:    make(chan []common.Address, 512),

		//liquidation channel
		liquidationCh: make(chan *AccountInfo, 1024),
	}
}

func (s *Scanner) Start() {
	logger.Printf("server start")

	s.wg.Add(11)
	go s.ScanLoop()

	//event processors
	go s.NewMarketLoop()
	go s.CloseFactorLoop()
	go s.CollateralFactorLoop()
	go s.EnterMarketLoop()
	go s.ExitMarketLoop()
	go s.RepayVaiAmountChangedLoop()
	go s.VTokenAmountChangedLoop()
	go s.PriceChangedLoop()

	//sync account
	go s.SyncAccountLoop()

	//liquidation
	go s.LiquidationLoop()
}

func (s *Scanner) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Scanner) ScanLoop() {
	defer s.wg.Done()
	t := time.NewTimer(0)
	defer t.Stop()
	db := s.db
	c := s.c

	for {
		select {
		case <-s.quitCh:
			return

		case <-t.C:
			currentHeight, err := c.BlockNumber(context.Background())
			if err != nil {
				t.Reset(time.Second * 3)
				continue
			}

			bz, err := db.Get(dbm.LatestHandledHeightStoreKey(), nil)
			if err != nil {
				t.Reset(time.Millisecond * 20)
				continue
			}

			latestHandledHeight := big.NewInt(0).SetBytes(bz).Int64()
			startHeight := latestHandledHeight + 1
			endHeight := min(int64(currentHeight), startHeight+100) //scan 100 blocks each round, give chance to respond quitCh
			logger.Printf("startHeight:%v, endHeight:%v, currentHeight:%v", startHeight, endHeight, currentHeight)

			if startHeight+ConfirmHeight >= endHeight {
				t.Reset(time.Second * 3)
				continue
			}

			query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.feeders)
			query2 := buildVTokenQueryWithoutHeight(s.markets)

			for height := startHeight; height < endHeight; height++ {
				err := s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				err = db.Put(dbm.LatestHandledHeightStoreKey(), big.NewInt(height).Bytes(), nil)
				if err != nil {
					goto EndWithoutUpdateHeight
				}
				log.Printf("successfully scanned block:%v", height)
			}

		EndWithoutUpdateHeight:
			t.Reset(time.Millisecond * 20)
		}
	}
}

func (s *Scanner) NewMarketLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.newMarketCh:
			//since there is no method to remove a listed market, a history listed market must exist now
			logger.Printf("NewMarketLoop, new market:%v\n", event)
			exist := false
			market := event.Market
			for _, _market := range s.markets {
				if _market == market {
					exist = true
					break
				}
			}

			if !exist {
				token, price, err1 := newMarket(s.c, s.comptroller, s.oracle, market)
				vbep20, err2 := venus.NewVbep20(market, s.c) //vBep20

				if err1 != nil || err2 != nil {
					logger.Printf("fail to add newMarket:%v err1:%v,err2:%v", market, err1, err2)
					time.Sleep(5 * time.Second)
					logger.Printf("retry to add newMarket:%v event:%v", market, event)
					s.newMarketCh <- event //retry
				} else {
					s.m.Lock()
					s.markets = append(s.markets, market)
					s.tokens[market] = token
					s.prices[market] = price
					s.vbep20s[market] = vbep20
					s.m.Unlock()
				}
			}
		}
	}
}

func (s *Scanner) CloseFactorLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return
		case event := <-s.closeFactorChangedCh:
			logger.Printf("CloseFactorLoop, new closeFactor:%v\n", event)
			if event.UpdatedHeight > s.closeFactor.UpdatedHeight {
				//only update closeFactor if event.UpdatedHeight > s.closeFactor.UpdatedHeight
				s.m.Lock()
				s.closeFactor = &CloseFactor{
					Factor:        event.CloseFactor,
					UpdatedHeight: event.UpdatedHeight,
				}
				s.m.Unlock()
			}
		}
	}
}

func (s *Scanner) CollateralFactorLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.collateralFactorChangedCh:
			logger.Printf("CollateralFactorLoop, new collateralFactor:%v\n", event)
			if s.tokens[event.Market] == nil {
				log.Fatalf("fail to find market %v", event.Market)
			}

			market := event.Market
			if event.UpdatedHeight > s.tokens[market].UpdatedHeight {
				if event.CollateralFactor.Cmp(s.tokens[market].CollateralFactor) == 0 {
					continue
				}

				//collect affected accounts, and recalculate their health factor
				var accounts []common.Address
				prefix := append(dbm.MarketMemberPrefix, market.Bytes()...)
				iter := db.NewIterator(util.BytesPrefix(prefix), nil)
				for iter.Next() {
					accounts = append(accounts, common.BytesToAddress(iter.Value()))
				}
				iter.Release()

				if event.CollateralFactor.Cmp(s.tokens[market].CollateralFactor) == -1 {
					s.highAccountSyncCh <- accounts
				} else {
					s.middleAccountSyncCh <- accounts
				}

				s.m.Lock()
				s.tokens[market].CollateralFactor = event.CollateralFactor
				s.tokens[market].UpdatedHeight = event.UpdatedHeight
				s.m.Unlock()
			}
		}
	}
}

func (s *Scanner) EnterMarketLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.enterMarketCh:
			logger.Printf("EnterMarketLoop:%v\n", event)
			market := event.Market.Bytes()
			account := event.Account.Bytes()
			db.Put(dbm.MarketMemberStoreKey(market, account), account, nil)
			s.middleAccountSyncCh <- []common.Address{event.Account}
		}
	}
}

func (s *Scanner) ExitMarketLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.exitMarketCh:
			logger.Printf("ExistMarketLoop:%v\n", event)
			market := event.Market.Bytes()
			account := event.Account.Bytes()
			db.Delete(dbm.MarketMemberStoreKey(market, account), nil)
			s.highAccountSyncCh <- []common.Address{event.Account}
		}
	}
}

func (s *Scanner) RepayVaiAmountChangedLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.repayVaiAmountChangedCh:
			logger.Printf("RepayVaiAmountChangedLoop:%v\n", event)
			db.Put(dbm.BorrowersStoreKey(event.Account.Bytes()), event.Account.Bytes(), nil)
			if event.Amount.IsPositive() {
				s.highAccountSyncCh <- []common.Address{event.Account}
			} else {
				s.middleAccountSyncCh <- []common.Address{event.Account}
			}
		}
	}
}

func (s *Scanner) VTokenAmountChangedLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case event := <-s.vTokenAmountChangedCh:
			logger.Printf("VTokenAmountChangedLoop:%v\n", event)
			market := event.Market
			from := event.From
			to := event.To
			if from == to {
				continue
			}

			had, err := db.Has(dbm.MarketMemberStoreKey(market.Bytes(), from.Bytes()), nil)
			if err != nil {
				log.Fatal(err)
			}
			if had {
				s.highAccountSyncCh <- []common.Address{from}
			}

			had, err = db.Has(dbm.MarketMemberStoreKey(market.Bytes(), to.Bytes()), nil)
			if err != nil {
				log.Fatal(err)
			}
			if had {
				s.middleAccountSyncCh <- []common.Address{to}
			}
		}
	}
}

func (s *Scanner) PriceChangedLoop() {
	db := s.db
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return
		case event := <-s.priceChangedCh:
			logger.Printf("PriceChangedLoop:%v\n", event)
			if event.UpdatedHeight > s.prices[event.Market].UpdatedHeight {
				market := event.Market

				//collect affected accounts, and recalculate their health factor
				var accounts []common.Address
				prefix := append(dbm.MarketMemberPrefix, market.Bytes()...)
				iter := db.NewIterator(util.BytesPrefix(prefix), nil)
				for iter.Next() {
					accounts = append(accounts, common.BytesToAddress(iter.Value()))
				}
				iter.Release()

				if event.Price.Cmp(s.prices[market].Price) == -1 {
					s.highAccountSyncCh <- accounts
				} else {
					s.middleAccountSyncCh <- accounts
				}

				s.m.Lock()
				s.prices[market].Price = event.Price
				s.prices[market].UpdatedHeight = event.UpdatedHeight
				s.m.Unlock()
			}
		}
	}
}

func (s *Scanner) ScanOneBlock(height int64, querys []ethereum.FilterQuery) error {
	c := s.c
	h := big.NewInt(height)

	logs := []types.Log{}
	for _, query := range querys {
		query.FromBlock = h
		query.ToBlock = h
		_logs, err := c.FilterLogs(context.Background(), query)
		if err != nil {
			return err
		}
		logs = append(logs, _logs...)
	}

	//IMPORTANT: Do not use thread to decode log, which may break the sequence of logs
	for _, log := range logs {
		s.DecodeLog(log)
	}

	return nil
}

func (s *Scanner) ScanBlockBySpan(from, to int64, querys []ethereum.FilterQuery) error {
	c := s.c
	_from := big.NewInt(from)
	_to := big.NewInt(to)

	logs := []types.Log{}
	for _, query := range querys {
		query.FromBlock = _from
		query.ToBlock = _to
		_logs, err := c.FilterLogs(context.Background(), query)
		if err != nil {
			return err
		}
		logs = append(logs, _logs...)
	}

	//IMPORTANT: Do not use thread to decode log, which may break the sequence of logs of
	for _, log := range logs {
		s.DecodeLog(log)
	}
	return nil
}

func newMarket(c *ethclient.Client, comptroller *venus.Comptroller, oracle *venus.Oracle, market common.Address) (*TokenInfo, *TokenPrice, error) {
	vbep20, err := venus.NewVbep20(market, c) //vBep20
	if err != nil {
		return nil, nil, err
	}

	symbol, err := vbep20.Symbol(nil)
	if err != nil {
		return nil, nil, err
	}

	marketDetail, err := comptroller.Markets(nil, market)
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}
	}

	bep20, err := venus.NewVbep20(underlyingAddress, c)
	underlyingDecimals, err := bep20.Decimals(nil)
	if err != nil {
		return nil, nil, err
	}

	collateralFactor := decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
	price := decimal.NewFromBigInt(bigPrice, 0)

	height, err := c.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, err
	}

	tokenInfo := &TokenInfo{
		Symbol:             symbol,
		Market:             market,
		UnderlyingAddress:  underlyingAddress,
		UnderlyingDecimals: underlyingDecimals,
		CollateralFactor:   collateralFactor,
		UpdatedHeight:      height,
	}
	tokenPrice := &TokenPrice{
		Price:         price,
		UpdatedHeight: height,
	}

	return tokenInfo, tokenPrice, nil
}
