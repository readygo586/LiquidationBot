package server

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

var syncer *Syncer

func TestGetvAAVEUnderlyingPrice(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	oracle, err := venus.NewOracle(common.HexToAddress(cfg.Oracle), c)
	if err != nil {
		panic(err)
	}
	_, err = oracle.GetUnderlyingPrice(nil, common.HexToAddress("0x26DA28954763B92139ED49283625ceCAf52C6f94"))
	require.NoError(t, err)
}

func TestNewSyncer(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org" //"http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	verifyTokens(t, sync)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	require.NoError(t, err)

	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(0), num.Int64())
	//print tokens info
	for symbol, token := range sync.tokens {
		logger.Printf("symbol:%v, token:%+v\n", symbol, token)
	}

	//check flashpool
	for symbol, pools := range sync.flashLoanPools {
		t.Logf("%v pools:%v\n", symbol, pools)

		bep20, _ := venus.NewBep20(VAIAddress, c)
		if symbol != "VAI" {
			bep20, _ = venus.NewBep20(sync.tokens[symbol].UnderlyingAddress, c)
		}
		var balances []*big.Int
		for _, pair := range pools {
			balance, _ := bep20.BalanceOf(nil, pair)
			balances = append(balances, balance)
		}

		for i := 0; i < len(balances)-1; i++ {
			for j := i + 1; j < len(balances); j++ {
				require.True(t, balances[i].Cmp(balances[j]) != -1)
			}
		}
	}

	// check paths
	for pair, paths := range sync.paths {
		t.Logf("%v:%v\n", pair, paths)

		symbols := strings.Split(pair, ":")
		_, amountOut, _ := sync.selectPathWithMaxAmountOut(symbols[0], symbols[1], decimal.New(1000, 18))
		_, amountIn, _ := sync.selectPathWithMinAmountIn(symbols[0], symbols[1], amountOut)
		t.Logf("amountIn:%v, amountOut:%v\n", amountIn, amountOut)
	}
}

func TestDoSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org" //"http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	t.Logf("begin do sync markets and prices\n")

	sync.doSyncMarketsAndPrices()
	verifyTokens(t, sync)
}

func TestSyncMarketsAndPricesLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	t.Logf("begin sync markets and prices\n")
	sync.wg.Add(1)
	go sync.SyncMarketsAndPricesLoop()

	time.Sleep(time.Second * 60)
	close(sync.quitCh)
	sync.wg.Wait()
	verifyTokens(t, sync)
}

func TestFilterAllCotractsBorrowEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	sync := NewSyncer(c, nil, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	var addresses []common.Address
	name := make(map[string]string)
	for _, token := range sync.tokens {
		addresses = append(addresses, token.Address)
	}

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(13000000),
		ToBlock:   big.NewInt(13002000),
		Addresses: addresses,
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	logger.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var borrowEvent venus.Vbep20Borrow
		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
		logger.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)
	}
	logger.Printf("end Time:%v\n", time.Now())
}

func TestStoreAndDeleteAccount(t *testing.T) {
	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := &Syncer{
		db: db,
	}

	healthFactor, _ := decimal.NewFromString("0.9")

	vusdtBalance, _ := decimal.NewFromString("1000000000.0")
	vusdtLoan, _ := decimal.NewFromString("0")

	vbtcBalance, _ := decimal.NewFromString("2.5")
	vbtctLoan, _ := decimal.NewFromString("0.2")

	vbusdBalance, _ := decimal.NewFromString("0")
	vbusdtLoan, _ := decimal.NewFromString("500.23")

	/*
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
	*/
	assets := []Asset{
		{
			Symbol:           "vUSDT",
			Balance:          vusdtBalance,
			Loan:             vusdtLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
		{
			Symbol:           "vBTC",
			Balance:          vbtcBalance,
			Loan:             vbtctLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
		{
			Symbol:           "vBUSD",
			Balance:          vbusdBalance,
			Loan:             vbusdtLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
	}
	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
	info := AccountInfo{
		Account:       account,
		HealthFactor:  healthFactor,
		MaxLoanSymbol: "vBUSD",
		MintedVAIS:    decimal.New(100, 0),
		Height:        100,
		MaxLoanValue:  MaxLoanValueThreshold.Mul(decimal.NewFromInt(2)),
		Assets:        assets,
	}

	sync.storeAccount(account, info)

	bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	var got AccountInfo
	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)
	t.Logf("got:%+v\n", got)

	has, err := db.Has(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		has, err = db.Has(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())

		prefix := append(dbm.MarketPrefix, []byte(asset.Symbol)...)
		var accounts []common.Address
		iter := db.NewIterator(util.BytesPrefix(prefix), nil)
		defer iter.Release()
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}

		require.Equal(t, 1, len(accounts))
	}

	had, err := db.Has(dbm.MarketMemberStoreKey([]byte("vETH"), account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, had)

	sync.deleteAccount(account)
	has, err = db.Has(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.Equal(t, leveldb.ErrNotFound, err)
}

func TestStoreAndDeleteAccount1(t *testing.T) {
	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	//sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	sync := &Syncer{
		db: db,
	}

	healthFactor, _ := decimal.NewFromString("1.1")
	vusdtBalance, _ := decimal.NewFromString("1000000000.0")
	vusdtLoan, _ := decimal.NewFromString("0")

	vbtcBalance, _ := decimal.NewFromString("2.5")
	vbtctLoan, _ := decimal.NewFromString("0.2")

	vbusdBalance, _ := decimal.NewFromString("0")
	vbusdtLoan, _ := decimal.NewFromString("500.23")

	assets := []Asset{
		{
			Symbol:           "vUSDT",
			Balance:          vusdtBalance,
			Loan:             vusdtLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
		{
			Symbol:           "vBTC",
			Balance:          vbtcBalance,
			Loan:             vbtctLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
		{
			Symbol:           "vBUSD",
			Balance:          vbusdBalance,
			Loan:             vbusdtLoan,
			CollateralFactor: decimal.New(8, -1),
			BalanceValue:     decimal.New(100, 0),
			CollateralValue:  decimal.New(80, 0),
			LoanValue:        decimal.New(50, 0),
			Price:            decimal.New(100, 0),
			ExchangeRate:     decimal.New(100, 0),
		},
	}
	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
	info := AccountInfo{
		Account:       account,
		HealthFactor:  healthFactor,
		MaxLoanValue:  MaxLoanValueThreshold.Mul(decimal.NewFromInt(2)),
		MaxLoanSymbol: "vBUSD",
		MintedVAIS:    decimal.New(100, 0),
		Height:        100,
		Assets:        assets,
	}

	sync.storeAccount(account, info)

	bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	var got AccountInfo
	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	has, err := db.Has(dbm.LiquidationBelow1P5StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P5StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		has, err = db.Has(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}

	had, err := db.Has(dbm.MarketMemberStoreKey([]byte("vETH"), account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, had)

	sync.deleteAccount(account)

	vsxpBalance, _ := decimal.NewFromString("236.5")
	vsxpLoan, _ := decimal.NewFromString("800.23")

	assets = append(assets, Asset{
		Symbol:  "vSXP",
		Balance: vsxpBalance,
		Loan:    vsxpLoan,
	})

	info = AccountInfo{
		HealthFactor: healthFactor,
		MaxLoanValue: MaxLoanValueThreshold.Div(decimal.NewFromInt(2)),
		Assets:       assets,
	}

	sync.storeAccount(account, info)
	bz, err = db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	has, err = db.Has(dbm.LiquidationNonProfitStoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationNonProfitStoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		//logger.Printf("symbol:%v\n", asset.Symbol)
		has, err = db.Has(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}
}

func TestFeedPricesWithUpdateDB(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	oldPrice := sync.tokens["vBTC"].Price
	oldHeight := sync.tokens["vBTC"].PriceUpdatedHeight
	newPrice := oldPrice.Mul(decimal.New(103, -2))

	time.Sleep(10 * time.Second)
	height, err := sync.c.BlockNumber(context.Background())

	feededPrice := &FeededPrice{
		Symbol:  "vBTC",
		Address: sync.tokens["vBTC"].Address,
		Price:   newPrice,
	}

	sync.processFeededPrice(feededPrice)

	require.EqualValues(t, sync.tokens["vBTC"].Price, oldPrice)
	require.Equal(t, sync.tokens["vBTC"].PriceUpdatedHeight, oldHeight)
	require.EqualValues(t, sync.tokens["vBTC"].FeededPrice, newPrice)
	require.Equal(t, sync.tokens["vBTC"].FeedPriceUpdateHeihgt, height)
}

func TestSyncOneAccount(t *testing.T) {
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
}

func TestSyncOneAccount1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	account := common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
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
}

func TestSyncOneAccount2(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	account := common.HexToAddress("0x05bbf0C12882FDEcd53FD734731ad578aF79621C") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
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
}

func TestSyncOneAccount3(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	account := common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
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
}

func TestSyncAccounts(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	accounts := []common.Address{
		common.HexToAddress("0xF5A008a26c8C06F0e778ac07A0db9a2f42423c84"),
		common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833"),
		common.HexToAddress("0x05bbf0C12882FDEcd53FD734731ad578aF79621C"),
		common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627"),
	}

	sync.syncAccounts(accounts)
	require.NoError(t, err)

	gotAccounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	defer iter.Release()
	for iter.Next() {
		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, len(accounts), len(gotAccounts))

	symbolCount := make(map[string]int)

	for _, account := range accounts {
		accountBytes := account.Bytes()
		exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
		require.NoError(t, err)
		require.True(t, exist)

		bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
		var info AccountInfo
		err = json.Unmarshal(bz, &info)
		t.Logf("info:%+v\n", info.toReadable())

		key := getLiquidationKey(info.MaxLoanValue, info.HealthFactor, accountBytes)
		bz, err = db.Get(key, nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		for _, asset := range info.Assets {
			symbol := asset.Symbol
			bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(symbol), accountBytes), nil)
			require.NoError(t, err)
			require.Equal(t, account, common.BytesToAddress(bz))
			symbolCount[symbol]++
		}
	}

	for symbol, count := range symbolCount {
		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		accounts = []common.Address{}
		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}
		require.Equal(t, count, len(accounts))
	}
}

func TestSyncOneAccountWithIncreaseAccountNumber(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	account := common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	accountBytes := account.Bytes()
	err = sync.syncOneAccountWithIncreaseAccountNumber(account)
	require.NoError(t, err)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	count := big.NewInt(0).SetBytes(bz).Int64()
	require.Equal(t, int64(1), count)

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.True(t, exist)

	bz, err = db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
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
}

func TestSyncOneAccountWithIncreaseAccountNumber1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	accounts := []common.Address{
		common.HexToAddress("0xF5A008a26c8C06F0e778ac07A0db9a2f42423c84"),
		common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833"),
		common.HexToAddress("0x05bbf0C12882FDEcd53FD734731ad578aF79621C"),
		common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627"),
	}

	for _, account := range accounts {
		sync.syncOneAccountWithIncreaseAccountNumber(account)
		require.NoError(t, err)
	}
	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	count := big.NewInt(0).SetBytes(bz).Int64()
	require.Equal(t, int64(4), count)

	//sync an already existed account
	sync.syncOneAccountWithIncreaseAccountNumber(common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833"))
	bz, err = db.Get(dbm.BorrowerNumberKey(), nil)
	count = big.NewInt(0).SetBytes(bz).Int64()
	require.Equal(t, int64(4), count)

	gotAccounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	defer iter.Release()
	for iter.Next() {
		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, len(accounts), len(gotAccounts))

	symbolCount := make(map[string]int)

	for _, account := range accounts {
		accountBytes := account.Bytes()
		exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
		require.NoError(t, err)
		require.True(t, exist)

		bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
		var info AccountInfo
		err = json.Unmarshal(bz, &info)
		t.Logf("info:%+v\n", info.toReadable())

		key := getLiquidationKey(info.MaxLoanValue, info.HealthFactor, accountBytes)
		bz, err = db.Get(key, nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		for _, asset := range info.Assets {
			symbol := asset.Symbol
			bz, err = db.Get(dbm.MarketMemberStoreKey([]byte(symbol), accountBytes), nil)
			require.NoError(t, err)
			require.Equal(t, account, common.BytesToAddress(bz))
			symbolCount[symbol]++
		}
	}

	for symbol, count := range symbolCount {
		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		accounts = []common.Address{}
		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}
		require.Equal(t, count, len(accounts))
	}
}

func TestSyncOneAccountWithFeededPrice(t *testing.T) {
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

	height, err := sync.c.BlockNumber(context.Background())
	require.NoError(t, err)

	feedPrice := &FeededPrice{
		Address: sync.tokens["vBTC"].Address,
		Price:   sync.tokens["vBTC"].Price.Div(decimal.NewFromInt(2)),
		Height:  height,
	}

	sync.syncOneAccountWithFeededPrice(account, feedPrice)

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	err = json.Unmarshal(bz, &info)
	t.Logf("info after feededPrice:%+v\n", info.toReadable())

	if info.HealthFactor.Cmp(Decimal1P0) == -1 {
		priorityliquidation := <-sync.priortyLiquidationCh
		t.Logf("liquiadtion:%+v\n", priorityliquidation)
	}

	if info.HealthFactor.Cmp(Decimal1P1) == -1 {
		cinfo := <-sync.concernedAccountInfoCh
		t.Logf("cinfo:%+v\n", cinfo.toReadable())
	}
}

func TestSyncOneAccountWithFeededPrice1(t *testing.T) {
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

	//height, err := sync.c.BlockNumber(context.Background())
	//require.NoError(t, err)

	feedPrice := &FeededPrice{
		Address: sync.tokens["vBTC"].Address,
		Price:   sync.tokens["vBTC"].Price.Div(decimal.NewFromInt(2)),
	}

	sync.syncOneAccountWithFeededPrice(account, feedPrice)

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	err = json.Unmarshal(bz, &info)
	t.Logf("info after feededPrice:%+v\n", info.toReadable())

	if info.HealthFactor.Cmp(Decimal1P0) == -1 {
		priorityliquidation := <-sync.priortyLiquidationCh
		t.Logf("liquiadtion:%+v\n", priorityliquidation)
	}

	if info.HealthFactor.Cmp(Decimal1P1) == -1 {
		cinfo := <-sync.concernedAccountInfoCh
		t.Logf("cinfo:%+v\n", cinfo.toReadable())
	}

}

func TestProcessFeedPrices(t *testing.T) {
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

	time.Sleep(10 * time.Second)

	height, err := sync.c.BlockNumber(context.Background())
	require.NoError(t, err)

	oldPrice := sync.tokens["vBTC"].Price
	oldPriceUpdateHeight := sync.tokens["vBTC"].PriceUpdatedHeight
	newPrice := oldPrice.Mul(decimal.New(104, -2))

	feededPrice := &FeededPrice{
		Symbol:  "vBTC",
		Address: sync.tokens["vBTC"].Address,
		Price:   newPrice,
	}

	sync.processFeededPrice(feededPrice)

	if len(sync.highPriorityAccountSyncCh) != 0 {
		accountWithFeededPrice := <-sync.highPriorityAccountSyncCh
		require.Equal(t, account, accountWithFeededPrice.Addresses[0])
		require.EqualValues(t, *feededPrice, *accountWithFeededPrice.FeededPrice)
		t.Logf("highPriorityAccountCh:%v", accountWithFeededPrice)
	}

	require.Equal(t, sync.tokens["vBTC"].Price, oldPrice)
	require.Equal(t, sync.tokens["vBTC"].PriceUpdatedHeight, oldPriceUpdateHeight)
	require.Equal(t, sync.tokens["vBTC"].FeededPrice, newPrice)
	require.Equal(t, sync.tokens["vBTC"].FeedPriceUpdateHeihgt, height)
}

func TestProcessFeedPrice1(t *testing.T) {
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

	time.Sleep(10 * time.Second)

	height, err := sync.c.BlockNumber(context.Background())
	require.NoError(t, err)

	oldPrice := sync.tokens["vBTC"].Price
	oldPriceUpdateHeight := sync.tokens["vBTC"].PriceUpdatedHeight
	newPrice := oldPrice.Mul(decimal.New(96, -2))

	feededPrice := &FeededPrice{
		Symbol:  "vBTC",
		Address: sync.tokens["vBTC"].Address,
		Price:   newPrice,
	}

	sync.processFeededPrice(feededPrice)

	if len(sync.highPriorityAccountSyncCh) != 0 {
		accountWithFeededPrice := <-sync.highPriorityAccountSyncCh
		require.Equal(t, account, accountWithFeededPrice.Addresses[0])
		require.EqualValues(t, *feededPrice, *accountWithFeededPrice.FeededPrice)
		t.Logf("highPriorityAccountCh:%v", accountWithFeededPrice)
	}

	if len(sync.lowPriorityAccountSyncCh) != 0 {
		accountWithFeededPrice := <-sync.lowPriorityAccountSyncCh
		require.Equal(t, account, accountWithFeededPrice.Addresses[0])
		require.EqualValues(t, *feededPrice, *accountWithFeededPrice.FeededPrice)
		t.Logf("lowPriorityAccountCh:%v", accountWithFeededPrice)
	}

	require.Equal(t, sync.tokens["vBTC"].Price, oldPrice)
	require.Equal(t, sync.tokens["vBTC"].PriceUpdatedHeight, oldPriceUpdateHeight)
	require.Equal(t, sync.tokens["vBTC"].FeededPrice, newPrice)
	require.Equal(t, sync.tokens["vBTC"].FeedPriceUpdateHeihgt, height)
}

func TestTestProcessFeedPriceVibrationExceed5Percent(t *testing.T) {
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

	time.Sleep(10 * time.Second)

	height, err := sync.c.BlockNumber(context.Background())
	require.NoError(t, err)

	oldPrice := sync.tokens["vBTC"].Price
	oldPriceUpdateHeight := sync.tokens["vBTC"].PriceUpdatedHeight

	feedPrice := &FeededPrice{
		Address: sync.tokens["vBTC"].Address,
		Price:   sync.tokens["vBTC"].Price.Div(decimal.NewFromInt(2)),
		Height:  height,
	}

	sync.processFeededPrice(feedPrice)
	require.Equal(t, 0, len(sync.highPriorityAccountSyncCh))
	require.Equal(t, 0, len(sync.lowPriorityAccountSyncCh))

	require.Equal(t, sync.tokens["vBTC"].Price, oldPrice)
	require.Equal(t, sync.tokens["vBTC"].PriceUpdatedHeight, oldPriceUpdateHeight)
	require.True(t, sync.tokens["vBTC"].FeededPrice.Cmp(decimal.Zero) == 0)
	require.EqualValues(t, sync.tokens["vBTC"].FeedPriceUpdateHeihgt, 0)
}

func TestSyncAccountLoopWithFeedPrice(t *testing.T) {
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

	time.Sleep(10 * time.Second)

	height, err := sync.c.BlockNumber(context.Background())
	require.NoError(t, err)

	oldPrice := sync.tokens["vBTC"].Price
	oldPriceUpdateHeight := sync.tokens["vBTC"].PriceUpdatedHeight
	newPrice := oldPrice.Mul(decimal.New(96, -2))

	feededPrice := &FeededPrice{
		Symbol:  "vBTC",
		Address: sync.tokens["vBTC"].Address,
		Price:   newPrice,
		Height:  height,
	}

	sync.processFeededPrice(feededPrice)

	sync.wg.Add(1)
	go sync.syncAccountLoop()
	time.Sleep(10 * time.Second)
	close(sync.quitCh)

	require.Equal(t, sync.tokens["vBTC"].Price, oldPrice)
	require.Equal(t, sync.tokens["vBTC"].PriceUpdatedHeight, oldPriceUpdateHeight)
	require.Equal(t, sync.tokens["vBTC"].FeededPrice, newPrice)
	require.Equal(t, sync.tokens["vBTC"].FeedPriceUpdateHeihgt, height)

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	err = json.Unmarshal(bz, &info)
	t.Logf("after process feededPriceinfo:%+v\n", info.toReadable())
}

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

func TestScanAllBorrowers1(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	height, err := c.BlockNumber(ctx)
	require.NoError(t, err)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	startHeight := big.NewInt(int64(height - 5000))
	db.Put(dbm.KeyLastHandledHeight, startHeight.Bytes(), nil)
	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)

	sync.Start()
	time.Sleep(time.Second * 60)
	sync.Stop()

	bz, err := db.Get(dbm.KeyLastHandledHeight, nil)
	end := big.NewInt(0).SetBytes(bz)
	t.Logf("end height:%v\n", end.Int64())

	bz, err = db.Get(dbm.KeyBorrowerNumber, nil)
	num := big.NewInt(0).SetBytes(bz).Int64()
	t.Logf("num:%v\n", num)

	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	defer iter.Release()
	t.Logf("borrows address")
	for iter.Next() {
		addr := common.BytesToAddress(iter.Value())
		t.Logf("%v\n", addr.String())
	}
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

func verifyTokens(t *testing.T, sync *Syncer) {
	require.Equal(t, common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"), sync.tokens["vETH"].Address)
	require.Equal(t, common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"), sync.tokens["vUSDT"].Address)
	require.Equal(t, common.HexToAddress("0x61eDcFe8Dd6bA3c891CB9bEc2dc7657B3B422E93"), sync.tokens["vTRX"].Address)
	require.Equal(t, common.HexToAddress("0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3"), sync.tokens["vTUSD"].Address)
	require.Equal(t, common.HexToAddress("0x26DA28954763B92139ED49283625ceCAf52C6f94"), sync.tokens["vAAVE"].Address)
	require.Equal(t, common.HexToAddress("0x86aC3974e2BD0d60825230fa6F355fF11409df5c"), sync.tokens["vCAKE"].Address)
	require.Equal(t, common.HexToAddress("0x5c9476FcD6a4F9a3654139721c949c2233bBbBc8"), sync.tokens["vMATIC"].Address)
	require.Equal(t, common.HexToAddress("0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71"), sync.tokens["vDOGE"].Address)
	require.Equal(t, common.HexToAddress("0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec"), sync.tokens["vADA"].Address)
	require.Equal(t, common.HexToAddress("0xeBD0070237a0713E8D94fEf1B728d3d993d290ef"), sync.tokens["vCAN"].Address)
	require.Equal(t, common.HexToAddress("0x972207A639CC1B374B893cc33Fa251b55CEB7c07"), sync.tokens["vBETH"].Address)
	require.Equal(t, common.HexToAddress("0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1"), sync.tokens["vDAI"].Address)
	require.Equal(t, common.HexToAddress("0x650b940a1033B8A1b1873f78730FcFC73ec11f1f"), sync.tokens["vLINK"].Address)
	require.Equal(t, common.HexToAddress("0x1610bc33319e9398de5f57B33a5b184c806aD217"), sync.tokens["vDOT"].Address)
	require.Equal(t, common.HexToAddress("0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176"), sync.tokens["vBCH"].Address)
	require.Equal(t, common.HexToAddress("0xB248a295732e0225acd3337607cc01068e3b9c10"), sync.tokens["vXRP"].Address)
	require.Equal(t, common.HexToAddress("0x57A5297F2cB2c0AaC9D554660acd6D385Ab50c6B"), sync.tokens["vLTC"].Address)
	require.Equal(t, common.HexToAddress("0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B"), sync.tokens["vBTC"].Address)
	require.Equal(t, common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36"), sync.tokens["vBNB"].Address)
	require.Equal(t, common.HexToAddress("0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D"), sync.tokens["vXVS"].Address)
	require.Equal(t, common.HexToAddress("0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0"), sync.tokens["vSXP"].Address)
	require.Equal(t, common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"), sync.tokens["vBUSD"].Address)
	require.Equal(t, common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"), sync.tokens["vETH"].Address)
	require.Equal(t, common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8"), sync.tokens["vUSDC"].Address)
}

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

func TestGetOracle(t *testing.T) {
	rpcURL := "ws://192.168.88.144:28546"
	c, _ := ethclient.Dial(rpcURL)

	oracle, _ := venus.NewPriceOracle(common.HexToAddress("0xd8b6da2bfec71d684d3e2a2fc9492ddad5c3787f"), c)
	tokens := [24]string{"ETH", "USDT", "TRX", "TUSD", "AAVE", "CAKE", "MATIC", "MATIC", "DOGE", "ADA", "CAN", "BETH", "DAI", "LINK", "DOT", "BCH", "XRP", "LTC", "BTCB", "BNB", "XVS", "SXP", "BUSD", "USDC"}
	for _, token := range tokens {
		feedAddr, _ := oracle.GetFeed(nil, token)
		priceFeed, _ := venus.NewPriceFeed(feedAddr, c)
		finalOracle, _ := priceFeed.Aggregator(nil)
		println(token, strings.ToLower(finalOracle.String()))
	}
}

//func TestMonitorTxPoolLoop(t *testing.T) {
//	rpcURL := "ws://192.168.88.144:28546"
//	client, _ := ethclient.Dial(rpcURL)
//	fmt.Println("We have a connection")
//	v := reflect.ValueOf(client).Elem()
//	f := v.FieldByName("c")
//	rf := reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
//	concrete_client, _ := rf.Interface().(*rpc.Client)
//	txPoolTXs := make(chan common.Hash, 1024)
//	concrete_client.EthSubscribe(
//		context.Background(), txPoolTXs, "newPendingTransactions",
//	)
//	targetMap := make(map[string]struct{}, 24)
//	targetMap["0x137924d7c36816e0dcaf016eb617cc2c92c05782"] = struct{}{} //BNB
//	targetMap["0x178ba789e24a1d51e9ea3cb1db3b52917963d71d"] = struct{}{} //BTCB
//	targetMap["0xfc3069296a691250ffdf21fe51340fdd415a76ed"] = struct{}{} //ETH
//	targetMap["0x7935a51addab8550d346feef34e02f67c9330109"] = struct{}{} //CAKE
//	aggregatorABI, _ := venus.AggregatorMetaData.GetAbi()
//	for txn := range txPoolTXs {
//
//		txn, is_pending, err := client.TransactionByHash(context.Background(), txn)
//		if err == nil && txn != nil && txn.To() != nil && is_pending == true {
//
//			_, ok := targetMap[strings.ToLower(txn.To().String())]
//			if ok {
//
//				if len(txn.Data()) < 5 {
//					//Error
//				}
//				method, err := aggregatorABI.MethodById(txn.Data()[0:4])
//				if err != nil {
//					//Error
//				}
//				if method.Name == "transmit" {
//					inputData := make(map[string]interface{})
//					err = method.Inputs.UnpackIntoMap(inputData, txn.Data()[4:])
//					data := inputData["_report"].([]byte)
//					numbering := data[32+32+32+32:]
//					numberingmid := numbering[len(numbering)/2 : len(numbering)/2+32]
//
//					if err != nil {
//						panic(err)
//					}
//					fmt.Println("==================")
//					fmt.Println(txn.Hash().String(), "is updateing price @", time.Now())
//					logger.Printf("% s % x \n", txn.Hash().String(), numberingmid)
//					result := big.NewInt(0).SetBytes(numberingmid)
//					fmt.Println(txn.Hash().String(), "price: ", result)
//				}
//			}
//		}
//	}
//}

func TestMonitorTransmitEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	topicNewTransmission := common.HexToHash("0xf6a97944f31ea060dfde0566e4167c1a1082551e64b60ecb14d599a9d023d451")

	var addresses []common.Address
	for _, token := range sync.tokens {
		addresses = append(addresses, token.Oracle)
	}

	aggregator, err := abi.JSON(strings.NewReader(venus.AggregatorMetaData.ABI))
	require.NoError(t, err)
	monitorStartHeight := uint64(15806919)
	monitorEndHeight := uint64(15812714)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(monitorStartHeight)),
		ToBlock:   big.NewInt(int64(monitorEndHeight)),
		Addresses: addresses, //usdc
		Topics:    [][]common.Hash{{topicNewTransmission}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	if err == nil {
		for _, log := range logs {
			var eve venus.AggregatorNewTransmission
			aggregator.UnpackIntoInterface(&eve, "NewTransmission", log.Data)
			logger.Printf("NewTransmission happen @ height:%v, address:%v txhash:%v, answer:%v\n", log.BlockNumber, log.Address, log.TxHash, eve.Answer)
		}

		monitorStartHeight = monitorEndHeight
	}
}

func TestGetTxGasPrice(t *testing.T) {
	rpcURL := "https://bsc-dataseed.binance.org" //"http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)
	t.Logf("height:%v", height)
	hash := common.HexToHash("0xc85f4b884067941966dd209eb663f9ac5c26d4ef69ab38de2e37acf47d46d3c8")
	tx, _, _ := c.TransactionByHash(context.Background(), hash)
	t.Logf("tx's gasPrice:%v", tx.GasPrice())
}

func TestRoutineException(t *testing.T) {
	inputCh := make(chan int, 100)
	quitCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-quitCh:
				return

			case data := <-inputCh:
				if data == 10 {
					break //in this case continue = break
				}
				logger.Printf("input:%v\n", data)
			}
		}
	}()

	for i := 0; i < 20; i++ {
		inputCh <- i
		time.Sleep(10 * time.Millisecond)
	}
	close(quitCh)
}
