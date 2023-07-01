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

// T54BridgeMetaData contains all meta data concerning the T54Bridge contract.
var T54BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nfts\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"ClaimNFTs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"Depositor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Withdrawer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"t54\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"claimNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"syncNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"t54\",\"outputs\":[{\"internalType\":\"contractT54NFTs\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokendID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// T54BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use T54BridgeMetaData.ABI instead.
var T54BridgeABI = T54BridgeMetaData.ABI

// T54Bridge is an auto generated Go binding around an Ethereum contract.
type T54Bridge struct {
	T54BridgeCaller     // Read-only binding to the contract
	T54BridgeTransactor // Write-only binding to the contract
	T54BridgeFilterer   // Log filterer for contract events
}

// T54BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type T54BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T54BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T54BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T54BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T54BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T54BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T54BridgeSession struct {
	Contract     *T54Bridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T54BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T54BridgeCallerSession struct {
	Contract *T54BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// T54BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T54BridgeTransactorSession struct {
	Contract     *T54BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// T54BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type T54BridgeRaw struct {
	Contract *T54Bridge // Generic contract binding to access the raw methods on
}

// T54BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T54BridgeCallerRaw struct {
	Contract *T54BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// T54BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T54BridgeTransactorRaw struct {
	Contract *T54BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT54Bridge creates a new instance of T54Bridge, bound to a specific deployed contract.
func NewT54Bridge(address common.Address, backend bind.ContractBackend) (*T54Bridge, error) {
	contract, err := bindT54Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T54Bridge{T54BridgeCaller: T54BridgeCaller{contract: contract}, T54BridgeTransactor: T54BridgeTransactor{contract: contract}, T54BridgeFilterer: T54BridgeFilterer{contract: contract}}, nil
}

// NewT54BridgeCaller creates a new read-only instance of T54Bridge, bound to a specific deployed contract.
func NewT54BridgeCaller(address common.Address, caller bind.ContractCaller) (*T54BridgeCaller, error) {
	contract, err := bindT54Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T54BridgeCaller{contract: contract}, nil
}

// NewT54BridgeTransactor creates a new write-only instance of T54Bridge, bound to a specific deployed contract.
func NewT54BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*T54BridgeTransactor, error) {
	contract, err := bindT54Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T54BridgeTransactor{contract: contract}, nil
}

// NewT54BridgeFilterer creates a new log filterer instance of T54Bridge, bound to a specific deployed contract.
func NewT54BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*T54BridgeFilterer, error) {
	contract, err := bindT54Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T54BridgeFilterer{contract: contract}, nil
}

// bindT54Bridge binds a generic wrapper to an already deployed contract.
func bindT54Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := T54BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T54Bridge *T54BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T54Bridge.Contract.T54BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T54Bridge *T54BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T54Bridge.Contract.T54BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T54Bridge *T54BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T54Bridge.Contract.T54BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T54Bridge *T54BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T54Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T54Bridge *T54BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T54Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T54Bridge *T54BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T54Bridge.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(uint256 t54, bool exists)
func (_T54Bridge *T54BridgeCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	T54    *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _T54Bridge.contract.Call(opts, &out, "accounts", arg0)

	outstruct := new(struct {
		T54    *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.T54 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(uint256 t54, bool exists)
func (_T54Bridge *T54BridgeSession) Accounts(arg0 common.Address) (struct {
	T54    *big.Int
	Exists bool
}, error) {
	return _T54Bridge.Contract.Accounts(&_T54Bridge.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(uint256 t54, bool exists)
func (_T54Bridge *T54BridgeCallerSession) Accounts(arg0 common.Address) (struct {
	T54    *big.Int
	Exists bool
}, error) {
	return _T54Bridge.Contract.Accounts(&_T54Bridge.CallOpts, arg0)
}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T54Bridge *T54BridgeCaller) GetBridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _T54Bridge.contract.Call(opts, &out, "getBridgeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T54Bridge *T54BridgeSession) GetBridgeAddr() (common.Address, error) {
	return _T54Bridge.Contract.GetBridgeAddr(&_T54Bridge.CallOpts)
}

// GetBridgeAddr is a free data retrieval call binding the contract method 0x6e0989d8.
//
// Solidity: function getBridgeAddr() view returns(address)
func (_T54Bridge *T54BridgeCallerSession) GetBridgeAddr() (common.Address, error) {
	return _T54Bridge.Contract.GetBridgeAddr(&_T54Bridge.CallOpts)
}

// T54 is a free data retrieval call binding the contract method 0x5ee6f34d.
//
// Solidity: function t54() view returns(address)
func (_T54Bridge *T54BridgeCaller) T54(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _T54Bridge.contract.Call(opts, &out, "t54")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// T54 is a free data retrieval call binding the contract method 0x5ee6f34d.
//
// Solidity: function t54() view returns(address)
func (_T54Bridge *T54BridgeSession) T54() (common.Address, error) {
	return _T54Bridge.Contract.T54(&_T54Bridge.CallOpts)
}

// T54 is a free data retrieval call binding the contract method 0x5ee6f34d.
//
// Solidity: function t54() view returns(address)
func (_T54Bridge *T54BridgeCallerSession) T54() (common.Address, error) {
	return _T54Bridge.Contract.T54(&_T54Bridge.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address _to, uint256 _tokenID, string url) returns()
func (_T54Bridge *T54BridgeTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.contract.Transact(opts, "claim", _to, _tokenID, url)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address _to, uint256 _tokenID, string url) returns()
func (_T54Bridge *T54BridgeSession) Claim(_to common.Address, _tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.Contract.Claim(&_T54Bridge.TransactOpts, _to, _tokenID, url)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address _to, uint256 _tokenID, string url) returns()
func (_T54Bridge *T54BridgeTransactorSession) Claim(_to common.Address, _tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.Contract.Claim(&_T54Bridge.TransactOpts, _to, _tokenID, url)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T54Bridge *T54BridgeTransactor) ClaimNFTs(opts *bind.TransactOpts, sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.contract.Transact(opts, "claimNFTs", sender, tokenID)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T54Bridge *T54BridgeSession) ClaimNFTs(sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.ClaimNFTs(&_T54Bridge.TransactOpts, sender, tokenID)
}

// ClaimNFTs is a paid mutator transaction binding the contract method 0x972668d8.
//
// Solidity: function claimNFTs(address sender, uint256 tokenID) returns()
func (_T54Bridge *T54BridgeTransactorSession) ClaimNFTs(sender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.ClaimNFTs(&_T54Bridge.TransactOpts, sender, tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T54Bridge *T54BridgeTransactor) Deposit(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.contract.Transact(opts, "deposit", tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T54Bridge *T54BridgeSession) Deposit(tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.Deposit(&_T54Bridge.TransactOpts, tokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenID) payable returns()
func (_T54Bridge *T54BridgeTransactorSession) Deposit(tokenID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.Deposit(&_T54Bridge.TransactOpts, tokenID)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T54Bridge *T54BridgeTransactor) SyncNFTs(opts *bind.TransactOpts, sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.contract.Transact(opts, "syncNFTs", sender, tokenID, url)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T54Bridge *T54BridgeSession) SyncNFTs(sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.Contract.SyncNFTs(&_T54Bridge.TransactOpts, sender, tokenID, url)
}

// SyncNFTs is a paid mutator transaction binding the contract method 0x54a180df.
//
// Solidity: function syncNFTs(address sender, uint256 tokenID, string url) returns()
func (_T54Bridge *T54BridgeTransactorSession) SyncNFTs(sender common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T54Bridge.Contract.SyncNFTs(&_T54Bridge.TransactOpts, sender, tokenID, url)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T54Bridge *T54BridgeTransactor) Withdraw(opts *bind.TransactOpts, tokendID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.contract.Transact(opts, "withdraw", tokendID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T54Bridge *T54BridgeSession) Withdraw(tokendID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.Withdraw(&_T54Bridge.TransactOpts, tokendID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokendID) returns()
func (_T54Bridge *T54BridgeTransactorSession) Withdraw(tokendID *big.Int) (*types.Transaction, error) {
	return _T54Bridge.Contract.Withdraw(&_T54Bridge.TransactOpts, tokendID)
}

// T54BridgeClaimNFTsIterator is returned from FilterClaimNFTs and is used to iterate over the raw logs and unpacked data for ClaimNFTs events raised by the T54Bridge contract.
type T54BridgeClaimNFTsIterator struct {
	Event *T54BridgeClaimNFTs // Event containing the contract specifics and raw log

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
func (it *T54BridgeClaimNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T54BridgeClaimNFTs)
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
		it.Event = new(T54BridgeClaimNFTs)
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
func (it *T54BridgeClaimNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T54BridgeClaimNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T54BridgeClaimNFTs represents a ClaimNFTs event raised by the T54Bridge contract.
type T54BridgeClaimNFTs struct {
	To      common.Address
	TokenID *big.Int
	Url     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimNFTs is a free log retrieval operation binding the contract event 0x9226f2bbd0e82cb356d24b49e9d072c1133fb173d419fe1de19f1d77956d031c.
//
// Solidity: event ClaimNFTs(address _to, uint256 tokenID, string url)
func (_T54Bridge *T54BridgeFilterer) FilterClaimNFTs(opts *bind.FilterOpts) (*T54BridgeClaimNFTsIterator, error) {

	logs, sub, err := _T54Bridge.contract.FilterLogs(opts, "ClaimNFTs")
	if err != nil {
		return nil, err
	}
	return &T54BridgeClaimNFTsIterator{contract: _T54Bridge.contract, event: "ClaimNFTs", logs: logs, sub: sub}, nil
}

// WatchClaimNFTs is a free log subscription operation binding the contract event 0x9226f2bbd0e82cb356d24b49e9d072c1133fb173d419fe1de19f1d77956d031c.
//
// Solidity: event ClaimNFTs(address _to, uint256 tokenID, string url)
func (_T54Bridge *T54BridgeFilterer) WatchClaimNFTs(opts *bind.WatchOpts, sink chan<- *T54BridgeClaimNFTs) (event.Subscription, error) {

	logs, sub, err := _T54Bridge.contract.WatchLogs(opts, "ClaimNFTs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T54BridgeClaimNFTs)
				if err := _T54Bridge.contract.UnpackLog(event, "ClaimNFTs", log); err != nil {
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
func (_T54Bridge *T54BridgeFilterer) ParseClaimNFTs(log types.Log) (*T54BridgeClaimNFTs, error) {
	event := new(T54BridgeClaimNFTs)
	if err := _T54Bridge.contract.UnpackLog(event, "ClaimNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T54BridgeDepositorIterator is returned from FilterDepositor and is used to iterate over the raw logs and unpacked data for Depositor events raised by the T54Bridge contract.
type T54BridgeDepositorIterator struct {
	Event *T54BridgeDepositor // Event containing the contract specifics and raw log

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
func (it *T54BridgeDepositorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T54BridgeDepositor)
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
		it.Event = new(T54BridgeDepositor)
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
func (it *T54BridgeDepositorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T54BridgeDepositorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T54BridgeDepositor represents a Depositor event raised by the T54Bridge contract.
type T54BridgeDepositor struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Url     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositor is a free log retrieval operation binding the contract event 0x1f976a514c0c2cee422cac8488652958134794e6cd3a9f4ec6ea6e7dad54531d.
//
// Solidity: event Depositor(address from, address to, uint256 tokenID, string url)
func (_T54Bridge *T54BridgeFilterer) FilterDepositor(opts *bind.FilterOpts) (*T54BridgeDepositorIterator, error) {

	logs, sub, err := _T54Bridge.contract.FilterLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return &T54BridgeDepositorIterator{contract: _T54Bridge.contract, event: "Depositor", logs: logs, sub: sub}, nil
}

// WatchDepositor is a free log subscription operation binding the contract event 0x1f976a514c0c2cee422cac8488652958134794e6cd3a9f4ec6ea6e7dad54531d.
//
// Solidity: event Depositor(address from, address to, uint256 tokenID, string url)
func (_T54Bridge *T54BridgeFilterer) WatchDepositor(opts *bind.WatchOpts, sink chan<- *T54BridgeDepositor) (event.Subscription, error) {

	logs, sub, err := _T54Bridge.contract.WatchLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T54BridgeDepositor)
				if err := _T54Bridge.contract.UnpackLog(event, "Depositor", log); err != nil {
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
func (_T54Bridge *T54BridgeFilterer) ParseDepositor(log types.Log) (*T54BridgeDepositor, error) {
	event := new(T54BridgeDepositor)
	if err := _T54Bridge.contract.UnpackLog(event, "Depositor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T54BridgeWithdrawerIterator is returned from FilterWithdrawer and is used to iterate over the raw logs and unpacked data for Withdrawer events raised by the T54Bridge contract.
type T54BridgeWithdrawerIterator struct {
	Event *T54BridgeWithdrawer // Event containing the contract specifics and raw log

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
func (it *T54BridgeWithdrawerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T54BridgeWithdrawer)
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
		it.Event = new(T54BridgeWithdrawer)
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
func (it *T54BridgeWithdrawerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T54BridgeWithdrawerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T54BridgeWithdrawer represents a Withdrawer event raised by the T54Bridge contract.
type T54BridgeWithdrawer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawer is a free log retrieval operation binding the contract event 0x80398927050bfe98d2ff1af291a2a089b8082e6078017f61abc5febd735bec5f.
//
// Solidity: event Withdrawer(address from, address to, uint256 tokenID)
func (_T54Bridge *T54BridgeFilterer) FilterWithdrawer(opts *bind.FilterOpts) (*T54BridgeWithdrawerIterator, error) {

	logs, sub, err := _T54Bridge.contract.FilterLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return &T54BridgeWithdrawerIterator{contract: _T54Bridge.contract, event: "Withdrawer", logs: logs, sub: sub}, nil
}

// WatchWithdrawer is a free log subscription operation binding the contract event 0x80398927050bfe98d2ff1af291a2a089b8082e6078017f61abc5febd735bec5f.
//
// Solidity: event Withdrawer(address from, address to, uint256 tokenID)
func (_T54Bridge *T54BridgeFilterer) WatchWithdrawer(opts *bind.WatchOpts, sink chan<- *T54BridgeWithdrawer) (event.Subscription, error) {

	logs, sub, err := _T54Bridge.contract.WatchLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T54BridgeWithdrawer)
				if err := _T54Bridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
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
func (_T54Bridge *T54BridgeFilterer) ParseWithdrawer(log types.Log) (*T54BridgeWithdrawer, error) {
	event := new(T54BridgeWithdrawer)
	if err := _T54Bridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
