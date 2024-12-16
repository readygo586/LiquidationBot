package scanner

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	"github.com/readygo586/LiquidationBot/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
