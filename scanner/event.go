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

//const Mint = "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"
//const Redeem = "0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929"
//const MintBehalf = "0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345"
//const Borrow = "0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80"
//const RepayBorrow = "0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1"
//const LiquidateBorrow = "0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52"

// topics in PriceOralce
//const PriceUpdated = ""

//const VaiController = "0x96ae4986D9ff19992dA84B5DBA9790cAE7246b80"
//const Comptroller = "0xB4Abb34e08094B1915Ac3f7882aed81d0104b121"
//const Oralce = "0x6B392885f26b718C149f759B591094a06787A289"

// build QueryFilter for comptroller, vaiController, oracle
func buildQueryWithoutHeight(comptroller, vaiController, oracle common.Address) ethereum.FilterQuery {
	addresses := []common.Address{comptroller, vaiController, oracle}

	var _topics []common.Hash
	_topics = append(_topics,
		common.HexToHash(MarketListed), common.HexToHash(NewCloseFactor), common.HexToHash(NewCollateralFactor), common.HexToHash(MarketEntered), common.HexToHash(MarketExited),
		common.HexToHash(MintVAI), common.HexToHash(RepayVAI), common.HexToHash(LiquidateVAI),
		common.HexToHash(PriceUpdated),
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

func decodeMarketListed(log types.Log) (*NewMarket, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketListed || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}
	market := common.BytesToAddress(data[0:32])
	return &NewMarket{
		Market:        market,
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeNewCloseFactor(log types.Log) (*CloseFactorChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != NewCloseFactor || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}
	closeFactor := big.NewInt(0).SetBytes(data[32:64])
	return &CloseFactorChanged{
		CloseFactor:   decimal.NewFromBigInt(closeFactor, 0),
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeNewCollateralFactor(log types.Log) (*CollateralFactorChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != NewCollateralFactor || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	collateralFactor := big.NewInt(0).SetBytes(data[64:96])
	return &CollateralFactorChanged{
		Market:           market,
		CollateralFactor: decimal.NewFromBigInt(collateralFactor, 0),
		UpdatedHeight:    log.BlockNumber,
	}, nil
}

func decodeMarketEntered(log types.Log) (*EnterMarket, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketEntered || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	account := common.BytesToAddress(data[32:64])
	return &EnterMarket{
		Market:        market,
		Account:       account,
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeMarketExited(log types.Log) (*ExitMarket, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MarketExited || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}

	market := common.BytesToAddress(data[0:32])
	account := common.BytesToAddress(data[32:64])
	return &ExitMarket{
		Market:        market,
		Account:       account,
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeMintVAI(log types.Log) (*RepayVaiAmountChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != MintVAI || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}

	account := common.BytesToAddress(data[0:32])
	amount := big.NewInt(0).SetBytes(data[32:64])
	return &RepayVaiAmountChanged{
		Account:       account,
		Amount:        decimal.NewFromBigInt(amount, 0),
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeRepayVAI(log types.Log) (*RepayVaiAmountChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != RepayVAI || len(topics) != 1 {
		return nil, fmt.Errorf("invalid topic")
	}

	account := common.BytesToAddress(data[32:64])
	amount := big.NewInt(0).SetBytes(data[64:96])
	amount = big.NewInt(0).Neg(amount)
	return &RepayVaiAmountChanged{
		Account:       account,
		Amount:        decimal.NewFromBigInt(amount, 0),
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodeLiquidateVAI(log types.Log) (*VTokenAmountChanged, error) {
	panic("not implemented, in liquidation, borrower's vToken changed is covered by vToken's transfer event")
	return nil, nil
}

func decodeVTokenTransfer(log types.Log) (*VTokenAmountChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != Transfer || len(topics) != 3 {
		return nil, fmt.Errorf("invalid topic")
	}

	market := log.Address
	from := common.HexToAddress(topics[1].Hex())
	to := common.HexToAddress(topics[2].Hex())
	amount := big.NewInt(0).SetBytes(data[0:32])

	return &VTokenAmountChanged{
		Market:        market,
		From:          from,
		To:            to,
		Amount:        decimal.NewFromBigInt(amount, 0),
		UpdatedHeight: log.BlockNumber,
	}, nil
}

func decodePriceUpdate(log types.Log) (*PriceChanged, error) {
	panic("not implemented")
	return nil, nil
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
	if log.Removed == true {
		logger.Printf("log was reverted due to a chain reorganisation: %v\n", log)
		return nil
	}

	switch log.Topics[0].Hex() {
	case MarketListed:
		market, err := decodeMarketListed(log)
		if err != nil {
			return err
		}
		s.newMarketCh <- market

	case NewCloseFactor:
		closeFactor, err := decodeNewCloseFactor(log)
		if err != nil {
			return err
		}
		s.closeFactorChangedCh <- closeFactor

	case NewCollateralFactor:
		collateralFactor, err := decodeNewCollateralFactor(log)
		if err != nil {
			return err
		}
		s.collateralFactorChangedCh <- collateralFactor

	case MarketEntered:
		enter, err := decodeMarketEntered(log)
		if err != nil {
			return err
		}
		s.enterMarketCh <- enter

	case MarketExited:
		exit, err := decodeMarketExited(log)
		if err != nil {
			return err
		}
		s.exitMarketCh <- exit

	case MintVAI:
		change, err := decodeMintVAI(log)
		if err != nil {
			return err
		}
		s.repayVaiAmountChangedCh <- change

	case RepayVAI:
		change, err := decodeRepayVAI(log)
		if err != nil {
			return err
		}
		s.repayVaiAmountChangedCh <- change

	case Transfer:
		change, err := decodeVTokenTransfer(log)
		if err != nil {
			return err
		}
		s.vTokenAmountChangedCh <- change

	case PriceUpdated:
		change, err := decodePriceUpdate(log)
		if err != nil {
			return err
		}
		s.priceChangedCh <- change

	default:
		return fmt.Errorf("invalid topic")
	}

	return nil
}

//func (s *Scanner) DecodeLog(log types.Log) error {
//	return nil
//}
