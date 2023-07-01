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

// T55NFTsMetaData contains all meta data concerning the T55NFTs contract.
var T55NFTsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"MintNFTs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// T55NFTsABI is the input ABI used to generate the binding from.
// Deprecated: Use T55NFTsMetaData.ABI instead.
var T55NFTsABI = T55NFTsMetaData.ABI

// T55NFTs is an auto generated Go binding around an Ethereum contract.
type T55NFTs struct {
	T55NFTsCaller     // Read-only binding to the contract
	T55NFTsTransactor // Write-only binding to the contract
	T55NFTsFilterer   // Log filterer for contract events
}

// T55NFTsCaller is an auto generated read-only Go binding around an Ethereum contract.
type T55NFTsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55NFTsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T55NFTsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55NFTsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T55NFTsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T55NFTsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T55NFTsSession struct {
	Contract     *T55NFTs          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T55NFTsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T55NFTsCallerSession struct {
	Contract *T55NFTsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// T55NFTsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T55NFTsTransactorSession struct {
	Contract     *T55NFTsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// T55NFTsRaw is an auto generated low-level Go binding around an Ethereum contract.
type T55NFTsRaw struct {
	Contract *T55NFTs // Generic contract binding to access the raw methods on
}

// T55NFTsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T55NFTsCallerRaw struct {
	Contract *T55NFTsCaller // Generic read-only contract binding to access the raw methods on
}

// T55NFTsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T55NFTsTransactorRaw struct {
	Contract *T55NFTsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT55NFTs creates a new instance of T55NFTs, bound to a specific deployed contract.
func NewT55NFTs(address common.Address, backend bind.ContractBackend) (*T55NFTs, error) {
	contract, err := bindT55NFTs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T55NFTs{T55NFTsCaller: T55NFTsCaller{contract: contract}, T55NFTsTransactor: T55NFTsTransactor{contract: contract}, T55NFTsFilterer: T55NFTsFilterer{contract: contract}}, nil
}

// NewT55NFTsCaller creates a new read-only instance of T55NFTs, bound to a specific deployed contract.
func NewT55NFTsCaller(address common.Address, caller bind.ContractCaller) (*T55NFTsCaller, error) {
	contract, err := bindT55NFTs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T55NFTsCaller{contract: contract}, nil
}

// NewT55NFTsTransactor creates a new write-only instance of T55NFTs, bound to a specific deployed contract.
func NewT55NFTsTransactor(address common.Address, transactor bind.ContractTransactor) (*T55NFTsTransactor, error) {
	contract, err := bindT55NFTs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T55NFTsTransactor{contract: contract}, nil
}

// NewT55NFTsFilterer creates a new log filterer instance of T55NFTs, bound to a specific deployed contract.
func NewT55NFTsFilterer(address common.Address, filterer bind.ContractFilterer) (*T55NFTsFilterer, error) {
	contract, err := bindT55NFTs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T55NFTsFilterer{contract: contract}, nil
}

// bindT55NFTs binds a generic wrapper to an already deployed contract.
func bindT55NFTs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := T55NFTsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T55NFTs *T55NFTsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T55NFTs.Contract.T55NFTsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T55NFTs *T55NFTsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T55NFTs.Contract.T55NFTsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T55NFTs *T55NFTsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T55NFTs.Contract.T55NFTsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T55NFTs *T55NFTsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _T55NFTs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T55NFTs *T55NFTsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T55NFTs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T55NFTs *T55NFTsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T55NFTs.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_T55NFTs *T55NFTsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_T55NFTs *T55NFTsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _T55NFTs.Contract.BalanceOf(&_T55NFTs.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_T55NFTs *T55NFTsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _T55NFTs.Contract.BalanceOf(&_T55NFTs.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _T55NFTs.Contract.GetApproved(&_T55NFTs.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _T55NFTs.Contract.GetApproved(&_T55NFTs.CallOpts, tokenId)
}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_T55NFTs *T55NFTsCaller) GetTokenID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "getTokenID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_T55NFTs *T55NFTsSession) GetTokenID() (*big.Int, error) {
	return _T55NFTs.Contract.GetTokenID(&_T55NFTs.CallOpts)
}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_T55NFTs *T55NFTsCallerSession) GetTokenID() (*big.Int, error) {
	return _T55NFTs.Contract.GetTokenID(&_T55NFTs.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_T55NFTs *T55NFTsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_T55NFTs *T55NFTsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _T55NFTs.Contract.IsApprovedForAll(&_T55NFTs.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_T55NFTs *T55NFTsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _T55NFTs.Contract.IsApprovedForAll(&_T55NFTs.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_T55NFTs *T55NFTsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_T55NFTs *T55NFTsSession) Name() (string, error) {
	return _T55NFTs.Contract.Name(&_T55NFTs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_T55NFTs *T55NFTsCallerSession) Name() (string, error) {
	return _T55NFTs.Contract.Name(&_T55NFTs.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _T55NFTs.Contract.OwnerOf(&_T55NFTs.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_T55NFTs *T55NFTsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _T55NFTs.Contract.OwnerOf(&_T55NFTs.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_T55NFTs *T55NFTsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_T55NFTs *T55NFTsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _T55NFTs.Contract.SupportsInterface(&_T55NFTs.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_T55NFTs *T55NFTsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _T55NFTs.Contract.SupportsInterface(&_T55NFTs.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_T55NFTs *T55NFTsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_T55NFTs *T55NFTsSession) Symbol() (string, error) {
	return _T55NFTs.Contract.Symbol(&_T55NFTs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_T55NFTs *T55NFTsCallerSession) Symbol() (string, error) {
	return _T55NFTs.Contract.Symbol(&_T55NFTs.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_T55NFTs *T55NFTsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _T55NFTs.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_T55NFTs *T55NFTsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _T55NFTs.Contract.TokenURI(&_T55NFTs.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_T55NFTs *T55NFTsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _T55NFTs.Contract.TokenURI(&_T55NFTs.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Approve(&_T55NFTs.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Approve(&_T55NFTs.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address player, string tokenURI) returns(uint256)
func (_T55NFTs *T55NFTsTransactor) Mint(opts *bind.TransactOpts, player common.Address, tokenURI string) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "mint", player, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address player, string tokenURI) returns(uint256)
func (_T55NFTs *T55NFTsSession) Mint(player common.Address, tokenURI string) (*types.Transaction, error) {
	return _T55NFTs.Contract.Mint(&_T55NFTs.TransactOpts, player, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address player, string tokenURI) returns(uint256)
func (_T55NFTs *T55NFTsTransactorSession) Mint(player common.Address, tokenURI string) (*types.Transaction, error) {
	return _T55NFTs.Contract.Mint(&_T55NFTs.TransactOpts, player, tokenURI)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address _to, uint256 tokenID, string url) returns()
func (_T55NFTs *T55NFTsTransactor) MintTo(opts *bind.TransactOpts, _to common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "mintTo", _to, tokenID, url)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address _to, uint256 tokenID, string url) returns()
func (_T55NFTs *T55NFTsSession) MintTo(_to common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55NFTs.Contract.MintTo(&_T55NFTs.TransactOpts, _to, tokenID, url)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address _to, uint256 tokenID, string url) returns()
func (_T55NFTs *T55NFTsTransactorSession) MintTo(_to common.Address, tokenID *big.Int, url string) (*types.Transaction, error) {
	return _T55NFTs.Contract.MintTo(&_T55NFTs.TransactOpts, _to, tokenID, url)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.SafeTransferFrom(&_T55NFTs.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.SafeTransferFrom(&_T55NFTs.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_T55NFTs *T55NFTsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_T55NFTs *T55NFTsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _T55NFTs.Contract.SafeTransferFrom0(&_T55NFTs.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_T55NFTs *T55NFTsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _T55NFTs.Contract.SafeTransferFrom0(&_T55NFTs.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_T55NFTs *T55NFTsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_T55NFTs *T55NFTsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _T55NFTs.Contract.SetApprovalForAll(&_T55NFTs.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_T55NFTs *T55NFTsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _T55NFTs.Contract.SetApprovalForAll(&_T55NFTs.TransactOpts, operator, approved)
}

// Sync is a paid mutator transaction binding the contract method 0xb1357bf9.
//
// Solidity: function sync(uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactor) Sync(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "sync", tokenId)
}

// Sync is a paid mutator transaction binding the contract method 0xb1357bf9.
//
// Solidity: function sync(uint256 tokenId) returns()
func (_T55NFTs *T55NFTsSession) Sync(tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Sync(&_T55NFTs.TransactOpts, tokenId)
}

// Sync is a paid mutator transaction binding the contract method 0xb1357bf9.
//
// Solidity: function sync(uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactorSession) Sync(tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Sync(&_T55NFTs.TransactOpts, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _tokenID) returns()
func (_T55NFTs *T55NFTsTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "transfer", _from, _to, _tokenID)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _tokenID) returns()
func (_T55NFTs *T55NFTsSession) Transfer(_from common.Address, _to common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Transfer(&_T55NFTs.TransactOpts, _from, _to, _tokenID)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _tokenID) returns()
func (_T55NFTs *T55NFTsTransactorSession) Transfer(_from common.Address, _to common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.Transfer(&_T55NFTs.TransactOpts, _from, _to, _tokenID)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.TransferFrom(&_T55NFTs.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_T55NFTs *T55NFTsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _T55NFTs.Contract.TransferFrom(&_T55NFTs.TransactOpts, from, to, tokenId)
}

// T55NFTsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the T55NFTs contract.
type T55NFTsApprovalIterator struct {
	Event *T55NFTsApproval // Event containing the contract specifics and raw log

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
func (it *T55NFTsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsApproval)
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
		it.Event = new(T55NFTsApproval)
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
func (it *T55NFTsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsApproval represents a Approval event raised by the T55NFTs contract.
type T55NFTsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*T55NFTsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &T55NFTsApprovalIterator{contract: _T55NFTs.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *T55NFTsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsApproval)
				if err := _T55NFTs.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) ParseApproval(log types.Log) (*T55NFTsApproval, error) {
	event := new(T55NFTsApproval)
	if err := _T55NFTs.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55NFTsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the T55NFTs contract.
type T55NFTsApprovalForAllIterator struct {
	Event *T55NFTsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *T55NFTsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsApprovalForAll)
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
		it.Event = new(T55NFTsApprovalForAll)
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
func (it *T55NFTsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsApprovalForAll represents a ApprovalForAll event raised by the T55NFTs contract.
type T55NFTsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_T55NFTs *T55NFTsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*T55NFTsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &T55NFTsApprovalForAllIterator{contract: _T55NFTs.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_T55NFTs *T55NFTsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *T55NFTsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsApprovalForAll)
				if err := _T55NFTs.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_T55NFTs *T55NFTsFilterer) ParseApprovalForAll(log types.Log) (*T55NFTsApprovalForAll, error) {
	event := new(T55NFTsApprovalForAll)
	if err := _T55NFTs.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55NFTsBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the T55NFTs contract.
type T55NFTsBatchMetadataUpdateIterator struct {
	Event *T55NFTsBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *T55NFTsBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsBatchMetadataUpdate)
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
		it.Event = new(T55NFTsBatchMetadataUpdate)
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
func (it *T55NFTsBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the T55NFTs contract.
type T55NFTsBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_T55NFTs *T55NFTsFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*T55NFTsBatchMetadataUpdateIterator, error) {

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &T55NFTsBatchMetadataUpdateIterator{contract: _T55NFTs.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_T55NFTs *T55NFTsFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *T55NFTsBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsBatchMetadataUpdate)
				if err := _T55NFTs.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_T55NFTs *T55NFTsFilterer) ParseBatchMetadataUpdate(log types.Log) (*T55NFTsBatchMetadataUpdate, error) {
	event := new(T55NFTsBatchMetadataUpdate)
	if err := _T55NFTs.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55NFTsMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the T55NFTs contract.
type T55NFTsMetadataUpdateIterator struct {
	Event *T55NFTsMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *T55NFTsMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsMetadataUpdate)
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
		it.Event = new(T55NFTsMetadataUpdate)
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
func (it *T55NFTsMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsMetadataUpdate represents a MetadataUpdate event raised by the T55NFTs contract.
type T55NFTsMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_T55NFTs *T55NFTsFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*T55NFTsMetadataUpdateIterator, error) {

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &T55NFTsMetadataUpdateIterator{contract: _T55NFTs.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_T55NFTs *T55NFTsFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *T55NFTsMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsMetadataUpdate)
				if err := _T55NFTs.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_T55NFTs *T55NFTsFilterer) ParseMetadataUpdate(log types.Log) (*T55NFTsMetadataUpdate, error) {
	event := new(T55NFTsMetadataUpdate)
	if err := _T55NFTs.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55NFTsMintNFTsIterator is returned from FilterMintNFTs and is used to iterate over the raw logs and unpacked data for MintNFTs events raised by the T55NFTs contract.
type T55NFTsMintNFTsIterator struct {
	Event *T55NFTsMintNFTs // Event containing the contract specifics and raw log

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
func (it *T55NFTsMintNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsMintNFTs)
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
		it.Event = new(T55NFTsMintNFTs)
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
func (it *T55NFTsMintNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsMintNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsMintNFTs represents a MintNFTs event raised by the T55NFTs contract.
type T55NFTsMintNFTs struct {
	Player  common.Address
	TokenID *big.Int
	Url     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintNFTs is a free log retrieval operation binding the contract event 0xe3bc396b08fcfcf8518187729864f05e31474d50948b7a8eae8e64b8cfa70884.
//
// Solidity: event MintNFTs(address player, uint256 tokenID, string url)
func (_T55NFTs *T55NFTsFilterer) FilterMintNFTs(opts *bind.FilterOpts) (*T55NFTsMintNFTsIterator, error) {

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "MintNFTs")
	if err != nil {
		return nil, err
	}
	return &T55NFTsMintNFTsIterator{contract: _T55NFTs.contract, event: "MintNFTs", logs: logs, sub: sub}, nil
}

// WatchMintNFTs is a free log subscription operation binding the contract event 0xe3bc396b08fcfcf8518187729864f05e31474d50948b7a8eae8e64b8cfa70884.
//
// Solidity: event MintNFTs(address player, uint256 tokenID, string url)
func (_T55NFTs *T55NFTsFilterer) WatchMintNFTs(opts *bind.WatchOpts, sink chan<- *T55NFTsMintNFTs) (event.Subscription, error) {

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "MintNFTs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsMintNFTs)
				if err := _T55NFTs.contract.UnpackLog(event, "MintNFTs", log); err != nil {
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

// ParseMintNFTs is a log parse operation binding the contract event 0xe3bc396b08fcfcf8518187729864f05e31474d50948b7a8eae8e64b8cfa70884.
//
// Solidity: event MintNFTs(address player, uint256 tokenID, string url)
func (_T55NFTs *T55NFTsFilterer) ParseMintNFTs(log types.Log) (*T55NFTsMintNFTs, error) {
	event := new(T55NFTsMintNFTs)
	if err := _T55NFTs.contract.UnpackLog(event, "MintNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// T55NFTsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the T55NFTs contract.
type T55NFTsTransferIterator struct {
	Event *T55NFTsTransfer // Event containing the contract specifics and raw log

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
func (it *T55NFTsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T55NFTsTransfer)
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
		it.Event = new(T55NFTsTransfer)
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
func (it *T55NFTsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T55NFTsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T55NFTsTransfer represents a Transfer event raised by the T55NFTs contract.
type T55NFTsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*T55NFTsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _T55NFTs.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &T55NFTsTransferIterator{contract: _T55NFTs.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *T55NFTsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _T55NFTs.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T55NFTsTransfer)
				if err := _T55NFTs.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_T55NFTs *T55NFTsFilterer) ParseTransfer(log types.Log) (*T55NFTsTransfer, error) {
	event := new(T55NFTsTransfer)
	if err := _T55NFTs.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
