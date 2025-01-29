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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeder\",\"type\":\"address\"}],\"name\":\"UpdateFeeder\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPriceOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeder\",\"type\":\"address\"}],\"name\":\"updateFeeder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Feeder is a free data retrieval call binding the contract method 0xbc6c0fa5.
//
// Solidity: function feeder(address ) view returns(address)
func (_Oracle *OracleCaller) Feeder(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "feeder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Feeder is a free data retrieval call binding the contract method 0xbc6c0fa5.
//
// Solidity: function feeder(address ) view returns(address)
func (_Oracle *OracleSession) Feeder(arg0 common.Address) (common.Address, error) {
	return _Oracle.Contract.Feeder(&_Oracle.CallOpts, arg0)
}

// Feeder is a free data retrieval call binding the contract method 0xbc6c0fa5.
//
// Solidity: function feeder(address ) view returns(address)
func (_Oracle *OracleCallerSession) Feeder(arg0 common.Address) (common.Address, error) {
	return _Oracle.Contract.Feeder(&_Oracle.CallOpts, arg0)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_Oracle *OracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, vToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUnderlyingPrice", vToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_Oracle *OracleSession) GetUnderlyingPrice(vToken common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.CallOpts, vToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_Oracle *OracleCallerSession) GetUnderlyingPrice(vToken common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.CallOpts, vToken)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleCaller) IsPriceOracle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isPriceOracle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleSession) IsPriceOracle() (bool, error) {
	return _Oracle.Contract.IsPriceOracle(&_Oracle.CallOpts)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleCallerSession) IsPriceOracle() (bool, error) {
	return _Oracle.Contract.IsPriceOracle(&_Oracle.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Oracle *OracleCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Oracle *OracleSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Wards(&_Oracle.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Oracle *OracleCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Wards(&_Oracle.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Oracle *OracleTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Oracle *OracleSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Deny(&_Oracle.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Oracle *OracleTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Deny(&_Oracle.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Oracle *OracleTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Oracle *OracleSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Rely(&_Oracle.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Oracle *OracleTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Rely(&_Oracle.TransactOpts, guy)
}

// UpdateFeeder is a paid mutator transaction binding the contract method 0x8b5d5743.
//
// Solidity: function updateFeeder(address vToken, address newFeeder) returns()
func (_Oracle *OracleTransactor) UpdateFeeder(opts *bind.TransactOpts, vToken common.Address, newFeeder common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateFeeder", vToken, newFeeder)
}

// UpdateFeeder is a paid mutator transaction binding the contract method 0x8b5d5743.
//
// Solidity: function updateFeeder(address vToken, address newFeeder) returns()
func (_Oracle *OracleSession) UpdateFeeder(vToken common.Address, newFeeder common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateFeeder(&_Oracle.TransactOpts, vToken, newFeeder)
}

// UpdateFeeder is a paid mutator transaction binding the contract method 0x8b5d5743.
//
// Solidity: function updateFeeder(address vToken, address newFeeder) returns()
func (_Oracle *OracleTransactorSession) UpdateFeeder(vToken common.Address, newFeeder common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateFeeder(&_Oracle.TransactOpts, vToken, newFeeder)
}

// OracleUpdateFeederIterator is returned from FilterUpdateFeeder and is used to iterate over the raw logs and unpacked data for UpdateFeeder events raised by the Oracle contract.
type OracleUpdateFeederIterator struct {
	Event *OracleUpdateFeeder // Event containing the contract specifics and raw log

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
func (it *OracleUpdateFeederIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUpdateFeeder)
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
		it.Event = new(OracleUpdateFeeder)
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
func (it *OracleUpdateFeederIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUpdateFeederIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUpdateFeeder represents a UpdateFeeder event raised by the Oracle contract.
type OracleUpdateFeeder struct {
	VToken    common.Address
	OldFeeder common.Address
	NewFeeder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeder is a free log retrieval operation binding the contract event 0x448e493dc8404f24e65ab6f0e5b8352c18005f91c49520f689340c11aeb107f7.
//
// Solidity: event UpdateFeeder(address indexed vToken, address oldFeeder, address newFeeder)
func (_Oracle *OracleFilterer) FilterUpdateFeeder(opts *bind.FilterOpts, vToken []common.Address) (*OracleUpdateFeederIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UpdateFeeder", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &OracleUpdateFeederIterator{contract: _Oracle.contract, event: "UpdateFeeder", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeder is a free log subscription operation binding the contract event 0x448e493dc8404f24e65ab6f0e5b8352c18005f91c49520f689340c11aeb107f7.
//
// Solidity: event UpdateFeeder(address indexed vToken, address oldFeeder, address newFeeder)
func (_Oracle *OracleFilterer) WatchUpdateFeeder(opts *bind.WatchOpts, sink chan<- *OracleUpdateFeeder, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UpdateFeeder", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUpdateFeeder)
				if err := _Oracle.contract.UnpackLog(event, "UpdateFeeder", log); err != nil {
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

// ParseUpdateFeeder is a log parse operation binding the contract event 0x448e493dc8404f24e65ab6f0e5b8352c18005f91c49520f689340c11aeb107f7.
//
// Solidity: event UpdateFeeder(address indexed vToken, address oldFeeder, address newFeeder)
func (_Oracle *OracleFilterer) ParseUpdateFeeder(log types.Log) (*OracleUpdateFeeder, error) {
	event := new(OracleUpdateFeeder)
	if err := _Oracle.contract.UnpackLog(event, "UpdateFeeder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
