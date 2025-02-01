package scanner

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSupportMarketEvent_46341448(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)
	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341448)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodeMarketListed(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", event.Market.Hex())
	assert.EqualValues(t, 46341448, event.UpdatedHeight)
}

func Test_NewCloseFactorEvent_46341424(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)
	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341424)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodeNewCloseFactor(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "500000000000000000", event.CloseFactor.String())
	assert.EqualValues(t, 46341424, event.UpdatedHeight)
}

func TestNewCollateralFactorEvent_46341454(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341454)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	event, err := decodeNewCollateralFactor(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", event.Market.Hex())
	assert.Equal(t, "800000000000000000", event.CollateralFactor.String())
	assert.EqualValues(t, 46341454, event.UpdatedHeight)
}

func TestMarketEnteredEvent_46388955(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46388955)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logs))

	event, err := decodeMarketEntered(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28", event.Market.Hex())
	assert.Equal(t, "0xc6B21654b936188158b788Ada6679f1c3463293c", event.Account.Hex())
	assert.EqualValues(t, 46388955, event.UpdatedHeight)
}

func TestMarketExited_46389092(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46389092)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logs))

	event, err := decodeMarketExited(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", event.Market.Hex())
	assert.Equal(t, "0xc6B21654b936188158b788Ada6679f1c3463293c", event.Account.Hex())
	assert.EqualValues(t, 46389092, event.UpdatedHeight)
}

func TestMintVaiEvent_46372737(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46372737)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	event, err := decodeMintVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.Account.Hex())
	assert.Equal(t, "32000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46372737, event.UpdatedHeight)
}

func TestRepayVaiEvent_46373178(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46373178)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	event, err := decodeRepayVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.Account.Hex())
	assert.Equal(t, "-16000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46373178, event.UpdatedHeight)
}

func TestLiquidateVaiEvent_46373178(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46373178)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	event, err := decodeRepayVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.Account.Hex())
	assert.Equal(t, "-16000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46373178, event.UpdatedHeight)
}

func TestVUSDTMintEvent_46359438(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46359438)

	vUSDT := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	filter := buildVTokenQueryWithoutHeight([]common.Address{vUSDT})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", event.From.Hex())
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", event.To.Hex())
	assert.Equal(t, "100000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46359438, event.UpdatedHeight)
}

func TestVUSDTRedeem_46372646(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46372646)

	vUSDT := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	filter := buildVTokenQueryWithoutHeight([]common.Address{vUSDT})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.From.Hex())
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", event.To.Hex())
	assert.Equal(t, "20000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46372646, event.UpdatedHeight)
}

func TestVUSDTTransfer_46486375(t *testing.T) {
	cfg, err := config.New("../config_test.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)
	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46486375)

	vUSDT := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	filter := buildVTokenQueryWithoutHeight([]common.Address{vUSDT})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.From.Hex())
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", event.To.Hex())
	assert.Equal(t, "100000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46486375, event.UpdatedHeight)
}

func TestBtcPriceUpdated_47700455(t *testing.T) {
	feederMap := make(map[common.Address]common.Address)
	btcFeeder := common.HexToAddress("0x33deb1bCDCC9ecc2056F87A20CFF3dcBd54a37f6")
	//ethFeeder := common.HexToAddress("0x11ffA6965b4c25790980897241100dA23b87C7f2")
	vBTCMarket := common.HexToAddress("0xaa46Fe4fc775A51117808b85f7b5D974040cdE0e")
	feederMap[btcFeeder] = vBTCMarket
	cfg, err := config.New("../config.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)
	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(47700455)

	filter := buildQueryWithoutHeight(common.Address{}, common.Address{}, []common.Address{btcFeeder})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodePriceUpdate(feederMap, logs[0])
	assert.NoError(t, err)
	assert.Equal(t, vBTCMarket.Hex(), event.Market.Hex())
	assert.Equal(t, "107593430000000000000000", event.Price.String())
	assert.EqualValues(t, 47700455, event.UpdatedHeight)
}

func TestEthPriceUpdated_47700433(t *testing.T) {
	feederMap := make(map[common.Address]common.Address)
	ethFeeder := common.HexToAddress("0x11ffA6965b4c25790980897241100dA23b87C7f2")
	vETHMarket := common.HexToAddress("0x5a57B04Bc33f7E22daED781fa32cB074241BeA09")
	feederMap[ethFeeder] = vETHMarket

	//ethFeeder := common.HexToAddress("0x11ffA6965b4c25790980897241100dA23b87C7f2")
	cfg, err := config.New("../config.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)
	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(47700433)

	filter := buildQueryWithoutHeight(common.Address{}, common.Address{}, []common.Address{ethFeeder})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	event, err := decodePriceUpdate(feederMap, logs[0])
	assert.NoError(t, err)
	assert.Equal(t, vETHMarket.Hex(), event.Market.Hex())
	assert.Equal(t, "3392450000000000000000", event.Price.String())
	assert.EqualValues(t, 47700433, event.UpdatedHeight)
}

func TestEventSignature(t *testing.T) {
	eventSignature := "MintVAI(address,uint256)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))
	fmt.Printf("Keccak256 hash of '%s': %s\n", eventSignature, hash.Hex())

	eventSignature = "RepayVAI(address,address,uint256)"
	hash = crypto.Keccak256Hash([]byte(eventSignature))
	fmt.Printf("Keccak256 hash of '%s': %s\n", eventSignature, hash.Hex())

	eventSignature = "LiquidateVAI(address,address,uint256,address,uint256)"
	hash = crypto.Keccak256Hash([]byte(eventSignature))
	fmt.Printf("Keccak256 hash of '%s': %s\n", eventSignature, hash.Hex())

	eventSignature = "LiquidateBorrow(address,address,uint256,address,uint256)"
	hash = crypto.Keccak256Hash([]byte(eventSignature))
	fmt.Printf("Keccak256 hash of '%s': %s\n", eventSignature, hash.Hex())

	//MintBehalf(address payer, address receiver, uint mintAmount, uint mintTokens);
	eventSignature = "MintBehalf(address,address,uint256,uint256)"
	hash = crypto.Keccak256Hash([]byte(eventSignature))
	fmt.Printf("Keccak256 hash of '%s': %s\n", eventSignature, hash.Hex())
}

func TestLiquidateVaiEvent_47797935(t *testing.T) {
	cfg, err := config.New("../config.yml")
	assert.NoError(t, err)
	c, err := ethclient.Dial(cfg.RpcUrl)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(47797935)

	filter := buildQueryWithoutHeight(common.HexToAddress(cfg.Comptroller), common.HexToAddress(cfg.VaiController), nil)
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	event, err := decodeRepayVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", event.Account.Hex())
	assert.Equal(t, "-16000000000000000000", event.Amount.String())
	assert.EqualValues(t, 46373178, event.UpdatedHeight)
}

func Test1E10(t *testing.T) {
	OneE10 := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(10), nil)
	fmt.Printf("1e10: %s\n", OneE10.String())
	multiplier := big.NewInt(10000000000)
	fmt.Printf("multiplier: %s\n", multiplier.String())
}
