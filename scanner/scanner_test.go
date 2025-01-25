package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	"github.com/readygo586/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"testing"
	"time"
)

/*
Unitroller deployed to: 0xB4Abb34e08094B1915Ac3f7882aed81d0104b121
Comptroller deployed to: 0x4039C2a906D5eEc6A8F036dF248Cf14FF4274Ef2
USDT deployed to: 0x39d770382A22cdb61AD47B6faFC76A872d4fb3e8
USDC deployed to: 0xB167B4136446a07fFbC83946C0F66Fa4289e2953
price oracle deployed to: 0x6B392885f26b718C149f759B591094a06787A289
access control deployed to: 0x259ae555eeeE48E91e70bf5035484F039c009167
VAI deployed to: 0x7C4f97bF4c28732F9E257B6dF24D12C8Bf43E1f8
VAIController deployed to: 0x96ae4986D9ff19992dA84B5DBA9790cAE7246b80
vUSDT deployed to: 0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591
vUSDC deployed to: 0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28
*/

// bsc net
const (
	Url           = "https://frequent-side-arrow.bsc-testnet.quiknode.pro/d53e466c6ac0b3adaf534a1c641d6264ee4f9886"
	Comptroller   = "0xB4Abb34e08094B1915Ac3f7882aed81d0104b121"
	VaiController = "0x96ae4986D9ff19992dA84B5DBA9790cAE7246b80"
	Oralce        = "0x6B392885f26b718C149f759B591094a06787A289"
)

func TestBlockByHash_46372737(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	blockHash := common.HexToHash("0x69cb410b1b98bf543d543a716f99f2b9f0c9e93619ed5188b37c73e5e1d22ddd")

	block, err := c.BlockByHash(context.Background(), blockHash)
	assert.NoError(t, err)
	fmt.Printf("blockhash:%v\n", block.Hash())
	assert.Equal(t, blockHash, block.Hash())
}

func Test_ScanOneBlock_Non_VToken_Events(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.oracleAddr)
	query2 := buildVTokenQueryWithoutHeight(s.markets)

	heights := []int64{46341424, 46341448, 46341450, 46341454, 46341456, 46359440, 46369030, 46371979,
		46373558, 46387970, 46388361, 46388930, 46388955, 46389092, 46484194}
	for _, height := range heights {
		err = s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
		time.Sleep(20 * time.Microsecond)
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.closeFactorChangedCh))
	assert.Equal(t, 2, len(s.newMarketCh))
	assert.Equal(t, 2, len(s.collateralFactorChangedCh))
	assert.Equal(t, 8, len(s.enterMarketCh))
	assert.Equal(t, 3, len(s.exitMarketCh))

}

func Test_ScanOneBlock_VToken_Events(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	query1 := buildQueryWithoutHeight(s.comptrollerAddr, s.vaiControllerAddr, s.oracleAddr)
	query2 := buildVTokenQueryWithoutHeight(s.markets)

	heights := []int64{46359436, 46359438, 46371967, 46372646, 46373897, 46388393, 46484129, 46484210, 46485843, 46486375}
	for _, height := range heights {
		err = s.ScanOneBlock(height, []ethereum.FilterQuery{query1, query2})
		time.Sleep(20 * time.Microsecond)
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 10, len(s.vTokenAmountChangedCh))
}

func Test_SyncOneAccount(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	account := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
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

	bz, err = s.db.Get(dbm.LiquidationBelow2P0StoreKey(account.Bytes()), nil)
	assert.Equal(t, account.Bytes(), bz)
}

func Test_ScanLoop(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
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
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: big.NewInt(46341420).Uint64(),
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: big.NewInt(46341421).Uint64(),
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account3,
		UpdatedHeight: big.NewInt(46341422).Uint64(),
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 3, len(s.middleAccountSyncCh))
	had, _ := s.db.Has(dbm.MarketMemberStoreKey(vUSDTMarket.Bytes(), account1.Bytes()), nil)
	assert.True(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDCMarket.Bytes(), account2.Bytes()), nil)
	assert.True(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDCMarket.Bytes(), account3.Bytes()), nil)
	assert.True(t, had)

	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDCMarket.Bytes(), account1.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDTMarket.Bytes(), account2.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDTMarket.Bytes(), account3.Bytes()), nil)
	assert.False(t, had)

	s.exitMarketCh <- &ExitMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: big.NewInt(46341520).Uint64(),
	}

	s.exitMarketCh <- &ExitMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: big.NewInt(46341521).Uint64(),
	}

	s.exitMarketCh <- &ExitMarket{
		Market:        vUSDCMarket,
		Account:       account3,
		UpdatedHeight: big.NewInt(46341522).Uint64(),
	}

	s.wg.Add(1)
	go s.ExitMarketLoop()
	time.Sleep(time.Second * 2)
	s.Stop()

	assert.Equal(t, 3, len(s.highAccountSyncCh))
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDTMarket.Bytes(), account1.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDCMarket.Bytes(), account2.Bytes()), nil)
	assert.False(t, had)
	had, _ = s.db.Has(dbm.MarketMemberStoreKey(vUSDCMarket.Bytes(), account3.Bytes()), nil)
	assert.False(t, had)
}

func Test_CollateralFactorLoop_CollateralFactorDecrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account3,
		UpdatedHeight: height + 1,
	}

	s.wg.Add(1)
	go s.EnterMarketLoop()
	time.Sleep(time.Second * 2)

	newFactor := decimal.NewFromInt(700000000000000000)
	s.collateralFactorChangedCh <- &CollateralFactorChanged{
		Market:           vUSDCMarket,
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
	assert.Equal(t, newFactor, s.tokens[vUSDCMarket].CollateralFactor)
}

func Test_CollateralFactorLoop_CollateralFactorIncrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:           vUSDCMarket,
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
	assert.Equal(t, newFactor, s.tokens[vUSDCMarket].CollateralFactor)
}

func Test_CollateralFactorLoop_CollateralFactorNotChange(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:           vUSDCMarket,
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
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:           vUSDCMarket,
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
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:        vUSDCMarket,
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
	assert.Equal(t, newPrice, s.prices[vUSDCMarket].Price)
}

func Test_PriceChangedLoop_PriceIncrease(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:        vUSDCMarket,
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
	assert.Equal(t, newPrice, s.prices[vUSDCMarket].Price)
}

func Test_PriceChangedLoop_PriceTooOld(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:        vUSDCMarket,
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
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)

	vUSDTMarket := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	vUSDCMarket := common.HexToAddress("0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28")
	account1 := common.HexToAddress("0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	account2 := common.HexToAddress("0x658a6c7962e64132d2487EB2bc431d8Bc285882F")
	account3 := common.HexToAddress("0x7A2Fc9dc53103f15ec43CC3D1e69eFB73b860562")

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDTMarket,
		Account:       account1,
		UpdatedHeight: height + 1,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
		Account:       account2,
		UpdatedHeight: height + 2,
	}

	s.enterMarketCh <- &EnterMarket{
		Market:        vUSDCMarket,
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
		Market:        vUSDCMarket,
		From:          account2,
		To:            account3,
		Amount:        amount,
		UpdatedHeight: height + 10,
	}

	s.vTokenAmountChangedCh <- &VTokenAmountChanged{
		Market:        vUSDTMarket,
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
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")
	height, err := c.BlockNumber(context.Background())
	require.NoError(t, err)

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
