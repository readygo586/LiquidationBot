package scanner

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"
)

var (
	vBNBAddress = common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36")
	wBNBAddress = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
)

// topics in VaiController
const MarketListed = "0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f"
const NewCloseFactor = "0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9"
const NewCollateralFactor = "0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5"
const MarketEntered = "0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5"
const MarketExited = "0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d"

// topics in VaiController
const MintVAI = "0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a"
const RepayVAI = "0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e"
const LiquidateVAI = "0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7"

// topics in vToken
const Transfer = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

//const Mint = "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"
//const Redeem = "0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929"
//const MintBehalf = "0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345"
//const Borrow = "0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80"
//const RepayBorrow = "0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1"
//const LiquidateBorrow = "0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52"

// topics in PriceOralce
const PriceUpdate = ""

const VaiController = "0x96ae4986D9ff19992dA84B5DBA9790cAE7246b80"
const Comptroller = "0xB4Abb34e08094B1915Ac3f7882aed81d0104b121"
const Oralce = "0x6B392885f26b718C149f759B591094a06787A289"

// build QueryFilter for comptroller, vaiController, oracle
func buildQueryWithoutHeight(comptroller, vaiController, oracle common.Address) ethereum.FilterQuery {
	addresses := []common.Address{comptroller, vaiController, oracle}

	var _topics []common.Hash
	_topics = append(_topics,
		common.HexToHash(MarketListed), common.HexToHash(NewCloseFactor), common.HexToHash(NewCollateralFactor), common.HexToHash(MarketEntered), common.HexToHash(MarketExited),
		common.HexToHash(MintVAI), common.HexToHash(RepayVAI), common.HexToHash(LiquidateVAI),
		common.HexToHash(PriceUpdate),
	)
	topics := [][]common.Hash{_topics}

	return ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topics,
	}
}

func buildVTokenQueryWithoutHeight(addresses []common.Address) ethereum.FilterQuery {
	var _topics []common.Hash
	_topics = append(_topics,
		common.HexToHash(Transfer),
	)
	topics := [][]common.Hash{_topics}
	return ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topics,
	}
}

func decodeMarketListed(log types.Log) (*common.Address, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketListed || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}
	market := common.BytesToAddress(data[0:32])
	return &market, nil
}

func decodeNewCloseFactor(log types.Log) (*big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != NewCloseFactor || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}
	closeFaFactor := big.NewInt(0).SetBytes(data[32:64])
	return closeFaFactor, nil
}

func decodeNewCollateralFactor(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != NewCollateralFactor || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	collateralFactor := big.NewInt(0).SetBytes(data[64:96])
	return &market, collateralFactor, nil
}

func decodeMarketEntered(log types.Log) (*common.Address, *common.Address, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketEntered || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	account := common.BytesToAddress(data[32:64])
	return &market, &account, nil
}

func decodeMarketExited(log types.Log) (*common.Address, *common.Address, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketExited || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	account := common.BytesToAddress(data[32:64])
	return &market, &account, nil

}

func decodeMintVAI(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MintVAI || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[0:32])
	amount := big.NewInt(0).SetBytes(data[32:64])
	return &address, amount, nil
}

func decodeRepayVAI(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != RepayVAI || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeLiquidateVAI(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != LiquidateVAI || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeVTokenTransfer(log types.Log) (*common.Address, *common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != Transfer || len(topics) != 3 {
		return nil, nil, nil, fmt.Errorf("invalid topic")
	}
	from := common.HexToAddress(topics[1].Hex())
	to := common.HexToAddress(topics[2].Hex())
	amount := big.NewInt(0).SetBytes(data[0:32])
	return &from, &to, amount, nil
}

/*
func decodeMint(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != Mint || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[0:32])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeMintBehalf(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MintBehalf || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[96:128])
	return &address, amount, nil
}

func decodeRedeem(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != Redeem || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[0:32])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeBorrow(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != Borrow || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[0:32])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeRepayBorrow(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != RepayBorrow || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}

func decodeLiquidateBorrow(log types.Log) (*common.Address, *big.Int, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != LiquidateBorrow || len(topics) != 1 {
		return nil, nil, fmt.Errorf("invalid topic")
	}

	address := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[64:96])
	return &address, amount, nil
}
*/

func (s *Scanner) DecodeLog(log types.Log) error {
	switch log.Topics[0].Hex() {
	case MarketListed:
		market, err := decodeMarketListed(log)
		if err != nil {
			return err
		}
		s.newMarketCh <- *market

	case NewCloseFactor:
		closeFactor, err := decodeNewCloseFactor(log)
		if err != nil {
			return err
		}
		s.closeFactor = decimal.NewFromBigInt(closeFactor, 0)
		s.closeFactorChangedCh <- s.closeFactor

	case NewCollateralFactor:
		market, _collateralFactor, err := decodeNewCollateralFactor(log)
		if err != nil {
			return err
		}
		collateralFactor := decimal.NewFromBigInt(_collateralFactor, 0)
		s.collateralFactorChangedCh <- &CollateralFactorChanged{
			Market:           *market,
			CollateralFactor: collateralFactor,
		}

	case MarketEntered:
		account, market, err := decodeMarketEntered(log)
		if err != nil {
			return err
		}
		s.enterMarketCh <- &EnterMarket{
			Market:  *market,
			Account: *account,
		}

	case MarketExited:
		account, market, err := decodeMarketExited(log)
		if err != nil {
			return err
		}
		s.exitMarketCh <- &ExitMarket{
			Market:  *market,
			Account: *account,
		}

	case MintVAI:
		account, amount, err := decodeMintVAI(log)
		if err != nil {
			return err
		}
		s.repayVaiAmountChangedCh <- &RepayVaiAmountChanged{
			Account: *account,
			Amount:  decimal.NewFromBigInt(amount, 0),
		}

	case RepayVAI:
		account, _amount, err := decodeRepayVAI(log)
		if err != nil {
			return err
		}
		amount := big.NewInt(0).Neg(_amount)
		s.repayVaiAmountChangedCh <- &RepayVaiAmountChanged{
			Account: *account,
			Amount:  decimal.NewFromBigInt(amount, 0),
		}

	case Transfer:
		from, to, amount, err := decodeVTokenTransfer(log)
		if err != nil {
			return err
		}
		s.vTokenAmountChangedCh <- &VTokenAmountChanged{
			Market: log.Address,
			From:   *from,
			To:     *to,
			Amount: decimal.NewFromBigInt(amount, 0),
		}

	case PriceUpdate:
		//TODO(keep)
		panic("TODO")
	default:
		return fmt.Errorf("invalid topic")
	}

	return nil
}

//func (s *Scanner) DecodeLog(log types.Log) error {
//	return nil
//}
