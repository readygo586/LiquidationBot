package scanner

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
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

type PendingLiquidation struct {
	Hash   common.Hash
	Height uint64
}

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
	var actualRepayValue decimal.Decimal
	//currently, repay VAI only
	if repayMarket == s.vaiControllerAddr {
		repayAmount = repayValue.Truncate(0).Mul(decimal.New(9999, -4)) //不能100%的repay,负责会报TOO_MUCH_REPAY的错误
		actualRepayValue = repayValue.Truncate(0).Mul(decimal.New(9999, -4))
		publicKey := s.PrivateKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		repayer := crypto.PubkeyToAddress(*publicKeyECDSA)

		_vaiBalance, err := s.vai.BalanceOf(nil, repayer)
		if err != nil {
			return err
		}

		vaiBalance := decimal.NewFromBigInt(_vaiBalance, 0)
		if vaiBalance.Cmp(repayAmount) == -1 {
			repayAmount = vaiBalance
		}

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

	ratio := seizedUnderlyingTokenValue.Div(actualRepayValue)
	fmt.Printf("seizedUnderlyingTokenValue:%v, actualRepayValue:%v, ratio:%v\n", seizedUnderlyingTokenValue, actualRepayValue, ratio) //
	//
	massProfit := seizedUnderlyingTokenValue.Sub(actualRepayValue)
	if massProfit.Cmp(ProfitThreshold) == -1 {
		logger.Printf("processLiquidationReq, profit:%v < 1 USD, omit\n", massProfit.Div(EXPSACLE))
		return nil
	}

	tx, err := s.doLiquidation(account, repayAmount.BigInt(), seizedMarket)
	if err != nil {
		logger.Printf("doLiquidation error:%v\n", err)
		db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
		return err
	}
	if tx != nil {
		pendingLiquidation := PendingLiquidation{
			Hash:   tx.Hash(),
			Height: currentHeight,
		}
		logger.Printf("sending liquidationtx success, hash:%v\n", tx.Hash())
		bz, err := json.Marshal(pendingLiquidation)
		if err != nil {
			logger.Printf("marshal pendingLiquidation error:%v\n", err)
			return err
		}
		db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), bz, nil)
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
		var pendingTx PendingLiquidation
		err = json.Unmarshal(bz, &pendingTx)
		if err != nil {
			logger.Printf("checkPendingLiquidation, fail to unmarshal PendingLiquidationTx,err:%v\n", err)
			return err
		}

		receipt, err := s.c.TransactionReceipt(context.Background(), pendingTx.Hash)
		if err != nil {
			logger.Printf("checkPendingLiquidation, fail to get PendingLiquidationTx:%v,err:%v\n", pendingTx, err)
			return err
		}
		if receipt != nil {
			//previous tx has been executed
			db.Delete(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
			if receipt.Status == types.ReceiptStatusFailed {
				db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(pendingTx.Height)).Bytes(), nil)
				err = fmt.Errorf("checkPendingLiquidation, %v fail", pendingTx)
				return err
			}
			return nil
		} else {
			err = fmt.Errorf("checkPendingLiquidation, PendingLiquidationTx:%v is still pending", pendingTx)
			return err
		}
	}
	return nil
}

func (s *Scanner) doLiquidation(borrower common.Address, repayVaiAmount *big.Int, seizedVToken common.Address) (*types.Transaction, error) {
	publicKey := s.PrivateKey.Public()
	vaiController := s.vaiController
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	repayer := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.c.PendingNonceAt(context.Background(), repayer)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.c.SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = big.NewInt(5000000000)
	}

	gasLimit := uint64(3000000)
	chainId, err := s.c.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, chainId)
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
