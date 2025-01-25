package scanner

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/readygo586/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"math/big"
	"sort"
)

func (s *Scanner) LiquidationLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case account := <-s.liquidationCh:
			logger.Printf("receive priority liquidation:%v\n", account)
			s.liquidate(account)
		}
	}
}

func (s *Scanner) liquidate(info *AccountInfo) error {
	comptroller := s.comptroller
	account := info.Account
	db := s.db

	closeFactor := s.closeFactor

	//current height
	currentHeight, err := s.c.BlockNumber(context.Background())
	if err != nil {
		logger.Printf("processLiquidationReq, fail to get blockNumber,err:%v\n", err)
		return err
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

	//select the repay token and seized collateral token
	maxLoanValue := info.MaxLoanValue
	maxLoanMarket := info.MaxLoanMarket

	maxRepayValue := maxLoanValue.Mul(closeFactor.Factor).Div(EXPSACLE)
	repayMarket := maxLoanMarket
	repayValue := decimal.NewFromInt(0)

	assets := info.Assets
	sort.Slice(assets, func(i, j int) bool {
		return assets[i].BalanceValue.Cmp(assets[j].BalanceValue) == 1
	})

	seizedMarket := assets[0].Market
	repayValue = maxRepayValue
	if assets[0].BalanceValue.Cmp(maxRepayValue) == -1 {
		repayValue = assets[0].BalanceValue
	}

	bigSeizedVTokenAmount := big.NewInt(0)
	errCode := big.NewInt(0)

	var repayAmount decimal.Decimal

	//currently, repay VAI only
	if repayMarket == s.vaiControllerAddr {
		repayAmount = repayValue.Truncate(0).Mul(decimal.New(9999, -4)) //不能100%的repay,负责会报TOO_MUCH_REPAY的错误
		errCode, bigSeizedVTokenAmount, err = comptroller.LiquidateVAICalculateSeizeTokens(nil, seizedMarket, repayAmount.BigInt())
		if err != nil || errCode.Cmp(BigZero) != 0 {
			logger.Printf("processLiquidationReq, fail to get LiquidateVAICalculateSeizeTokens, account:%v, err:%v, errCode:%v\n", account, err, errCode)
			return err
		}
	} else {
		panic("not implemented")
	}

	seizedVTokenAmount := decimal.NewFromBigInt(bigSeizedVTokenAmount, 0)
	seizedUnderlyingTokenAmount := seizedVTokenAmount.Mul(assets[0].ExchangeRate).Div(EXPSACLE)
	seizedUnderlyingTokenValue := seizedUnderlyingTokenAmount.Mul(assets[0].Price).Div(EXPSACLE)

	ratio := seizedUnderlyingTokenValue.Div(repayValue)
	fmt.Printf("seizedUnderlyingTokenValue:%v, repayValue:%v, ratio:%v\n", seizedUnderlyingTokenValue, repayValue, ratio) //
	//
	massProfit := seizedUnderlyingTokenValue.Sub(repayValue)
	if massProfit.Cmp(ProfitThreshold) == -1 {
		logger.Printf("processLiquidationReq, profit:%v < 5 USD, omit\n", massProfit.Div(EXPSACLE))
		return nil
	}

	tx, err := s.doLiquidation(account, repayAmount.BigInt(), seizedMarket)
	if err != nil {
		logger.Printf("doLiquidation error:%v\n", err)
		db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
		return err
	}
	if tx != nil {
		logger.Printf("tx success, hash:%v\n", tx.Hash())
		db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
	}

	return nil
}

func (s *Scanner) checkBadLiquidation(account common.Address, currentHeight uint64) error {
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

func (s *Scanner) checkPendingLiquidation(account common.Address, currentHeight uint64) error {
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

func (s *Scanner) doLiquidation(borrower common.Address, repayVaiAmount *big.Int, seizedVToken common.Address) (*types.Transaction, error) {
	publicKey := s.PrivateKey.Public()
	vaiController := s.vaiController
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.c.SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = big.NewInt(5000000000)
	}

	gasLimit := uint64(3000000)

	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, big.NewInt(ChainID))
	auth.Value = big.NewInt(0)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = gasLimit

	tx, err := vaiController.LiquidateVAI(auth, borrower, repayVaiAmount, seizedVToken)
	if err != nil {
		return nil, err
	}

	if tx == nil {
		return nil, fmt.Errorf("empty tx")
	}

	return tx, nil
}
