package scanner

import (
	"github.com/shopspring/decimal"
	"math/big"
)

//const (
//	ChainID = 10086
//)

const (
	ForbiddenPeriodForBadLiquidation     = 200 //200 block
	ForbiddenPeriodForPendingLiquidation = 200
)

const (
	//topics in Comptroller
	MarketListed        = "0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f"
	NewCloseFactor      = "0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9"
	NewCollateralFactor = "0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5"
	MarketEntered       = "0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5"
	MarketExited        = "0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d"

	//topics in VaiController
	MintVAI  = "0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a"
	RepayVAI = "0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e"
	
	//topics in vToken
	Transfer = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	//PriceUpdated
	PriceUpdated = "0x7d8cee5d1217e47a14a662098e84a7758580aaf78f430c07c543249234e867bf"
)

const (
	ConfirmHeight = 0
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
	ProfitThreshold       = decimal.New(1, 18)           //1 USDT
	MaxLoanValueThreshold = decimal.New(100, 18)         //100 USDT
)
