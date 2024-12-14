package scanner

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/venus"
	"github.com/readygo67/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"github.com/syndtr/goleveldb/leveldb"
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
	Address            common.Address
	UnderlyingAddress  common.Address
	UnderlyingDecimals uint8
	CollateralFactor   decimal.Decimal
}

type TokenPrice struct {
	Price              decimal.Decimal
	PriceUpdatedHeight uint64
}

type PriceChanged struct {
	Address common.Address
	Price   decimal.Decimal
	Height  uint64
}

type CollateralFactorChanged struct {
	Address          common.Address
	CollateralFactor decimal.Decimal
}

type RepayVaiAmountChanged struct {
	Account common.Address
	Amount  decimal.Decimal
}

type VTokenAmountChanged struct {
	Address common.Address
	From    common.Address
	To      common.Address
	Amount  decimal.Decimal
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
	repayVaiAmountChangedCh   chan *RepayVaiAmountChanged
	vTokenAmountChangedCh     chan *VTokenAmountChanged //collateralAmount change, including mint, redeem, transfer
	priceChangedCh            chan *PriceChanged

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

	exist, err := db.Has(dbm.BorrowerNumberKey(), nil)
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
			Address:            common.HexToAddress(_vaiController),
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
		Address:            market,
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
