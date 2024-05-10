// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingIStrategyManager

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

// IStrategyManagerDeprecatedStructQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Staker               common.Address
	WithdrawerAndNonce   IStrategyManagerDeprecatedStructWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IStrategyManagerDeprecatedStructWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// BindingIStrategyManagerMetaData contains all meta data concerning the BindingIStrategyManager contract.
var BindingIStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateQueuedWithdrawal\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// BindingIStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingIStrategyManagerMetaData.ABI instead.
var BindingIStrategyManagerABI = BindingIStrategyManagerMetaData.ABI

// BindingIStrategyManager is an auto generated Go binding around an Ethereum contract.
type BindingIStrategyManager struct {
	BindingIStrategyManagerCaller     // Read-only binding to the contract
	BindingIStrategyManagerTransactor // Write-only binding to the contract
	BindingIStrategyManagerFilterer   // Log filterer for contract events
}

// BindingIStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingIStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingIStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingIStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingIStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingIStrategyManagerSession struct {
	Contract     *BindingIStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BindingIStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingIStrategyManagerCallerSession struct {
	Contract *BindingIStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BindingIStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingIStrategyManagerTransactorSession struct {
	Contract     *BindingIStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BindingIStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingIStrategyManagerRaw struct {
	Contract *BindingIStrategyManager // Generic contract binding to access the raw methods on
}

// BindingIStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingIStrategyManagerCallerRaw struct {
	Contract *BindingIStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BindingIStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingIStrategyManagerTransactorRaw struct {
	Contract *BindingIStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingIStrategyManager creates a new instance of BindingIStrategyManager, bound to a specific deployed contract.
func NewBindingIStrategyManager(address common.Address, backend bind.ContractBackend) (*BindingIStrategyManager, error) {
	contract, err := bindBindingIStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManager{BindingIStrategyManagerCaller: BindingIStrategyManagerCaller{contract: contract}, BindingIStrategyManagerTransactor: BindingIStrategyManagerTransactor{contract: contract}, BindingIStrategyManagerFilterer: BindingIStrategyManagerFilterer{contract: contract}}, nil
}

// NewBindingIStrategyManagerCaller creates a new read-only instance of BindingIStrategyManager, bound to a specific deployed contract.
func NewBindingIStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*BindingIStrategyManagerCaller, error) {
	contract, err := bindBindingIStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerCaller{contract: contract}, nil
}

// NewBindingIStrategyManagerTransactor creates a new write-only instance of BindingIStrategyManager, bound to a specific deployed contract.
func NewBindingIStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingIStrategyManagerTransactor, error) {
	contract, err := bindBindingIStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerTransactor{contract: contract}, nil
}

// NewBindingIStrategyManagerFilterer creates a new log filterer instance of BindingIStrategyManager, bound to a specific deployed contract.
func NewBindingIStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingIStrategyManagerFilterer, error) {
	contract, err := bindBindingIStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerFilterer{contract: contract}, nil
}

// bindBindingIStrategyManager binds a generic wrapper to an already deployed contract.
func bindBindingIStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingIStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingIStrategyManager *BindingIStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingIStrategyManager.Contract.BindingIStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingIStrategyManager *BindingIStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.BindingIStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingIStrategyManager *BindingIStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.BindingIStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingIStrategyManager *BindingIStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingIStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "calculateWithdrawalRoot", queuedWithdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _BindingIStrategyManager.Contract.CalculateWithdrawalRoot(&_BindingIStrategyManager.CallOpts, queuedWithdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _BindingIStrategyManager.Contract.CalculateWithdrawalRoot(&_BindingIStrategyManager.CallOpts, queuedWithdrawal)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) Delegation() (common.Address, error) {
	return _BindingIStrategyManager.Contract.Delegation(&_BindingIStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _BindingIStrategyManager.Contract.Delegation(&_BindingIStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) EigenPodManager() (common.Address, error) {
	return _BindingIStrategyManager.Contract.EigenPodManager(&_BindingIStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _BindingIStrategyManager.Contract.EigenPodManager(&_BindingIStrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_BindingIStrategyManager *BindingIStrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _BindingIStrategyManager.Contract.GetDeposits(&_BindingIStrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _BindingIStrategyManager.Contract.GetDeposits(&_BindingIStrategyManager.CallOpts, staker)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) Slasher() (common.Address, error) {
	return _BindingIStrategyManager.Contract.Slasher(&_BindingIStrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) Slasher() (common.Address, error) {
	return _BindingIStrategyManager.Contract.Slasher(&_BindingIStrategyManager.CallOpts)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _BindingIStrategyManager.Contract.StakerStrategyListLength(&_BindingIStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _BindingIStrategyManager.Contract.StakerStrategyListLength(&_BindingIStrategyManager.CallOpts, staker)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) StakerStrategyShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "stakerStrategyShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _BindingIStrategyManager.Contract.StakerStrategyShares(&_BindingIStrategyManager.CallOpts, user, strategy)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _BindingIStrategyManager.Contract.StakerStrategyShares(&_BindingIStrategyManager.CallOpts, user, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _BindingIStrategyManager.Contract.StrategyWhitelister(&_BindingIStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _BindingIStrategyManager.Contract.StrategyWhitelister(&_BindingIStrategyManager.CallOpts)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_BindingIStrategyManager *BindingIStrategyManagerCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _BindingIStrategyManager.contract.Call(opts, &out, "thirdPartyTransfersForbidden", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) ThirdPartyTransfersForbidden(strategy common.Address) (bool, error) {
	return _BindingIStrategyManager.Contract.ThirdPartyTransfersForbidden(&_BindingIStrategyManager.CallOpts, strategy)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_BindingIStrategyManager *BindingIStrategyManagerCallerSession) ThirdPartyTransfersForbidden(strategy common.Address) (bool, error) {
	return _BindingIStrategyManager.Contract.ThirdPartyTransfersForbidden(&_BindingIStrategyManager.CallOpts, strategy)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "addShares", staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.AddShares(&_BindingIStrategyManager.TransactOpts, staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.AddShares(&_BindingIStrategyManager.TransactOpts, staker, token, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_BindingIStrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_BindingIStrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.DepositIntoStrategy(&_BindingIStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.DepositIntoStrategy(&_BindingIStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.DepositIntoStrategyWithSignature(&_BindingIStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.DepositIntoStrategyWithSignature(&_BindingIStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) MigrateQueuedWithdrawal(opts *bind.TransactOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "migrateQueuedWithdrawal", queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.MigrateQueuedWithdrawal(&_BindingIStrategyManager.TransactOpts, queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.MigrateQueuedWithdrawal(&_BindingIStrategyManager.TransactOpts, queuedWithdrawal)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.RemoveShares(&_BindingIStrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.RemoveShares(&_BindingIStrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_BindingIStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_BindingIStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.WithdrawSharesAsTokens(&_BindingIStrategyManager.TransactOpts, recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_BindingIStrategyManager *BindingIStrategyManagerTransactorSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _BindingIStrategyManager.Contract.WithdrawSharesAsTokens(&_BindingIStrategyManager.TransactOpts, recipient, strategy, shares, token)
}

// BindingIStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerDepositIterator struct {
	Event *BindingIStrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *BindingIStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingIStrategyManagerDeposit)
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
		it.Event = new(BindingIStrategyManagerDeposit)
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
func (it *BindingIStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingIStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingIStrategyManagerDeposit represents a Deposit event raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerDeposit struct {
	Staker   common.Address
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*BindingIStrategyManagerDepositIterator, error) {

	logs, sub, err := _BindingIStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerDepositIterator{contract: _BindingIStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BindingIStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _BindingIStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingIStrategyManagerDeposit)
				if err := _BindingIStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) ParseDeposit(log types.Log) (*BindingIStrategyManagerDeposit, error) {
	event := new(BindingIStrategyManagerDeposit)
	if err := _BindingIStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *BindingIStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingIStrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(BindingIStrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingIStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _BindingIStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _BindingIStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *BindingIStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _BindingIStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingIStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*BindingIStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(BindingIStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *BindingIStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingIStrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(BindingIStrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingIStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _BindingIStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _BindingIStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *BindingIStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _BindingIStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingIStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*BindingIStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(BindingIStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingIStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *BindingIStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *BindingIStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingIStrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(BindingIStrategyManagerStrategyWhitelisterChanged)
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
func (it *BindingIStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingIStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingIStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*BindingIStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _BindingIStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerStrategyWhitelisterChangedIterator{contract: _BindingIStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *BindingIStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _BindingIStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingIStrategyManagerStrategyWhitelisterChanged)
				if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*BindingIStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(BindingIStrategyManagerStrategyWhitelisterChanged)
	if err := _BindingIStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden)
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
func (it *BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the BindingIStrategyManager contract.
type BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden struct {
	Strategy common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _BindingIStrategyManager.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &BindingIStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator{contract: _BindingIStrategyManager.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _BindingIStrategyManager.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden)
				if err := _BindingIStrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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

// ParseUpdatedThirdPartyTransfersForbidden is a log parse operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_BindingIStrategyManager *BindingIStrategyManagerFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden, error) {
	event := new(BindingIStrategyManagerUpdatedThirdPartyTransfersForbidden)
	if err := _BindingIStrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
