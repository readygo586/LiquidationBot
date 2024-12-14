package scanner

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSupportMarketEvent_46341448(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341448)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	market, err := decodeMarketListed(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", market.Hex())
}

func TestNewCloseFactorEvent_46341424(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341424)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	closeFactor, err := decodeNewCloseFactor(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "500000000000000000", closeFactor.String())
}

func TestNewCollateralFactorEvent_46341454(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46341454)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	address, closeFactor, err := decodeNewCollateralFactor(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", address.Hex())
	assert.Equal(t, "800000000000000000", closeFactor.String())
}

func TestMarketEnteredEvent_46388955(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46388955)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logs))

	address, account, err := decodeMarketEntered(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x0dB931cE74a54Ed1c04Bef1ad2459F829dC4fa28", address.Hex())
	assert.Equal(t, "0xc6B21654b936188158b788Ada6679f1c3463293c", account.Hex())
}

func TestMarketExited_46389092(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46389092)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logs))

	address, account, err := decodeMarketExited(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", address.Hex())
	assert.Equal(t, "0xc6B21654b936188158b788Ada6679f1c3463293c", account.Hex())
}

func TestMintVaiEvent_46372737(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46372737)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	address, amount, err := decodeMintVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", address.Hex())
	assert.Equal(t, "32000000000000000000", amount.String())
}

func TestRepayVaiEvent_46373178(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46373178)

	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce))
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)

	address, amount, err := decodeRepayVAI(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", address.Hex())
	assert.Equal(t, "16000000000000000000", amount.String())
}

func TestLiquidateVaiEvent_46373178(t *testing.T) {
	//FIXME(keep), no testdata
	//c, err := ethclient.Dial(Url)
	//assert.NoError(t, err)
	//
	//_, err = c.BlockNumber(context.Background())
	//assert.NoError(t, err)
	//
	//blockHeight := big.NewInt(46373178)
	//
	//filter := buildQueryWithoutHeight(nil)
	//filter.FromBlock = blockHeight
	//filter.ToBlock = blockHeight
	//
	//logs, err := c.FilterLogs(context.Background(), filter)
	//assert.NoError(t, err)
	//assert.Equal(t, len(logs), 1)
	//
	//address, amount, err := decodeRepayVAI(logs[0])
	//assert.NoError(t, err)
	//assert.Equal(t, address.Hex(), "0x1EE399b35337505DAFCE451a3311ed23Ee023885")
	//assert.Equal(t, amount.String(), "16000000000000000000")
}

func TestVUSDTMintEvent_46359438(t *testing.T) {
	c, err := ethclient.Dial(Url)
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

	from, to, amount, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", from.Hex())
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", to.Hex())
	assert.Equal(t, "100000000000000000000", amount.String())
}

func TestVUSDTRedeem_46372646(t *testing.T) {
	c, err := ethclient.Dial(Url)
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

	from, to, amount, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", from.Hex())
	assert.Equal(t, "0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591", to.Hex())
	assert.Equal(t, "20000000000000000000", amount.String())
}

func TestVUSDTTransfer_46486375(t *testing.T) {
	c, err := ethclient.Dial(Url)
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

	from, to, amount, err := decodeVTokenTransfer(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", from.Hex())
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", to.Hex())
	assert.Equal(t, "100000000000000000000", amount.String())
}

/*


func TestMintEvent_46359438(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46359438)

	vUSDT := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce), []common.Address{vUSDT})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	address, amount, err := decodeMint(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x658a6c7962e64132d2487EB2bc431d8Bc285882F", address.Hex())
	assert.Equal(t, "100000000000000000000", amount.String())
}

func TestRedeemEvent_46372646(t *testing.T) {
	c, err := ethclient.Dial(Url)
	assert.NoError(t, err)

	_, err = c.BlockNumber(context.Background())
	assert.NoError(t, err)

	blockHeight := big.NewInt(46372646)

	vUSDT := common.HexToAddress("0xEAB5387c7d9280eC791cdF46921cF4b3C62fd591")
	filter := buildQueryWithoutHeight(common.HexToAddress(Comptroller), common.HexToAddress(VaiController), common.HexToAddress(Oralce), []common.Address{vUSDT})
	filter.FromBlock = blockHeight
	filter.ToBlock = blockHeight

	logs, err := c.FilterLogs(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 1)

	address, amount, err := decodeRedeem(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, "0x1EE399b35337505DAFCE451a3311ed23Ee023885", address.Hex())
	assert.Equal(t, "20000000000000000000", amount.String())
}
*/

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
