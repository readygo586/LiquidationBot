package server

import (
	"context"
	"encoding/json"
	"fmt"
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

func TestMapStructAssignment(t *testing.T) {
	testmap := make(map[string]*TokenInfo)
	tokenInfo := &TokenInfo{
		Price: decimal.Zero,
	}
	testmap["usdt"] = tokenInfo
	testmap["usdt"].Price = decimal.NewFromInt(1)
}

func TestGetvAAVEUnderlyingPrice(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	oracle, err := venus.NewOracle(common.HexToAddress(cfg.Oracle), c)
	if err != nil {
		panic(err)
	}
	//fail to get vAAVE prices @ 20220201
	_, err = oracle.GetUnderlyingPrice(nil, common.HexToAddress("0x26DA28954763B92139ED49283625ceCAf52C6f94"))
	require.Equal(t, "execution reverted: REF_DATA_NOT_AVAILABLE", err.Error())
}

func TestGetUnderlyingDecimal(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	comptroller, err := venus.NewComptroller(common.HexToAddress(cfg.Comptroller), c)
	require.NoError(t, err)

	markets, err := comptroller.GetAllMarkets(nil)
	require.NoError(t, err)

	var underlyingAddress common.Address
	for _, market := range markets {

		vbep20, err := venus.NewVbep20(market, c)
		require.NoError(t, err)

		symbol, err := vbep20.Symbol(nil)
		require.NoError(t, err)
		fmt.Printf("market:%v, symbol:%v\n", market, symbol)
		if market == vBNBAddress {
			underlyingAddress = wBNBAddress
		} else {
			underlyingAddress, err = vbep20.Underlying(nil)
		}
		require.NoError(t, err)

		bep20, err := venus.NewVbep20(underlyingAddress, c)
		underlyingDecimals, err := bep20.Decimals(nil)
		require.NoError(t, err)

		underlyingSybmol, err := bep20.Symbol(nil)
		require.NoError(t, err)

		fmt.Printf("symbol:%v, underlyingSymbol:%v, underlyingDecimals:%v\n", symbol, underlyingSybmol, underlyingDecimals)
	}

}

func TestNewSyncer(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	verifyTokens(t, sync)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	require.NoError(t, err)

	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(0), num.Int64())

	for symbol, pair := range sync.flashLoanPool {
		t.Logf("%v's flashloan pool is:%v\n", symbol, pair)
	}

	for symbolPair, paths := range sync.uniswapPaths {
		t.Logf("%v's paths is:%v\n", symbolPair, paths)
	}
}

func TestDoSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	t.Logf("begin do sync markets and prices\n")

	sync.doSyncMarketsAndPrices()
	verifyTokens(t, sync)
}

func TestSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	t.Logf("begin sync markets and prices\n")
	sync.wg.Add(1)
	go sync.syncMarketsAndPrices()

	time.Sleep(time.Second * 60)
	close(sync.quitCh)
	sync.wg.Wait()
	verifyTokens(t, sync)
}

func TestFormulateUniswapPath1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	tokens := sync.tokens
	//pancakeRouter := sync.pancakeRouter
	pancakeFactory := sync.pancakeFactory

	pair, err := pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vXVS"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("pair:%v\n", pair)
	pair, err = pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vBNB"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("vLTCvBNB pair:%v\n", pair)
	pair, err = pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vUSDT"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("vLTCvUSDT pair:%v\n", pair)
	pair, err = pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vDAI"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("vLTCvDAI pair:%v\n", pair)
	pair, err = pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vUSDC"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("vLTCvUSDC pair:%v\n", pair)
	pair, err = pancakeFactory.GetPair(nil, tokens["vLTC"].UnderlyingAddress, tokens["vTUSD"].UnderlyingAddress)
	require.NoError(t, err)
	fmt.Printf("vLTCvTUSD pair:%v\n", pair)
}

func TestFormulateUniswapPath2(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	tokens := sync.tokens
	//pancakeRouter := sync.pancakeRouter
	pancakeFactory := sync.pancakeFactory

	interSymbols := []string{"vBNB", "vUSDT"}
	connection := make(map[string]int)

	for _, interSymbol := range interSymbols {
		for symbol, _ := range tokens {
			if symbol == interSymbol {
				continue
			}
			pair, _ := pancakeFactory.GetPair(nil, tokens[interSymbol].UnderlyingAddress, tokens[symbol].UnderlyingAddress)
			if pair.String() != "0x0000000000000000000000000000000000000000" {
				connection[interSymbol]++
			} else {
				fmt.Printf("missed %v%v path\n", interSymbol, symbol)
			}
		}

	}

	for _, interSymbol := range interSymbols {
		fmt.Printf("%v's connection %v\n", interSymbol, connection[interSymbol])
	}

}

func TestFormulateUniswapPath3(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	tokens := sync.tokens
	//pancakeRouter := sync.pancakeRouter
	pancakeFactory := sync.pancakeFactory

	interSymbols := []string{"vCAN"}
	connection := make(map[string]int)

	for _, interSymbol := range interSymbols {
		for symbol, _ := range tokens {
			if symbol == interSymbol {
				continue
			}
			pair, _ := pancakeFactory.GetPair(nil, tokens[interSymbol].UnderlyingAddress, tokens[symbol].UnderlyingAddress)
			if pair.String() != "0x0000000000000000000000000000000000000000" {
				connection[interSymbol]++
			} else {
				fmt.Printf("missed %v%v path\n", interSymbol, symbol)
			}
		}

	}

	for _, interSymbol := range interSymbols {
		fmt.Printf("%v's connection %v\n", interSymbol, connection[interSymbol])
	}

}

func TestFormulateUniswapPath4(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	tokens := sync.tokens
	pancakeRouter := sync.pancakeRouter
	//pancakeFactory := sync.pancakeFactory

	tmpPaths := make([]common.Address, 3)
	tmpPaths[0] = tokens["vSXP"].UnderlyingAddress
	tmpPaths[1] = tokens["vBNB"].UnderlyingAddress
	tmpPaths[2] = tokens["vTRX"].UnderlyingAddress
	amountOut := big.NewInt(10000000000000000)
	amountsIn, err := pancakeRouter.GetAmountsIn(nil, amountOut, tmpPaths)
	require.NoError(t, err)
	t.Logf("amountsIn%v", amountsIn)

	//for _, interSymbol := range interSymbols {
	//	for symbol, _ := range tokens {
	//		if symbol == interSymbol {
	//			continue
	//		}
	//
	//		pair, _ := pancakeFactory.GetPair(nil, tokens[interSymbol].UnderlyingAddress, tokens[symbol].UnderlyingAddress)
	//		if pair.String() != "0x0000000000000000000000000000000000000000" {
	//			connection[interSymbol]++
	//		} else {
	//			fmt.Printf("missed %v%v path\n", interSymbol, symbol)
	//		}
	//	}
	//
	//}
	//
	//for _, interSymbol := range interSymbols {
	//	fmt.Printf("%v's connection %v\n", interSymbol, connection[interSymbol])
	//}

}

func TestFormulateUniswapPath(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)

	//pancakeRouter := sync.pancakeRouter
	pancakeFactory := sync.pancakeFactory

	tokens := sync.tokens
	//hash := crypto.Keccak256Hash
	paths := make(map[string][]common.Address)
	flashLoanMarkets := make(map[string]common.Address)

	//interSymbols := []string{"vBNB", "vUSDT"}

	for srcSymbol, srcToken := range tokens {
		srcBep20, err := venus.NewBep20(tokens[srcSymbol].UnderlyingAddress, sync.c)
		require.NoError(t, err)

		maxSrcAmount := big.NewInt(0)
		maxSrcMaret := common.Address{}

		for dstSymbol, dstToken := range tokens {
			if srcSymbol == dstSymbol {
				continue
			}

			//fmt.Printf("srcSymbol:%v, dstSymol:%v\n", srcSymbol, dstSymbol)
			pair, err := pancakeFactory.GetPair(nil, srcToken.UnderlyingAddress, dstToken.UnderlyingAddress)
			if err != nil || pair.String() == "0x0000000000000000000000000000000000000000" {
				//amountOut := big.NewInt(1000000000000000000)
				tmpPaths := make([]common.Address, 3)
				tmpPaths[0] = srcToken.UnderlyingAddress
				tmpPaths[1] = tokens["vBNB"].UnderlyingAddress
				tmpPaths[2] = dstToken.UnderlyingAddress
				paths[srcSymbol+dstSymbol] = tmpPaths
				fmt.Printf("%v%v%v: paths:%v\n", srcSymbol, "vBNB", dstSymbol, paths)

				//minAmountIn := big.NewInt(big.MaxPrec)
				//selectedInterSymbol := ""
				//
				//for _, interSymbol := range interSymbols {
				//	fmt.Printf("srcSymbol %v, dstSymbol:%v, iterSymbol:%v, srcAddress:%v, iterAddress:%v\n", srcSymbol, dstSymbol, interSymbol, srcToken.UnderlyingAddress, tokens[interSymbol].UnderlyingAddress)
				//	_, err := pancakeFactory.GetPair(nil, srcToken.UnderlyingAddress, tokens[interSymbol].UnderlyingAddress)
				//	if err != nil || pair.String() == "0x0000000000000000000000000000000000000000" {
				//		continue
				//	}
				//
				//	tmpPaths[1] = tokens[interSymbol].UnderlyingAddress
				//	amountsIn, err := pancakeRouter.GetAmountsIn(nil, amountOut, tmpPaths)
				//	if err != nil {
				//		fmt.Printf("getAmountsIn, %v%v:path:%v\n", srcSymbol, dstSymbol, paths)
				//		continue
				//	}
				//
				//	if amountsIn[0].Cmp(minAmountIn) == -1 {
				//		minAmountIn = amountsIn[0]
				//		selectedInterSymbol = interSymbol
				//	}
				//}
				//
				//if selectedInterSymbol != "" {
				//	paths[srcSymbol+dstSymbol] = tmpPaths
				//	fmt.Printf("%v%v%v: paths:%v\n", srcSymbol, selectedInterSymbol, dstSymbol, paths)
				//}

			} else {
				//select the deepest market as flashloan from
				srcAmout, err := srcBep20.BalanceOf(nil, pair)
				if err != nil {
					srcAmout = big.NewInt(0)
				}
				if srcAmout.Cmp(maxSrcAmount) == 1 {
					maxSrcAmount = srcAmout
					maxSrcMaret = pair
				}

				//formulate the path
				tmpPaths := make([]common.Address, 2)
				tmpPaths[0] = tokens[srcSymbol].UnderlyingAddress
				tmpPaths[1] = tokens[dstSymbol].UnderlyingAddress
				paths[srcSymbol+dstSymbol] = tmpPaths
			}
			//fmt.Printf("paths[%v%v]= %v\n", srcSymbol, dstSymbol, paths[srcSymbol+dstSymbol])
		}
		flashLoanMarkets[srcSymbol] = maxSrcMaret
		//fmt.Printf("flashLoanMarket[%v] = %v\n", srcSymbol, flashLoanMarkets[srcSymbol])
	}

	for srcSymbol, _ := range tokens {
		for dstSymbol, _ := range tokens {
			fmt.Printf("%v\n", srcSymbol)
			fmt.Printf("flashLoanMarket[%v] = %v\n", srcSymbol, flashLoanMarkets[srcSymbol])
			fmt.Printf("paths[%v%v]= %v\n", srcSymbol, dstSymbol, paths[srcSymbol+dstSymbol])
		}
	}

}
func TestFilterAllCotractsBorrowEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, nil, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)

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
	fmt.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var borrowEvent venus.Vbep20Borrow
		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
		fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)
	}
	fmt.Printf("end Time:%v\n", time.Now())
}

func TestCalculateHealthFactor(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	comptroller := sync.comptroller
	oracle := sync.oracle

	accounts := []string{
		//"0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B",
		//"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
		//"0xc528045078Ff53eA289fA42aF3e12D8eF901cABD",
		"0xF2ddE5689B0e13344231D3459B03432E97a48E0c",
	}

	for _, account := range accounts {
		_, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, common.HexToAddress(account))
		require.NoError(t, err)

		assets, err := comptroller.GetAssetsIn(nil, common.HexToAddress(account))
		fmt.Printf("assets:%v\n", assets)
		require.NoError(t, err)

		totalCollateral := decimal.NewFromInt(0)
		totalLoan := decimal.NewFromInt(0)
		bigMintedVAIS, err := comptroller.MintedVAIs(nil, common.HexToAddress(account))

		mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)

		for _, asset := range assets {
			//fmt.Printf("asset:%v\n", asset)
			marketInfo, err := comptroller.Markets(nil, asset)
			require.NoError(t, err)

			token, err := venus.NewVbep20(asset, c)
			require.NoError(t, err)

			errCode, bigBalance, bigBorrow, bigExchangeRate, err := token.GetAccountSnapshot(nil, common.HexToAddress(account))
			require.NoError(t, err)
			require.True(t, errCode.Cmp(BigZero) == 0)

			if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
				continue
			}

			bigPrice, err := oracle.GetUnderlyingPrice(nil, asset)
			if bigPrice.Cmp(BigZero) == 0 {
				continue
			}

			exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
			collateralFactor := decimal.NewFromBigInt(marketInfo.CollateralFactorMantissa, 0)
			price := decimal.NewFromBigInt(bigPrice, 0)
			balance := decimal.NewFromBigInt(bigBalance, 0)
			borrow := decimal.NewFromBigInt(bigBorrow, 0)
			fmt.Printf("collateralFactor:%v, price:%v, exchangeRate:%v, balance:%v, borrow:%v\n", collateralFactor, bigPrice, bigExchangeRate, bigBalance, bigBorrow)

			multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
			multiplier = multiplier.Mul(price).Div(EXPSACLE)
			collateral := balance.Mul(multiplier).Div(EXPSACLE)
			totalCollateral = totalCollateral.Add(collateral.Truncate(0))

			loan := borrow.Mul(price).Div(EXPSACLE)
			totalLoan = totalLoan.Add(loan.Truncate(0))
		}

		totalLoan = totalLoan.Add(mintedVAIS)
		fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral.String(), totalLoan)
		healthFactor := decimal.NewFromInt(100)
		if totalLoan.Cmp(decimal.Zero) == 1 {
			healthFactor = totalCollateral.Div(totalLoan)
		}

		fmt.Printf("healthFactor：%v\n", healthFactor)
		calculatedLiquidity := decimal.NewFromInt(0)
		calculatedShortfall := decimal.NewFromInt(0)
		if totalLoan.Cmp(totalCollateral) == 1 {
			calculatedShortfall = totalLoan.Sub(totalCollateral)
		} else {
			calculatedLiquidity = totalCollateral.Sub(totalLoan)
		}

		fmt.Printf("liquidity:%v, calculatedLiquidity:%v\n", liquidity.String(), calculatedLiquidity.String())
		fmt.Printf("shortfall:%v, calculatedShortfall:%v\n", shortfall, calculatedShortfall)
	}
}

func TestStoreAndDeleteAccount(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)

	healthFactor, _ := decimal.NewFromString("0.9")

	vusdtBalance, _ := decimal.NewFromString("1000000000.0")
	vusdtLoan, _ := decimal.NewFromString("0")

	vbtcBalance, _ := decimal.NewFromString("2.5")
	vbtctLoan, _ := decimal.NewFromString("0.2")

	vbusdBalance, _ := decimal.NewFromString("0")
	vbusdtLoan, _ := decimal.NewFromString("500.23")

	assets := []Asset{
		{
			Symbol:  "vUSDT",
			Balance: vusdtBalance,
			Loan:    vusdtLoan,
		},
		{
			Symbol:  "vBTC",
			Balance: vbtcBalance,
			Loan:    vbtctLoan,
		},
		{
			Symbol:  "vBUSD",
			Balance: vbusdBalance,
			Loan:    vbusdtLoan,
		},
	}
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}

	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
	sync.storeAccount(account, info)

	bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	var got AccountInfo
	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	has, err := db.Has(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())

		prefix := append(dbm.MarketPrefix, []byte(asset.Symbol)...)
		var accounts []common.Address
		iter := db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}

		require.Equal(t, 1, len(accounts))
	}

	had, err := db.Has(dbm.MarketStoreKey([]byte("vETH"), account.Bytes()), nil)
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
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)

	healthFactor, _ := decimal.NewFromString("1.1")
	vusdtBalance, _ := decimal.NewFromString("1000000000.0")
	vusdtLoan, _ := decimal.NewFromString("0")

	vbtcBalance, _ := decimal.NewFromString("2.5")
	vbtctLoan, _ := decimal.NewFromString("0.2")

	vbusdBalance, _ := decimal.NewFromString("0")
	vbusdtLoan, _ := decimal.NewFromString("500.23")

	assets := []Asset{
		{
			Symbol:  "vUSDT",
			Balance: vusdtBalance,
			Loan:    vusdtLoan,
		},
		{
			Symbol:  "vBTC",
			Balance: vbtcBalance,
			Loan:    vbtctLoan,
		},
		{
			Symbol:  "vBUSD",
			Balance: vbusdBalance,
			Loan:    vbusdtLoan,
		},
	}
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}

	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
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
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}

	had, err := db.Has(dbm.MarketStoreKey([]byte("vETH"), account.Bytes()), nil)
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
		Assets:       assets,
	}

	sync.storeAccount(account, info)
	bz, err = db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	for _, asset := range assets {
		//fmt.Printf("symbol:%v\n", asset.Symbol)
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}
}

// 从compound通过getExchangeRateStored方法获得的exchangeRat是乘了10^18的结果，实际使用时需要除10^18,
func TestCalculateExchangeRate(t *testing.T) {
	//exchangeRateStored: 202001285536565656590891932
	//totalSupply: 76384766592957
	//totalBorrow: 2298168762317337651162
	//totalReserver:  4713643651873292071
	//cash: 13136365928522364031146
	borrow, _ := decimal.NewFromString("2298168762317337651162")
	supply, _ := decimal.NewFromString("76384766592957")
	reserve, _ := decimal.NewFromString("4713643651873292071")
	cash, _ := decimal.NewFromString("13136365928522364031146")
	sum := cash.Add(borrow)
	sum = sum.Sub(reserve)
	rate := sum.Div(supply)
	fmt.Printf("rate:%v\n", rate)

	rateExp := sum.Mul(EXPSACLE).Div(supply)
	//ExpScale, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	//sumExp := big.NewInt(0).Mul(sum, ExpScale)
	//rateExp := big.NewInt(0).Div(sumExp, supply)
	////fmt.Printf("rateExp:%v\n", rateExp)
	require.Equal(t, "202001285536565656590891932", rateExp.Truncate(0).String())
}

func TestSyncOneAccount(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
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

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	fmt.Printf("info:%v\n", info)
	require.True(t, decimal.NewFromInt(100).Cmp(info.HealthFactor) == 1)

	bz, err = db.Get(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))
}

func TestSyncOneAccount1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0x1E73902Ab4144299DFc2ac5a3765122c02CE889f") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	//accountBytes := account.Bytes()
	err = sync.syncOneAccount(account)
	require.NoError(t, err)
}

//
//func TestSyncAccounts(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	liquidationCh := make(chan *Liquidation, 64)
//	priorityliquidationCh := make(chan *Liquidation, 64)
//	feededPricesCh := make(chan *FeededPrices, 64)
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
//	accountStrs := []string{
//		"0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B",
//		"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
//		"0xc528045078Ff53eA289fA42aF3e12D8eF901cABD",
//		"0xF2ddE5689B0e13344231D3459B03432E97a48E0c",
//	}
//	var accounts []common.Address
//	for _, accountStr := range accountStrs {
//		accounts = append(accounts, common.HexToAddress(accountStr))
//	}
//
//	sync.syncAccounts(accounts)
//	require.NoError(t, err)
//
//	count := make(map[string]int)
//
//	for _, account := range accounts {
//		accountBytes := account.Bytes()
//		bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//		require.Equal(t, account, common.BytesToAddress(bz))
//
//		bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//
//		var got AccountInfo
//		err = json.Unmarshal(bz, &got)
//		require.NoError(t, err)
//
//		for _, token := range got.Assets {
//			count[token.Symbol] += 1
//			bz, err = db.Get(dbm.MarketStoreKey([]byte(token.Symbol), accountBytes), nil)
//			require.NoError(t, err)
//			require.Equal(t, account, common.BytesToAddress(bz))
//		}
//	}
//
//	//total account number = 4
//	gotAccounts := []common.Address{}
//	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 4, len(gotAccounts))
//
//	for symbol, cnt := range count {
//		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
//		gotAccounts = []common.Address{}
//		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
//		for iter.Next() {
//			gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//		}
//		require.Equal(t, cnt, len(gotAccounts))
//	}
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P5Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 2, len(gotAccounts))
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow3P0Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 1, len(gotAccounts))
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationAbove3P0Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 1, len(gotAccounts))
//}

func TestSyncOneAccountWithIncreaseAccountNumer(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B")
	accountBytes := account.Bytes()
	err = sync.syncOneAccountWithIncreaseAccountNumber(account)
	require.NoError(t, err)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(1), num.Int64())

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

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	fmt.Printf("info:%v\n", info)
	require.True(t, decimal.NewFromInt(100).Cmp(info.HealthFactor) == 0)

	bz, err = db.Get(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))
}

func TestSyncOneAccountWithFeededPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B")
	accountBytes := account.Bytes()

	feededUsdtPrice, _ := decimal.NewFromString("1.02")
	feededDogePrice, _ := decimal.NewFromString("0.04")
	feededPrices := &FeededPrices{
		Prices: []FeededPrice{
			{
				Address: common.HexToAddress("0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71"),
				Price:   feededUsdtPrice,
			},
			{
				Address: common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"),
				Price:   feededDogePrice,
			},
		},
		Height: 100000,
	}
	err = sync.syncOneAccountWithFeededPrices(account, feededPrices)
	require.NoError(t, err)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(0), num.Int64())

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	exist, err = db.Has(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	exist, err = db.Has(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	exist, err = db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)
}

//func TestScanAllBorrowers(t *testing.T) {
//	ctx := context.Background()
//
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	_, err = c.BlockNumber(ctx)
//	require.NoError(t, err)
//
//	sync := NewSyncer(c, db, cfg)
//	start := big.NewInt(13000000) //big.NewInt(14747565)
//	db.Put(dbm.KeyLastHandledHeight, start.Bytes(), nil)
//	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)
//
//	sync.Start()
//	time.Sleep(time.Second * 120)
//	sync.Stop()
//
//	bz, err := db.Get(dbm.KeyLastHandledHeight, nil)
//	end := big.NewInt(0).SetBytes(bz)
//	t.Logf("end height:%v\n", end.Int64())
//
//	bz, err = db.Get(dbm.KeyBorrowerNumber, nil)
//	num := big.NewInt(0).SetBytes(bz).Int64()
//	t.Logf("num:%v\n", num)
//
//	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
//	defer iter.Release()
//	t.Logf("borrows address")
//	for iter.Next() {
//		addr := common.BytesToAddress(iter.Value())
//		t.Logf("%v\n", addr.String())
//	}
//}
//
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

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
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

func TestCalculateSeizedToken(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	liquidation := Liquidation{
		Address: common.HexToAddress("0x1E73902Ab4144299DFc2ac5a3765122c02CE889f"),
	}
	sync.calculateSeizedTokenAmount(&liquidation)
}

func TestCalculateSeizedToken1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	liquidation := Liquidation{
		Address: common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627"),
	}

	err = sync.calculateSeizedTokenAmount(&liquidation)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestCalculateSeizedTokenGetAmountsOutWithMulOverFlow(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	liquidation := Liquidation{
		Address: common.HexToAddress("1e73902ab4144299dfc2ac5a3765122c02ce889f"),
	}

	err = sync.calculateSeizedTokenAmount(&liquidation)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestCalculateSeizedTokenGetAmountsInExecutionRevert(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)
	liquidation := Liquidation{
		Address: common.HexToAddress("ba3b9a3ecf19e1139c78c4718d45fb99f7a838cd"),
	}

	err = sync.calculateSeizedTokenAmount(&liquidation)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestCalculateSeizedTokens(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, feededPricesCh, liquidationCh, priorityliquidationCh)

	accounts := []string{
		"0x1E73902Ab4144299DFc2ac5a3765122c02CE889f",
		"0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627",
		"0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC",
		"0x0e0c57Ae65739394b405bC3afC5003bE9f858fDB",
		"0x2eB71e5335d5328e76fa0755Db27E184Be834D31",
		"0x4F41889788528e213692af181B582519BF4Cd30E",
		"0x564EE8bF0bA977A1ccc92fe3D683AbF4569c9f5E",
		"0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD",
		"0x89fa3aec0A7632dDBbdBaf448534f26BA4B771F1",
		"0xFAbE4C180b6eDad32eA0Cf56587c54417189e422",
		"0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7",
		"0xdcC896d48B17ECC88a9011057294EB0905bCb240",
		"0xfDA2b6948E01525633B4058297bb89656609e6Ad",
		"0xEAFb5e9E52A865D7BF1D3a9C17e0d29710928b8b",
		"0x05bbf0C12882FDEcd53FD734731ad578aF79621C",
	}

	for _, account := range accounts {
		liquidation := Liquidation{
			Address: common.HexToAddress(account),
		}
		err := sync.calculateSeizedTokenAmount(&liquidation)
		if err != nil {
			t.Logf(err.Error())
		}
	}
}

/*
verify pending liquidation:&{0xFAbE4C180b6eDad32eA0Cf56587c54417189e422 0.974535755200296 15008266 2022-02-06 11:47:03.292206 +0800 CST m=+33578.787466126}
verify pending liquidation:&{0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7 0.8819686150405764 15008266 2022-02-06 11:47:05.618654 +0800 CST m=+33581.113938293}
verify pending liquidation:&{0xdcC896d48B17ECC88a9011057294EB0905bCb240 0.9879989476061114 15008267 2022-02-06 11:47:05.94491 +0800 CST m=+33581.440198085}
verify pending liquidation:&{0xfDA2b6948E01525633B4058297bb89656609e6Ad 0.9570252601324154 15008267 2022-02-06 11:47:06.487259 +0800 CST m=+33581.982551793}
verify pending liquidation:&{0xEAFb5e9E52A865D7BF1D3a9C17e0d29710928b8b 0.9699014328167632 15008267 2022-02-06 11:47:08.815577 +0800 CST m=+33584.310894293}
verify pending liquidation:&{0x05bbf0C12882FDEcd53FD734731ad578aF79621C 0 15008270 2022-02-06 11:47:14.605148 +0800 CST m=+33590.100524751}
verify pending liquidation:&{0x07d1c21878C2f84BAE1DD3bA2C674d92133cc282 0.8938938376798766 15008270 2022-02-06 11:47:14.614635 +0800 CST m=+33590.110011876}
verify pending liquidation:&{0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627 0.9643391777901693 15008270 2022-02-06 11:47:15.675667 +0800 CST m=+33591.171055668}
verify pending liquidation:&{0x02360b97bBC9729916B470F699DF75Ff651bF926 0.3290733449378455 15008270 2022-02-06 11:47:16.200425 +0800 CST m=+33591.695818168}
verify pending liquidation:&{0x0fe11130B1819e2E3E5e5308b9EA16fFDa2032a6 0.9653441663232362 15008270 2022-02-06 11:47:16.343301 +0800 CST m=+33591.838696501}
verify pending liquidation:&{0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC 0.8580776654144922 15008270 2022-02-06 11:47:16.722097 +0800 CST m=+33592.217495960}
verify pending liquidation:&{0x0e0c57Ae65739394b405bC3afC5003bE9f858fDB 0.8568370199332438 15008270 2022-02-06 11:47:17.401952 +0800 CST m=+33592.897358293}
verify pending liquidation:&{0x1E73902Ab4144299DFc2ac5a3765122c02CE889f 0.7494185449809593 15008271 2022-02-06 11:47:18.643962 +0800 CST m=+33594.139380501}
verify pending liquidation:&{0x1743F248e67c810c8851f70B39b6578f36e9dD10 0.658660147678469 15008271 2022-02-06 11:47:18.841001 +0800 CST m=+33594.336422460}
verify pending liquidation:&{0x271f80305d43f6617840285ADC57A9D39d6d9F62 0 15008271 2022-02-06 11:47:19.177304 +0800 CST m=+33594.672728710}
verify pending liquidation:&{0x2eB71e5335d5328e76fa0755Db27E184Be834D31 0.9048364603440113 15008271 2022-02-06 11:47:19.900623 +0800 CST m=+33595.396054960}
verify pending liquidation:&{0x0C13Fafb81AAbA173547eD5D1941bD8b1f182962 0.7943135451562215 15008271 2022-02-06 11:47:20.441521 +0800 CST m=+33595.936958001}
*/

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
