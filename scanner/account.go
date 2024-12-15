package scanner

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"math/big"
	"runtime"
	"sync"
)

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
	ProfitThreshold       = decimal.New(5, 18)           //5 USDT
	MaxLoanValueThreshold = decimal.New(100, 18)         //100 USDT
)

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
	MaxLoanMarket common.Address
	MaxLoanValue  decimal.Decimal
	VaiLoan       decimal.Decimal
	Height        uint64
	Assets        []Asset
}

type Liquidation struct {
	AccountInfo AccountInfo
}

func (s *Scanner) SyncAccountLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case accounts := <-s.topAccountSyncCh:
			s.processAccounts(accounts)

		case accounts := <-s.highAccountSyncCh:
			if len(s.topAccountSyncCh) != 0 {
				continue
			}
			s.processAccounts(accounts)

		case accounts := <-s.middleAccountSyncCh:
			if len(s.highAccountSyncCh) != 0 {
				continue
			}
			s.processAccounts(accounts)

		case accounts := <-s.lowAccountSyncCh:
			if len(s.middleAccountSyncCh) != 0 {
				continue
			}
			s.processAccounts(accounts)
		}
	}
}

func (s *Scanner) processAccounts(accounts []common.Address) {
	var wg sync.WaitGroup
	wg.Add(len(accounts))
	sem := make(semaphore, runtime.NumCPU())
	for _, account := range accounts {
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			err := s.syncOneAccount(account)
			if err != nil {
				logger.Printf("fail to syncOneAccount, account:%v err:%v\n", account, err)
			}
		}()
	}
	wg.Wait()
}

func (s *Scanner) syncOneAccount(account common.Address) error {
	//ctx := context.Background()
	comptroller := s.comptroller
	vaiController := s.vaiController

	_vaiLoan, err := vaiController.GetVAIRepayAmount(nil, account)
	if err != nil {
		logger.Printf("syncOneAccount, fail to get MintedVAIs, err:%v\n", err)
		return err
	}
	if _vaiLoan.Cmp(BigZero) == 0 {
		//shortcut, in jupiter, vai is the only borrowable asset
		return nil
	}
	vaiLoan := decimal.NewFromBigInt(_vaiLoan, 0)

	s.m.Lock()
	tokens := copyTokens(s.tokens)
	prices := copyPrices(s.prices)
	vbep20s := s.vbep20s
	s.m.Unlock()

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	var assets []Asset
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil || len(markets) == 0 {
		logger.Printf("syncOneAccount, fail to get GetAssetsIn or account is not in any markets, err:%v\n", err)
		return err
	}

	maxLoanValue := decimal.NewFromInt(0)
	maxLoanMarket := s.vaiMarket
	for _, market := range markets {
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil || errCode.Cmp(BigZero) != 0 {
			logger.Printf("syncOneAccount, fail to get GetAccountSnapshot, err:%v, errCode:%v\n", err, errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			//shortcut, no collateral and loan in this market, skip it
			continue
		}

		collateralFactor := tokens[market].CollateralFactor
		price := prices[market].Price

		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)

		multiplier := price.Mul(exchangeRate).Div(EXPSACLE).Div(EXPSACLE)
		balanceValue := balance.Mul(multiplier)
		collateral := balanceValue.Mul(collateralFactor).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateral.Truncate(0))

		borrow := decimal.Zero
		loan := decimal.Zero
		if bigBorrow.Cmp(BigZero) == 1 {
			borrow = decimal.NewFromBigInt(bigBorrow, 0)
			loan = borrow.Mul(price).Div(EXPSACLE)
			totalLoan = totalLoan.Add(loan.Truncate(0))
		}

		asset := Asset{
			Symbol:           tokens[market].Symbol,
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
			maxLoanMarket = market
		}
	}

	totalLoan = totalLoan.Add(vaiLoan)
	healthFactor := decimal.New(100, 0)
	if totalLoan.Cmp(decimal.Zero) == 1 {
		healthFactor = totalCollateral.Div(totalLoan)
	}

	if vaiLoan.Cmp(maxLoanValue) == 1 {
		maxLoanValue = vaiLoan
		maxLoanMarket = s.vaiMarket
	}

	currentHeight, _ := s.c.BlockNumber(context.Background())
	info := AccountInfo{
		Account:       account,
		HealthFactor:  healthFactor,
		MaxLoanValue:  maxLoanValue,
		MaxLoanMarket: maxLoanMarket,
		VaiLoan:       vaiLoan,
		Height:        currentHeight,
		Assets:        assets,
	}
	s.UpdateAccount(account, info)
	logger.Printf("syncOneAccount,account:%v, height:%v,totalCollateral:%v, totalLoan:%v,info:%+v\n", account, currentHeight, totalCollateral, totalLoan, info.toReadable())

	//trigger liquidation immediately
	errCode, _, shortfall, err := comptroller.GetAccountLiquidity(nil, account)
	if err == nil && errCode.Cmp(BigZero) == 0 && shortfall.Cmp(BigZero) == 1 {
		liquidation := &Liquidation{
			AccountInfo: info,
		}
		s.liquidationCh <- liquidation
	}

	return nil
}

func (s *Scanner) UpdateAccount(account common.Address, info AccountInfo) {
	s.deleteAccount(account)
	s.storeAccount(account, info)
}

func (s *Scanner) deleteAccount(account common.Address) {
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

		db.Delete(dbm.AccountStoreKey(accountBytes), nil)
	}
}

func (s *Scanner) storeAccount(account common.Address, info AccountInfo) {
	db := s.db
	accountBytes := account.Bytes()
	healthFactor := info.HealthFactor

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

	bz, _ := json.Marshal(info)
	db.Put(dbm.AccountStoreKey(accountBytes), bz, nil)
}

func copyTokens(src map[common.Address]*TokenInfo) map[common.Address]*TokenInfo {
	dst := make(map[common.Address]*TokenInfo)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyPrices(src map[common.Address]*TokenPrice) map[common.Address]*TokenPrice {
	dst := make(map[common.Address]*TokenPrice)
	for k, v := range src {
		dst[k] = v
	}
	return dst
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
