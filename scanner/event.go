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

// build QueryFilter for comptroller, vaiController, oracle
func buildQueryWithoutHeight(comptroller, vaiController common.Address, feeders []common.Address) ethereum.FilterQuery {
	addresses := []common.Address{comptroller, vaiController}
	addresses = append(addresses, feeders...)

	var _topics []common.Hash
	_topics = append(_topics,
		common.HexToHash(MarketListed), common.HexToHash(NewCloseFactor), common.HexToHash(NewCollateralFactor), common.HexToHash(MarketEntered), common.HexToHash(MarketExited),
		common.HexToHash(MintVAI), common.HexToHash(RepayVAI), //vaiController event
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

func decodePriceUpdate(feederMap map[common.Address]common.Address, log types.Log) (*PriceChanged, error) {
	topics := log.Topics
	data := log.Data

	if topics[0].Hex() != PriceUpdated || len(topics) != 2 {
		return nil, fmt.Errorf("invalid topic")
	}
	price := big.NewInt(0).SetBytes(data[0:32])
	market := feederMap[log.Address]
	return &PriceChanged{
		Market:        market,
		Price:         decimal.NewFromBigInt(price, 0),
		UpdatedHeight: log.BlockNumber,
	}, nil

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
		change, err := decodePriceUpdate(s.feederMap, log)
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
