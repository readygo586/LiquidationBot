package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"testing"
	"time"
)

/*
let BTCAdderss = `0xac2ceee60e79f572cae3d2ea9c1c7f0b03934f5e`
let ETHAddress = `0x4354230038d0C3120B8756f1AbA72E9F4FC94979`
let BTCFeeder = `0x33deb1bCDCC9ecc2056F87A20CFF3dcBd54a37f6`
let ETHFeeder = `0x11ffA6965b4c25790980897241100dA23b87C7f2`
Unitroller deployed to: 0xbECEC0b03123C8f2A3C87fa720d70D2014070ec7
Comptroller deployed to: 0x7660F4B3E0AA407E89532F6E6674581FC15e51E5
price oracle deployed to: 0xd62aB2b1792F536bea70b75db41D9C8B6b3c1dee
access control deployed to: 0x3E2c5Aff6585b79D1E6230e678401b5b616cC8aa
VAI deployed to: 0x27D5638c474A2530FA9771E6d901E8A87De48dC1
VAIController deployed to: 0x0E882A2A9FaFD0dDdeEf618485A67Eb66e906303
vBTC deployed to: 0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e
vETH deployed to: 0x5a57B04Bc33f7E22daED781fa32cB074241BeA09
*/

func Test_ScanOneBlock_Non_VToken_Events(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.feeders)
	query2 := buildVTokenQueryWithoutHeight(s.markets)

	heights := []int64{47719833, 47719850, 47719851, 47719852, 47719853, 47768075, 47768750, 47769460,
		47769465}
	for _, height := range heights {
		err = s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
		time.Sleep(20 * time.Microsecond)
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.closeFactorChangedCh))
	assert.Equal(t, 2, len(s.newMarketCh))
	assert.Equal(t, 2, len(s.collateralFactorChangedCh))
	assert.Equal(t, 3, len(s.enterMarketCh))
	assert.Equal(t, 1, len(s.exitMarketCh))

}

func Test_ScanOneBlock_VToken_Events(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.feeders)
	query2 := buildVTokenQueryWithoutHeight(s.markets)

	heights := []int64{47768056, 47768708, 47769450}
	for _, height := range heights {
		err = s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
		time.Sleep(20 * time.Microsecond)
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 3, len(s.vTokenAmountChangedCh))
}

func Test_ScanOneBlock_Liquidate_Events(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.feeders)
	query2 := buildVTokenQueryWithoutHeight(s.markets)

	heights := []int64{47797935}
	for _, height := range heights {
		err = s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
		time.Sleep(20 * time.Microsecond)
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
	vTokenAmountChange := <-s.vTokenAmountChangedCh
	assert.Equal(t, "0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e", vTokenAmountChange.Market.Hex())
	assert.Equal(t, "0x4e3CC26bce18b0F420155DCE102c976aF057867E", vTokenAmountChange.From.Hex())
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", vTokenAmountChange.To.Hex())
	assert.Equal(t, "51508848", vTokenAmountChange.Amount.String())

	repayVaiAmountChanged := <-s.repayVaiAmountChangedCh
	assert.Equal(t, "0x4e3CC26bce18b0F420155DCE102c976aF057867E", repayVaiAmountChanged.Account.Hex())
	assert.Equal(t, "-42995700000000000000000", repayVaiAmountChanged.Amount.String())
}

func Test_SyncOneAccount1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	account := common.HexToAddress("0x4e3CC26bce18b0F420155DCE102c976aF057867E")
	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)
	err = s.syncOneAccount(account)
	assert.NoError(t, err)

	bz, err := s.db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	assert.NoError(t, err)

	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	assert.NoError(t, err)
	assert.Equal(t, account, info.Account)
	assert.Equal(t, 1, len(info.Assets))
	fmt.Printf("info: %+v\n", info.toReadable())

	bz, err = s.db.Get(dbm.LiquidationBelow1P1StoreKey(account.Bytes()), nil)
	assert.Equal(t, account.Bytes(), bz)
}

func Test_ScanLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)
	t.Logf("begin scan loop\n")
	s.db.Put(dbm.LatestHandledHeightStoreKey(), big.NewInt(46341420).Bytes(), nil)

	s.wg.Add(1)
	go s.ScanLoop()

	time.Sleep(time.Second * 30)
	s.Stop()
	t.Logf("end scan loop\n")

	assert.Equal(t, 1, len(s.closeFactorChangedCh))
	assert.Equal(t, 2, len(s.newMarketCh))
	assert.Equal(t, 2, len(s.collateralFactorChangedCh))
	assert.Equal(t, 0, len(s.enterMarketCh))
	assert.Equal(t, 0, len(s.exitMarketCh))
	assert.Equal(t, 0, len(s.repayVaiAmountChangedCh))
	assert.Equal(t, 0, len(s.vTokenAmountChangedCh))
	assert.Equal(t, 0, len(s.priceChangedCh))
}

//CollateralFactorLoop
//EnterMarketLoop
//ExitMarketLoop
//PriceChangeLoop
//VTokenAmountChangedLoop
//RepayVaiAmountChangedLoop

func Test_EnterMarketLoop_ExitMarketLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: big.NewInt(46341420).Uint64(),
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: big.NewInt(46341421).Uint64(),
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: big.NewInt(46341422).Uint64(),
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)

	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	had, _ := s.db.Has(dbm.MarketMemberStoreKey(vBTCMarket.Bytes(), account1.Bytes()), nil)
	assert.True(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vETHMarket.Bytes(), account2.Bytes()), nil)
	assert.True(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vETHMarket.Bytes(), account3.Bytes()), nil)
	assert.True(t, had)

	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vETHMarket.Bytes(), account1.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vBTCMarket.Bytes(), account2.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vBTCMarket.Bytes(), account3.Bytes()), nil)
	assert.False(t, had)

	s.exitMarketCh <- &ExitMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: big.NewInt(46341520).Uint64(),
	}

	s.exitMarketCh <- &ExitMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: big.NewInt(46341521).Uint64(),
	}

	s.exitMarketCh <- &ExitMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: big.NewInt(46341522).Uint64(),
	}

	s.wg.Add(1)
	go s.ExitMarketLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 3, len(s.highAccountSyncCh))
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vBTCMarket.Bytes(), account1.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vETHMarket.Bytes(), account2.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vETHMarket.Bytes(), account3.Bytes()), nil)
	assert.False(t, had)
}

func Test_CollateralFactorLoop_CollateralFactorDecrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 1,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)

	newFactor := decimal.NewFromInt(700000000000000000)
	s.collateralFactorChangedCh <- &CollateralFactorChanged{
		Market:           vETHMarket,
		CollateralFactor: newFactor,
		UpdatedHeight:    height + 10,
	}

	s.wg.Add(1)
	go s.CollateralFactorLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 1, len(s.highAccountSyncCh))

	accounts := <-s.highAccountSyncCh
	assert.Equal(t, []common.Address{account2, account3}, accounts)
	assert.Equal(t, newFactor, s.tokens[vETHMarket].CollateralFactor)
}

func Test_CollateralFactorLoop_CollateralFactorIncrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newFactor := decimal.NewFromInt(950000000000000000)
	s.collateralFactorChangedCh <- &CollateralFactorChanged{
		Market:           vETHMarket,
		CollateralFactor: newFactor,
		UpdatedHeight:    height + 10,
	}

	s.wg.Add(1)
	go s.CollateralFactorLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 1, len(s.middleAccountSyncCh))

	accounts := <-s.middleAccountSyncCh
	assert.Equal(t, []common.Address{account2, account3}, accounts)
	assert.Equal(t, newFactor, s.tokens[vETHMarket].CollateralFactor)
}

func Test_CollateralFactorLoop_CollateralFactorNotChange(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newFactor := decimal.NewFromInt(800000000000000000)
	s.collateralFactorChangedCh <- &CollateralFactorChanged{
		Market:           vETHMarket,
		CollateralFactor: newFactor,
		UpdatedHeight:    height + 10,
	}

	s.wg.Add(1)
	go s.CollateralFactorLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 0, len(s.middleAccountSyncCh))
}

func Test_CollateralFactorLoop_CollateralFactorTooOld(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newFactor := decimal.NewFromInt(800000000000000000)
	s.collateralFactorChangedCh <- &CollateralFactorChanged{
		Market:           vETHMarket,
		CollateralFactor: newFactor,
		UpdatedHeight:    height - 10,
	}

	s.wg.Add(1)
	go s.CollateralFactorLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 0, len(s.middleAccountSyncCh))
}

func Test_PriceChangedLoop_PriceDecrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newPrice := decimal.NewFromInt(900000000000000000)
	s.priceChangedCh <- &PriceChanged{
		Market:        vETHMarket,
		Price:         newPrice,
		UpdatedHeight: height + 10,
	}

	s.wg.Add(1)
	go s.PriceChangedLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 1, len(s.highAccountSyncCh))
	accounts := <-s.highAccountSyncCh
	assert.Equal(t, []common.Address{account2, account3}, accounts)
	assert.Equal(t, newPrice, s.prices[vETHMarket].Price)
}

func Test_PriceChangedLoop_PriceIncrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newPrice := decimal.NewFromInt(1100000000000000000)
	s.priceChangedCh <- &PriceChanged{
		Market:        vETHMarket,
		Price:         newPrice,
		UpdatedHeight: height + 10,
	}

	s.wg.Add(1)
	go s.PriceChangedLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 1, len(s.middleAccountSyncCh))
	accounts := <-s.middleAccountSyncCh
	assert.Equal(t, []common.Address{account2, account3}, accounts)
	assert.Equal(t, newPrice, s.prices[vETHMarket].Price)
}

func Test_PriceChangedLoop_PriceTooOld(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	newPrice := decimal.NewFromInt(1100000000000000000)
	s.priceChangedCh <- &PriceChanged{
		Market:        vETHMarket,
		Price:         newPrice,
		UpdatedHeight: height - 10,
	}

	s.wg.Add(1)
	go s.PriceChangedLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 0, len(s.middleAccountSyncCh))
}

func Test_VTokenAmountChangedLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vBTCMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vETHMarket,
		Account:       account3,
		UpdatedHeight: height + 3,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	//drawn out
	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh
	<-s.middleAccountSyncCh

	amount := decimal.NewFromInt(1100000000000000000)
	s.vTokenAmountChangedCh <- &VTokenAmountChanged{
		Market:        vETHMarket,
		From:          account2,
		To:            account3,
		Amount:        amount,
		UpdatedHeight: height + 10,
	}

	s.vTokenAmountChangedCh <- &VTokenAmountChanged{
		Market:        vBTCMarket,
		From:          account1,
		To:            account2,
		Amount:        amount,
		UpdatedHeight: height + 11,
	}

	s.wg.Add(1)
	go s.VTokenAmountChangedLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 2, len(s.highAccountSyncCh))
	accounts := <-s.highAccountSyncCh
	assert.Equal(t, []common.Address{account2}, accounts)

	accounts = <-s.highAccountSyncCh
	assert.Equal(t, []common.Address{account1}, accounts)

	assert.Equal(t, 1, len(s.middleAccountSyncCh))
	accounts = <-s.middleAccountSyncCh
	assert.Equal(t, []common.Address{account3}, accounts)
}

func Test_RepayVaiAmountChangedLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.repayVaiAmountChangedCh <- &RepayVaiAmountChanged{
		Account:       account1,
		Amount:        decimal.NewFromInt(1000000000000000000),
		UpdatedHeight: height + 1,
	}

	s.repayVaiAmountChangedCh <- &RepayVaiAmountChanged{
		Account:       account2,
		Amount:        decimal.NewFromInt(-1000000000000000000),
		UpdatedHeight: height + 2,
	}

	s.repayVaiAmountChangedCh <- &RepayVaiAmountChanged{
		Account:       account3,
		Amount:        decimal.NewFromInt(-2000000000000000000),
		UpdatedHeight: height + 2,
	}

	s.wg.Add(1)
	go s.RepayVaiAmountChangedLoop()
	time.Sleep(time.Second * 2)
	s.Stop()
	//drawn out
	assert.Equal(t, 1, len(s.highAccountSyncCh))
	assert.Equal(t, 2, len(s.middleAccountSyncCh))

	accounts := <-s.highAccountSyncCh
	assert.Equal(t, []common.Address{account1}, accounts)

	accounts = <-s.middleAccountSyncCh
	assert.Equal(t, []common.Address{account2}, accounts)

	accounts = <-s.middleAccountSyncCh
	assert.Equal(t, []common.Address{account3}, accounts)

	had, _ := s.db.Has(dbm.BorrowersStoreKey(account1.Bytes()), nil)
	assert.True(t, had)

	had, _ = s.db.Has(dbm.BorrowersStoreKey(account2.Bytes()), nil)
	assert.True(t, had)

	had, _ = s.db.Has(dbm.BorrowersStoreKey(account3.Bytes()), nil)
	assert.True(t, had)
}
