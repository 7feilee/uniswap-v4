// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LpRouter

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

// IPoolManagerModifyLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerModifyLiquidityParams struct {
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// LpRouterMetaData contains all meta data concerning the LpRouter contract.
var LpRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolManager\",\"name\":\"_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractIPoolManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"liquidityDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"settleUsingBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"takeClaims\",\"type\":\"bool\"}],\"name\":\"modifyLiquidity\",\"outputs\":[{\"internalType\":\"BalanceDelta\",\"name\":\"delta\",\"type\":\"int256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"liquidityDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"modifyLiquidity\",\"outputs\":[{\"internalType\":\"BalanceDelta\",\"name\":\"delta\",\"type\":\"int256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawData\",\"type\":\"bytes\"}],\"name\":\"unlockCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LpRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use LpRouterMetaData.ABI instead.
var LpRouterABI = LpRouterMetaData.ABI

// LpRouter is an auto generated Go binding around an Ethereum contract.
type LpRouter struct {
	LpRouterCaller     // Read-only binding to the contract
	LpRouterTransactor // Write-only binding to the contract
	LpRouterFilterer   // Log filterer for contract events
}

// LpRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpRouterSession struct {
	Contract     *LpRouter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpRouterCallerSession struct {
	Contract *LpRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LpRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpRouterTransactorSession struct {
	Contract     *LpRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LpRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpRouterRaw struct {
	Contract *LpRouter // Generic contract binding to access the raw methods on
}

// LpRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpRouterCallerRaw struct {
	Contract *LpRouterCaller // Generic read-only contract binding to access the raw methods on
}

// LpRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpRouterTransactorRaw struct {
	Contract *LpRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLpRouter creates a new instance of LpRouter, bound to a specific deployed contract.
func NewLpRouter(address common.Address, backend bind.ContractBackend) (*LpRouter, error) {
	contract, err := bindLpRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LpRouter{LpRouterCaller: LpRouterCaller{contract: contract}, LpRouterTransactor: LpRouterTransactor{contract: contract}, LpRouterFilterer: LpRouterFilterer{contract: contract}}, nil
}

// NewLpRouterCaller creates a new read-only instance of LpRouter, bound to a specific deployed contract.
func NewLpRouterCaller(address common.Address, caller bind.ContractCaller) (*LpRouterCaller, error) {
	contract, err := bindLpRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpRouterCaller{contract: contract}, nil
}

// NewLpRouterTransactor creates a new write-only instance of LpRouter, bound to a specific deployed contract.
func NewLpRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*LpRouterTransactor, error) {
	contract, err := bindLpRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpRouterTransactor{contract: contract}, nil
}

// NewLpRouterFilterer creates a new log filterer instance of LpRouter, bound to a specific deployed contract.
func NewLpRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*LpRouterFilterer, error) {
	contract, err := bindLpRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpRouterFilterer{contract: contract}, nil
}

// bindLpRouter binds a generic wrapper to an already deployed contract.
func bindLpRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LpRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpRouter *LpRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpRouter.Contract.LpRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpRouter *LpRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpRouter.Contract.LpRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpRouter *LpRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpRouter.Contract.LpRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpRouter *LpRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpRouter *LpRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpRouter *LpRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpRouter.Contract.contract.Transact(opts, method, params...)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LpRouter *LpRouterCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpRouter.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LpRouter *LpRouterSession) Manager() (common.Address, error) {
	return _LpRouter.Contract.Manager(&_LpRouter.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LpRouter *LpRouterCallerSession) Manager() (common.Address, error) {
	return _LpRouter.Contract.Manager(&_LpRouter.CallOpts)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x0a5b11e4.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData, bool settleUsingBurn, bool takeClaims) payable returns(int256 delta)
func (_LpRouter *LpRouterTransactor) ModifyLiquidity(opts *bind.TransactOpts, key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte, settleUsingBurn bool, takeClaims bool) (*types.Transaction, error) {
	return _LpRouter.contract.Transact(opts, "modifyLiquidity", key, params, hookData, settleUsingBurn, takeClaims)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x0a5b11e4.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData, bool settleUsingBurn, bool takeClaims) payable returns(int256 delta)
func (_LpRouter *LpRouterSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte, settleUsingBurn bool, takeClaims bool) (*types.Transaction, error) {
	return _LpRouter.Contract.ModifyLiquidity(&_LpRouter.TransactOpts, key, params, hookData, settleUsingBurn, takeClaims)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x0a5b11e4.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData, bool settleUsingBurn, bool takeClaims) payable returns(int256 delta)
func (_LpRouter *LpRouterTransactorSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte, settleUsingBurn bool, takeClaims bool) (*types.Transaction, error) {
	return _LpRouter.Contract.ModifyLiquidity(&_LpRouter.TransactOpts, key, params, hookData, settleUsingBurn, takeClaims)
}

// ModifyLiquidity0 is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) payable returns(int256 delta)
func (_LpRouter *LpRouterTransactor) ModifyLiquidity0(opts *bind.TransactOpts, key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _LpRouter.contract.Transact(opts, "modifyLiquidity0", key, params, hookData)
}

// ModifyLiquidity0 is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) payable returns(int256 delta)
func (_LpRouter *LpRouterSession) ModifyLiquidity0(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _LpRouter.Contract.ModifyLiquidity0(&_LpRouter.TransactOpts, key, params, hookData)
}

// ModifyLiquidity0 is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) payable returns(int256 delta)
func (_LpRouter *LpRouterTransactorSession) ModifyLiquidity0(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _LpRouter.Contract.ModifyLiquidity0(&_LpRouter.TransactOpts, key, params, hookData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_LpRouter *LpRouterTransactor) UnlockCallback(opts *bind.TransactOpts, rawData []byte) (*types.Transaction, error) {
	return _LpRouter.contract.Transact(opts, "unlockCallback", rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_LpRouter *LpRouterSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _LpRouter.Contract.UnlockCallback(&_LpRouter.TransactOpts, rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_LpRouter *LpRouterTransactorSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _LpRouter.Contract.UnlockCallback(&_LpRouter.TransactOpts, rawData)
}
