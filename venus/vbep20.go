// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Vbep20MetaData contains all meta data concerning the Vbep20 contract.
var Vbep20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"MintBehalf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"RedeemFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Vbep20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Vbep20MetaData.ABI instead.
var Vbep20ABI = Vbep20MetaData.ABI

// Vbep20 is an auto generated Go binding around an Ethereum contract.
type Vbep20 struct {
	Vbep20Caller     // Read-only binding to the contract
	Vbep20Transactor // Write-only binding to the contract
	Vbep20Filterer   // Log filterer for contract events
}

// Vbep20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Vbep20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vbep20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Vbep20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vbep20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Vbep20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vbep20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Vbep20Session struct {
	Contract     *Vbep20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Vbep20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Vbep20CallerSession struct {
	Contract *Vbep20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Vbep20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Vbep20TransactorSession struct {
	Contract     *Vbep20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Vbep20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Vbep20Raw struct {
	Contract *Vbep20 // Generic contract binding to access the raw methods on
}

// Vbep20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Vbep20CallerRaw struct {
	Contract *Vbep20Caller // Generic read-only contract binding to access the raw methods on
}

// Vbep20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Vbep20TransactorRaw struct {
	Contract *Vbep20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVbep20 creates a new instance of Vbep20, bound to a specific deployed contract.
func NewVbep20(address common.Address, backend bind.ContractBackend) (*Vbep20, error) {
	contract, err := bindVbep20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vbep20{Vbep20Caller: Vbep20Caller{contract: contract}, Vbep20Transactor: Vbep20Transactor{contract: contract}, Vbep20Filterer: Vbep20Filterer{contract: contract}}, nil
}

// NewVbep20Caller creates a new read-only instance of Vbep20, bound to a specific deployed contract.
func NewVbep20Caller(address common.Address, caller bind.ContractCaller) (*Vbep20Caller, error) {
	contract, err := bindVbep20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Vbep20Caller{contract: contract}, nil
}

// NewVbep20Transactor creates a new write-only instance of Vbep20, bound to a specific deployed contract.
func NewVbep20Transactor(address common.Address, transactor bind.ContractTransactor) (*Vbep20Transactor, error) {
	contract, err := bindVbep20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Vbep20Transactor{contract: contract}, nil
}

// NewVbep20Filterer creates a new log filterer instance of Vbep20, bound to a specific deployed contract.
func NewVbep20Filterer(address common.Address, filterer bind.ContractFilterer) (*Vbep20Filterer, error) {
	contract, err := bindVbep20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Vbep20Filterer{contract: contract}, nil
}

// bindVbep20 binds a generic wrapper to an already deployed contract.
func bindVbep20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Vbep20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vbep20 *Vbep20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vbep20.Contract.Vbep20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vbep20 *Vbep20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.Contract.Vbep20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vbep20 *Vbep20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vbep20.Contract.Vbep20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vbep20 *Vbep20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vbep20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vbep20 *Vbep20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vbep20 *Vbep20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vbep20.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbep20 *Vbep20Caller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbep20 *Vbep20Session) AccrualBlockNumber() (*big.Int, error) {
	return _Vbep20.Contract.AccrualBlockNumber(&_Vbep20.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Vbep20.Contract.AccrualBlockNumber(&_Vbep20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbep20 *Vbep20Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbep20 *Vbep20Session) Admin() (common.Address, error) {
	return _Vbep20.Contract.Admin(&_Vbep20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbep20 *Vbep20CallerSession) Admin() (common.Address, error) {
	return _Vbep20.Contract.Admin(&_Vbep20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbep20 *Vbep20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbep20 *Vbep20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vbep20.Contract.Allowance(&_Vbep20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vbep20.Contract.Allowance(&_Vbep20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbep20 *Vbep20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbep20 *Vbep20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vbep20.Contract.BalanceOf(&_Vbep20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vbep20.Contract.BalanceOf(&_Vbep20.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbep20 *Vbep20Caller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbep20 *Vbep20Session) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Vbep20.Contract.BorrowBalanceStored(&_Vbep20.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Vbep20.Contract.BorrowBalanceStored(&_Vbep20.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbep20 *Vbep20Caller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbep20 *Vbep20Session) BorrowIndex() (*big.Int, error) {
	return _Vbep20.Contract.BorrowIndex(&_Vbep20.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) BorrowIndex() (*big.Int, error) {
	return _Vbep20.Contract.BorrowIndex(&_Vbep20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20Caller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20Session) BorrowRatePerBlock() (*big.Int, error) {
	return _Vbep20.Contract.BorrowRatePerBlock(&_Vbep20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Vbep20.Contract.BorrowRatePerBlock(&_Vbep20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbep20 *Vbep20Caller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbep20 *Vbep20Session) Comptroller() (common.Address, error) {
	return _Vbep20.Contract.Comptroller(&_Vbep20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbep20 *Vbep20CallerSession) Comptroller() (common.Address, error) {
	return _Vbep20.Contract.Comptroller(&_Vbep20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbep20 *Vbep20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbep20 *Vbep20Session) Decimals() (uint8, error) {
	return _Vbep20.Contract.Decimals(&_Vbep20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbep20 *Vbep20CallerSession) Decimals() (uint8, error) {
	return _Vbep20.Contract.Decimals(&_Vbep20.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Vbep20 *Vbep20Caller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Vbep20 *Vbep20Session) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Vbep20.Contract.DelegateToViewImplementation(&_Vbep20.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Vbep20 *Vbep20CallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Vbep20.Contract.DelegateToViewImplementation(&_Vbep20.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbep20 *Vbep20Caller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbep20 *Vbep20Session) ExchangeRateStored() (*big.Int, error) {
	return _Vbep20.Contract.ExchangeRateStored(&_Vbep20.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Vbep20.Contract.ExchangeRateStored(&_Vbep20.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbep20 *Vbep20Caller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbep20 *Vbep20Session) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Vbep20.Contract.GetAccountSnapshot(&_Vbep20.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbep20 *Vbep20CallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Vbep20.Contract.GetAccountSnapshot(&_Vbep20.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbep20 *Vbep20Caller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbep20 *Vbep20Session) GetCash() (*big.Int, error) {
	return _Vbep20.Contract.GetCash(&_Vbep20.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) GetCash() (*big.Int, error) {
	return _Vbep20.Contract.GetCash(&_Vbep20.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Vbep20 *Vbep20Caller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Vbep20 *Vbep20Session) Implementation() (common.Address, error) {
	return _Vbep20.Contract.Implementation(&_Vbep20.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Vbep20 *Vbep20CallerSession) Implementation() (common.Address, error) {
	return _Vbep20.Contract.Implementation(&_Vbep20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbep20 *Vbep20Caller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbep20 *Vbep20Session) InterestRateModel() (common.Address, error) {
	return _Vbep20.Contract.InterestRateModel(&_Vbep20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbep20 *Vbep20CallerSession) InterestRateModel() (common.Address, error) {
	return _Vbep20.Contract.InterestRateModel(&_Vbep20.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbep20 *Vbep20Caller) IsVToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "isVToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbep20 *Vbep20Session) IsVToken() (bool, error) {
	return _Vbep20.Contract.IsVToken(&_Vbep20.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbep20 *Vbep20CallerSession) IsVToken() (bool, error) {
	return _Vbep20.Contract.IsVToken(&_Vbep20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbep20 *Vbep20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbep20 *Vbep20Session) Name() (string, error) {
	return _Vbep20.Contract.Name(&_Vbep20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbep20 *Vbep20CallerSession) Name() (string, error) {
	return _Vbep20.Contract.Name(&_Vbep20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbep20 *Vbep20Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbep20 *Vbep20Session) PendingAdmin() (common.Address, error) {
	return _Vbep20.Contract.PendingAdmin(&_Vbep20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbep20 *Vbep20CallerSession) PendingAdmin() (common.Address, error) {
	return _Vbep20.Contract.PendingAdmin(&_Vbep20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbep20 *Vbep20Caller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbep20 *Vbep20Session) ReserveFactorMantissa() (*big.Int, error) {
	return _Vbep20.Contract.ReserveFactorMantissa(&_Vbep20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Vbep20.Contract.ReserveFactorMantissa(&_Vbep20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20Caller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20Session) SupplyRatePerBlock() (*big.Int, error) {
	return _Vbep20.Contract.SupplyRatePerBlock(&_Vbep20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Vbep20.Contract.SupplyRatePerBlock(&_Vbep20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbep20 *Vbep20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbep20 *Vbep20Session) Symbol() (string, error) {
	return _Vbep20.Contract.Symbol(&_Vbep20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbep20 *Vbep20CallerSession) Symbol() (string, error) {
	return _Vbep20.Contract.Symbol(&_Vbep20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbep20 *Vbep20Caller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbep20 *Vbep20Session) TotalBorrows() (*big.Int, error) {
	return _Vbep20.Contract.TotalBorrows(&_Vbep20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) TotalBorrows() (*big.Int, error) {
	return _Vbep20.Contract.TotalBorrows(&_Vbep20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbep20 *Vbep20Caller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbep20 *Vbep20Session) TotalReserves() (*big.Int, error) {
	return _Vbep20.Contract.TotalReserves(&_Vbep20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) TotalReserves() (*big.Int, error) {
	return _Vbep20.Contract.TotalReserves(&_Vbep20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbep20 *Vbep20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbep20 *Vbep20Session) TotalSupply() (*big.Int, error) {
	return _Vbep20.Contract.TotalSupply(&_Vbep20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbep20 *Vbep20CallerSession) TotalSupply() (*big.Int, error) {
	return _Vbep20.Contract.TotalSupply(&_Vbep20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Vbep20 *Vbep20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbep20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Vbep20 *Vbep20Session) Underlying() (common.Address, error) {
	return _Vbep20.Contract.Underlying(&_Vbep20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Vbep20 *Vbep20CallerSession) Underlying() (common.Address, error) {
	return _Vbep20.Contract.Underlying(&_Vbep20.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbep20 *Vbep20Transactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbep20 *Vbep20Session) AcceptAdmin() (*types.Transaction, error) {
	return _Vbep20.Contract.AcceptAdmin(&_Vbep20.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Vbep20.Contract.AcceptAdmin(&_Vbep20.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.AddReserves(&_Vbep20.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.AddReserves(&_Vbep20.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.ReduceReserves(&_Vbep20.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.ReduceReserves(&_Vbep20.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbep20 *Vbep20Transactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbep20 *Vbep20Session) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetComptroller(&_Vbep20.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetComptroller(&_Vbep20.TransactOpts, newComptroller)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Vbep20 *Vbep20Transactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Vbep20 *Vbep20Session) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.SetImplementation(&_Vbep20.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Vbep20 *Vbep20TransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.SetImplementation(&_Vbep20.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbep20 *Vbep20Transactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbep20 *Vbep20Session) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetInterestRateModel(&_Vbep20.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetInterestRateModel(&_Vbep20.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbep20 *Vbep20Transactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbep20 *Vbep20Session) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetPendingAdmin(&_Vbep20.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.SetPendingAdmin(&_Vbep20.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbep20 *Vbep20Transactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbep20 *Vbep20Session) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.SetReserveFactor(&_Vbep20.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.SetReserveFactor(&_Vbep20.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbep20 *Vbep20Transactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbep20 *Vbep20Session) AccrueInterest() (*types.Transaction, error) {
	return _Vbep20.Contract.AccrueInterest(&_Vbep20.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Vbep20.Contract.AccrueInterest(&_Vbep20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Approve(&_Vbep20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Approve(&_Vbep20.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbep20 *Vbep20Transactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbep20 *Vbep20Session) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.BalanceOfUnderlying(&_Vbep20.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.BalanceOfUnderlying(&_Vbep20.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Borrow(&_Vbep20.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Borrow(&_Vbep20.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbep20 *Vbep20Transactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbep20 *Vbep20Session) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.BorrowBalanceCurrent(&_Vbep20.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.BorrowBalanceCurrent(&_Vbep20.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Vbep20 *Vbep20Transactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Vbep20 *Vbep20Session) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.DelegateToImplementation(&_Vbep20.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Vbep20 *Vbep20TransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.DelegateToImplementation(&_Vbep20.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbep20 *Vbep20Transactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbep20 *Vbep20Session) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Vbep20.Contract.ExchangeRateCurrent(&_Vbep20.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Vbep20.Contract.ExchangeRateCurrent(&_Vbep20.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Vbep20 *Vbep20Transactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Vbep20 *Vbep20Session) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.LiquidateBorrow(&_Vbep20.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbep20.Contract.LiquidateBorrow(&_Vbep20.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Mint(&_Vbep20.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Mint(&_Vbep20.TransactOpts, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) MintBehalf(opts *bind.TransactOpts, receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "mintBehalf", receiver, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) MintBehalf(receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.MintBehalf(&_Vbep20.TransactOpts, receiver, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) MintBehalf(receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.MintBehalf(&_Vbep20.TransactOpts, receiver, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbep20 *Vbep20Transactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbep20 *Vbep20Session) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Redeem(&_Vbep20.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Redeem(&_Vbep20.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RedeemUnderlying(&_Vbep20.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RedeemUnderlying(&_Vbep20.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RepayBorrow(&_Vbep20.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RepayBorrow(&_Vbep20.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20Transactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20Session) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RepayBorrowBehalf(&_Vbep20.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.RepayBorrowBehalf(&_Vbep20.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbep20 *Vbep20Transactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbep20 *Vbep20Session) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Seize(&_Vbep20.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Seize(&_Vbep20.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbep20 *Vbep20Transactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbep20 *Vbep20Session) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Vbep20.Contract.TotalBorrowsCurrent(&_Vbep20.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbep20 *Vbep20TransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Vbep20.Contract.TotalBorrowsCurrent(&_Vbep20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Session) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Transfer(&_Vbep20.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20TransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.Transfer(&_Vbep20.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20Session) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.TransferFrom(&_Vbep20.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbep20 *Vbep20TransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbep20.Contract.TransferFrom(&_Vbep20.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbep20 *Vbep20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Vbep20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbep20 *Vbep20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.Fallback(&_Vbep20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbep20 *Vbep20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vbep20.Contract.Fallback(&_Vbep20.TransactOpts, calldata)
}

// Vbep20AccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Vbep20 contract.
type Vbep20AccrueInterestIterator struct {
	Event *Vbep20AccrueInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20AccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20AccrueInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20AccrueInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20AccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20AccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20AccrueInterest represents a AccrueInterest event raised by the Vbep20 contract.
type Vbep20AccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) FilterAccrueInterest(opts *bind.FilterOpts) (*Vbep20AccrueInterestIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &Vbep20AccrueInterestIterator{contract: _Vbep20.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *Vbep20AccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20AccrueInterest)
				if err := _Vbep20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) ParseAccrueInterest(log types.Log) (*Vbep20AccrueInterest, error) {
	event := new(Vbep20AccrueInterest)
	if err := _Vbep20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vbep20 contract.
type Vbep20ApprovalIterator struct {
	Event *Vbep20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Approval represents a Approval event raised by the Vbep20 contract.
type Vbep20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbep20 *Vbep20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Vbep20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Vbep20ApprovalIterator{contract: _Vbep20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbep20 *Vbep20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Vbep20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Approval)
				if err := _Vbep20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbep20 *Vbep20Filterer) ParseApproval(log types.Log) (*Vbep20Approval, error) {
	event := new(Vbep20Approval)
	if err := _Vbep20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Vbep20 contract.
type Vbep20BorrowIterator struct {
	Event *Vbep20Borrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Borrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Borrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Borrow represents a Borrow event raised by the Vbep20 contract.
type Vbep20Borrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) FilterBorrow(opts *bind.FilterOpts) (*Vbep20BorrowIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &Vbep20BorrowIterator{contract: _Vbep20.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Vbep20Borrow) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Borrow)
				if err := _Vbep20.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) ParseBorrow(log types.Log) (*Vbep20Borrow, error) {
	event := new(Vbep20Borrow)
	if err := _Vbep20.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Vbep20 contract.
type Vbep20FailureIterator struct {
	Event *Vbep20Failure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Failure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Failure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Failure represents a Failure event raised by the Vbep20 contract.
type Vbep20Failure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Vbep20 *Vbep20Filterer) FilterFailure(opts *bind.FilterOpts) (*Vbep20FailureIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &Vbep20FailureIterator{contract: _Vbep20.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Vbep20 *Vbep20Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *Vbep20Failure) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Failure)
				if err := _Vbep20.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Vbep20 *Vbep20Filterer) ParseFailure(log types.Log) (*Vbep20Failure, error) {
	event := new(Vbep20Failure)
	if err := _Vbep20.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20LiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Vbep20 contract.
type Vbep20LiquidateBorrowIterator struct {
	Event *Vbep20LiquidateBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20LiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20LiquidateBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20LiquidateBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20LiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20LiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20LiquidateBorrow represents a LiquidateBorrow event raised by the Vbep20 contract.
type Vbep20LiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	VTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbep20 *Vbep20Filterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*Vbep20LiquidateBorrowIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &Vbep20LiquidateBorrowIterator{contract: _Vbep20.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbep20 *Vbep20Filterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *Vbep20LiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20LiquidateBorrow)
				if err := _Vbep20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbep20 *Vbep20Filterer) ParseLiquidateBorrow(log types.Log) (*Vbep20LiquidateBorrow, error) {
	event := new(Vbep20LiquidateBorrow)
	if err := _Vbep20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Vbep20 contract.
type Vbep20MintIterator struct {
	Event *Vbep20Mint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Mint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Mint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Mint represents a Mint event raised by the Vbep20 contract.
type Vbep20Mint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) FilterMint(opts *bind.FilterOpts) (*Vbep20MintIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Vbep20MintIterator{contract: _Vbep20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Vbep20Mint) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Mint)
				if err := _Vbep20.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) ParseMint(log types.Log) (*Vbep20Mint, error) {
	event := new(Vbep20Mint)
	if err := _Vbep20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20MintBehalfIterator is returned from FilterMintBehalf and is used to iterate over the raw logs and unpacked data for MintBehalf events raised by the Vbep20 contract.
type Vbep20MintBehalfIterator struct {
	Event *Vbep20MintBehalf // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20MintBehalfIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20MintBehalf)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20MintBehalf)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20MintBehalfIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20MintBehalfIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20MintBehalf represents a MintBehalf event raised by the Vbep20 contract.
type Vbep20MintBehalf struct {
	Payer      common.Address
	Receiver   common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintBehalf is a free log retrieval operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) FilterMintBehalf(opts *bind.FilterOpts) (*Vbep20MintBehalfIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return &Vbep20MintBehalfIterator{contract: _Vbep20.contract, event: "MintBehalf", logs: logs, sub: sub}, nil
}

// WatchMintBehalf is a free log subscription operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) WatchMintBehalf(opts *bind.WatchOpts, sink chan<- *Vbep20MintBehalf) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20MintBehalf)
				if err := _Vbep20.contract.UnpackLog(event, "MintBehalf", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintBehalf is a log parse operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbep20 *Vbep20Filterer) ParseMintBehalf(log types.Log) (*Vbep20MintBehalf, error) {
	event := new(Vbep20MintBehalf)
	if err := _Vbep20.contract.UnpackLog(event, "MintBehalf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Vbep20 contract.
type Vbep20NewAdminIterator struct {
	Event *Vbep20NewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewAdmin represents a NewAdmin event raised by the Vbep20 contract.
type Vbep20NewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbep20 *Vbep20Filterer) FilterNewAdmin(opts *bind.FilterOpts) (*Vbep20NewAdminIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewAdminIterator{contract: _Vbep20.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbep20 *Vbep20Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Vbep20NewAdmin) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewAdmin)
				if err := _Vbep20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbep20 *Vbep20Filterer) ParseNewAdmin(log types.Log) (*Vbep20NewAdmin, error) {
	event := new(Vbep20NewAdmin)
	if err := _Vbep20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Vbep20 contract.
type Vbep20NewComptrollerIterator struct {
	Event *Vbep20NewComptroller // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewComptroller)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewComptroller)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewComptroller represents a NewComptroller event raised by the Vbep20 contract.
type Vbep20NewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbep20 *Vbep20Filterer) FilterNewComptroller(opts *bind.FilterOpts) (*Vbep20NewComptrollerIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewComptrollerIterator{contract: _Vbep20.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbep20 *Vbep20Filterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *Vbep20NewComptroller) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewComptroller)
				if err := _Vbep20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbep20 *Vbep20Filterer) ParseNewComptroller(log types.Log) (*Vbep20NewComptroller, error) {
	event := new(Vbep20NewComptroller)
	if err := _Vbep20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Vbep20 contract.
type Vbep20NewImplementationIterator struct {
	Event *Vbep20NewImplementation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewImplementation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewImplementation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewImplementation represents a NewImplementation event raised by the Vbep20 contract.
type Vbep20NewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Vbep20 *Vbep20Filterer) FilterNewImplementation(opts *bind.FilterOpts) (*Vbep20NewImplementationIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewImplementationIterator{contract: _Vbep20.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Vbep20 *Vbep20Filterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *Vbep20NewImplementation) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewImplementation)
				if err := _Vbep20.contract.UnpackLog(event, "NewImplementation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Vbep20 *Vbep20Filterer) ParseNewImplementation(log types.Log) (*Vbep20NewImplementation, error) {
	event := new(Vbep20NewImplementation)
	if err := _Vbep20.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Vbep20 contract.
type Vbep20NewMarketInterestRateModelIterator struct {
	Event *Vbep20NewMarketInterestRateModel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewMarketInterestRateModel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewMarketInterestRateModel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Vbep20 contract.
type Vbep20NewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbep20 *Vbep20Filterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*Vbep20NewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewMarketInterestRateModelIterator{contract: _Vbep20.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbep20 *Vbep20Filterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *Vbep20NewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewMarketInterestRateModel)
				if err := _Vbep20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbep20 *Vbep20Filterer) ParseNewMarketInterestRateModel(log types.Log) (*Vbep20NewMarketInterestRateModel, error) {
	event := new(Vbep20NewMarketInterestRateModel)
	if err := _Vbep20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Vbep20 contract.
type Vbep20NewPendingAdminIterator struct {
	Event *Vbep20NewPendingAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewPendingAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewPendingAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewPendingAdmin represents a NewPendingAdmin event raised by the Vbep20 contract.
type Vbep20NewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbep20 *Vbep20Filterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*Vbep20NewPendingAdminIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewPendingAdminIterator{contract: _Vbep20.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbep20 *Vbep20Filterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *Vbep20NewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewPendingAdmin)
				if err := _Vbep20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbep20 *Vbep20Filterer) ParseNewPendingAdmin(log types.Log) (*Vbep20NewPendingAdmin, error) {
	event := new(Vbep20NewPendingAdmin)
	if err := _Vbep20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20NewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Vbep20 contract.
type Vbep20NewReserveFactorIterator struct {
	Event *Vbep20NewReserveFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20NewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20NewReserveFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20NewReserveFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20NewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20NewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20NewReserveFactor represents a NewReserveFactor event raised by the Vbep20 contract.
type Vbep20NewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbep20 *Vbep20Filterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*Vbep20NewReserveFactorIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &Vbep20NewReserveFactorIterator{contract: _Vbep20.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbep20 *Vbep20Filterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *Vbep20NewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20NewReserveFactor)
				if err := _Vbep20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbep20 *Vbep20Filterer) ParseNewReserveFactor(log types.Log) (*Vbep20NewReserveFactor, error) {
	event := new(Vbep20NewReserveFactor)
	if err := _Vbep20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Vbep20 contract.
type Vbep20RedeemIterator struct {
	Event *Vbep20Redeem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Redeem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Redeem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Redeem represents a Redeem event raised by the Vbep20 contract.
type Vbep20Redeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) FilterRedeem(opts *bind.FilterOpts) (*Vbep20RedeemIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Vbep20RedeemIterator{contract: _Vbep20.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Vbep20Redeem) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Redeem)
				if err := _Vbep20.contract.UnpackLog(event, "Redeem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) ParseRedeem(log types.Log) (*Vbep20Redeem, error) {
	event := new(Vbep20Redeem)
	if err := _Vbep20.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20RedeemFeeIterator is returned from FilterRedeemFee and is used to iterate over the raw logs and unpacked data for RedeemFee events raised by the Vbep20 contract.
type Vbep20RedeemFeeIterator struct {
	Event *Vbep20RedeemFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20RedeemFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20RedeemFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20RedeemFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20RedeemFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20RedeemFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20RedeemFee represents a RedeemFee event raised by the Vbep20 contract.
type Vbep20RedeemFee struct {
	Redeemer     common.Address
	FeeAmount    *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemFee is a free log retrieval operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) FilterRedeemFee(opts *bind.FilterOpts) (*Vbep20RedeemFeeIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return &Vbep20RedeemFeeIterator{contract: _Vbep20.contract, event: "RedeemFee", logs: logs, sub: sub}, nil
}

// WatchRedeemFee is a free log subscription operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) WatchRedeemFee(opts *bind.WatchOpts, sink chan<- *Vbep20RedeemFee) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20RedeemFee)
				if err := _Vbep20.contract.UnpackLog(event, "RedeemFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemFee is a log parse operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbep20 *Vbep20Filterer) ParseRedeemFee(log types.Log) (*Vbep20RedeemFee, error) {
	event := new(Vbep20RedeemFee)
	if err := _Vbep20.contract.UnpackLog(event, "RedeemFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20RepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Vbep20 contract.
type Vbep20RepayBorrowIterator struct {
	Event *Vbep20RepayBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20RepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20RepayBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20RepayBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20RepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20RepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20RepayBorrow represents a RepayBorrow event raised by the Vbep20 contract.
type Vbep20RepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) FilterRepayBorrow(opts *bind.FilterOpts) (*Vbep20RepayBorrowIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &Vbep20RepayBorrowIterator{contract: _Vbep20.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *Vbep20RepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20RepayBorrow)
				if err := _Vbep20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbep20 *Vbep20Filterer) ParseRepayBorrow(log types.Log) (*Vbep20RepayBorrow, error) {
	event := new(Vbep20RepayBorrow)
	if err := _Vbep20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20ReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Vbep20 contract.
type Vbep20ReservesAddedIterator struct {
	Event *Vbep20ReservesAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20ReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20ReservesAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20ReservesAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20ReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20ReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20ReservesAdded represents a ReservesAdded event raised by the Vbep20 contract.
type Vbep20ReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) FilterReservesAdded(opts *bind.FilterOpts) (*Vbep20ReservesAddedIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &Vbep20ReservesAddedIterator{contract: _Vbep20.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *Vbep20ReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20ReservesAdded)
				if err := _Vbep20.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) ParseReservesAdded(log types.Log) (*Vbep20ReservesAdded, error) {
	event := new(Vbep20ReservesAdded)
	if err := _Vbep20.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20ReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Vbep20 contract.
type Vbep20ReservesReducedIterator struct {
	Event *Vbep20ReservesReduced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20ReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20ReservesReduced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20ReservesReduced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20ReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20ReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20ReservesReduced represents a ReservesReduced event raised by the Vbep20 contract.
type Vbep20ReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) FilterReservesReduced(opts *bind.FilterOpts) (*Vbep20ReservesReducedIterator, error) {

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &Vbep20ReservesReducedIterator{contract: _Vbep20.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *Vbep20ReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20ReservesReduced)
				if err := _Vbep20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbep20 *Vbep20Filterer) ParseReservesReduced(log types.Log) (*Vbep20ReservesReduced, error) {
	event := new(Vbep20ReservesReduced)
	if err := _Vbep20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vbep20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vbep20 contract.
type Vbep20TransferIterator struct {
	Event *Vbep20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vbep20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vbep20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Vbep20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Vbep20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vbep20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vbep20Transfer represents a Transfer event raised by the Vbep20 contract.
type Vbep20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbep20 *Vbep20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Vbep20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vbep20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Vbep20TransferIterator{contract: _Vbep20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbep20 *Vbep20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Vbep20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vbep20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vbep20Transfer)
				if err := _Vbep20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbep20 *Vbep20Filterer) ParseTransfer(log types.Log) (*Vbep20Transfer, error) {
	event := new(Vbep20Transfer)
	if err := _Vbep20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
