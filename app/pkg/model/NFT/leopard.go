// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// LeopardMetaData contains all meta data concerning the Leopard contract.
var LeopardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"mintItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"mintProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"transferItemTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LeopardABI is the input ABI used to generate the binding from.
// Deprecated: Use LeopardMetaData.ABI instead.
var LeopardABI = LeopardMetaData.ABI

// Leopard is an auto generated Go binding around an Ethereum contract.
type Leopard struct {
	LeopardCaller     // Read-only binding to the contract
	LeopardTransactor // Write-only binding to the contract
	LeopardFilterer   // Log filterer for contract events
}

// LeopardCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeopardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeopardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeopardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeopardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeopardSession struct {
	Contract     *Leopard          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeopardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeopardCallerSession struct {
	Contract *LeopardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LeopardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeopardTransactorSession struct {
	Contract     *LeopardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LeopardRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeopardRaw struct {
	Contract *Leopard // Generic contract binding to access the raw methods on
}

// LeopardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeopardCallerRaw struct {
	Contract *LeopardCaller // Generic read-only contract binding to access the raw methods on
}

// LeopardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeopardTransactorRaw struct {
	Contract *LeopardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeopard creates a new instance of Leopard, bound to a specific deployed contract.
func NewLeopard(address common.Address, backend bind.ContractBackend) (*Leopard, error) {
	contract, err := bindLeopard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Leopard{LeopardCaller: LeopardCaller{contract: contract}, LeopardTransactor: LeopardTransactor{contract: contract}, LeopardFilterer: LeopardFilterer{contract: contract}}, nil
}

// NewLeopardCaller creates a new read-only instance of Leopard, bound to a specific deployed contract.
func NewLeopardCaller(address common.Address, caller bind.ContractCaller) (*LeopardCaller, error) {
	contract, err := bindLeopard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeopardCaller{contract: contract}, nil
}

// NewLeopardTransactor creates a new write-only instance of Leopard, bound to a specific deployed contract.
func NewLeopardTransactor(address common.Address, transactor bind.ContractTransactor) (*LeopardTransactor, error) {
	contract, err := bindLeopard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeopardTransactor{contract: contract}, nil
}

// NewLeopardFilterer creates a new log filterer instance of Leopard, bound to a specific deployed contract.
func NewLeopardFilterer(address common.Address, filterer bind.ContractFilterer) (*LeopardFilterer, error) {
	contract, err := bindLeopard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeopardFilterer{contract: contract}, nil
}

// bindLeopard binds a generic wrapper to an already deployed contract.
func bindLeopard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeopardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leopard *LeopardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leopard.Contract.LeopardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leopard *LeopardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leopard.Contract.LeopardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leopard *LeopardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leopard.Contract.LeopardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leopard *LeopardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leopard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leopard *LeopardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leopard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leopard *LeopardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leopard.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leopard *LeopardCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leopard *LeopardSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Leopard.Contract.BalanceOf(&_Leopard.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leopard *LeopardCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Leopard.Contract.BalanceOf(&_Leopard.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leopard *LeopardCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leopard *LeopardSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Leopard.Contract.GetApproved(&_Leopard.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leopard *LeopardCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Leopard.Contract.GetApproved(&_Leopard.CallOpts, tokenId)
}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_Leopard *LeopardCaller) GetTokenID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "getTokenID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_Leopard *LeopardSession) GetTokenID() (*big.Int, error) {
	return _Leopard.Contract.GetTokenID(&_Leopard.CallOpts)
}

// GetTokenID is a free data retrieval call binding the contract method 0x6ba367a3.
//
// Solidity: function getTokenID() view returns(uint256)
func (_Leopard *LeopardCallerSession) GetTokenID() (*big.Int, error) {
	return _Leopard.Contract.GetTokenID(&_Leopard.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leopard *LeopardCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leopard *LeopardSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Leopard.Contract.IsApprovedForAll(&_Leopard.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leopard *LeopardCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Leopard.Contract.IsApprovedForAll(&_Leopard.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leopard *LeopardCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leopard *LeopardSession) Name() (string, error) {
	return _Leopard.Contract.Name(&_Leopard.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leopard *LeopardCallerSession) Name() (string, error) {
	return _Leopard.Contract.Name(&_Leopard.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leopard *LeopardCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leopard *LeopardSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Leopard.Contract.OwnerOf(&_Leopard.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leopard *LeopardCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Leopard.Contract.OwnerOf(&_Leopard.CallOpts, tokenId)
}

// Signature is a free data retrieval call binding the contract method 0x851f7a74.
//
// Solidity: function signature(uint256 tokenId, string url) view returns(bytes32)
func (_Leopard *LeopardCaller) Signature(opts *bind.CallOpts, tokenId *big.Int, url string) ([32]byte, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "signature", tokenId, url)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Signature is a free data retrieval call binding the contract method 0x851f7a74.
//
// Solidity: function signature(uint256 tokenId, string url) view returns(bytes32)
func (_Leopard *LeopardSession) Signature(tokenId *big.Int, url string) ([32]byte, error) {
	return _Leopard.Contract.Signature(&_Leopard.CallOpts, tokenId, url)
}

// Signature is a free data retrieval call binding the contract method 0x851f7a74.
//
// Solidity: function signature(uint256 tokenId, string url) view returns(bytes32)
func (_Leopard *LeopardCallerSession) Signature(tokenId *big.Int, url string) ([32]byte, error) {
	return _Leopard.Contract.Signature(&_Leopard.CallOpts, tokenId, url)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leopard *LeopardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leopard *LeopardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Leopard.Contract.SupportsInterface(&_Leopard.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leopard *LeopardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Leopard.Contract.SupportsInterface(&_Leopard.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leopard *LeopardCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leopard *LeopardSession) Symbol() (string, error) {
	return _Leopard.Contract.Symbol(&_Leopard.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leopard *LeopardCallerSession) Symbol() (string, error) {
	return _Leopard.Contract.Symbol(&_Leopard.CallOpts)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_Leopard *LeopardCaller) TokenExists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "tokenExists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_Leopard *LeopardSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _Leopard.Contract.TokenExists(&_Leopard.CallOpts, tokenId)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_Leopard *LeopardCallerSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _Leopard.Contract.TokenExists(&_Leopard.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leopard *LeopardCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Leopard.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leopard *LeopardSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Leopard.Contract.TokenURI(&_Leopard.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leopard *LeopardCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Leopard.Contract.TokenURI(&_Leopard.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Leopard *LeopardSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.Approve(&_Leopard.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.Approve(&_Leopard.TransactOpts, to, tokenId)
}

// MintItem is a paid mutator transaction binding the contract method 0xc7545849.
//
// Solidity: function mintItem(string url) returns()
func (_Leopard *LeopardTransactor) MintItem(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "mintItem", url)
}

// MintItem is a paid mutator transaction binding the contract method 0xc7545849.
//
// Solidity: function mintItem(string url) returns()
func (_Leopard *LeopardSession) MintItem(url string) (*types.Transaction, error) {
	return _Leopard.Contract.MintItem(&_Leopard.TransactOpts, url)
}

// MintItem is a paid mutator transaction binding the contract method 0xc7545849.
//
// Solidity: function mintItem(string url) returns()
func (_Leopard *LeopardTransactorSession) MintItem(url string) (*types.Transaction, error) {
	return _Leopard.Contract.MintItem(&_Leopard.TransactOpts, url)
}

// MintProtocol is a paid mutator transaction binding the contract method 0x7e96fd64.
//
// Solidity: function mintProtocol(address _to, uint256 tokenId, string url) returns()
func (_Leopard *LeopardTransactor) MintProtocol(opts *bind.TransactOpts, _to common.Address, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "mintProtocol", _to, tokenId, url)
}

// MintProtocol is a paid mutator transaction binding the contract method 0x7e96fd64.
//
// Solidity: function mintProtocol(address _to, uint256 tokenId, string url) returns()
func (_Leopard *LeopardSession) MintProtocol(_to common.Address, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _Leopard.Contract.MintProtocol(&_Leopard.TransactOpts, _to, tokenId, url)
}

// MintProtocol is a paid mutator transaction binding the contract method 0x7e96fd64.
//
// Solidity: function mintProtocol(address _to, uint256 tokenId, string url) returns()
func (_Leopard *LeopardTransactorSession) MintProtocol(_to common.Address, tokenId *big.Int, url string) (*types.Transaction, error) {
	return _Leopard.Contract.MintProtocol(&_Leopard.TransactOpts, _to, tokenId, url)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.SafeTransferFrom(&_Leopard.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.SafeTransferFrom(&_Leopard.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Leopard *LeopardTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Leopard *LeopardSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Leopard.Contract.SafeTransferFrom0(&_Leopard.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Leopard *LeopardTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Leopard.Contract.SafeTransferFrom0(&_Leopard.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leopard *LeopardTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leopard *LeopardSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leopard.Contract.SetApprovalForAll(&_Leopard.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leopard *LeopardTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leopard.Contract.SetApprovalForAll(&_Leopard.TransactOpts, operator, approved)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_Leopard *LeopardTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_Leopard *LeopardSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _Leopard.Contract.SetBridge(&_Leopard.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_Leopard *LeopardTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _Leopard.Contract.SetBridge(&_Leopard.TransactOpts, _bridge)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.TransferFrom(&_Leopard.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Leopard *LeopardTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.TransferFrom(&_Leopard.TransactOpts, from, to, tokenId)
}

// TransferItemTo is a paid mutator transaction binding the contract method 0x943cd49a.
//
// Solidity: function transferItemTo(address _from, address _to, uint256 tokenID) returns()
func (_Leopard *LeopardTransactor) TransferItemTo(opts *bind.TransactOpts, _from common.Address, _to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _Leopard.contract.Transact(opts, "transferItemTo", _from, _to, tokenID)
}

// TransferItemTo is a paid mutator transaction binding the contract method 0x943cd49a.
//
// Solidity: function transferItemTo(address _from, address _to, uint256 tokenID) returns()
func (_Leopard *LeopardSession) TransferItemTo(_from common.Address, _to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.TransferItemTo(&_Leopard.TransactOpts, _from, _to, tokenID)
}

// TransferItemTo is a paid mutator transaction binding the contract method 0x943cd49a.
//
// Solidity: function transferItemTo(address _from, address _to, uint256 tokenID) returns()
func (_Leopard *LeopardTransactorSession) TransferItemTo(_from common.Address, _to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _Leopard.Contract.TransferItemTo(&_Leopard.TransactOpts, _from, _to, tokenID)
}

// LeopardApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Leopard contract.
type LeopardApprovalIterator struct {
	Event *LeopardApproval // Event containing the contract specifics and raw log

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
func (it *LeopardApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardApproval)
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
		it.Event = new(LeopardApproval)
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
func (it *LeopardApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardApproval represents a Approval event raised by the Leopard contract.
type LeopardApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Leopard *LeopardFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LeopardApprovalIterator, error) {

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

	logs, sub, err := _Leopard.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LeopardApprovalIterator{contract: _Leopard.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Leopard *LeopardFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LeopardApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Leopard.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardApproval)
				if err := _Leopard.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Leopard *LeopardFilterer) ParseApproval(log types.Log) (*LeopardApproval, error) {
	event := new(LeopardApproval)
	if err := _Leopard.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeopardApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Leopard contract.
type LeopardApprovalForAllIterator struct {
	Event *LeopardApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LeopardApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardApprovalForAll)
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
		it.Event = new(LeopardApprovalForAll)
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
func (it *LeopardApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardApprovalForAll represents a ApprovalForAll event raised by the Leopard contract.
type LeopardApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Leopard *LeopardFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LeopardApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Leopard.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LeopardApprovalForAllIterator{contract: _Leopard.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Leopard *LeopardFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LeopardApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Leopard.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardApprovalForAll)
				if err := _Leopard.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Leopard *LeopardFilterer) ParseApprovalForAll(log types.Log) (*LeopardApprovalForAll, error) {
	event := new(LeopardApprovalForAll)
	if err := _Leopard.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeopardBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Leopard contract.
type LeopardBatchMetadataUpdateIterator struct {
	Event *LeopardBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LeopardBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardBatchMetadataUpdate)
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
		it.Event = new(LeopardBatchMetadataUpdate)
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
func (it *LeopardBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Leopard contract.
type LeopardBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Leopard *LeopardFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*LeopardBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Leopard.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LeopardBatchMetadataUpdateIterator{contract: _Leopard.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Leopard *LeopardFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LeopardBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Leopard.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardBatchMetadataUpdate)
				if err := _Leopard.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Leopard *LeopardFilterer) ParseBatchMetadataUpdate(log types.Log) (*LeopardBatchMetadataUpdate, error) {
	event := new(LeopardBatchMetadataUpdate)
	if err := _Leopard.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeopardMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Leopard contract.
type LeopardMetadataUpdateIterator struct {
	Event *LeopardMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LeopardMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardMetadataUpdate)
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
		it.Event = new(LeopardMetadataUpdate)
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
func (it *LeopardMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardMetadataUpdate represents a MetadataUpdate event raised by the Leopard contract.
type LeopardMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Leopard *LeopardFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*LeopardMetadataUpdateIterator, error) {

	logs, sub, err := _Leopard.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LeopardMetadataUpdateIterator{contract: _Leopard.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Leopard *LeopardFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LeopardMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Leopard.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardMetadataUpdate)
				if err := _Leopard.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Leopard *LeopardFilterer) ParseMetadataUpdate(log types.Log) (*LeopardMetadataUpdate, error) {
	event := new(LeopardMetadataUpdate)
	if err := _Leopard.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeopardTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Leopard contract.
type LeopardTransferIterator struct {
	Event *LeopardTransfer // Event containing the contract specifics and raw log

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
func (it *LeopardTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeopardTransfer)
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
		it.Event = new(LeopardTransfer)
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
func (it *LeopardTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeopardTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeopardTransfer represents a Transfer event raised by the Leopard contract.
type LeopardTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Leopard *LeopardFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LeopardTransferIterator, error) {

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

	logs, sub, err := _Leopard.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LeopardTransferIterator{contract: _Leopard.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Leopard *LeopardFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LeopardTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Leopard.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeopardTransfer)
				if err := _Leopard.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Leopard *LeopardFilterer) ParseTransfer(log types.Log) (*LeopardTransfer, error) {
	event := new(LeopardTransfer)
	if err := _Leopard.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
