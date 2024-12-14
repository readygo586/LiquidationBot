package scanner

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	priceUpdate
	MintedVAI
	RepayVAI

	collateralFactor
	bigFactor
	Supply
	Redeem

 	AddMarkets
	EnterMarket
	ExitMarket

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
const Url = "https://frequent-side-arrow.bsc-testnet.quiknode.pro/d53e466c6ac0b3adaf534a1c641d6264ee4f9886"

// venus@bsc mainnet, https://docs-v4.venus.io/deployed-contracts/core-pool
//const Url = "https://bsc-dataseed4.bnbchain.org"
//const VaiController = "0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE"
//const Comptroller = "0xfD36E2c2a6789Db23113685031d7F16329158384"
//const Oralce = "0x6B392885f26b718C149f759B591094a06787A289"

func TestBlockByHash_46372737(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	blockHash := common.HexToHash("0x69cb410b1b98bf543d543a716f99f2b9f0c9e93619ed5188b37c73e5e1d22ddd")

	block, err := c.BlockByHash(context.Background(), blockHash)
	assert.NoError(t, err)
	fmt.Printf("blockhash:%v\n", block.Hash())
	assert.Equal(t, blockHash, block.Hash())
}
