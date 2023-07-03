// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
	_ = abi.ConvertType
)

// LeopardBridgeMetaData contains all meta data concerning the LeopardBridge contract.
var LeopardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deposit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURL\",\"type\":\"string\"}],\"name\":\"Depositor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Withdrawer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifyString\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LeopardBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use LeopardBridgeMetaData.ABI instead.
var LeopardBridgeABI = LeopardBridgeMetaData.ABI

// LeopardBridge is an auto generated Go binding around an Ethereum contract.
type LeopardBridge struct {
	LeopardBridgeCaller     // Read-only binding to the contract
	LeopardBridgeTransactor // Write-only binding to the contract
	LeopardBridgeFilterer   // Log filterer for contract events
}

// LeopardBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeopardBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeopardBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeopardBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeopardBridgeSession struct {
	Contract     *LeopardBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeopardBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeopardBridgeCallerSession struct {
	Contract *LeopardBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LeopardBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeopardBridgeTransactorSession struct {
	Contract     *LeopardBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LeopardBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeopardBridgeRaw struct {
	Contract *LeopardBridge // Generic contract binding to access the raw methods on
}

// LeopardBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeopardBridgeCallerRaw struct {
	Contract *LeopardBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// LeopardBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeopardBridgeTransactorRaw struct {
	Contract *LeopardBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeopardBridge creates a new instance of LeopardBridge, bound to a specific deployed contract.
func NewLeopardBridge(address common.Address, backend bind.ContractBackend) (*LeopardBridge, error) {
	contract, err := bindLeopardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeopardBridge{LeopardBridgeCaller: LeopardBridgeCaller{contract: contract}, LeopardBridgeTransactor: LeopardBridgeTransactor{contract: contract}, LeopardBridgeFilterer: LeopardBridgeFilterer{contract: contract}}, nil
}

// NewLeopardBridgeCaller creates a new read-only instance of LeopardBridge, bound to a specific deployed contract.
func NewLeopardBridgeCaller(address common.Address, caller bind.ContractCaller) (*LeopardBridgeCaller, error) {
	contract, err := bindLeopardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeopardBridgeCaller{contract: contract}, nil
}

// NewLeopardBridgeTransactor creates a new write-only instance of LeopardBridge, bound to a specific deployed contract.
func NewLeopardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*LeopardBridgeTransactor, error) {
	contract, err := bindLeopardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeopardBridgeTransactor{contract: contract}, nil
}

// NewLeopardBridgeFilterer creates a new log filterer instance of LeopardBridge, bound to a specific deployed contract.
func NewLeopardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*LeopardBridgeFilterer, error) {
	contract, err := bindLeopardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeopardBridgeFilterer{contract: contract}, nil
}

// bindLeopardBridge binds a generic wrapper to an already deployed contract.
func bindLeopardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeopardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeopardBridge *LeopardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeopardBridge.Contract.LeopardBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeopardBridge *LeopardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeopardBridge.Contract.LeopardBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeopardBridge *LeopardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeopardBridge.Contract.LeopardBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeopardBridge *LeopardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeopardBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeopardBridge *LeopardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeopardBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeopardBridge *LeopardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeopardBridge.Contract.contract.Transact(opts, method, params...)
}

// VerifyString is a free data retrieval call binding the contract method 0xd2c3fbb7.
//
// Solidity: function verifyString(string message, uint8 v, bytes32 r, bytes32 s) pure returns(address signer)
func (_LeopardBridge *LeopardBridgeCaller) VerifyString(opts *bind.CallOpts, message string, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LeopardBridge.contract.Call(opts, &out, "verifyString", message, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyString is a free data retrieval call binding the contract method 0xd2c3fbb7.
//
// Solidity: function verifyString(string message, uint8 v, bytes32 r, bytes32 s) pure returns(address signer)
func (_LeopardBridge *LeopardBridgeSession) VerifyString(message string, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _LeopardBridge.Contract.VerifyString(&_LeopardBridge.CallOpts, message, v, r, s)
}

// VerifyString is a free data retrieval call binding the contract method 0xd2c3fbb7.
//
// Solidity: function verifyString(string message, uint8 v, bytes32 r, bytes32 s) pure returns(address signer)
func (_LeopardBridge *LeopardBridgeCallerSession) VerifyString(message string, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _LeopardBridge.Contract.VerifyString(&_LeopardBridge.CallOpts, message, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_LeopardBridge *LeopardBridgeTransactor) Deposit(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.contract.Transact(opts, "deposit", tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_LeopardBridge *LeopardBridgeSession) Deposit(tokenId *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.Contract.Deposit(&_LeopardBridge.TransactOpts, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_LeopardBridge *LeopardBridgeTransactorSession) Deposit(tokenId *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.Contract.Deposit(&_LeopardBridge.TransactOpts, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_LeopardBridge *LeopardBridgeTransactor) Withdraw(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.contract.Transact(opts, "withdraw", tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_LeopardBridge *LeopardBridgeSession) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.Contract.Withdraw(&_LeopardBridge.TransactOpts, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_LeopardBridge *LeopardBridgeTransactorSession) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _LeopardBridge.Contract.Withdraw(&_LeopardBridge.TransactOpts, tokenID)
}

// LeopardBridgeDepositorIterator is returned from FilterDepositor and is used to iterate over the raw logs and unpacked data for Depositor events raised by the LeopardBridge contract.
type LeopardBridgeDepositorIterator struct {
	Event *LeopardBridgeDepositor // Event containing the contract specifics and raw log

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
func (it *LeopardBridgeDepositorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardBridgeDepositor)
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
		it.Event = new(LeopardBridgeDepositor)
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
func (it *LeopardBridgeDepositorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardBridgeDepositorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardBridgeDepositor represents a Depositor event raised by the LeopardBridge contract.
type LeopardBridgeDepositor struct {
	Deposit  common.Address
	TokenID  *big.Int
	TokenURL string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositor is a free log retrieval operation binding the contract event 0xf10b860f691c47b47ec33530a9bf1a4f632ab3dc1f62dbba2cccfa3ecdb1e31e.
//
// Solidity: event Depositor(address deposit, uint256 tokenID, string tokenURL)
func (_LeopardBridge *LeopardBridgeFilterer) FilterDepositor(opts *bind.FilterOpts) (*LeopardBridgeDepositorIterator, error) {

	logs, sub, err := _LeopardBridge.contract.FilterLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return &LeopardBridgeDepositorIterator{contract: _LeopardBridge.contract, event: "Depositor", logs: logs, sub: sub}, nil
}

// WatchDepositor is a free log subscription operation binding the contract event 0xf10b860f691c47b47ec33530a9bf1a4f632ab3dc1f62dbba2cccfa3ecdb1e31e.
//
// Solidity: event Depositor(address deposit, uint256 tokenID, string tokenURL)
func (_LeopardBridge *LeopardBridgeFilterer) WatchDepositor(opts *bind.WatchOpts, sink chan<- *LeopardBridgeDepositor) (event.Subscription, error) {

	logs, sub, err := _LeopardBridge.contract.WatchLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardBridgeDepositor)
				if err := _LeopardBridge.contract.UnpackLog(event, "Depositor", log); err != nil {
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

// ParseDepositor is a log parse operation binding the contract event 0xf10b860f691c47b47ec33530a9bf1a4f632ab3dc1f62dbba2cccfa3ecdb1e31e.
//
// Solidity: event Depositor(address deposit, uint256 tokenID, string tokenURL)
func (_LeopardBridge *LeopardBridgeFilterer) ParseDepositor(log types.Log) (*LeopardBridgeDepositor, error) {
	event := new(LeopardBridgeDepositor)
	if err := _LeopardBridge.contract.UnpackLog(event, "Depositor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeopardBridgeWithdrawerIterator is returned from FilterWithdrawer and is used to iterate over the raw logs and unpacked data for Withdrawer events raised by the LeopardBridge contract.
type LeopardBridgeWithdrawerIterator struct {
	Event *LeopardBridgeWithdrawer // Event containing the contract specifics and raw log

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
func (it *LeopardBridgeWithdrawerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardBridgeWithdrawer)
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
		it.Event = new(LeopardBridgeWithdrawer)
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
func (it *LeopardBridgeWithdrawerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardBridgeWithdrawerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardBridgeWithdrawer represents a Withdrawer event raised by the LeopardBridge contract.
type LeopardBridgeWithdrawer struct {
	Withdrawer common.Address
	TokenID    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawer is a free log retrieval operation binding the contract event 0x8cb11afff86f7ff1d42446d5e53c4be953f1f2cfb679cc213c66ee01a984123b.
//
// Solidity: event Withdrawer(address withdrawer, uint256 tokenID)
func (_LeopardBridge *LeopardBridgeFilterer) FilterWithdrawer(opts *bind.FilterOpts) (*LeopardBridgeWithdrawerIterator, error) {

	logs, sub, err := _LeopardBridge.contract.FilterLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return &LeopardBridgeWithdrawerIterator{contract: _LeopardBridge.contract, event: "Withdrawer", logs: logs, sub: sub}, nil
}

// WatchWithdrawer is a free log subscription operation binding the contract event 0x8cb11afff86f7ff1d42446d5e53c4be953f1f2cfb679cc213c66ee01a984123b.
//
// Solidity: event Withdrawer(address withdrawer, uint256 tokenID)
func (_LeopardBridge *LeopardBridgeFilterer) WatchWithdrawer(opts *bind.WatchOpts, sink chan<- *LeopardBridgeWithdrawer) (event.Subscription, error) {

	logs, sub, err := _LeopardBridge.contract.WatchLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardBridgeWithdrawer)
				if err := _LeopardBridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
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

// ParseWithdrawer is a log parse operation binding the contract event 0x8cb11afff86f7ff1d42446d5e53c4be953f1f2cfb679cc213c66ee01a984123b.
//
// Solidity: event Withdrawer(address withdrawer, uint256 tokenID)
func (_LeopardBridge *LeopardBridgeFilterer) ParseWithdrawer(log types.Log) (*LeopardBridgeWithdrawer, error) {
	event := new(LeopardBridgeWithdrawer)
	if err := _LeopardBridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
