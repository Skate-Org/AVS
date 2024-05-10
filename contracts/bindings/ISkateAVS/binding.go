// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingISkateAVS

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ISkateAVSOperator is an auto generated low-level Go binding around an user-defined struct.
type ISkateAVSOperator struct {
	Addr      common.Address
	Delegated *big.Int
	Staked    *big.Int
}

// ISkateAVSSignatureTuple is an auto generated low-level Go binding around an user-defined struct.
type ISkateAVSSignatureTuple struct {
	Operator  common.Address
	Signature []byte
}

// ISkateAVSStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type ISkateAVSStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// BindingISkateAVSMetaData contains all meta data concerning the BindingISkateAVS contract.
var BindingISkateAVSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addToAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSubmitData\",\"inputs\":[{\"name\":\"taskIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"messageDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"signaturesTuples\",\"type\":\"tuple[][]\",\"internalType\":\"structISkateAVS.SignatureTuple[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canRegister\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minOperatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structISkateAVS.Operator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegated\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"staked\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxOperatorCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinOperatorStake\",\"inputs\":[{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategies\",\"inputs\":[{\"name\":\"strategies_\",\"type\":\"tuple[]\",\"internalType\":\"structISkateAVS.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structISkateAVS.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitData\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"messageData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatureTuples\",\"type\":\"tuple[]\",\"internalType\":\"structISkateAVS.SignatureTuple[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowlistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataSubmitted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"messageData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxOperatorCountSet\",\"inputs\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinOperatorStakeSet\",\"inputs\":[{\"name\":\"minOperatorStake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAllowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDisallowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategiesSet\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structISkateAVS.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"anonymous\":false}]",
}

// BindingISkateAVSABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingISkateAVSMetaData.ABI instead.
var BindingISkateAVSABI = BindingISkateAVSMetaData.ABI

// BindingISkateAVS is an auto generated Go binding around an Ethereum contract.
type BindingISkateAVS struct {
	BindingISkateAVSCaller     // Read-only binding to the contract
	BindingISkateAVSTransactor // Write-only binding to the contract
	BindingISkateAVSFilterer   // Log filterer for contract events
}

// BindingISkateAVSCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingISkateAVSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingISkateAVSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingISkateAVSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingISkateAVSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingISkateAVSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingISkateAVSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingISkateAVSSession struct {
	Contract     *BindingISkateAVS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingISkateAVSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingISkateAVSCallerSession struct {
	Contract *BindingISkateAVSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BindingISkateAVSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingISkateAVSTransactorSession struct {
	Contract     *BindingISkateAVSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BindingISkateAVSRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingISkateAVSRaw struct {
	Contract *BindingISkateAVS // Generic contract binding to access the raw methods on
}

// BindingISkateAVSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingISkateAVSCallerRaw struct {
	Contract *BindingISkateAVSCaller // Generic read-only contract binding to access the raw methods on
}

// BindingISkateAVSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingISkateAVSTransactorRaw struct {
	Contract *BindingISkateAVSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingISkateAVS creates a new instance of BindingISkateAVS, bound to a specific deployed contract.
func NewBindingISkateAVS(address common.Address, backend bind.ContractBackend) (*BindingISkateAVS, error) {
	contract, err := bindBindingISkateAVS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVS{BindingISkateAVSCaller: BindingISkateAVSCaller{contract: contract}, BindingISkateAVSTransactor: BindingISkateAVSTransactor{contract: contract}, BindingISkateAVSFilterer: BindingISkateAVSFilterer{contract: contract}}, nil
}

// NewBindingISkateAVSCaller creates a new read-only instance of BindingISkateAVS, bound to a specific deployed contract.
func NewBindingISkateAVSCaller(address common.Address, caller bind.ContractCaller) (*BindingISkateAVSCaller, error) {
	contract, err := bindBindingISkateAVS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSCaller{contract: contract}, nil
}

// NewBindingISkateAVSTransactor creates a new write-only instance of BindingISkateAVS, bound to a specific deployed contract.
func NewBindingISkateAVSTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingISkateAVSTransactor, error) {
	contract, err := bindBindingISkateAVS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSTransactor{contract: contract}, nil
}

// NewBindingISkateAVSFilterer creates a new log filterer instance of BindingISkateAVS, bound to a specific deployed contract.
func NewBindingISkateAVSFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingISkateAVSFilterer, error) {
	contract, err := bindBindingISkateAVS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSFilterer{contract: contract}, nil
}

// bindBindingISkateAVS binds a generic wrapper to an already deployed contract.
func bindBindingISkateAVS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingISkateAVSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingISkateAVS *BindingISkateAVSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingISkateAVS.Contract.BindingISkateAVSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingISkateAVS *BindingISkateAVSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.BindingISkateAVSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingISkateAVS *BindingISkateAVSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.BindingISkateAVSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingISkateAVS *BindingISkateAVSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingISkateAVS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingISkateAVS *BindingISkateAVSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingISkateAVS *BindingISkateAVSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.contract.Transact(opts, method, params...)
}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCaller) AllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "allowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSSession) AllowlistEnabled() (bool, error) {
	return _BindingISkateAVS.Contract.AllowlistEnabled(&_BindingISkateAVS.CallOpts)
}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) AllowlistEnabled() (bool, error) {
	return _BindingISkateAVS.Contract.AllowlistEnabled(&_BindingISkateAVS.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingISkateAVS *BindingISkateAVSCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingISkateAVS *BindingISkateAVSSession) AvsDirectory() (common.Address, error) {
	return _BindingISkateAVS.Contract.AvsDirectory(&_BindingISkateAVS.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) AvsDirectory() (common.Address, error) {
	return _BindingISkateAVS.Contract.AvsDirectory(&_BindingISkateAVS.CallOpts)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCaller) CanRegister(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "canRegister", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSSession) CanRegister(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.CanRegister(&_BindingISkateAVS.CallOpts, operator)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) CanRegister(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.CanRegister(&_BindingISkateAVS.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BindingISkateAVS.Contract.GetOperatorRestakedStrategies(&_BindingISkateAVS.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BindingISkateAVS.Contract.GetOperatorRestakedStrategies(&_BindingISkateAVS.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BindingISkateAVS.Contract.GetRestakeableStrategies(&_BindingISkateAVS.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingISkateAVS *BindingISkateAVSCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BindingISkateAVS.Contract.GetRestakeableStrategies(&_BindingISkateAVS.CallOpts)
}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCaller) IsInAllowlist(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "isInAllowlist", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSSession) IsInAllowlist(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.IsInAllowlist(&_BindingISkateAVS.CallOpts, operator)
}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) IsInAllowlist(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.IsInAllowlist(&_BindingISkateAVS.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSSession) IsOperator(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.IsOperator(&_BindingISkateAVS.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _BindingISkateAVS.Contract.IsOperator(&_BindingISkateAVS.CallOpts, operator)
}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_BindingISkateAVS *BindingISkateAVSCaller) MaxOperatorCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "maxOperatorCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_BindingISkateAVS *BindingISkateAVSSession) MaxOperatorCount() (uint32, error) {
	return _BindingISkateAVS.Contract.MaxOperatorCount(&_BindingISkateAVS.CallOpts)
}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) MaxOperatorCount() (uint32, error) {
	return _BindingISkateAVS.Contract.MaxOperatorCount(&_BindingISkateAVS.CallOpts)
}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint96)
func (_BindingISkateAVS *BindingISkateAVSCaller) MinOperatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "minOperatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint96)
func (_BindingISkateAVS *BindingISkateAVSSession) MinOperatorStake() (*big.Int, error) {
	return _BindingISkateAVS.Contract.MinOperatorStake(&_BindingISkateAVS.CallOpts)
}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint96)
func (_BindingISkateAVS *BindingISkateAVSCallerSession) MinOperatorStake() (*big.Int, error) {
	return _BindingISkateAVS.Contract.MinOperatorStake(&_BindingISkateAVS.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSCaller) Operators(opts *bind.CallOpts) ([]ISkateAVSOperator, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "operators")

	if err != nil {
		return *new([]ISkateAVSOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISkateAVSOperator)).(*[]ISkateAVSOperator)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSSession) Operators() ([]ISkateAVSOperator, error) {
	return _BindingISkateAVS.Contract.Operators(&_BindingISkateAVS.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSCallerSession) Operators() ([]ISkateAVSOperator, error) {
	return _BindingISkateAVS.Contract.Operators(&_BindingISkateAVS.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns((address,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSCaller) Strategies(opts *bind.CallOpts) ([]ISkateAVSStrategyParams, error) {
	var out []interface{}
	err := _BindingISkateAVS.contract.Call(opts, &out, "strategies")

	if err != nil {
		return *new([]ISkateAVSStrategyParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISkateAVSStrategyParams)).(*[]ISkateAVSStrategyParams)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns((address,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSSession) Strategies() ([]ISkateAVSStrategyParams, error) {
	return _BindingISkateAVS.Contract.Strategies(&_BindingISkateAVS.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns((address,uint96)[])
func (_BindingISkateAVS *BindingISkateAVSCallerSession) Strategies() ([]ISkateAVSStrategyParams, error) {
	return _BindingISkateAVS.Contract.Strategies(&_BindingISkateAVS.CallOpts)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) AddToAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "addToAllowlist", operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.AddToAllowlist(&_BindingISkateAVS.TransactOpts, operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.AddToAllowlist(&_BindingISkateAVS.TransactOpts, operator)
}

// BatchSubmitData is a paid mutator transaction binding the contract method 0x118b48f0.
//
// Solidity: function batchSubmitData(uint256[] taskIds, bytes[] messageDatas, (address,bytes)[][] signaturesTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) BatchSubmitData(opts *bind.TransactOpts, taskIds []*big.Int, messageDatas [][]byte, signaturesTuples [][]ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "batchSubmitData", taskIds, messageDatas, signaturesTuples)
}

// BatchSubmitData is a paid mutator transaction binding the contract method 0x118b48f0.
//
// Solidity: function batchSubmitData(uint256[] taskIds, bytes[] messageDatas, (address,bytes)[][] signaturesTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) BatchSubmitData(taskIds []*big.Int, messageDatas [][]byte, signaturesTuples [][]ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.BatchSubmitData(&_BindingISkateAVS.TransactOpts, taskIds, messageDatas, signaturesTuples)
}

// BatchSubmitData is a paid mutator transaction binding the contract method 0x118b48f0.
//
// Solidity: function batchSubmitData(uint256[] taskIds, bytes[] messageDatas, (address,bytes)[][] signaturesTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) BatchSubmitData(taskIds []*big.Int, messageDatas [][]byte, signaturesTuples [][]ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.BatchSubmitData(&_BindingISkateAVS.TransactOpts, taskIds, messageDatas, signaturesTuples)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.DeregisterOperatorFromAVS(&_BindingISkateAVS.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.DeregisterOperatorFromAVS(&_BindingISkateAVS.TransactOpts, operator)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) DisableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "disableAllowlist")
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSSession) DisableAllowlist() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.DisableAllowlist(&_BindingISkateAVS.TransactOpts)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) DisableAllowlist() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.DisableAllowlist(&_BindingISkateAVS.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) EnableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "enableAllowlist")
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSSession) EnableAllowlist() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.EnableAllowlist(&_BindingISkateAVS.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) EnableAllowlist() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.EnableAllowlist(&_BindingISkateAVS.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BindingISkateAVS *BindingISkateAVSSession) Pause() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.Pause(&_BindingISkateAVS.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) Pause() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.Pause(&_BindingISkateAVS.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.RegisterOperatorToAVS(&_BindingISkateAVS.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.RegisterOperatorToAVS(&_BindingISkateAVS.TransactOpts, operator, operatorSignature)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) RemoveFromAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "removeFromAllowlist", operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.RemoveFromAllowlist(&_BindingISkateAVS.TransactOpts, operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.RemoveFromAllowlist(&_BindingISkateAVS.TransactOpts, operator)
}

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) SetMaxOperatorCount(opts *bind.TransactOpts, count uint32) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "setMaxOperatorCount", count)
}

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) SetMaxOperatorCount(count uint32) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetMaxOperatorCount(&_BindingISkateAVS.TransactOpts, count)
}

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) SetMaxOperatorCount(count uint32) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetMaxOperatorCount(&_BindingISkateAVS.TransactOpts, count)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0xeb316235.
//
// Solidity: function setMinOperatorStake(uint96 stake) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) SetMinOperatorStake(opts *bind.TransactOpts, stake *big.Int) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "setMinOperatorStake", stake)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0xeb316235.
//
// Solidity: function setMinOperatorStake(uint96 stake) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) SetMinOperatorStake(stake *big.Int) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetMinOperatorStake(&_BindingISkateAVS.TransactOpts, stake)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0xeb316235.
//
// Solidity: function setMinOperatorStake(uint96 stake) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) SetMinOperatorStake(stake *big.Int) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetMinOperatorStake(&_BindingISkateAVS.TransactOpts, stake)
}

// SetStrategies is a paid mutator transaction binding the contract method 0xff5c02f2.
//
// Solidity: function setStrategies((address,uint96)[] strategies_) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) SetStrategies(opts *bind.TransactOpts, strategies_ []ISkateAVSStrategyParams) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "setStrategies", strategies_)
}

// SetStrategies is a paid mutator transaction binding the contract method 0xff5c02f2.
//
// Solidity: function setStrategies((address,uint96)[] strategies_) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) SetStrategies(strategies_ []ISkateAVSStrategyParams) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetStrategies(&_BindingISkateAVS.TransactOpts, strategies_)
}

// SetStrategies is a paid mutator transaction binding the contract method 0xff5c02f2.
//
// Solidity: function setStrategies((address,uint96)[] strategies_) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) SetStrategies(strategies_ []ISkateAVSStrategyParams) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SetStrategies(&_BindingISkateAVS.TransactOpts, strategies_)
}

// SubmitData is a paid mutator transaction binding the contract method 0x5917dcca.
//
// Solidity: function submitData(uint256 taskId, bytes messageData, (address,bytes)[] signatureTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) SubmitData(opts *bind.TransactOpts, taskId *big.Int, messageData []byte, signatureTuples []ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "submitData", taskId, messageData, signatureTuples)
}

// SubmitData is a paid mutator transaction binding the contract method 0x5917dcca.
//
// Solidity: function submitData(uint256 taskId, bytes messageData, (address,bytes)[] signatureTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) SubmitData(taskId *big.Int, messageData []byte, signatureTuples []ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SubmitData(&_BindingISkateAVS.TransactOpts, taskId, messageData, signatureTuples)
}

// SubmitData is a paid mutator transaction binding the contract method 0x5917dcca.
//
// Solidity: function submitData(uint256 taskId, bytes messageData, (address,bytes)[] signatureTuples) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) SubmitData(taskId *big.Int, messageData []byte, signatureTuples []ISkateAVSSignatureTuple) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.SubmitData(&_BindingISkateAVS.TransactOpts, taskId, messageData, signatureTuples)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BindingISkateAVS *BindingISkateAVSSession) Unpause() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.Unpause(&_BindingISkateAVS.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) Unpause() (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.Unpause(&_BindingISkateAVS.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _BindingISkateAVS.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_BindingISkateAVS *BindingISkateAVSSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.UpdateAVSMetadataURI(&_BindingISkateAVS.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_BindingISkateAVS *BindingISkateAVSTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _BindingISkateAVS.Contract.UpdateAVSMetadataURI(&_BindingISkateAVS.TransactOpts, metadataURI)
}

// BindingISkateAVSAllowlistDisabledIterator is returned from FilterAllowlistDisabled and is used to iterate over the raw logs and unpacked data for AllowlistDisabled events raised by the BindingISkateAVS contract.
type BindingISkateAVSAllowlistDisabledIterator struct {
	Event *BindingISkateAVSAllowlistDisabled // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSAllowlistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSAllowlistDisabled)
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
		it.Event = new(BindingISkateAVSAllowlistDisabled)
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
func (it *BindingISkateAVSAllowlistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSAllowlistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSAllowlistDisabled represents a AllowlistDisabled event raised by the BindingISkateAVS contract.
type BindingISkateAVSAllowlistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistDisabled is a free log retrieval operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterAllowlistDisabled(opts *bind.FilterOpts) (*BindingISkateAVSAllowlistDisabledIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSAllowlistDisabledIterator{contract: _BindingISkateAVS.contract, event: "AllowlistDisabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistDisabled is a free log subscription operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchAllowlistDisabled(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSAllowlistDisabled) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSAllowlistDisabled)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
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

// ParseAllowlistDisabled is a log parse operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseAllowlistDisabled(log types.Log) (*BindingISkateAVSAllowlistDisabled, error) {
	event := new(BindingISkateAVSAllowlistDisabled)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSAllowlistEnabledIterator is returned from FilterAllowlistEnabled and is used to iterate over the raw logs and unpacked data for AllowlistEnabled events raised by the BindingISkateAVS contract.
type BindingISkateAVSAllowlistEnabledIterator struct {
	Event *BindingISkateAVSAllowlistEnabled // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSAllowlistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSAllowlistEnabled)
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
		it.Event = new(BindingISkateAVSAllowlistEnabled)
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
func (it *BindingISkateAVSAllowlistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSAllowlistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSAllowlistEnabled represents a AllowlistEnabled event raised by the BindingISkateAVS contract.
type BindingISkateAVSAllowlistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistEnabled is a free log retrieval operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterAllowlistEnabled(opts *bind.FilterOpts) (*BindingISkateAVSAllowlistEnabledIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSAllowlistEnabledIterator{contract: _BindingISkateAVS.contract, event: "AllowlistEnabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistEnabled is a free log subscription operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchAllowlistEnabled(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSAllowlistEnabled) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSAllowlistEnabled)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
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

// ParseAllowlistEnabled is a log parse operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseAllowlistEnabled(log types.Log) (*BindingISkateAVSAllowlistEnabled, error) {
	event := new(BindingISkateAVSAllowlistEnabled)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSDataSubmittedIterator is returned from FilterDataSubmitted and is used to iterate over the raw logs and unpacked data for DataSubmitted events raised by the BindingISkateAVS contract.
type BindingISkateAVSDataSubmittedIterator struct {
	Event *BindingISkateAVSDataSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSDataSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSDataSubmitted)
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
		it.Event = new(BindingISkateAVSDataSubmitted)
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
func (it *BindingISkateAVSDataSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSDataSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSDataSubmitted represents a DataSubmitted event raised by the BindingISkateAVS contract.
type BindingISkateAVSDataSubmitted struct {
	TaskId      *big.Int
	MessageData []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataSubmitted is a free log retrieval operation binding the contract event 0x0e5bceb96aa8298cee518cb21454c4b1f3abf125133836e647d5007b33f0e751.
//
// Solidity: event DataSubmitted(uint256 taskId, bytes messageData)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterDataSubmitted(opts *bind.FilterOpts) (*BindingISkateAVSDataSubmittedIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "DataSubmitted")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSDataSubmittedIterator{contract: _BindingISkateAVS.contract, event: "DataSubmitted", logs: logs, sub: sub}, nil
}

// WatchDataSubmitted is a free log subscription operation binding the contract event 0x0e5bceb96aa8298cee518cb21454c4b1f3abf125133836e647d5007b33f0e751.
//
// Solidity: event DataSubmitted(uint256 taskId, bytes messageData)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchDataSubmitted(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSDataSubmitted) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "DataSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSDataSubmitted)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
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

// ParseDataSubmitted is a log parse operation binding the contract event 0x0e5bceb96aa8298cee518cb21454c4b1f3abf125133836e647d5007b33f0e751.
//
// Solidity: event DataSubmitted(uint256 taskId, bytes messageData)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseDataSubmitted(log types.Log) (*BindingISkateAVSDataSubmitted, error) {
	event := new(BindingISkateAVSDataSubmitted)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSMaxOperatorCountSetIterator is returned from FilterMaxOperatorCountSet and is used to iterate over the raw logs and unpacked data for MaxOperatorCountSet events raised by the BindingISkateAVS contract.
type BindingISkateAVSMaxOperatorCountSetIterator struct {
	Event *BindingISkateAVSMaxOperatorCountSet // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSMaxOperatorCountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSMaxOperatorCountSet)
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
		it.Event = new(BindingISkateAVSMaxOperatorCountSet)
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
func (it *BindingISkateAVSMaxOperatorCountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSMaxOperatorCountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSMaxOperatorCountSet represents a MaxOperatorCountSet event raised by the BindingISkateAVS contract.
type BindingISkateAVSMaxOperatorCountSet struct {
	MaxOperatorCount uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxOperatorCountSet is a free log retrieval operation binding the contract event 0x4867705ef2f7341363cb4d4bb0e1501ad37f84ae371b9f31e73a1c25c39840e4.
//
// Solidity: event MaxOperatorCountSet(uint32 maxOperatorCount)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterMaxOperatorCountSet(opts *bind.FilterOpts) (*BindingISkateAVSMaxOperatorCountSetIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "MaxOperatorCountSet")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSMaxOperatorCountSetIterator{contract: _BindingISkateAVS.contract, event: "MaxOperatorCountSet", logs: logs, sub: sub}, nil
}

// WatchMaxOperatorCountSet is a free log subscription operation binding the contract event 0x4867705ef2f7341363cb4d4bb0e1501ad37f84ae371b9f31e73a1c25c39840e4.
//
// Solidity: event MaxOperatorCountSet(uint32 maxOperatorCount)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchMaxOperatorCountSet(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSMaxOperatorCountSet) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "MaxOperatorCountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSMaxOperatorCountSet)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "MaxOperatorCountSet", log); err != nil {
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

// ParseMaxOperatorCountSet is a log parse operation binding the contract event 0x4867705ef2f7341363cb4d4bb0e1501ad37f84ae371b9f31e73a1c25c39840e4.
//
// Solidity: event MaxOperatorCountSet(uint32 maxOperatorCount)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseMaxOperatorCountSet(log types.Log) (*BindingISkateAVSMaxOperatorCountSet, error) {
	event := new(BindingISkateAVSMaxOperatorCountSet)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "MaxOperatorCountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSMinOperatorStakeSetIterator is returned from FilterMinOperatorStakeSet and is used to iterate over the raw logs and unpacked data for MinOperatorStakeSet events raised by the BindingISkateAVS contract.
type BindingISkateAVSMinOperatorStakeSetIterator struct {
	Event *BindingISkateAVSMinOperatorStakeSet // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSMinOperatorStakeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSMinOperatorStakeSet)
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
		it.Event = new(BindingISkateAVSMinOperatorStakeSet)
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
func (it *BindingISkateAVSMinOperatorStakeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSMinOperatorStakeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSMinOperatorStakeSet represents a MinOperatorStakeSet event raised by the BindingISkateAVS contract.
type BindingISkateAVSMinOperatorStakeSet struct {
	MinOperatorStake *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinOperatorStakeSet is a free log retrieval operation binding the contract event 0x9b13fd38c94948514c63f3132d22b60ecb8ed37521ba05fdfad046cfbc3a772a.
//
// Solidity: event MinOperatorStakeSet(uint96 minOperatorStake)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterMinOperatorStakeSet(opts *bind.FilterOpts) (*BindingISkateAVSMinOperatorStakeSetIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "MinOperatorStakeSet")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSMinOperatorStakeSetIterator{contract: _BindingISkateAVS.contract, event: "MinOperatorStakeSet", logs: logs, sub: sub}, nil
}

// WatchMinOperatorStakeSet is a free log subscription operation binding the contract event 0x9b13fd38c94948514c63f3132d22b60ecb8ed37521ba05fdfad046cfbc3a772a.
//
// Solidity: event MinOperatorStakeSet(uint96 minOperatorStake)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchMinOperatorStakeSet(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSMinOperatorStakeSet) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "MinOperatorStakeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSMinOperatorStakeSet)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "MinOperatorStakeSet", log); err != nil {
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

// ParseMinOperatorStakeSet is a log parse operation binding the contract event 0x9b13fd38c94948514c63f3132d22b60ecb8ed37521ba05fdfad046cfbc3a772a.
//
// Solidity: event MinOperatorStakeSet(uint96 minOperatorStake)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseMinOperatorStakeSet(log types.Log) (*BindingISkateAVSMinOperatorStakeSet, error) {
	event := new(BindingISkateAVSMinOperatorStakeSet)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "MinOperatorStakeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorAddedIterator struct {
	Event *BindingISkateAVSOperatorAdded // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSOperatorAdded)
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
		it.Event = new(BindingISkateAVSOperatorAdded)
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
func (it *BindingISkateAVSOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSOperatorAdded represents a OperatorAdded event raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operator []common.Address) (*BindingISkateAVSOperatorAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSOperatorAddedIterator{contract: _BindingISkateAVS.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSOperatorAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSOperatorAdded)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseOperatorAdded(log types.Log) (*BindingISkateAVSOperatorAdded, error) {
	event := new(BindingISkateAVSOperatorAdded)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSOperatorAllowedIterator is returned from FilterOperatorAllowed and is used to iterate over the raw logs and unpacked data for OperatorAllowed events raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorAllowedIterator struct {
	Event *BindingISkateAVSOperatorAllowed // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSOperatorAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSOperatorAllowed)
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
		it.Event = new(BindingISkateAVSOperatorAllowed)
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
func (it *BindingISkateAVSOperatorAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSOperatorAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSOperatorAllowed represents a OperatorAllowed event raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorAllowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAllowed is a free log retrieval operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterOperatorAllowed(opts *bind.FilterOpts, operator []common.Address) (*BindingISkateAVSOperatorAllowedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "OperatorAllowed", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSOperatorAllowedIterator{contract: _BindingISkateAVS.contract, event: "OperatorAllowed", logs: logs, sub: sub}, nil
}

// WatchOperatorAllowed is a free log subscription operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchOperatorAllowed(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSOperatorAllowed, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "OperatorAllowed", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSOperatorAllowed)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
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

// ParseOperatorAllowed is a log parse operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseOperatorAllowed(log types.Log) (*BindingISkateAVSOperatorAllowed, error) {
	event := new(BindingISkateAVSOperatorAllowed)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSOperatorDisallowedIterator is returned from FilterOperatorDisallowed and is used to iterate over the raw logs and unpacked data for OperatorDisallowed events raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorDisallowedIterator struct {
	Event *BindingISkateAVSOperatorDisallowed // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSOperatorDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSOperatorDisallowed)
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
		it.Event = new(BindingISkateAVSOperatorDisallowed)
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
func (it *BindingISkateAVSOperatorDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSOperatorDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSOperatorDisallowed represents a OperatorDisallowed event raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorDisallowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDisallowed is a free log retrieval operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterOperatorDisallowed(opts *bind.FilterOpts, operator []common.Address) (*BindingISkateAVSOperatorDisallowedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "OperatorDisallowed", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSOperatorDisallowedIterator{contract: _BindingISkateAVS.contract, event: "OperatorDisallowed", logs: logs, sub: sub}, nil
}

// WatchOperatorDisallowed is a free log subscription operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchOperatorDisallowed(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSOperatorDisallowed, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "OperatorDisallowed", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSOperatorDisallowed)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
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

// ParseOperatorDisallowed is a log parse operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseOperatorDisallowed(log types.Log) (*BindingISkateAVSOperatorDisallowed, error) {
	event := new(BindingISkateAVSOperatorDisallowed)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorRemovedIterator struct {
	Event *BindingISkateAVSOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSOperatorRemoved)
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
		it.Event = new(BindingISkateAVSOperatorRemoved)
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
func (it *BindingISkateAVSOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSOperatorRemoved represents a OperatorRemoved event raised by the BindingISkateAVS contract.
type BindingISkateAVSOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*BindingISkateAVSOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSOperatorRemovedIterator{contract: _BindingISkateAVS.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSOperatorRemoved)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseOperatorRemoved(log types.Log) (*BindingISkateAVSOperatorRemoved, error) {
	event := new(BindingISkateAVSOperatorRemoved)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingISkateAVSStrategiesSetIterator is returned from FilterStrategiesSet and is used to iterate over the raw logs and unpacked data for StrategiesSet events raised by the BindingISkateAVS contract.
type BindingISkateAVSStrategiesSetIterator struct {
	Event *BindingISkateAVSStrategiesSet // Event containing the contract specifics and raw log

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
func (it *BindingISkateAVSStrategiesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingISkateAVSStrategiesSet)
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
		it.Event = new(BindingISkateAVSStrategiesSet)
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
func (it *BindingISkateAVSStrategiesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingISkateAVSStrategiesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingISkateAVSStrategiesSet represents a StrategiesSet event raised by the BindingISkateAVS contract.
type BindingISkateAVSStrategiesSet struct {
	Strategies []ISkateAVSStrategyParams
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategiesSet is a free log retrieval operation binding the contract event 0x5022676774eb0d9d4ee6bdfda494800505636c9c64dfd32bcf32cec97c563802.
//
// Solidity: event StrategiesSet((address,uint96)[] strategies)
func (_BindingISkateAVS *BindingISkateAVSFilterer) FilterStrategiesSet(opts *bind.FilterOpts) (*BindingISkateAVSStrategiesSetIterator, error) {

	logs, sub, err := _BindingISkateAVS.contract.FilterLogs(opts, "StrategiesSet")
	if err != nil {
		return nil, err
	}
	return &BindingISkateAVSStrategiesSetIterator{contract: _BindingISkateAVS.contract, event: "StrategiesSet", logs: logs, sub: sub}, nil
}

// WatchStrategiesSet is a free log subscription operation binding the contract event 0x5022676774eb0d9d4ee6bdfda494800505636c9c64dfd32bcf32cec97c563802.
//
// Solidity: event StrategiesSet((address,uint96)[] strategies)
func (_BindingISkateAVS *BindingISkateAVSFilterer) WatchStrategiesSet(opts *bind.WatchOpts, sink chan<- *BindingISkateAVSStrategiesSet) (event.Subscription, error) {

	logs, sub, err := _BindingISkateAVS.contract.WatchLogs(opts, "StrategiesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingISkateAVSStrategiesSet)
				if err := _BindingISkateAVS.contract.UnpackLog(event, "StrategiesSet", log); err != nil {
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

// ParseStrategiesSet is a log parse operation binding the contract event 0x5022676774eb0d9d4ee6bdfda494800505636c9c64dfd32bcf32cec97c563802.
//
// Solidity: event StrategiesSet((address,uint96)[] strategies)
func (_BindingISkateAVS *BindingISkateAVSFilterer) ParseStrategiesSet(log types.Log) (*BindingISkateAVSStrategiesSet, error) {
	event := new(BindingISkateAVSStrategiesSet)
	if err := _BindingISkateAVS.contract.UnpackLog(event, "StrategiesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
