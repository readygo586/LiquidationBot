package server

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"testing"
)

//
//func TestUpdateDB(t *testing.T) {
//	db, err := dbm.NewDB("../bin/maindb")
//	require.NoError(t, err)
//	defer db.Close()
//
//	var iter = db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
//	defer iter.Release()
//	count0 := 0
//	count1 := 0
//
//	for iter.Next() {
//		count0++
//		accountBytes := iter.Value()
//		logger.Printf("account:%v\n", common.BytesToAddress(accountBytes))
//		bz, err := db.Get(dbm.AccountStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//
//		info := &AccountInfo{}
//		err = json.Unmarshal(bz, info)
//		require.NoError(t, err)
//
//		maxLoan := decimal.NewFromInt(0)
//		maxRepay := decimal.NewFromInt(0)
//		for _, asset := range info.Assets {
//			if asset.Loan.Cmp(maxLoan) == 1 {
//				maxLoan = asset.Loan
//			}
//
//			if asset.Balance.Cmp(maxRepay) == 1 {
//				maxRepay = asset.Balance
//			}
//		}
//
//		logger.Printf("account:%v\n", common.BytesToAddress(accountBytes))
//		logger.Printf("before info:%+v\n", info)
//
//		info.MaxLoan = maxLoan
//		info.MaxRepay = maxRepay
//		healthFactor := info.HealthFactor
//
//		bz1, err := json.Marshal(&info)
//		require.NoError(t, err)
//
//		db.Put(dbm.AccountStoreKey(accountBytes), bz1, nil)
//
//		bz2, err := db.Get(dbm.AccountStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//		newInfo := &AccountInfo{}
//		err = json.Unmarshal(bz2, newInfo)
//		require.NoError(t, err)
//
//		logger.Printf("after info:%+v\n", newInfo)
//
//		if maxRepay.Cmp(MaxLoanValueThreshold) == -1 {
//			count1++
//			if healthFactor.Cmp(Decimal1P0) == -1 {
//				db.Delete(dbm.LiquidationBelow1P0StoreKey(accountBytes), nil)
//			} else if healthFactor.Cmp(Decimal1P1) == -1 {
//				db.Delete(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
//			} else if healthFactor.Cmp(Decimal1P5) == -1 {
//				db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
//			} else if healthFactor.Cmp(Decimal2P0) == -1 {
//				db.Delete(dbm.LiquidationBelow2P0StoreKey(accountBytes), nil)
//			} else {
//				db.Delete(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
//			}
//
//			//db.Put(dbm.LiquidationNoAssetStoreKey(accountBytes), accountBytes, nil)
//		}
//	}
//
//	logger.Printf("totalCount:%v, nonProfitCount:%v\n", count0, count1)
//}

func TestDeleteAllKeysExceptBorrowers(t *testing.T) {
	//cfg, err := config.New("../config.yml")
	//rpcURL := "http://42.3.146.198:21993"
	//c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("maindb")
	require.NoError(t, err)
	defer db.Close()

	//liquidationCh := make(chan *Liquidation, 64)
	//priorityliquidationCh := make(chan *Liquidation, 64)
	//feededPriceCh := make(chan *FeededPrices, 64)
	//
	//syncer := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPriceCh, liquidationCh, priorityliquidationCh)
	//logger.Printf("db:%v", syncer.db)
	/*
		AccountPrefix              = []byte("account")
			MarketPrefix               = []byte("market")
			LiquidationBelow1P0Prefix  = []byte("liquidation_below_1p0")
			LiquidationBelow1P1Prefix  = []byte("liquidation_below_1p1")
			LiquidationBelow1P5Prefix  = []byte("liquidation_below_1p5")
			LiquidationBelow2P0Prefix  = []byte("liquidation_below_2p0")
			LiquidationAbove2P0Prefix  = []byte("liquidation_above_2p0")
			LiquidationNoAssetPrefix = []byte("liquidation_no_asset") //
	*/
	var addresses []common.Address
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	count := 0
	for iter.Next() {
		count++
		addresses = append(addresses, common.BytesToAddress(iter.Value()))
	}
	logger.Printf("total:%v address:%v\n", count, addresses)
	iter.Release()

	count0 := 0
	iter1 := db.NewIterator(util.BytesPrefix(dbm.AccountPrefix), nil)
	for iter1.Next() {
		count0++
		accountBytes := bytes.TrimPrefix(iter1.Key(), dbm.AccountPrefix)
		var info AccountInfo
		err := json.Unmarshal(iter1.Value(), &info)
		require.NoError(t, err)
		assets := info.Assets
		for _, asset := range assets {
			db.Delete(dbm.MarketMemberStoreKey([]byte(asset.Symbol), accountBytes), nil)
		}
		db.Delete(dbm.AccountStoreKey(accountBytes), nil)
	}
	//logger.Printf("delete %v account\n", count0)
	iter1.Release()

	countBelow1P0 := 0
	iter2 := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P0Prefix), nil)
	for iter2.Next() {
		countBelow1P0++
		accountBytes := iter2.Value()
		db.Delete(dbm.LiquidationBelow1P0StoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v below1P0 account\n", countBelow1P0)
	iter2.Release()

	countBelow1P1 := 0
	iter3 := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P1Prefix), nil)
	for iter3.Next() {
		countBelow1P1++
		accountBytes := iter3.Value()
		db.Delete(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v below1P1 account\n", countBelow1P1)
	iter3.Release()

	countBelow1P5 := 0
	iter4 := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P5Prefix), nil)
	for iter4.Next() {
		countBelow1P5++
		accountBytes := iter4.Value()
		db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v below1P5 account\n", countBelow1P5)
	iter4.Release()

	countBelow2P0 := 0
	iter5 := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow2P0Prefix), nil)
	for iter5.Next() {
		countBelow2P0++
		accountBytes := iter5.Value()
		db.Delete(dbm.LiquidationBelow2P0StoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v below2P0 account\n", countBelow2P0)
	iter5.Release()

	countAbove2P0 := 0
	iter6 := db.NewIterator(util.BytesPrefix(dbm.LiquidationAbove2P0Prefix), nil)
	for iter6.Next() {
		countAbove2P0++
		accountBytes := iter6.Value()
		db.Delete(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v above2P0 account\n", countAbove2P0)
	iter6.Release()

	countNonProfit := 0
	iter7 := db.NewIterator(util.BytesPrefix(dbm.LiquidationNonProfitPrefix), nil)
	for iter7.Next() {
		countNonProfit++
		accountBytes := iter7.Value()
		db.Delete(dbm.LiquidationNonProfitStoreKey(accountBytes), nil)
	}
	logger.Printf("delete %v nonprofit account\n", countNonProfit)
	iter7.Release()
	db.Close()
}

func TestRebuildDB(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("maindb")
	require.NoError(t, err)
	defer db.Close()

	syncer := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	logger.Printf("db:%v", syncer.db)

	var addresses []common.Address
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	count := 0
	for iter.Next() {
		count++
		addresses = append(addresses, common.BytesToAddress(iter.Value()))
	}
	//logger.Printf("total:%v address:%v\n", count, addresses)
	iter.Release()
	syncer.syncAccounts(addresses)

}

func TestDeleteAllNonProfitAccounts(t *testing.T) {
	db, err := dbm.NewDB("../bin/maindb")
	require.NoError(t, err)
	defer db.Close()

	var iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationNonProfitPrefix), nil)
	defer iter.Release()

	for iter.Next() {
		db.Delete(dbm.LiquidationNonProfitStoreKey(iter.Value()), nil)
	}
}
