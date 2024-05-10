// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingIStrategy

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

// BindingIStrategyMetaData contains all meta data concerning the BindingIStrategy contract.
var BindingIStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BindingIStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingIStrategyMetaData.ABI instead.
var BindingIStrategyABI = BindingIStrategyMetaData.ABI

// BindingIStrategy is an auto generated Go binding around an Ethereum contract.
type BindingIStrategy struct {
	BindingIStrategyCaller     // Read-only binding to the contract
	BindingIStrategyTransactor // Write-only binding to the contract
	BindingIStrategyFilterer   // Log filterer for contract events
}

// BindingIStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingIStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingIStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingIStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingIStrategySession struct {
	Contract     *BindingIStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingIStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingIStrategyCallerSession struct {
	Contract *BindingIStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BindingIStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingIStrategyTransactorSession struct {
	Contract     *BindingIStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BindingIStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingIStrategyRaw struct {
	Contract *BindingIStrategy // Generic contract binding to access the raw methods on
}

// BindingIStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingIStrategyCallerRaw struct {
	Contract *BindingIStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// BindingIStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingIStrategyTransactorRaw struct {
	Contract *BindingIStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingIStrategy creates a new instance of BindingIStrategy, bound to a specific deployed contract.
func NewBindingIStrategy(address common.Address, backend bind.ContractBackend) (*BindingIStrategy, error) {
	contract, err := bindBindingIStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategy{BindingIStrategyCaller: BindingIStrategyCaller{contract: contract}, BindingIStrategyTransactor: BindingIStrategyTransactor{contract: contract}, BindingIStrategyFilterer: BindingIStrategyFilterer{contract: contract}}, nil
}

// NewBindingIStrategyCaller creates a new read-only instance of BindingIStrategy, bound to a specific deployed contract.
func NewBindingIStrategyCaller(address common.Address, caller bind.ContractCaller) (*BindingIStrategyCaller, error) {
	contract, err := bindBindingIStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyCaller{contract: contract}, nil
}

// NewBindingIStrategyTransactor creates a new write-only instance of BindingIStrategy, bound to a specific deployed contract.
func NewBindingIStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingIStrategyTransactor, error) {
	contract, err := bindBindingIStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyTransactor{contract: contract}, nil
}

// NewBindingIStrategyFilterer creates a new log filterer instance of BindingIStrategy, bound to a specific deployed contract.
func NewBindingIStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingIStrategyFilterer, error) {
	contract, err := bindBindingIStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyFilterer{contract: contract}, nil
}

// bindBindingIStrategy binds a generic wrapper to an already deployed contract.
func bindBindingIStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingIStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingIStrategy *BindingIStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingIStrategy.Contract.BindingIStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingIStrategy *BindingIStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.BindingIStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingIStrategy *BindingIStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.BindingIStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingIStrategy *BindingIStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingIStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingIStrategy *BindingIStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingIStrategy *BindingIStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_BindingIStrategy *BindingIStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_BindingIStrategy *BindingIStrategySession) Explanation() (string, error) {
	return _BindingIStrategy.Contract.Explanation(&_BindingIStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_BindingIStrategy *BindingIStrategyCallerSession) Explanation() (string, error) {
	return _BindingIStrategy.Contract.Explanation(&_BindingIStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _BindingIStrategy.Contract.Shares(&_BindingIStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _BindingIStrategy.Contract.Shares(&_BindingIStrategy.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _BindingIStrategy.Contract.SharesToUnderlyingView(&_BindingIStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _BindingIStrategy.Contract.SharesToUnderlyingView(&_BindingIStrategy.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) TotalShares() (*big.Int, error) {
	return _BindingIStrategy.Contract.TotalShares(&_BindingIStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _BindingIStrategy.Contract.TotalShares(&_BindingIStrategy.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _BindingIStrategy.Contract.UnderlyingToSharesView(&_BindingIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _BindingIStrategy.Contract.UnderlyingToSharesView(&_BindingIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_BindingIStrategy *BindingIStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_BindingIStrategy *BindingIStrategySession) UnderlyingToken() (common.Address, error) {
	return _BindingIStrategy.Contract.UnderlyingToken(&_BindingIStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_BindingIStrategy *BindingIStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _BindingIStrategy.Contract.UnderlyingToken(&_BindingIStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _BindingIStrategy.Contract.UserUnderlyingView(&_BindingIStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_BindingIStrategy *BindingIStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _BindingIStrategy.Contract.UserUnderlyingView(&_BindingIStrategy.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.Deposit(&_BindingIStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.Deposit(&_BindingIStrategy.TransactOpts, token, amount)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.SharesToUnderlying(&_BindingIStrategy.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.SharesToUnderlying(&_BindingIStrategy.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.UnderlyingToShares(&_BindingIStrategy.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.UnderlyingToShares(&_BindingIStrategy.TransactOpts, amountUnderlying)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _BindingIStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_BindingIStrategy *BindingIStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.UserUnderlying(&_BindingIStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_BindingIStrategy *BindingIStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.UserUnderlying(&_BindingIStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_BindingIStrategy *BindingIStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_BindingIStrategy *BindingIStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.Withdraw(&_BindingIStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_BindingIStrategy *BindingIStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategy.Contract.Withdraw(&_BindingIStrategy.TransactOpts, recipient, token, amountShares)
}
