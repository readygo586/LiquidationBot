package scanner

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
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
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

type semaphore chan struct{}

type TokenInfo struct {
	Symbol             string
	Market             common.Address
	UnderlyingAddress  common.Address
	UnderlyingDecimals uint8
	CollateralFactor   decimal.Decimal
}

type TokenPrice struct {
	Price              decimal.Decimal
	PriceUpdatedHeight uint64
}

type PriceChanged struct {
	Market common.Address
	Price  decimal.Decimal
	Height uint64
}

type CollateralFactorChanged struct {
	Market           common.Address
	CollateralFactor decimal.Decimal
}

type EnterMarket struct {
	Market  common.Address
	Account common.Address
}

type ExitMarket struct {
	Market  common.Address
	Account common.Address
}

type RepayVaiAmountChanged struct {
	Account common.Address
	Amount  decimal.Decimal
}

type VTokenAmountChanged struct {
	Market common.Address
	From   common.Address
	To     common.Address
	Amount decimal.Decimal
}

type Scanner struct {
	c  *ethclient.Client
	db *leveldb.DB

	//global setting
	comptroller   *venus.Comptroller
	vaiController *venus.VaiController
	vai           *venus.Vai
	oracle        *venus.PriceOracle
	closeFactor   decimal.Decimal

	//token info
	markets []common.Address
	tokens  map[common.Address]*TokenInfo
	prices  map[common.Address]*TokenPrice
	vbep20s map[common.Address]*venus.Vbep20

	//self privateKey and address
	liquidator *venus.IQingsuan
	PrivateKey *ecdsa.PrivateKey
	Account    common.Address
	vaiBalance decimal.Decimal

	//mutex, wg and channel
	m                         sync.Mutex
	wg                        sync.WaitGroup
	quitCh                    chan struct{}
	newMarketCh               chan common.Address
	closeFactorChangedCh      chan decimal.Decimal
	collateralFactorChangedCh chan *CollateralFactorChanged
	enterMarketCh             chan *EnterMarket
	exitMarketCh              chan *ExitMarket
	repayVaiAmountChangedCh   chan *RepayVaiAmountChanged
	vTokenAmountChangedCh     chan *VTokenAmountChanged //collateralAmount change, including mint, redeem, transfer
	priceChangedCh            chan *PriceChanged

	topAccountSyncCh    chan common.Address
	highAccountSyncCh   chan common.Address
	middleAccountSyncCh chan common.Address
	lowAccountSyncCh    chan common.Address

	//liquidationCh             chan *Liquidation
	//priortyLiquidationCh      chan *Liquidation
	//concernedAccountInfoCh    chan *ConcernedAccountInfo
	//backgroundAccountSyncCh   chan common.Address
	//lowPriorityAccountSyncCh  chan *AccountsWithFeededPrice
	//highPriorityAccountSyncCh chan *AccountsWithFeededPrice
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

	bigCloseFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		panic(err)
	}

	closeFactor := decimal.NewFromBigInt(bigCloseFactor, 0)

	vaiController, err := venus.NewVaiController(common.HexToAddress(_vaiController), c)
	if err != nil {
		panic(err)
	}

	vai, err := venus.NewVai(common.HexToAddress(_vai), c)
	if err != nil {
		panic(err)
	}

	oracle, err := venus.NewPriceOracle(common.HexToAddress(_oracle), c)
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

	//collect all markets information parrallel
	tokens := make(map[common.Address]*TokenInfo)
	prices := make(map[common.Address]*TokenPrice)
	vbep20s := make(map[common.Address]*venus.Vbep20)

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(markets))

	sem := make(semaphore, runtime.NumCPU())
	for _, _market := range markets {
		market := _market
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

		tokens[common.HexToAddress(_vai)] = &TokenInfo{
			Symbol:             vaiSymbol,
			Market:             common.HexToAddress(_vaiController),
			UnderlyingAddress:  common.HexToAddress(_vai),
			UnderlyingDecimals: vaiDecimals,
			CollateralFactor:   decimal.Zero,
		}

		prices[common.HexToAddress(_vai)] = &TokenPrice{
			Price: decimal.New(1, 18),
		}
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

	//TODO(fix), calculate health factor ??

	return &Scanner{
		c:                         c,
		db:                        db,
		comptroller:               comptroller,
		vaiController:             vaiController,
		vai:                       vai,
		oracle:                    oracle,
		closeFactor:               closeFactor,
		markets:                   markets,
		tokens:                    tokens,
		prices:                    prices,
		vbep20s:                   vbep20s,
		PrivateKey:                privateKey,
		Account:                   account,
		vaiBalance:                vaiBalance,
		m:                         m,
		quitCh:                    make(chan struct{}),
		newMarketCh:               make(chan common.Address, 64),
		closeFactorChangedCh:      make(chan decimal.Decimal, 8),
		collateralFactorChangedCh: make(chan *CollateralFactorChanged, 8),
		repayVaiAmountChangedCh:   make(chan *RepayVaiAmountChanged, 1024),
		vTokenAmountChangedCh:     make(chan *VTokenAmountChanged, 1024),
		priceChangedCh:            make(chan *PriceChanged, 1024),
		//liquidationCh:             make(chan *Liquidation, 1024),
		//priortyLiquidationCh:      make(chan *Liquidation, 1024),
		//concernedAccountInfoCh:    make(chan *ConcernedAccountInfo, 4096),
		//backgroundAccountSyncCh:   make(chan common.Address, 8192),
		//lowPriorityAccountSyncCh:  make(chan *AccountsWithFeededPrice, 1024),
		//highPriorityAccountSyncCh: make(chan *AccountsWithFeededPrice, 248),
	}
}

func (s *Scanner) Start() {
	logger.Printf("server start")

	//s.wg.Add(11)
	//go s.SyncCloseFactorLoop()
	//go s.SyncMarketsLoop()
	//go s.SyncPriceLoop()
	////go s.SyncMarketsAndPricesLoop()
	//go s.ProcessFeededPriceLoop()
	//go s.SearchNewBorrowerLoop()
	//go s.BackgroundSyncLoop()
	//go s.syncAccountLoop()
	//go s.MonitorLiquidationEventLoop()
	//go s.PrintConcernedAccountInfoLoop()
	//go s.MonitorTxPoolLoop()
	//go s.ProcessLiquidationLoop()
}

func (s *Scanner) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Scanner) NewMarketLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case market := <-s.newMarketCh:
			exist := false
			for _, _market := range s.markets {
				if _market == market {
					exist = true
					break
				}
			}

			if !exist {
				token, price, err := newMarket(s.c, s.comptroller, s.oracle, market)
				if err != nil {
					logger.Fatalf("fail to add newMarket %v err:%v", market, err)
				}

				//TODO(keep), store the markets or not
				//db := s.db
				//db.Put(dbm.MarketStoreKey(market.Bytes()), market.Bytes(), nil)

				s.m.Lock()
				s.markets = append(s.markets, market)
				s.tokens[market] = token
				s.prices[market] = price
				s.m.Unlock()
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
		case newCloseFactor := <-s.closeFactorChangedCh:
			s.m.Lock()
			s.closeFactor = newCloseFactor
			s.m.Unlock()

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

		case newFactor := <-s.collateralFactorChangedCh:
			if s.tokens[newFactor.Market] == nil {
				log.Fatalf("fail to find market %v", newFactor.Market)
			}

			s.m.Lock()
			oldCollateralFactor := s.tokens[newFactor.Market].CollateralFactor
			s.tokens[newFactor.Market].CollateralFactor = newFactor.CollateralFactor
			s.m.Unlock()

			//if new collateralFactor is less than oldCollateralFactor, then recalculate all affected accounts health factor
			if oldCollateralFactor.GreaterThan(newFactor.CollateralFactor) {
				var accounts []common.Address
				iter := db.NewIterator(util.BytesPrefix(dbm.MarketMemberPrefix), nil)
				for iter.Next() {
					accounts = append(accounts, common.BytesToAddress(iter.Value()))
				}
				iter.Release()

				for _, account := range accounts {
					s.highAccountSyncCh <- account
				}
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

		case enter := <-s.enterMarketCh:
			market := enter.Market.Bytes()
			account := enter.Account.Bytes()
			db.Put(dbm.MarketMemberStoreKey(market, account), account, nil)
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

		case exit := <-s.exitMarketCh:
			market := exit.Market.Bytes()
			account := exit.Account.Bytes()
			db.Delete(dbm.MarketMemberStoreKey(market, account), nil)
		}
	}
}

func (s *Scanner) RepayVaiAmountChangedLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return
		case change := <-s.repayVaiAmountChangedCh:
			if change.Amount.IsPositive() {
				s.highAccountSyncCh <- change.Account
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
		case change := <-s.vTokenAmountChangedCh:
			//market := change.Market
			from := change.From
			//to := change.To

			had, err := db.Has(dbm.BorrowersStoreKey(from.Bytes()), nil)
			if err != nil {
				log.Fatal(err)
			}
			if had {
				s.highAccountSyncCh <- from
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
		case change := <-s.priceChangedCh:
			market := change.Market
			var accounts []common.Address
			prefix := append(dbm.MarketMemberPrefix, market.Bytes()...)
			iter := db.NewIterator(util.BytesPrefix(prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			for _, account := range accounts {
				s.highAccountSyncCh <- account
			}
		}
	}
}

func newMarket(c *ethclient.Client, comptroller *venus.Comptroller, oracle *venus.PriceOracle, market common.Address) (*TokenInfo, *TokenPrice, error) {
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
	}
	tokenPrice := &TokenPrice{
		Price:              price,
		PriceUpdatedHeight: height,
	}

	return tokenInfo, tokenPrice, nil
}
