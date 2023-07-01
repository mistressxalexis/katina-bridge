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

// SentinelsBridgeMetaData contains all meta data concerning the SentinelsBridge contract.
var SentinelsBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deposit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURL\",\"type\":\"string\"}],\"name\":\"Depositor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Withdrawer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"awaitSyncData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"awaitWithdrawEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SentinelsBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use SentinelsBridgeMetaData.ABI instead.
var SentinelsBridgeABI = SentinelsBridgeMetaData.ABI

// SentinelsBridge is an auto generated Go binding around an Ethereum contract.
type SentinelsBridge struct {
	SentinelsBridgeCaller     // Read-only binding to the contract
	SentinelsBridgeTransactor // Write-only binding to the contract
	SentinelsBridgeFilterer   // Log filterer for contract events
}

// SentinelsBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SentinelsBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentinelsBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SentinelsBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentinelsBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SentinelsBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentinelsBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SentinelsBridgeSession struct {
	Contract     *SentinelsBridge  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SentinelsBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SentinelsBridgeCallerSession struct {
	Contract *SentinelsBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SentinelsBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SentinelsBridgeTransactorSession struct {
	Contract     *SentinelsBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SentinelsBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SentinelsBridgeRaw struct {
	Contract *SentinelsBridge // Generic contract binding to access the raw methods on
}

// SentinelsBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SentinelsBridgeCallerRaw struct {
	Contract *SentinelsBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// SentinelsBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SentinelsBridgeTransactorRaw struct {
	Contract *SentinelsBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSentinelsBridge creates a new instance of SentinelsBridge, bound to a specific deployed contract.
func NewSentinelsBridge(address common.Address, backend bind.ContractBackend) (*SentinelsBridge, error) {
	contract, err := bindSentinelsBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SentinelsBridge{SentinelsBridgeCaller: SentinelsBridgeCaller{contract: contract}, SentinelsBridgeTransactor: SentinelsBridgeTransactor{contract: contract}, SentinelsBridgeFilterer: SentinelsBridgeFilterer{contract: contract}}, nil
}

// NewSentinelsBridgeCaller creates a new read-only instance of SentinelsBridge, bound to a specific deployed contract.
func NewSentinelsBridgeCaller(address common.Address, caller bind.ContractCaller) (*SentinelsBridgeCaller, error) {
	contract, err := bindSentinelsBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SentinelsBridgeCaller{contract: contract}, nil
}

// NewSentinelsBridgeTransactor creates a new write-only instance of SentinelsBridge, bound to a specific deployed contract.
func NewSentinelsBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*SentinelsBridgeTransactor, error) {
	contract, err := bindSentinelsBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SentinelsBridgeTransactor{contract: contract}, nil
}

// NewSentinelsBridgeFilterer creates a new log filterer instance of SentinelsBridge, bound to a specific deployed contract.
func NewSentinelsBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*SentinelsBridgeFilterer, error) {
	contract, err := bindSentinelsBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SentinelsBridgeFilterer{contract: contract}, nil
}

// bindSentinelsBridge binds a generic wrapper to an already deployed contract.
func bindSentinelsBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SentinelsBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SentinelsBridge *SentinelsBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SentinelsBridge.Contract.SentinelsBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SentinelsBridge *SentinelsBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.SentinelsBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SentinelsBridge *SentinelsBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.SentinelsBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SentinelsBridge *SentinelsBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SentinelsBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SentinelsBridge *SentinelsBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SentinelsBridge *SentinelsBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.contract.Transact(opts, method, params...)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _SentinelsBridge.Contract.GetEthSignedMessageHash(&_SentinelsBridge.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _SentinelsBridge.Contract.GetEthSignedMessageHash(&_SentinelsBridge.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeCaller) GetMessageHash(opts *bind.CallOpts, _message string) ([32]byte, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "getMessageHash", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeSession) GetMessageHash(_message string) ([32]byte, error) {
	return _SentinelsBridge.Contract.GetMessageHash(&_SentinelsBridge.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_SentinelsBridge *SentinelsBridgeCallerSession) GetMessageHash(_message string) ([32]byte, error) {
	return _SentinelsBridge.Contract.GetMessageHash(&_SentinelsBridge.CallOpts, _message)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_SentinelsBridge *SentinelsBridgeCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_SentinelsBridge *SentinelsBridgeSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SentinelsBridge.Contract.RecoverSigner(&_SentinelsBridge.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_SentinelsBridge *SentinelsBridgeCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SentinelsBridge.Contract.RecoverSigner(&_SentinelsBridge.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_SentinelsBridge *SentinelsBridgeCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_SentinelsBridge *SentinelsBridgeSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _SentinelsBridge.Contract.SplitSignature(&_SentinelsBridge.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_SentinelsBridge *SentinelsBridgeCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _SentinelsBridge.Contract.SplitSignature(&_SentinelsBridge.CallOpts, sig)
}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes signature) pure returns(bool)
func (_SentinelsBridge *SentinelsBridgeCaller) Verify(opts *bind.CallOpts, _signer common.Address, _message string, signature []byte) (bool, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "verify", _signer, _message, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes signature) pure returns(bool)
func (_SentinelsBridge *SentinelsBridgeSession) Verify(_signer common.Address, _message string, signature []byte) (bool, error) {
	return _SentinelsBridge.Contract.Verify(&_SentinelsBridge.CallOpts, _signer, _message, signature)
}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes signature) pure returns(bool)
func (_SentinelsBridge *SentinelsBridgeCallerSession) Verify(_signer common.Address, _message string, signature []byte) (bool, error) {
	return _SentinelsBridge.Contract.Verify(&_SentinelsBridge.CallOpts, _signer, _message, signature)
}

// VerifySignature is a free data retrieval call binding the contract method 0x23e470ac.
//
// Solidity: function verifySignature(string message, bytes32 r, bytes32 s, uint8 v) view returns(bool)
func (_SentinelsBridge *SentinelsBridgeCaller) VerifySignature(opts *bind.CallOpts, message string, r [32]byte, s [32]byte, v uint8) (bool, error) {
	var out []interface{}
	err := _SentinelsBridge.contract.Call(opts, &out, "verifySignature", message, r, s, v)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0x23e470ac.
//
// Solidity: function verifySignature(string message, bytes32 r, bytes32 s, uint8 v) view returns(bool)
func (_SentinelsBridge *SentinelsBridgeSession) VerifySignature(message string, r [32]byte, s [32]byte, v uint8) (bool, error) {
	return _SentinelsBridge.Contract.VerifySignature(&_SentinelsBridge.CallOpts, message, r, s, v)
}

// VerifySignature is a free data retrieval call binding the contract method 0x23e470ac.
//
// Solidity: function verifySignature(string message, bytes32 r, bytes32 s, uint8 v) view returns(bool)
func (_SentinelsBridge *SentinelsBridgeCallerSession) VerifySignature(message string, r [32]byte, s [32]byte, v uint8) (bool, error) {
	return _SentinelsBridge.Contract.VerifySignature(&_SentinelsBridge.CallOpts, message, r, s, v)
}

// AwaitSyncData is a paid mutator transaction binding the contract method 0xdb13d37f.
//
// Solidity: function awaitSyncData(bytes signature, uint256 tokenId, string url) returns()
func (_SentinelsBridge *SentinelsBridgeTransactor) AwaitSyncData(opts *bind.TransactOpts, signature []byte, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _SentinelsBridge.contract.Transact(opts, "awaitSyncData", signature, tokenId, url)
}

// AwaitSyncData is a paid mutator transaction binding the contract method 0xdb13d37f.
//
// Solidity: function awaitSyncData(bytes signature, uint256 tokenId, string url) returns()
func (_SentinelsBridge *SentinelsBridgeSession) AwaitSyncData(signature []byte, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.AwaitSyncData(&_SentinelsBridge.TransactOpts, signature, tokenId, url)
}

// AwaitSyncData is a paid mutator transaction binding the contract method 0xdb13d37f.
//
// Solidity: function awaitSyncData(bytes signature, uint256 tokenId, string url) returns()
func (_SentinelsBridge *SentinelsBridgeTransactorSession) AwaitSyncData(signature []byte, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.AwaitSyncData(&_SentinelsBridge.TransactOpts, signature, tokenId, url)
}

// AwaitWithdrawEvent is a paid mutator transaction binding the contract method 0xf487da4c.
//
// Solidity: function awaitWithdrawEvent(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeTransactor) AwaitWithdrawEvent(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.contract.Transact(opts, "awaitWithdrawEvent", tokenId)
}

// AwaitWithdrawEvent is a paid mutator transaction binding the contract method 0xf487da4c.
//
// Solidity: function awaitWithdrawEvent(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeSession) AwaitWithdrawEvent(tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.AwaitWithdrawEvent(&_SentinelsBridge.TransactOpts, tokenId)
}

// AwaitWithdrawEvent is a paid mutator transaction binding the contract method 0xf487da4c.
//
// Solidity: function awaitWithdrawEvent(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeTransactorSession) AwaitWithdrawEvent(tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.AwaitWithdrawEvent(&_SentinelsBridge.TransactOpts, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeTransactor) Deposit(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.contract.Transact(opts, "deposit", tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeSession) Deposit(tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.Deposit(&_SentinelsBridge.TransactOpts, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 tokenId) returns()
func (_SentinelsBridge *SentinelsBridgeTransactorSession) Deposit(tokenId *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.Deposit(&_SentinelsBridge.TransactOpts, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_SentinelsBridge *SentinelsBridgeTransactor) Withdraw(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.contract.Transact(opts, "withdraw", tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_SentinelsBridge *SentinelsBridgeSession) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.Withdraw(&_SentinelsBridge.TransactOpts, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_SentinelsBridge *SentinelsBridgeTransactorSession) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _SentinelsBridge.Contract.Withdraw(&_SentinelsBridge.TransactOpts, tokenID)
}

// SentinelsBridgeDepositorIterator is returned from FilterDepositor and is used to iterate over the raw logs and unpacked data for Depositor events raised by the SentinelsBridge contract.
type SentinelsBridgeDepositorIterator struct {
	Event *SentinelsBridgeDepositor // Event containing the contract specifics and raw log

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
func (it *SentinelsBridgeDepositorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentinelsBridgeDepositor)
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
		it.Event = new(SentinelsBridgeDepositor)
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
func (it *SentinelsBridgeDepositorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentinelsBridgeDepositorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentinelsBridgeDepositor represents a Depositor event raised by the SentinelsBridge contract.
type SentinelsBridgeDepositor struct {
	Deposit  common.Address
	TokenID  *big.Int
	TokenURL string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositor is a free log retrieval operation binding the contract event 0xf10b860f691c47b47ec33530a9bf1a4f632ab3dc1f62dbba2cccfa3ecdb1e31e.
//
// Solidity: event Depositor(address deposit, uint256 tokenID, string tokenURL)
func (_SentinelsBridge *SentinelsBridgeFilterer) FilterDepositor(opts *bind.FilterOpts) (*SentinelsBridgeDepositorIterator, error) {

	logs, sub, err := _SentinelsBridge.contract.FilterLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return &SentinelsBridgeDepositorIterator{contract: _SentinelsBridge.contract, event: "Depositor", logs: logs, sub: sub}, nil
}

// WatchDepositor is a free log subscription operation binding the contract event 0xf10b860f691c47b47ec33530a9bf1a4f632ab3dc1f62dbba2cccfa3ecdb1e31e.
//
// Solidity: event Depositor(address deposit, uint256 tokenID, string tokenURL)
func (_SentinelsBridge *SentinelsBridgeFilterer) WatchDepositor(opts *bind.WatchOpts, sink chan<- *SentinelsBridgeDepositor) (event.Subscription, error) {

	logs, sub, err := _SentinelsBridge.contract.WatchLogs(opts, "Depositor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentinelsBridgeDepositor)
				if err := _SentinelsBridge.contract.UnpackLog(event, "Depositor", log); err != nil {
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
func (_SentinelsBridge *SentinelsBridgeFilterer) ParseDepositor(log types.Log) (*SentinelsBridgeDepositor, error) {
	event := new(SentinelsBridgeDepositor)
	if err := _SentinelsBridge.contract.UnpackLog(event, "Depositor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SentinelsBridgeWithdrawerIterator is returned from FilterWithdrawer and is used to iterate over the raw logs and unpacked data for Withdrawer events raised by the SentinelsBridge contract.
type SentinelsBridgeWithdrawerIterator struct {
	Event *SentinelsBridgeWithdrawer // Event containing the contract specifics and raw log

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
func (it *SentinelsBridgeWithdrawerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentinelsBridgeWithdrawer)
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
		it.Event = new(SentinelsBridgeWithdrawer)
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
func (it *SentinelsBridgeWithdrawerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentinelsBridgeWithdrawerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentinelsBridgeWithdrawer represents a Withdrawer event raised by the SentinelsBridge contract.
type SentinelsBridgeWithdrawer struct {
	Withdrawer common.Address
	TokenID    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawer is a free log retrieval operation binding the contract event 0x8cb11afff86f7ff1d42446d5e53c4be953f1f2cfb679cc213c66ee01a984123b.
//
// Solidity: event Withdrawer(address withdrawer, uint256 tokenID)
func (_SentinelsBridge *SentinelsBridgeFilterer) FilterWithdrawer(opts *bind.FilterOpts) (*SentinelsBridgeWithdrawerIterator, error) {

	logs, sub, err := _SentinelsBridge.contract.FilterLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return &SentinelsBridgeWithdrawerIterator{contract: _SentinelsBridge.contract, event: "Withdrawer", logs: logs, sub: sub}, nil
}

// WatchWithdrawer is a free log subscription operation binding the contract event 0x8cb11afff86f7ff1d42446d5e53c4be953f1f2cfb679cc213c66ee01a984123b.
//
// Solidity: event Withdrawer(address withdrawer, uint256 tokenID)
func (_SentinelsBridge *SentinelsBridgeFilterer) WatchWithdrawer(opts *bind.WatchOpts, sink chan<- *SentinelsBridgeWithdrawer) (event.Subscription, error) {

	logs, sub, err := _SentinelsBridge.contract.WatchLogs(opts, "Withdrawer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentinelsBridgeWithdrawer)
				if err := _SentinelsBridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
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
func (_SentinelsBridge *SentinelsBridgeFilterer) ParseWithdrawer(log types.Log) (*SentinelsBridgeWithdrawer, error) {
	event := new(SentinelsBridgeWithdrawer)
	if err := _SentinelsBridge.contract.UnpackLog(event, "Withdrawer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
