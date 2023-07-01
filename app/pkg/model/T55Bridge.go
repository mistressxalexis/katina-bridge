// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model

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

// T55BridgeMetaData contains all meta data concerning the T55Bridge contract.
var T55BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nfts\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"ClaimNFTs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"Depositor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Withdrawer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"claimNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"syncNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"t55\",\"outputs\":[{\"internalType\":\"contractT55NFTs\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokendID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// T55BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use T55BridgeMetaData.ABI instead.
var T55BridgeABI = T55BridgeMetaData.ABI

// T55Bridge is an auto generated Go binding around an Ethereum contract.
type T55Bridge struct {
	T55BridgeCaller     // Read-only binding to the contract
	T55BridgeTransactor // Write-only binding to the contract
	T55BridgeFilterer   // Log filterer for contract events
}

// T55BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type T55BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T55BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T55BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T55BridgeSession struct {
	Contract     *T55Bridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T55BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T55BridgeCallerSession struct {
	Contract *T55BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// T55BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T55BridgeTransactorSession struct {
	Contract     *T55BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// T55BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type T55BridgeRaw struct {
	Contract *T55Bridge // Generic contract binding to access the raw methods on
}

// T55BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T55BridgeCallerRaw struct {
	Contract *T55BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// T55BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T55BridgeTransactorRaw struct {
	Contract *T55BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT55Bridge creates a new instance of T55Bridge, bound to a specific deployed contract.
func NewT55Bridge(address common.Address, backend bind.ContractBackend) (*T55Bridge, error) {
	contract, err := bindT55Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T55Bridge{T55BridgeCaller: T55BridgeCaller{contract: contract}, T55BridgeTransactor: T55BridgeTransactor{contract: contract}, T55BridgeFilterer: T55BridgeFilterer{contract: contract}}, nil
}

// NewT55BridgeCaller creates a new read-only instance of T55Bridge, bound to a specific deployed contract.
func NewT55BridgeCaller(address common.Address, caller bind.ContractCaller) (*T55BridgeCaller, error) {
	contract, err := bindT55Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T55BridgeCaller{contract: contract}, nil
}

// NewT55BridgeTransactor creates a new write-only instance of T55Bridge, bound to a specific deployed contract.
func NewT55BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*T55BridgeTransactor, error) {
	contract, err := bindT55Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T55BridgeTransactor{contract: contract}, nil
}

// NewT55BridgeFilterer creates a new log filterer instance of T55Bridge, bound to a specific deployed contract.
func NewT55BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*T55BridgeFilterer, error) {
	contract, err := bindT55Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T55BridgeFilterer{contract: contract}, nil
}

// bindT55Bridge binds a generic wrapper to an already deployed contract.
func bindT55Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := T55BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T55Bridge *T55BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T55Bridge.Contract.T55BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T55Bridge *T55BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T55Bridge.Contract.T55BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T55Bridge *T55BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T55Bridge.Contract.T55BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T55Bridge *T55BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T55Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T55Bridge *T55BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T55Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T55Bridge *T55BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T55Bridge.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T55Bridge *T55BridgeCaller) GetBridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _T55Bridge.contract.Call(opts, &out, "getBridgeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T55Bridge *T55BridgeSession) GetBridgeAddr() (common.Address, error) {
	return _T55Bridge.Contract.GetBridgeAddr(&_T55Bridge.CallOpts)
}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T55Bridge *T55BridgeCallerSession) GetBridgeAddr() (common.Address, error) {
	return _T55Bridge.Contract.GetBridgeAddr(&_T55Bridge.CallOpts)
}

// T55 is a free data retrieval call binding the contract method 0xa08867ed.
//
// Solidity: function t55() view returns(address)
func (_T55Bridge *T55BridgeCaller) T55(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _T55Bridge.contract.Call(opts, &out, "t55")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// T55 is a free data retrieval call binding the contract method 0xa08867ed.
//
// Solidity: function t55() view returns(address)
func (_T55Bridge *T55BridgeSession) T55() (common.Address, error) {
	return _T55Bridge.Contract.T55(&_T55Bridge.CallOpts)
}

// T55 is a free data retrieval call binding the contract method 0xa08867ed.
//
// Solidity: function t55() view returns(address)
func (_T55Bridge *T55BridgeCallerSession) T55() (common.Address, error) {
	return _T55Bridge.Contract.T55(&_T55Bridge.CallOpts)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T55Bridge *T55BridgeTransactor) ClaimNFTs(opts *bind.TransactOpts, sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.contract.Transact(opts, "claimNFTs", sender, tokenID)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T55Bridge *T55BridgeSession) ClaimNFTs(sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.ClaimNFTs(&_T55Bridge.TransactOpts, sender, tokenID)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T55Bridge *T55BridgeTransactorSession) ClaimNFTs(sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.ClaimNFTs(&_T55Bridge.TransactOpts, sender, tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T55Bridge *T55BridgeTransactor) Deposit(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.contract.Transact(opts, "deposit", tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T55Bridge *T55BridgeSession) Deposit(tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.Deposit(&_T55Bridge.TransactOpts, tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T55Bridge *T55BridgeTransactorSession) Deposit(tokenID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.Deposit(&_T55Bridge.TransactOpts, tokenID)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T55Bridge *T55BridgeTransactor) SyncNFTs(opts *bind.TransactOpts, sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55Bridge.contract.Transact(opts, "syncNFTs", sender, tokenID, url)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T55Bridge *T55BridgeSession) SyncNFTs(sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55Bridge.Contract.SyncNFTs(&_T55Bridge.TransactOpts, sender, tokenID, url)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T55Bridge *T55BridgeTransactorSession) SyncNFTs(sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55Bridge.Contract.SyncNFTs(&_T55Bridge.TransactOpts, sender, tokenID, url)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T55Bridge *T55BridgeTransactor) Withdraw(opts *bind.TransactOpts, tokendID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.contract.Transact(opts, "withdraw", tokendID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T55Bridge *T55BridgeSession) Withdraw(tokendID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.Withdraw(&_T55Bridge.TransactOpts, tokendID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T55Bridge *T55BridgeTransactorSession) Withdraw(tokendID *big.Int) (*types.Transaction, error) {
	return _T55Bridge.Contract.Withdraw(&_T55Bridge.TransactOpts, tokendID)
}

// T55BridgeClaimNFTsIterator is returned from FilterClaimNFTs and is used to iterate over the raw logs and unpacked data for ClaimNFTs events raised by the T55Bridge contract.
type T55BridgeClaimNFTsIterator struct {
	Event *T55BridgeClaimNFTs // Event containing the contract specifics and raw log

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
func (it *T55BridgeClaimNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55BridgeClaimNFTs)
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
		it.Event = new(T55BridgeClaimNFTs)
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
func (it *T55BridgeClaimNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55BridgeClaimNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55BridgeClaimNFTs represents a ClaimNFTs event raised by the T55Bridge contract.
type T55BridgeClaimNFTs struct {
	To      common.Address
	TokenID *big.Int
	Url     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimNFTs is a free log retrieval operation binding the contract event 0x9226f2bbd0e82cb356d24b49e9d072c1133fb173d419fe1de19f1d77956d031c.
//
// Solidity: event ClaimNFTs(address _to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) FilterClaimNFTs(opts *bind.FilterOpts) (*T55BridgeClaimNFTsIterator, error) {

	logs, sub, err := _T55Bridge.contract.FilterLogs(opts, "ClaimNFTs")
	if err != nil {
		return nil, err
	}
	return &T55BridgeClaimNFTsIterator{contract: _T55Bridge.contract, event: "ClaimNFTs", logs: logs, sub: sub}, nil
}

// WatchClaimNFTs is a free log subscription operation binding the contract event 0x9226f2bbd0e82cb356d24b49e9d072c1133fb173d419fe1de19f1d77956d031c.
//
// Solidity: event ClaimNFTs(address _to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) WatchClaimNFTs(opts *bind.WatchOpts, sink chan<- *T55BridgeClaimNFTs) (event.Subscription, error) {

	logs, sub, err := _T55Bridge.contract.WatchLogs(opts, "ClaimNFTs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55BridgeClaimNFTs)
				if err := _T55Bridge.contract.UnpackLog(event, "ClaimNFTs", log); err != nil {
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

// ParseClaimNFTs is a log parse operation binding the contract event 0x9226f2bbd0e82cb356d24b49e9d072c1133fb173d419fe1de19f1d77956d031c.
//
// Solidity: event ClaimNFTs(address _to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) ParseClaimNFTs(log types.Log) (*T55BridgeClaimNFTs, error) {
	event := new(T55BridgeClaimNFTs)
	if err := _T55Bridge.contract.UnpackLog(event, "ClaimNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55BridgeDepositorIterator is returned from FilterDepositor and is used to iterate over the raw logs and unpacked data for Depositor events raised by the T55Bridge contract.
type T55BridgeDepositorIterator struct {
	Event *T55BridgeDepositor // Event containing the contract specifics and raw log

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
func (it *T55BridgeDepositorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55BridgeDepositor)
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
		it.Event = new(T55BridgeDepositor)
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
func (it *T55BridgeDepositorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55BridgeDepositorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55BridgeDepositor represents a Depositor event raised by the T55Bridge contract.
type T55BridgeDepositor struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Url     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositor is a free log retrieval operation binding the contract event 0x1f976a514c0c2cee422cac8488652958134794e6cd3a9f4ec6ea6e7dad54531d.
//
// Solidity: event Depositor(address from, address to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) FilterDepositor(opts *bind.FilterOpts) (*T55BridgeDepositorIterator, error) {

	logs, sub, err := _T55Bridge.contract.FilterLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return &T55BridgeDepositorIterator{contract: _T55Bridge.contract, event: "Depositor", logs: logs, sub: sub}, nil
}

// WatchDepositor is a free log subscription operation binding the contract event 0x1f976a514c0c2cee422cac8488652958134794e6cd3a9f4ec6ea6e7dad54531d.
//
// Solidity: event Depositor(address from, address to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) WatchDepositor(opts *bind.WatchOpts, sink chan<- *T55BridgeDepositor) (event.Subscription, error) {

	logs, sub, err := _T55Bridge.contract.WatchLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55BridgeDepositor)
				if err := _T55Bridge.contract.UnpackLog(event, "Depositor", log); err != nil {
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

// ParseDepositor is a log parse operation binding the contract event 0x1f976a514c0c2cee422cac8488652958134794e6cd3a9f4ec6ea6e7dad54531d.
//
// Solidity: event Depositor(address from, address to, uint256 tokenID, string url)
func (_T55Bridge *T55BridgeFilterer) ParseDepositor(log types.Log) (*T55BridgeDepositor, error) {
	event := new(T55BridgeDepositor)
	if err := _T55Bridge.contract.UnpackLog(event, "Depositor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55BridgeWithdrawerIterator is returned from FilterWithdrawer and is used to iterate over the raw logs and unpacked data for Withdrawer events raised by the T55Bridge contract.
type T55BridgeWithdrawerIterator struct {
	Event *T55BridgeWithdrawer // Event containing the contract specifics and raw log

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
func (it *T55BridgeWithdrawerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55BridgeWithdrawer)
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
		it.Event = new(T55BridgeWithdrawer)
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
func (it *T55BridgeWithdrawerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55BridgeWithdrawerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55BridgeWithdrawer represents a Withdrawer event raised by the T55Bridge contract.
type T55BridgeWithdrawer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawer is a free log retrieval operation binding the contract event 0x80398927050bfe98d2ff1af291a2a089b8082e6078017f61abc5febd735bec5f.
//
// Solidity: event Withdrawer(address from, address to, uint256 tokenID)
func (_T55Bridge *T55BridgeFilterer) FilterWithdrawer(opts *bind.FilterOpts) (*T55BridgeWithdrawerIterator, error) {

	logs, sub, err := _T55Bridge.contract.FilterLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return &T55BridgeWithdrawerIterator{contract: _T55Bridge.contract, event: "Withdrawer", logs: logs, sub: sub}, nil
}

// WatchWithdrawer is a free log subscription operation binding the contract event 0x80398927050bfe98d2ff1af291a2a089b8082e6078017f61abc5febd735bec5f.
//
// Solidity: event Withdrawer(address from, address to, uint256 tokenID)
func (_T55Bridge *T55BridgeFilterer) WatchWithdrawer(opts *bind.WatchOpts, sink chan<- *T55BridgeWithdrawer) (event.Subscription, error) {

	logs, sub, err := _T55Bridge.contract.WatchLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55BridgeWithdrawer)
				if err := _T55Bridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
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

// ParseWithdrawer is a log parse operation binding the contract event 0x80398927050bfe98d2ff1af291a2a089b8082e6078017f61abc5febd735bec5f.
//
// Solidity: event Withdrawer(address from, address to, uint256 tokenID)
func (_T55Bridge *T55BridgeFilterer) ParseWithdrawer(log types.Log) (*T55BridgeWithdrawer, error) {
	event := new(T55BridgeWithdrawer)
	if err := _T55Bridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
