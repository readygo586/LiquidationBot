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

// VaiControllerMetaData contains all meta data concerning the VaiController contract.
var VaiControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateVAI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MintFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintVAIAmount\",\"type\":\"uint256\"}],\"name\":\"MintVAI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBaseRateMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBaseRateMantissa\",\"type\":\"uint256\"}],\"name\":\"NewVAIBaseRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFloatRateMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlatRateMantissa\",\"type\":\"uint256\"}],\"name\":\"NewVAIFloatRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMintCap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMintCap\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"NewVAIReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayVAIAmount\",\"type\":\"uint256\"}],\"name\":\"RepayVAI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueVAIInterest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRateMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"floatRateMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlocksPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintableVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVAIAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"getVAICalculateRepayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getVAIMinterInterestIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVAIRepayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVAIRepayRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVAIRepayRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVenusVAIInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintVAIAmount\",\"type\":\"uint256\"}],\"name\":\"mintVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pastVAIInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingVAIControllerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayVAIAmount\",\"type\":\"uint256\"}],\"name\":\"repayVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBaseRateMantissa\",\"type\":\"uint256\"}],\"name\":\"setBaseRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFloatRateMantissa\",\"type\":\"uint256\"}],\"name\":\"setFloatRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintCap\",\"type\":\"uint256\"}],\"name\":\"setMintCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vai_\",\"type\":\"address\"}],\"name\":\"setVAIAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiControllerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusVAIMinterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaiControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use VaiControllerMetaData.ABI instead.
var VaiControllerABI = VaiControllerMetaData.ABI

// VaiController is an auto generated Go binding around an Ethereum contract.
type VaiController struct {
	VaiControllerCaller     // Read-only binding to the contract
	VaiControllerTransactor // Write-only binding to the contract
	VaiControllerFilterer   // Log filterer for contract events
}

// VaiControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaiControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaiControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaiControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaiControllerSession struct {
	Contract     *VaiController    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaiControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaiControllerCallerSession struct {
	Contract *VaiControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VaiControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaiControllerTransactorSession struct {
	Contract     *VaiControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VaiControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaiControllerRaw struct {
	Contract *VaiController // Generic contract binding to access the raw methods on
}

// VaiControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaiControllerCallerRaw struct {
	Contract *VaiControllerCaller // Generic read-only contract binding to access the raw methods on
}

// VaiControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaiControllerTransactorRaw struct {
	Contract *VaiControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaiController creates a new instance of VaiController, bound to a specific deployed contract.
func NewVaiController(address common.Address, backend bind.ContractBackend) (*VaiController, error) {
	contract, err := bindVaiController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaiController{VaiControllerCaller: VaiControllerCaller{contract: contract}, VaiControllerTransactor: VaiControllerTransactor{contract: contract}, VaiControllerFilterer: VaiControllerFilterer{contract: contract}}, nil
}

// NewVaiControllerCaller creates a new read-only instance of VaiController, bound to a specific deployed contract.
func NewVaiControllerCaller(address common.Address, caller bind.ContractCaller) (*VaiControllerCaller, error) {
	contract, err := bindVaiController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaiControllerCaller{contract: contract}, nil
}

// NewVaiControllerTransactor creates a new write-only instance of VaiController, bound to a specific deployed contract.
func NewVaiControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaiControllerTransactor, error) {
	contract, err := bindVaiController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaiControllerTransactor{contract: contract}, nil
}

// NewVaiControllerFilterer creates a new log filterer instance of VaiController, bound to a specific deployed contract.
func NewVaiControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaiControllerFilterer, error) {
	contract, err := bindVaiController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaiControllerFilterer{contract: contract}, nil
}

// bindVaiController binds a generic wrapper to an already deployed contract.
func bindVaiController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaiControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaiController *VaiControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaiController.Contract.VaiControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaiController *VaiControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiController.Contract.VaiControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaiController *VaiControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaiController.Contract.VaiControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaiController *VaiControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaiController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaiController *VaiControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaiController *VaiControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaiController.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiController *VaiControllerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiController *VaiControllerSession) Admin() (common.Address, error) {
	return _VaiController.Contract.Admin(&_VaiController.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiController *VaiControllerCallerSession) Admin() (common.Address, error) {
	return _VaiController.Contract.Admin(&_VaiController.CallOpts)
}

// BaseRateMantissa is a free data retrieval call binding the contract method 0x3b72fbef.
//
// Solidity: function baseRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerCaller) BaseRateMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "baseRateMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRateMantissa is a free data retrieval call binding the contract method 0x3b72fbef.
//
// Solidity: function baseRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerSession) BaseRateMantissa() (*big.Int, error) {
	return _VaiController.Contract.BaseRateMantissa(&_VaiController.CallOpts)
}

// BaseRateMantissa is a free data retrieval call binding the contract method 0x3b72fbef.
//
// Solidity: function baseRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) BaseRateMantissa() (*big.Int, error) {
	return _VaiController.Contract.BaseRateMantissa(&_VaiController.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiController *VaiControllerCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiController *VaiControllerSession) Comptroller() (common.Address, error) {
	return _VaiController.Contract.Comptroller(&_VaiController.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiController *VaiControllerCallerSession) Comptroller() (common.Address, error) {
	return _VaiController.Contract.Comptroller(&_VaiController.CallOpts)
}

// FloatRateMantissa is a free data retrieval call binding the contract method 0x5ce73240.
//
// Solidity: function floatRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerCaller) FloatRateMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "floatRateMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FloatRateMantissa is a free data retrieval call binding the contract method 0x5ce73240.
//
// Solidity: function floatRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerSession) FloatRateMantissa() (*big.Int, error) {
	return _VaiController.Contract.FloatRateMantissa(&_VaiController.CallOpts)
}

// FloatRateMantissa is a free data retrieval call binding the contract method 0x5ce73240.
//
// Solidity: function floatRateMantissa() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) FloatRateMantissa() (*big.Int, error) {
	return _VaiController.Contract.FloatRateMantissa(&_VaiController.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiController *VaiControllerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiController *VaiControllerSession) GetBlockNumber() (*big.Int, error) {
	return _VaiController.Contract.GetBlockNumber(&_VaiController.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _VaiController.Contract.GetBlockNumber(&_VaiController.CallOpts)
}

// GetBlocksPerYear is a free data retrieval call binding the contract method 0x741de148.
//
// Solidity: function getBlocksPerYear() pure returns(uint256)
func (_VaiController *VaiControllerCaller) GetBlocksPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getBlocksPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlocksPerYear is a free data retrieval call binding the contract method 0x741de148.
//
// Solidity: function getBlocksPerYear() pure returns(uint256)
func (_VaiController *VaiControllerSession) GetBlocksPerYear() (*big.Int, error) {
	return _VaiController.Contract.GetBlocksPerYear(&_VaiController.CallOpts)
}

// GetBlocksPerYear is a free data retrieval call binding the contract method 0x741de148.
//
// Solidity: function getBlocksPerYear() pure returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetBlocksPerYear() (*big.Int, error) {
	return _VaiController.Contract.GetBlocksPerYear(&_VaiController.CallOpts)
}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiController *VaiControllerCaller) GetMintableVAI(opts *bind.CallOpts, minter common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getMintableVAI", minter)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiController *VaiControllerSession) GetMintableVAI(minter common.Address) (*big.Int, *big.Int, error) {
	return _VaiController.Contract.GetMintableVAI(&_VaiController.CallOpts, minter)
}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiController *VaiControllerCallerSession) GetMintableVAI(minter common.Address) (*big.Int, *big.Int, error) {
	return _VaiController.Contract.GetMintableVAI(&_VaiController.CallOpts, minter)
}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiController *VaiControllerCaller) GetVAIAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAIAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiController *VaiControllerSession) GetVAIAddress() (common.Address, error) {
	return _VaiController.Contract.GetVAIAddress(&_VaiController.CallOpts)
}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiController *VaiControllerCallerSession) GetVAIAddress() (common.Address, error) {
	return _VaiController.Contract.GetVAIAddress(&_VaiController.CallOpts)
}

// GetVAICalculateRepayAmount is a free data retrieval call binding the contract method 0x691e45ac.
//
// Solidity: function getVAICalculateRepayAmount(address borrower, uint256 repayAmount) view returns(uint256, uint256, uint256)
func (_VaiController *VaiControllerCaller) GetVAICalculateRepayAmount(opts *bind.CallOpts, borrower common.Address, repayAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAICalculateRepayAmount", borrower, repayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetVAICalculateRepayAmount is a free data retrieval call binding the contract method 0x691e45ac.
//
// Solidity: function getVAICalculateRepayAmount(address borrower, uint256 repayAmount) view returns(uint256, uint256, uint256)
func (_VaiController *VaiControllerSession) GetVAICalculateRepayAmount(borrower common.Address, repayAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VaiController.Contract.GetVAICalculateRepayAmount(&_VaiController.CallOpts, borrower, repayAmount)
}

// GetVAICalculateRepayAmount is a free data retrieval call binding the contract method 0x691e45ac.
//
// Solidity: function getVAICalculateRepayAmount(address borrower, uint256 repayAmount) view returns(uint256, uint256, uint256)
func (_VaiController *VaiControllerCallerSession) GetVAICalculateRepayAmount(borrower common.Address, repayAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VaiController.Contract.GetVAICalculateRepayAmount(&_VaiController.CallOpts, borrower, repayAmount)
}

// GetVAIMinterInterestIndex is a free data retrieval call binding the contract method 0x234f8977.
//
// Solidity: function getVAIMinterInterestIndex(address minter) view returns(uint256)
func (_VaiController *VaiControllerCaller) GetVAIMinterInterestIndex(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAIMinterInterestIndex", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVAIMinterInterestIndex is a free data retrieval call binding the contract method 0x234f8977.
//
// Solidity: function getVAIMinterInterestIndex(address minter) view returns(uint256)
func (_VaiController *VaiControllerSession) GetVAIMinterInterestIndex(minter common.Address) (*big.Int, error) {
	return _VaiController.Contract.GetVAIMinterInterestIndex(&_VaiController.CallOpts, minter)
}

// GetVAIMinterInterestIndex is a free data retrieval call binding the contract method 0x234f8977.
//
// Solidity: function getVAIMinterInterestIndex(address minter) view returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetVAIMinterInterestIndex(minter common.Address) (*big.Int, error) {
	return _VaiController.Contract.GetVAIMinterInterestIndex(&_VaiController.CallOpts, minter)
}

// GetVAIRepayAmount is a free data retrieval call binding the contract method 0x78c2f922.
//
// Solidity: function getVAIRepayAmount(address account) view returns(uint256)
func (_VaiController *VaiControllerCaller) GetVAIRepayAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAIRepayAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVAIRepayAmount is a free data retrieval call binding the contract method 0x78c2f922.
//
// Solidity: function getVAIRepayAmount(address account) view returns(uint256)
func (_VaiController *VaiControllerSession) GetVAIRepayAmount(account common.Address) (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayAmount(&_VaiController.CallOpts, account)
}

// GetVAIRepayAmount is a free data retrieval call binding the contract method 0x78c2f922.
//
// Solidity: function getVAIRepayAmount(address account) view returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetVAIRepayAmount(account common.Address) (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayAmount(&_VaiController.CallOpts, account)
}

// GetVAIRepayRate is a free data retrieval call binding the contract method 0xb9ee8726.
//
// Solidity: function getVAIRepayRate() view returns(uint256)
func (_VaiController *VaiControllerCaller) GetVAIRepayRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAIRepayRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVAIRepayRate is a free data retrieval call binding the contract method 0xb9ee8726.
//
// Solidity: function getVAIRepayRate() view returns(uint256)
func (_VaiController *VaiControllerSession) GetVAIRepayRate() (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayRate(&_VaiController.CallOpts)
}

// GetVAIRepayRate is a free data retrieval call binding the contract method 0xb9ee8726.
//
// Solidity: function getVAIRepayRate() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetVAIRepayRate() (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayRate(&_VaiController.CallOpts)
}

// GetVAIRepayRatePerBlock is a free data retrieval call binding the contract method 0x75c3de43.
//
// Solidity: function getVAIRepayRatePerBlock() view returns(uint256)
func (_VaiController *VaiControllerCaller) GetVAIRepayRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "getVAIRepayRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVAIRepayRatePerBlock is a free data retrieval call binding the contract method 0x75c3de43.
//
// Solidity: function getVAIRepayRatePerBlock() view returns(uint256)
func (_VaiController *VaiControllerSession) GetVAIRepayRatePerBlock() (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayRatePerBlock(&_VaiController.CallOpts)
}

// GetVAIRepayRatePerBlock is a free data retrieval call binding the contract method 0x75c3de43.
//
// Solidity: function getVAIRepayRatePerBlock() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) GetVAIRepayRatePerBlock() (*big.Int, error) {
	return _VaiController.Contract.GetVAIRepayRatePerBlock(&_VaiController.CallOpts)
}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiController *VaiControllerCaller) IsVenusVAIInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "isVenusVAIInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiController *VaiControllerSession) IsVenusVAIInitialized() (bool, error) {
	return _VaiController.Contract.IsVenusVAIInitialized(&_VaiController.CallOpts)
}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiController *VaiControllerCallerSession) IsVenusVAIInitialized() (bool, error) {
	return _VaiController.Contract.IsVenusVAIInitialized(&_VaiController.CallOpts)
}

// PastVAIInterest is a free data retrieval call binding the contract method 0xf20fd8f4.
//
// Solidity: function pastVAIInterest(address ) view returns(uint256)
func (_VaiController *VaiControllerCaller) PastVAIInterest(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "pastVAIInterest", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PastVAIInterest is a free data retrieval call binding the contract method 0xf20fd8f4.
//
// Solidity: function pastVAIInterest(address ) view returns(uint256)
func (_VaiController *VaiControllerSession) PastVAIInterest(arg0 common.Address) (*big.Int, error) {
	return _VaiController.Contract.PastVAIInterest(&_VaiController.CallOpts, arg0)
}

// PastVAIInterest is a free data retrieval call binding the contract method 0xf20fd8f4.
//
// Solidity: function pastVAIInterest(address ) view returns(uint256)
func (_VaiController *VaiControllerCallerSession) PastVAIInterest(arg0 common.Address) (*big.Int, error) {
	return _VaiController.Contract.PastVAIInterest(&_VaiController.CallOpts, arg0)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiController *VaiControllerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiController *VaiControllerSession) PendingAdmin() (common.Address, error) {
	return _VaiController.Contract.PendingAdmin(&_VaiController.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiController *VaiControllerCallerSession) PendingAdmin() (common.Address, error) {
	return _VaiController.Contract.PendingAdmin(&_VaiController.CallOpts)
}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiController *VaiControllerCaller) PendingVAIControllerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "pendingVAIControllerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiController *VaiControllerSession) PendingVAIControllerImplementation() (common.Address, error) {
	return _VaiController.Contract.PendingVAIControllerImplementation(&_VaiController.CallOpts)
}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiController *VaiControllerCallerSession) PendingVAIControllerImplementation() (common.Address, error) {
	return _VaiController.Contract.PendingVAIControllerImplementation(&_VaiController.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_VaiController *VaiControllerCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_VaiController *VaiControllerSession) Receiver() (common.Address, error) {
	return _VaiController.Contract.Receiver(&_VaiController.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_VaiController *VaiControllerCallerSession) Receiver() (common.Address, error) {
	return _VaiController.Contract.Receiver(&_VaiController.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiController *VaiControllerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiController *VaiControllerSession) TreasuryAddress() (common.Address, error) {
	return _VaiController.Contract.TreasuryAddress(&_VaiController.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiController *VaiControllerCallerSession) TreasuryAddress() (common.Address, error) {
	return _VaiController.Contract.TreasuryAddress(&_VaiController.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiController *VaiControllerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiController *VaiControllerSession) TreasuryGuardian() (common.Address, error) {
	return _VaiController.Contract.TreasuryGuardian(&_VaiController.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiController *VaiControllerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _VaiController.Contract.TreasuryGuardian(&_VaiController.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiController *VaiControllerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiController *VaiControllerSession) TreasuryPercent() (*big.Int, error) {
	return _VaiController.Contract.TreasuryPercent(&_VaiController.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiController *VaiControllerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _VaiController.Contract.TreasuryPercent(&_VaiController.CallOpts)
}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiController *VaiControllerCaller) VaiControllerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "vaiControllerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiController *VaiControllerSession) VaiControllerImplementation() (common.Address, error) {
	return _VaiController.Contract.VaiControllerImplementation(&_VaiController.CallOpts)
}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiController *VaiControllerCallerSession) VaiControllerImplementation() (common.Address, error) {
	return _VaiController.Contract.VaiControllerImplementation(&_VaiController.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiController *VaiControllerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiController *VaiControllerSession) VenusInitialIndex() (*big.Int, error) {
	return _VaiController.Contract.VenusInitialIndex(&_VaiController.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiController *VaiControllerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _VaiController.Contract.VenusInitialIndex(&_VaiController.CallOpts)
}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiController *VaiControllerCaller) VenusVAIMinterIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "venusVAIMinterIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiController *VaiControllerSession) VenusVAIMinterIndex(arg0 common.Address) (*big.Int, error) {
	return _VaiController.Contract.VenusVAIMinterIndex(&_VaiController.CallOpts, arg0)
}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiController *VaiControllerCallerSession) VenusVAIMinterIndex(arg0 common.Address) (*big.Int, error) {
	return _VaiController.Contract.VenusVAIMinterIndex(&_VaiController.CallOpts, arg0)
}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiController *VaiControllerCaller) VenusVAIState(opts *bind.CallOpts) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _VaiController.contract.Call(opts, &out, "venusVAIState")

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiController *VaiControllerSession) VenusVAIState() (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VaiController.Contract.VenusVAIState(&_VaiController.CallOpts)
}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiController *VaiControllerCallerSession) VenusVAIState() (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VaiController.Contract.VenusVAIState(&_VaiController.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiController *VaiControllerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiController *VaiControllerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.Become(&_VaiController.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiController *VaiControllerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.Become(&_VaiController.TransactOpts, unitroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiController *VaiControllerTransactor) SetComptroller(opts *bind.TransactOpts, comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "_setComptroller", comptroller_)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiController *VaiControllerSession) SetComptroller(comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetComptroller(&_VaiController.TransactOpts, comptroller_)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiController *VaiControllerTransactorSession) SetComptroller(comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetComptroller(&_VaiController.TransactOpts, comptroller_)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiController *VaiControllerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiController *VaiControllerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetTreasuryData(&_VaiController.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiController *VaiControllerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetTreasuryData(&_VaiController.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// AccrueVAIInterest is a paid mutator transaction binding the contract method 0xb49b1005.
//
// Solidity: function accrueVAIInterest() returns()
func (_VaiController *VaiControllerTransactor) AccrueVAIInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "accrueVAIInterest")
}

// AccrueVAIInterest is a paid mutator transaction binding the contract method 0xb49b1005.
//
// Solidity: function accrueVAIInterest() returns()
func (_VaiController *VaiControllerSession) AccrueVAIInterest() (*types.Transaction, error) {
	return _VaiController.Contract.AccrueVAIInterest(&_VaiController.TransactOpts)
}

// AccrueVAIInterest is a paid mutator transaction binding the contract method 0xb49b1005.
//
// Solidity: function accrueVAIInterest() returns()
func (_VaiController *VaiControllerTransactorSession) AccrueVAIInterest() (*types.Transaction, error) {
	return _VaiController.Contract.AccrueVAIInterest(&_VaiController.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiController *VaiControllerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiController *VaiControllerSession) Initialize() (*types.Transaction, error) {
	return _VaiController.Contract.Initialize(&_VaiController.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiController *VaiControllerTransactorSession) Initialize() (*types.Transaction, error) {
	return _VaiController.Contract.Initialize(&_VaiController.TransactOpts)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiController *VaiControllerTransactor) LiquidateVAI(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "liquidateVAI", borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiController *VaiControllerSession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.LiquidateVAI(&_VaiController.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiController *VaiControllerTransactorSession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.LiquidateVAI(&_VaiController.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiController *VaiControllerTransactor) MintVAI(opts *bind.TransactOpts, mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "mintVAI", mintVAIAmount)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiController *VaiControllerSession) MintVAI(mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.MintVAI(&_VaiController.TransactOpts, mintVAIAmount)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiController *VaiControllerTransactorSession) MintVAI(mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.MintVAI(&_VaiController.TransactOpts, mintVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiController *VaiControllerTransactor) RepayVAI(opts *bind.TransactOpts, repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "repayVAI", repayVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiController *VaiControllerSession) RepayVAI(repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.RepayVAI(&_VaiController.TransactOpts, repayVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiController *VaiControllerTransactorSession) RepayVAI(repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.RepayVAI(&_VaiController.TransactOpts, repayVAIAmount)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 newBaseRateMantissa) returns()
func (_VaiController *VaiControllerTransactor) SetBaseRate(opts *bind.TransactOpts, newBaseRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "setBaseRate", newBaseRateMantissa)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 newBaseRateMantissa) returns()
func (_VaiController *VaiControllerSession) SetBaseRate(newBaseRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetBaseRate(&_VaiController.TransactOpts, newBaseRateMantissa)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 newBaseRateMantissa) returns()
func (_VaiController *VaiControllerTransactorSession) SetBaseRate(newBaseRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetBaseRate(&_VaiController.TransactOpts, newBaseRateMantissa)
}

// SetFloatRate is a paid mutator transaction binding the contract method 0x3b5a0a64.
//
// Solidity: function setFloatRate(uint256 newFloatRateMantissa) returns()
func (_VaiController *VaiControllerTransactor) SetFloatRate(opts *bind.TransactOpts, newFloatRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "setFloatRate", newFloatRateMantissa)
}

// SetFloatRate is a paid mutator transaction binding the contract method 0x3b5a0a64.
//
// Solidity: function setFloatRate(uint256 newFloatRateMantissa) returns()
func (_VaiController *VaiControllerSession) SetFloatRate(newFloatRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetFloatRate(&_VaiController.TransactOpts, newFloatRateMantissa)
}

// SetFloatRate is a paid mutator transaction binding the contract method 0x3b5a0a64.
//
// Solidity: function setFloatRate(uint256 newFloatRateMantissa) returns()
func (_VaiController *VaiControllerTransactorSession) SetFloatRate(newFloatRateMantissa *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetFloatRate(&_VaiController.TransactOpts, newFloatRateMantissa)
}

// SetMintCap is a paid mutator transaction binding the contract method 0x4070a0c9.
//
// Solidity: function setMintCap(uint256 _mintCap) returns()
func (_VaiController *VaiControllerTransactor) SetMintCap(opts *bind.TransactOpts, _mintCap *big.Int) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "setMintCap", _mintCap)
}

// SetMintCap is a paid mutator transaction binding the contract method 0x4070a0c9.
//
// Solidity: function setMintCap(uint256 _mintCap) returns()
func (_VaiController *VaiControllerSession) SetMintCap(_mintCap *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetMintCap(&_VaiController.TransactOpts, _mintCap)
}

// SetMintCap is a paid mutator transaction binding the contract method 0x4070a0c9.
//
// Solidity: function setMintCap(uint256 _mintCap) returns()
func (_VaiController *VaiControllerTransactorSession) SetMintCap(_mintCap *big.Int) (*types.Transaction, error) {
	return _VaiController.Contract.SetMintCap(&_VaiController.TransactOpts, _mintCap)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address newReceiver) returns()
func (_VaiController *VaiControllerTransactor) SetReceiver(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "setReceiver", newReceiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address newReceiver) returns()
func (_VaiController *VaiControllerSession) SetReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetReceiver(&_VaiController.TransactOpts, newReceiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address newReceiver) returns()
func (_VaiController *VaiControllerTransactorSession) SetReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetReceiver(&_VaiController.TransactOpts, newReceiver)
}

// SetVAIAddress is a paid mutator transaction binding the contract method 0x919b22b4.
//
// Solidity: function setVAIAddress(address vai_) returns()
func (_VaiController *VaiControllerTransactor) SetVAIAddress(opts *bind.TransactOpts, vai_ common.Address) (*types.Transaction, error) {
	return _VaiController.contract.Transact(opts, "setVAIAddress", vai_)
}

// SetVAIAddress is a paid mutator transaction binding the contract method 0x919b22b4.
//
// Solidity: function setVAIAddress(address vai_) returns()
func (_VaiController *VaiControllerSession) SetVAIAddress(vai_ common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetVAIAddress(&_VaiController.TransactOpts, vai_)
}

// SetVAIAddress is a paid mutator transaction binding the contract method 0x919b22b4.
//
// Solidity: function setVAIAddress(address vai_) returns()
func (_VaiController *VaiControllerTransactorSession) SetVAIAddress(vai_ common.Address) (*types.Transaction, error) {
	return _VaiController.Contract.SetVAIAddress(&_VaiController.TransactOpts, vai_)
}

// VaiControllerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the VaiController contract.
type VaiControllerFailureIterator struct {
	Event *VaiControllerFailure // Event containing the contract specifics and raw log

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
func (it *VaiControllerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerFailure)
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
		it.Event = new(VaiControllerFailure)
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
func (it *VaiControllerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerFailure represents a Failure event raised by the VaiController contract.
type VaiControllerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VaiController *VaiControllerFilterer) FilterFailure(opts *bind.FilterOpts) (*VaiControllerFailureIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VaiControllerFailureIterator{contract: _VaiController.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VaiController *VaiControllerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VaiControllerFailure) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerFailure)
				if err := _VaiController.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_VaiController *VaiControllerFilterer) ParseFailure(log types.Log) (*VaiControllerFailure, error) {
	event := new(VaiControllerFailure)
	if err := _VaiController.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerLiquidateVAIIterator is returned from FilterLiquidateVAI and is used to iterate over the raw logs and unpacked data for LiquidateVAI events raised by the VaiController contract.
type VaiControllerLiquidateVAIIterator struct {
	Event *VaiControllerLiquidateVAI // Event containing the contract specifics and raw log

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
func (it *VaiControllerLiquidateVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerLiquidateVAI)
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
		it.Event = new(VaiControllerLiquidateVAI)
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
func (it *VaiControllerLiquidateVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerLiquidateVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerLiquidateVAI represents a LiquidateVAI event raised by the VaiController contract.
type VaiControllerLiquidateVAI struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	VTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateVAI is a free log retrieval operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiController *VaiControllerFilterer) FilterLiquidateVAI(opts *bind.FilterOpts) (*VaiControllerLiquidateVAIIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "LiquidateVAI")
	if err != nil {
		return nil, err
	}
	return &VaiControllerLiquidateVAIIterator{contract: _VaiController.contract, event: "LiquidateVAI", logs: logs, sub: sub}, nil
}

// WatchLiquidateVAI is a free log subscription operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiController *VaiControllerFilterer) WatchLiquidateVAI(opts *bind.WatchOpts, sink chan<- *VaiControllerLiquidateVAI) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "LiquidateVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerLiquidateVAI)
				if err := _VaiController.contract.UnpackLog(event, "LiquidateVAI", log); err != nil {
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

// ParseLiquidateVAI is a log parse operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiController *VaiControllerFilterer) ParseLiquidateVAI(log types.Log) (*VaiControllerLiquidateVAI, error) {
	event := new(VaiControllerLiquidateVAI)
	if err := _VaiController.contract.UnpackLog(event, "LiquidateVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerMintFeeIterator is returned from FilterMintFee and is used to iterate over the raw logs and unpacked data for MintFee events raised by the VaiController contract.
type VaiControllerMintFeeIterator struct {
	Event *VaiControllerMintFee // Event containing the contract specifics and raw log

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
func (it *VaiControllerMintFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerMintFee)
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
		it.Event = new(VaiControllerMintFee)
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
func (it *VaiControllerMintFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerMintFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerMintFee represents a MintFee event raised by the VaiController contract.
type VaiControllerMintFee struct {
	Minter    common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintFee is a free log retrieval operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiController *VaiControllerFilterer) FilterMintFee(opts *bind.FilterOpts) (*VaiControllerMintFeeIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "MintFee")
	if err != nil {
		return nil, err
	}
	return &VaiControllerMintFeeIterator{contract: _VaiController.contract, event: "MintFee", logs: logs, sub: sub}, nil
}

// WatchMintFee is a free log subscription operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiController *VaiControllerFilterer) WatchMintFee(opts *bind.WatchOpts, sink chan<- *VaiControllerMintFee) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "MintFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerMintFee)
				if err := _VaiController.contract.UnpackLog(event, "MintFee", log); err != nil {
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

// ParseMintFee is a log parse operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiController *VaiControllerFilterer) ParseMintFee(log types.Log) (*VaiControllerMintFee, error) {
	event := new(VaiControllerMintFee)
	if err := _VaiController.contract.UnpackLog(event, "MintFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerMintVAIIterator is returned from FilterMintVAI and is used to iterate over the raw logs and unpacked data for MintVAI events raised by the VaiController contract.
type VaiControllerMintVAIIterator struct {
	Event *VaiControllerMintVAI // Event containing the contract specifics and raw log

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
func (it *VaiControllerMintVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerMintVAI)
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
		it.Event = new(VaiControllerMintVAI)
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
func (it *VaiControllerMintVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerMintVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerMintVAI represents a MintVAI event raised by the VaiController contract.
type VaiControllerMintVAI struct {
	Minter        common.Address
	MintVAIAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintVAI is a free log retrieval operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiController *VaiControllerFilterer) FilterMintVAI(opts *bind.FilterOpts) (*VaiControllerMintVAIIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "MintVAI")
	if err != nil {
		return nil, err
	}
	return &VaiControllerMintVAIIterator{contract: _VaiController.contract, event: "MintVAI", logs: logs, sub: sub}, nil
}

// WatchMintVAI is a free log subscription operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiController *VaiControllerFilterer) WatchMintVAI(opts *bind.WatchOpts, sink chan<- *VaiControllerMintVAI) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "MintVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerMintVAI)
				if err := _VaiController.contract.UnpackLog(event, "MintVAI", log); err != nil {
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

// ParseMintVAI is a log parse operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiController *VaiControllerFilterer) ParseMintVAI(log types.Log) (*VaiControllerMintVAI, error) {
	event := new(VaiControllerMintVAI)
	if err := _VaiController.contract.UnpackLog(event, "MintVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the VaiController contract.
type VaiControllerNewComptrollerIterator struct {
	Event *VaiControllerNewComptroller // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewComptroller)
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
		it.Event = new(VaiControllerNewComptroller)
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
func (it *VaiControllerNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewComptroller represents a NewComptroller event raised by the VaiController contract.
type VaiControllerNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_VaiController *VaiControllerFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*VaiControllerNewComptrollerIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewComptrollerIterator{contract: _VaiController.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_VaiController *VaiControllerFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *VaiControllerNewComptroller) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewComptroller)
				if err := _VaiController.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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
func (_VaiController *VaiControllerFilterer) ParseNewComptroller(log types.Log) (*VaiControllerNewComptroller, error) {
	event := new(VaiControllerNewComptroller)
	if err := _VaiController.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the VaiController contract.
type VaiControllerNewTreasuryAddressIterator struct {
	Event *VaiControllerNewTreasuryAddress // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewTreasuryAddress)
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
		it.Event = new(VaiControllerNewTreasuryAddress)
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
func (it *VaiControllerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewTreasuryAddress represents a NewTreasuryAddress event raised by the VaiController contract.
type VaiControllerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiController *VaiControllerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*VaiControllerNewTreasuryAddressIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewTreasuryAddressIterator{contract: _VaiController.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiController *VaiControllerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *VaiControllerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewTreasuryAddress)
				if err := _VaiController.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
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

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiController *VaiControllerFilterer) ParseNewTreasuryAddress(log types.Log) (*VaiControllerNewTreasuryAddress, error) {
	event := new(VaiControllerNewTreasuryAddress)
	if err := _VaiController.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the VaiController contract.
type VaiControllerNewTreasuryGuardianIterator struct {
	Event *VaiControllerNewTreasuryGuardian // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewTreasuryGuardian)
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
		it.Event = new(VaiControllerNewTreasuryGuardian)
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
func (it *VaiControllerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the VaiController contract.
type VaiControllerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiController *VaiControllerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*VaiControllerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewTreasuryGuardianIterator{contract: _VaiController.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiController *VaiControllerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *VaiControllerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewTreasuryGuardian)
				if err := _VaiController.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
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

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiController *VaiControllerFilterer) ParseNewTreasuryGuardian(log types.Log) (*VaiControllerNewTreasuryGuardian, error) {
	event := new(VaiControllerNewTreasuryGuardian)
	if err := _VaiController.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the VaiController contract.
type VaiControllerNewTreasuryPercentIterator struct {
	Event *VaiControllerNewTreasuryPercent // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewTreasuryPercent)
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
		it.Event = new(VaiControllerNewTreasuryPercent)
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
func (it *VaiControllerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewTreasuryPercent represents a NewTreasuryPercent event raised by the VaiController contract.
type VaiControllerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiController *VaiControllerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*VaiControllerNewTreasuryPercentIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewTreasuryPercentIterator{contract: _VaiController.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiController *VaiControllerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *VaiControllerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewTreasuryPercent)
				if err := _VaiController.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
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

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiController *VaiControllerFilterer) ParseNewTreasuryPercent(log types.Log) (*VaiControllerNewTreasuryPercent, error) {
	event := new(VaiControllerNewTreasuryPercent)
	if err := _VaiController.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewVAIBaseRateIterator is returned from FilterNewVAIBaseRate and is used to iterate over the raw logs and unpacked data for NewVAIBaseRate events raised by the VaiController contract.
type VaiControllerNewVAIBaseRateIterator struct {
	Event *VaiControllerNewVAIBaseRate // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewVAIBaseRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewVAIBaseRate)
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
		it.Event = new(VaiControllerNewVAIBaseRate)
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
func (it *VaiControllerNewVAIBaseRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewVAIBaseRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewVAIBaseRate represents a NewVAIBaseRate event raised by the VaiController contract.
type VaiControllerNewVAIBaseRate struct {
	OldBaseRateMantissa *big.Int
	NewBaseRateMantissa *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewVAIBaseRate is a free log retrieval operation binding the contract event 0xc84c32795e68685ec107b0e94ae126ef464095f342c7e2e0fec06a23d2e8677e.
//
// Solidity: event NewVAIBaseRate(uint256 oldBaseRateMantissa, uint256 newBaseRateMantissa)
func (_VaiController *VaiControllerFilterer) FilterNewVAIBaseRate(opts *bind.FilterOpts) (*VaiControllerNewVAIBaseRateIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewVAIBaseRate")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewVAIBaseRateIterator{contract: _VaiController.contract, event: "NewVAIBaseRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIBaseRate is a free log subscription operation binding the contract event 0xc84c32795e68685ec107b0e94ae126ef464095f342c7e2e0fec06a23d2e8677e.
//
// Solidity: event NewVAIBaseRate(uint256 oldBaseRateMantissa, uint256 newBaseRateMantissa)
func (_VaiController *VaiControllerFilterer) WatchNewVAIBaseRate(opts *bind.WatchOpts, sink chan<- *VaiControllerNewVAIBaseRate) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewVAIBaseRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewVAIBaseRate)
				if err := _VaiController.contract.UnpackLog(event, "NewVAIBaseRate", log); err != nil {
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

// ParseNewVAIBaseRate is a log parse operation binding the contract event 0xc84c32795e68685ec107b0e94ae126ef464095f342c7e2e0fec06a23d2e8677e.
//
// Solidity: event NewVAIBaseRate(uint256 oldBaseRateMantissa, uint256 newBaseRateMantissa)
func (_VaiController *VaiControllerFilterer) ParseNewVAIBaseRate(log types.Log) (*VaiControllerNewVAIBaseRate, error) {
	event := new(VaiControllerNewVAIBaseRate)
	if err := _VaiController.contract.UnpackLog(event, "NewVAIBaseRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewVAIFloatRateIterator is returned from FilterNewVAIFloatRate and is used to iterate over the raw logs and unpacked data for NewVAIFloatRate events raised by the VaiController contract.
type VaiControllerNewVAIFloatRateIterator struct {
	Event *VaiControllerNewVAIFloatRate // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewVAIFloatRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewVAIFloatRate)
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
		it.Event = new(VaiControllerNewVAIFloatRate)
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
func (it *VaiControllerNewVAIFloatRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewVAIFloatRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewVAIFloatRate represents a NewVAIFloatRate event raised by the VaiController contract.
type VaiControllerNewVAIFloatRate struct {
	OldFloatRateMantissa *big.Int
	NewFlatRateMantissa  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVAIFloatRate is a free log retrieval operation binding the contract event 0x546fb35dbbd92233aecc22b5a11a6791e5db7ec14f62e49cbac2a10c0437f561.
//
// Solidity: event NewVAIFloatRate(uint256 oldFloatRateMantissa, uint256 newFlatRateMantissa)
func (_VaiController *VaiControllerFilterer) FilterNewVAIFloatRate(opts *bind.FilterOpts) (*VaiControllerNewVAIFloatRateIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewVAIFloatRate")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewVAIFloatRateIterator{contract: _VaiController.contract, event: "NewVAIFloatRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIFloatRate is a free log subscription operation binding the contract event 0x546fb35dbbd92233aecc22b5a11a6791e5db7ec14f62e49cbac2a10c0437f561.
//
// Solidity: event NewVAIFloatRate(uint256 oldFloatRateMantissa, uint256 newFlatRateMantissa)
func (_VaiController *VaiControllerFilterer) WatchNewVAIFloatRate(opts *bind.WatchOpts, sink chan<- *VaiControllerNewVAIFloatRate) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewVAIFloatRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewVAIFloatRate)
				if err := _VaiController.contract.UnpackLog(event, "NewVAIFloatRate", log); err != nil {
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

// ParseNewVAIFloatRate is a log parse operation binding the contract event 0x546fb35dbbd92233aecc22b5a11a6791e5db7ec14f62e49cbac2a10c0437f561.
//
// Solidity: event NewVAIFloatRate(uint256 oldFloatRateMantissa, uint256 newFlatRateMantissa)
func (_VaiController *VaiControllerFilterer) ParseNewVAIFloatRate(log types.Log) (*VaiControllerNewVAIFloatRate, error) {
	event := new(VaiControllerNewVAIFloatRate)
	if err := _VaiController.contract.UnpackLog(event, "NewVAIFloatRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewVAIMintCapIterator is returned from FilterNewVAIMintCap and is used to iterate over the raw logs and unpacked data for NewVAIMintCap events raised by the VaiController contract.
type VaiControllerNewVAIMintCapIterator struct {
	Event *VaiControllerNewVAIMintCap // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewVAIMintCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewVAIMintCap)
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
		it.Event = new(VaiControllerNewVAIMintCap)
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
func (it *VaiControllerNewVAIMintCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewVAIMintCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewVAIMintCap represents a NewVAIMintCap event raised by the VaiController contract.
type VaiControllerNewVAIMintCap struct {
	OldMintCap *big.Int
	NewMintCap *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintCap is a free log retrieval operation binding the contract event 0x43862b3eea2df8fce70329f3f84cbcad220f47a73be46c5e00df25165a6e1695.
//
// Solidity: event NewVAIMintCap(uint256 oldMintCap, uint256 newMintCap)
func (_VaiController *VaiControllerFilterer) FilterNewVAIMintCap(opts *bind.FilterOpts) (*VaiControllerNewVAIMintCapIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewVAIMintCap")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewVAIMintCapIterator{contract: _VaiController.contract, event: "NewVAIMintCap", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintCap is a free log subscription operation binding the contract event 0x43862b3eea2df8fce70329f3f84cbcad220f47a73be46c5e00df25165a6e1695.
//
// Solidity: event NewVAIMintCap(uint256 oldMintCap, uint256 newMintCap)
func (_VaiController *VaiControllerFilterer) WatchNewVAIMintCap(opts *bind.WatchOpts, sink chan<- *VaiControllerNewVAIMintCap) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewVAIMintCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewVAIMintCap)
				if err := _VaiController.contract.UnpackLog(event, "NewVAIMintCap", log); err != nil {
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

// ParseNewVAIMintCap is a log parse operation binding the contract event 0x43862b3eea2df8fce70329f3f84cbcad220f47a73be46c5e00df25165a6e1695.
//
// Solidity: event NewVAIMintCap(uint256 oldMintCap, uint256 newMintCap)
func (_VaiController *VaiControllerFilterer) ParseNewVAIMintCap(log types.Log) (*VaiControllerNewVAIMintCap, error) {
	event := new(VaiControllerNewVAIMintCap)
	if err := _VaiController.contract.UnpackLog(event, "NewVAIMintCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerNewVAIReceiverIterator is returned from FilterNewVAIReceiver and is used to iterate over the raw logs and unpacked data for NewVAIReceiver events raised by the VaiController contract.
type VaiControllerNewVAIReceiverIterator struct {
	Event *VaiControllerNewVAIReceiver // Event containing the contract specifics and raw log

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
func (it *VaiControllerNewVAIReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerNewVAIReceiver)
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
		it.Event = new(VaiControllerNewVAIReceiver)
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
func (it *VaiControllerNewVAIReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerNewVAIReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerNewVAIReceiver represents a NewVAIReceiver event raised by the VaiController contract.
type VaiControllerNewVAIReceiver struct {
	OldReceiver common.Address
	NewReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewVAIReceiver is a free log retrieval operation binding the contract event 0x4df62dd7d9cc4f480a167c19c616ae5d5bb40db6d0c2bc66dba57068225f00d8.
//
// Solidity: event NewVAIReceiver(address oldReceiver, address newReceiver)
func (_VaiController *VaiControllerFilterer) FilterNewVAIReceiver(opts *bind.FilterOpts) (*VaiControllerNewVAIReceiverIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "NewVAIReceiver")
	if err != nil {
		return nil, err
	}
	return &VaiControllerNewVAIReceiverIterator{contract: _VaiController.contract, event: "NewVAIReceiver", logs: logs, sub: sub}, nil
}

// WatchNewVAIReceiver is a free log subscription operation binding the contract event 0x4df62dd7d9cc4f480a167c19c616ae5d5bb40db6d0c2bc66dba57068225f00d8.
//
// Solidity: event NewVAIReceiver(address oldReceiver, address newReceiver)
func (_VaiController *VaiControllerFilterer) WatchNewVAIReceiver(opts *bind.WatchOpts, sink chan<- *VaiControllerNewVAIReceiver) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "NewVAIReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerNewVAIReceiver)
				if err := _VaiController.contract.UnpackLog(event, "NewVAIReceiver", log); err != nil {
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

// ParseNewVAIReceiver is a log parse operation binding the contract event 0x4df62dd7d9cc4f480a167c19c616ae5d5bb40db6d0c2bc66dba57068225f00d8.
//
// Solidity: event NewVAIReceiver(address oldReceiver, address newReceiver)
func (_VaiController *VaiControllerFilterer) ParseNewVAIReceiver(log types.Log) (*VaiControllerNewVAIReceiver, error) {
	event := new(VaiControllerNewVAIReceiver)
	if err := _VaiController.contract.UnpackLog(event, "NewVAIReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiControllerRepayVAIIterator is returned from FilterRepayVAI and is used to iterate over the raw logs and unpacked data for RepayVAI events raised by the VaiController contract.
type VaiControllerRepayVAIIterator struct {
	Event *VaiControllerRepayVAI // Event containing the contract specifics and raw log

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
func (it *VaiControllerRepayVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiControllerRepayVAI)
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
		it.Event = new(VaiControllerRepayVAI)
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
func (it *VaiControllerRepayVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiControllerRepayVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiControllerRepayVAI represents a RepayVAI event raised by the VaiController contract.
type VaiControllerRepayVAI struct {
	Payer          common.Address
	Borrower       common.Address
	RepayVAIAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayVAI is a free log retrieval operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiController *VaiControllerFilterer) FilterRepayVAI(opts *bind.FilterOpts) (*VaiControllerRepayVAIIterator, error) {

	logs, sub, err := _VaiController.contract.FilterLogs(opts, "RepayVAI")
	if err != nil {
		return nil, err
	}
	return &VaiControllerRepayVAIIterator{contract: _VaiController.contract, event: "RepayVAI", logs: logs, sub: sub}, nil
}

// WatchRepayVAI is a free log subscription operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiController *VaiControllerFilterer) WatchRepayVAI(opts *bind.WatchOpts, sink chan<- *VaiControllerRepayVAI) (event.Subscription, error) {

	logs, sub, err := _VaiController.contract.WatchLogs(opts, "RepayVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiControllerRepayVAI)
				if err := _VaiController.contract.UnpackLog(event, "RepayVAI", log); err != nil {
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

// ParseRepayVAI is a log parse operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiController *VaiControllerFilterer) ParseRepayVAI(log types.Log) (*VaiControllerRepayVAI, error) {
	event := new(VaiControllerRepayVAI)
	if err := _VaiController.contract.UnpackLog(event, "RepayVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
