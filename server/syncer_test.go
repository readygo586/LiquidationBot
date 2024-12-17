package server

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
	"testing"
	"time"
)

var syncer *Syncer

func TestSyncAccountLoopWithBackgroundSync(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	account := common.HexToAddress("0xF5A008a26c8C06F0e778ac07A0db9a2f42423c84") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	accountBytes := account.Bytes()
	err = sync.syncOneAccount(account)
	require.NoError(t, err)

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.True(t, exist)

	bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	t.Logf("info:%+v\n", info.toReadable())

	for _, asset := range info.Assets {
		symbol := asset.Symbol

		bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(symbol), accountBytes), nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		accounts = []common.Address{}
		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}
		require.Equal(t, 1, len(accounts))
	}

	key := getLiquidationKey(info.MaxLoanValue, info.HealthFactor, accountBytes)
	bz, err = db.Get(key, nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	sync.backgroundAccountSyncCh <- account
	sync.wg.Add(1)
	go sync.syncAccountLoop()
	time.Sleep(10 * time.Second)
	close(sync.quitCh)

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	err = json.Unmarshal(bz, &info)
	t.Logf("after background sync info:%+v\n", info.toReadable())
}

//
//func TestProcessLiquidationCase1(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//	liquidation := Liquidation{
//		Address: common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD"),
//	}
//	sync.processLiquidationReq(&liquidation)
//}
//
//func TestProcessLiquidation1(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//	liquidation := Liquidation{
//		Address: common.HexToAddress("0x05bbf0C12882FDEcd53FD734731ad578aF79621C"),
//	}
//
//	err = sync.processLiquidationReq(&liquidation)
//	if err != nil {
//		t.Logf(err.Error())
//	}
//}
//
//func TestProcessLiquidationWithBadLiquidationTxInForbiddenPeriod(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//
//	account := common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD")
//	liquidation := Liquidation{
//		Address: account,
//	}
//	currentHeight, err := sync.c.BlockNumber(context.Background())
//	db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
//	bz, err := db.Get(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
//	require.NoError(t, err)
//	gotHeight := big.NewInt(0).SetBytes(bz).Uint64()
//	require.Equal(t, currentHeight, gotHeight)
//
//	err = sync.processLiquidationReq(&liquidation)
//	require.Error(t, err)
//	t.Logf("%v", err)
//}
//
//func TestProcessLiquidationWithBadLiquidationTxForbiddenPeriodExpire(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//
//	account := common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD")
//	liquidation := Liquidation{
//		Address: account,
//	}
//	currentHeight, err := sync.c.BlockNumber(context.Background())
//	currentHeight -= (ForbiddenPeriodForBadLiquidation + 1)
//	db.Put(dbm.BadLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
//	bz, err := db.Get(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
//	require.NoError(t, err)
//	gotHeight := big.NewInt(0).SetBytes(bz).Uint64()
//	require.Equal(t, currentHeight, gotHeight)
//
//	err = sync.processLiquidationReq(&liquidation)
//	require.NoError(t, err)
//
//	exist, err := db.Has(dbm.BadLiquidationTxStoreKey(account.Bytes()), nil)
//	require.False(t, exist)
//}
//
//func TestProcessLiquidationWithPedningLiquidationTxInForbiddenPeriod(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//
//	account := common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD")
//	liquidation := Liquidation{
//		Address: account,
//	}
//	currentHeight, err := sync.c.BlockNumber(context.Background())
//	db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
//	bz, err := db.Get(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
//	require.NoError(t, err)
//	gotHeight := big.NewInt(0).SetBytes(bz).Uint64()
//	require.Equal(t, currentHeight, gotHeight)
//
//	err = sync.processLiquidationReq(&liquidation)
//	require.Error(t, err)
//	t.Logf("%v", err)
//}
//
//func TestProcessLiquidationWithPendingLiquidationTxForbiddenPeriodExpire(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//
//	account := common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD")
//	liquidation := Liquidation{
//		Address: account,
//	}
//	currentHeight, err := sync.c.BlockNumber(context.Background())
//	currentHeight -= (ForbiddenPeriodForPendingLiquidation + 1)
//	db.Put(dbm.PendingLiquidationTxStoreKey(account.Bytes()), big.NewInt(int64(currentHeight)).Bytes(), nil)
//	bz, err := db.Get(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
//	require.NoError(t, err)
//	gotHeight := big.NewInt(0).SetBytes(bz).Uint64()
//	require.Equal(t, currentHeight, gotHeight)
//
//	err = sync.processLiquidationReq(&liquidation)
//	require.NoError(t, err)
//
//	exist, err := db.Has(dbm.PendingLiquidationTxStoreKey(account.Bytes()), nil)
//	require.False(t, exist)
//}
//
//func TestCalculateSeizedTokenGetAmountsOutWithMulOverFlow(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//	liquidation := Liquidation{
//		Address: common.HexToAddress("1e73902ab4144299dfc2ac5a3765122c02ce889f"),
//	}
//
//	err = sync.processLiquidationReq(&liquidation)
//	if err != nil {
//		t.Logf(err.Error())
//	}
//}
//
//func TestCalculateSeizedTokenGetAmountsInExecutionRevert(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//	liquidation := Liquidation{
//		Address: common.HexToAddress("ba3b9a3ecf19e1139c78c4718d45fb99f7a838cd"),
//	}
//
//	err = sync.processLiquidationReq(&liquidation)
//	if err != nil {
//		t.Logf(err.Error())
//	}
//}
//
//func TestCalculateSeizedTokens(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
//
//	accounts := []string{
//		"0x1E73902Ab4144299DFc2ac5a3765122c02CE889f",
//		"0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627",
//		"0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC",
//		"0x0e0c57Ae65739394b405bC3afC5003bE9f858fDB",
//		"0x2eB71e5335d5328e76fa0755Db27E184Be834D31",
//		"0x4F41889788528e213692af181B582519BF4Cd30E",
//		"0x564EE8bF0bA977A1ccc92fe3D683AbF4569c9f5E",
//		"0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD",
//		"0x89fa3aec0A7632dDBbdBaf448534f26BA4B771F1",
//		"0xFAbE4C180b6eDad32eA0Cf56587c54417189e422",
//		"0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7",
//		"0xdcC896d48B17ECC88a9011057294EB0905bCb240",
//		"0xfDA2b6948E01525633B4058297bb89656609e6Ad",
//		"0xEAFb5e9E52A865D7BF1D3a9C17e0d29710928b8b",
//		"0x05bbf0C12882FDEcd53FD734731ad578aF79621C",
//	}
//
//	for _, account := range accounts {
//		liquidation := Liquidation{
//			Address: common.HexToAddress(account),
//		}
//		err := sync.processLiquidationReq(&liquidation)
//		if err != nil {
//			t.Logf(err.Error())
//		}
//	}
//}

func getLiquidationKey(maxLoanValue, healthFactor decimal.Decimal, accountBytes []byte) []byte {
	var key []byte
	if maxLoanValue.Cmp(MaxLoanValueThreshold) == -1 {
		key = dbm.LiquidationNonProfitStoreKey(accountBytes)
	} else {
		if healthFactor.Cmp(Decimal1P0) == -1 {
			key = dbm.LiquidationBelow1P0StoreKey(accountBytes)
		} else if healthFactor.Cmp(Decimal1P1) == -1 {
			key = dbm.LiquidationBelow1P1StoreKey(accountBytes)
		} else if healthFactor.Cmp(Decimal1P5) == -1 {
			key = dbm.LiquidationBelow1P5StoreKey(accountBytes)
		} else if healthFactor.Cmp(Decimal2P0) == -1 {
			key = dbm.LiquidationBelow2P0StoreKey(accountBytes)
		} else {
			key = dbm.LiquidationAbove2P0StoreKey(accountBytes)
		}
	}
	return key
}
