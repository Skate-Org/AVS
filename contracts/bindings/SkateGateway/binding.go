// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingSkateGateway

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

// BindingSkateGatewayMetaData contains all meta data concerning the BindingSkateGateway contract.
var BindingSkateGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMsg\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messages\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postMsg\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRelayer\",\"inputs\":[{\"name\":\"newRelayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106cf3803806106cf83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61063c806100936000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630d80fefd1461005c5780632b82735b146100865780633fcfdc1d146100a657806356f1612f146100bb5780638406c079146100ce575b600080fd5b61006f61006a366004610302565b6100f9565b60405161007d929190610361565b60405180910390f35b610099610094366004610302565b6101a6565b60405161007d919061038b565b6100b96100b43660046103d7565b610248565b005b6100b96100c93660046104a2565b6102c9565b6002546100e1906001600160a01b031681565b6040516001600160a01b03909116815260200161007d565b600160205260009081526040902080548190610114906104bd565b80601f0160208091040260200160405190810160405280929190818152602001828054610140906104bd565b801561018d5780601f106101625761010080835404028352916020019161018d565b820191906000526020600020905b81548152906001019060200180831161017057829003601f168201915b505050600190930154919250506001600160a01b031682565b60008181526001602052604090208054606091906101c3906104bd565b80601f01602080910402602001604051908101604052809291908181526020018280546101ef906104bd565b801561023c5780601f106102115761010080835404028352916020019161023c565b820191906000526020600020905b81548152906001019060200180831161021f57829003601f168201915b50505050509050919050565b6002546001600160a01b0316331461025f57600080fd5b6040805180820182528381526001600160a01b03831660208083019190915260008681526001909152919091208151819061029a9082610546565b5060209190910151600190910180546001600160a01b0319166001600160a01b03909216919091179055505050565b6000546001600160a01b031633146102e057600080fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60006020828403121561031457600080fd5b5035919050565b6000815180845260005b8181101561034157602081850181015186830182015201610325565b506000602082860101526020601f19601f83011685010191505092915050565b604081526000610374604083018561031b565b905060018060a01b03831660208301529392505050565b60208152600061039e602083018461031b565b9392505050565b634e487b7160e01b600052604160045260246000fd5b80356001600160a01b03811681146103d257600080fd5b919050565b6000806000606084860312156103ec57600080fd5b83359250602084013567ffffffffffffffff8082111561040b57600080fd5b818601915086601f83011261041f57600080fd5b813581811115610431576104316103a5565b604051601f8201601f19908116603f01168101908382118183101715610459576104596103a5565b8160405282815289602084870101111561047257600080fd5b826020860160208301376000602084830101528096505050505050610499604085016103bb565b90509250925092565b6000602082840312156104b457600080fd5b61039e826103bb565b600181811c908216806104d157607f821691505b6020821081036104f157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561054157600081815260208120601f850160051c8101602086101561051e5750805b601f850160051c820191505b8181101561053d5782815560010161052a565b5050505b505050565b815167ffffffffffffffff811115610560576105606103a5565b6105748161056e84546104bd565b846104f7565b602080601f8311600181146105a957600084156105915750858301515b600019600386901b1c1916600185901b17855561053d565b600085815260208120601f198616915b828110156105d8578886015182559484019460019091019084016105b9565b50858210156105f65787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212204584a9e9a460b2bce571c3ae5756cf81b280e346d3486837797c8cfeb1f3f30264736f6c63430008140033",
}

// BindingSkateGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateGatewayMetaData.ABI instead.
var BindingSkateGatewayABI = BindingSkateGatewayMetaData.ABI

// BindingSkateGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateGatewayMetaData.Bin instead.
var BindingSkateGatewayBin = BindingSkateGatewayMetaData.Bin

// DeployBindingSkateGateway deploys a new Ethereum contract, binding an instance of BindingSkateGateway to it.
func DeployBindingSkateGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *BindingSkateGateway, error) {
	parsed, err := BindingSkateGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateGatewayBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingSkateGateway{BindingSkateGatewayCaller: BindingSkateGatewayCaller{contract: contract}, BindingSkateGatewayTransactor: BindingSkateGatewayTransactor{contract: contract}, BindingSkateGatewayFilterer: BindingSkateGatewayFilterer{contract: contract}}, nil
}

// BindingSkateGateway is an auto generated Go binding around an Ethereum contract.
type BindingSkateGateway struct {
	BindingSkateGatewayCaller     // Read-only binding to the contract
	BindingSkateGatewayTransactor // Write-only binding to the contract
	BindingSkateGatewayFilterer   // Log filterer for contract events
}

// BindingSkateGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingSkateGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingSkateGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingSkateGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSkateGatewaySession struct {
	Contract     *BindingSkateGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BindingSkateGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingSkateGatewayCallerSession struct {
	Contract *BindingSkateGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BindingSkateGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingSkateGatewayTransactorSession struct {
	Contract     *BindingSkateGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BindingSkateGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingSkateGatewayRaw struct {
	Contract *BindingSkateGateway // Generic contract binding to access the raw methods on
}

// BindingSkateGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingSkateGatewayCallerRaw struct {
	Contract *BindingSkateGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// BindingSkateGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingSkateGatewayTransactorRaw struct {
	Contract *BindingSkateGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingSkateGateway creates a new instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGateway(address common.Address, backend bind.ContractBackend) (*BindingSkateGateway, error) {
	contract, err := bindBindingSkateGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGateway{BindingSkateGatewayCaller: BindingSkateGatewayCaller{contract: contract}, BindingSkateGatewayTransactor: BindingSkateGatewayTransactor{contract: contract}, BindingSkateGatewayFilterer: BindingSkateGatewayFilterer{contract: contract}}, nil
}

// NewBindingSkateGatewayCaller creates a new read-only instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayCaller(address common.Address, caller bind.ContractCaller) (*BindingSkateGatewayCaller, error) {
	contract, err := bindBindingSkateGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayCaller{contract: contract}, nil
}

// NewBindingSkateGatewayTransactor creates a new write-only instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingSkateGatewayTransactor, error) {
	contract, err := bindBindingSkateGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayTransactor{contract: contract}, nil
}

// NewBindingSkateGatewayFilterer creates a new log filterer instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingSkateGatewayFilterer, error) {
	contract, err := bindBindingSkateGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayFilterer{contract: contract}, nil
}

// bindBindingSkateGateway binds a generic wrapper to an already deployed contract.
func bindBindingSkateGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingSkateGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateGateway.Contract.BindingSkateGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.BindingSkateGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.BindingSkateGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateGateway *BindingSkateGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateGateway *BindingSkateGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateGateway *BindingSkateGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.contract.Transact(opts, method, params...)
}

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewayCaller) GetMsg(opts *bind.CallOpts, taskId *big.Int) (string, error) {
	var out []interface{}
	err := _BindingSkateGateway.contract.Call(opts, &out, "getMsg", taskId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewaySession) GetMsg(taskId *big.Int) (string, error) {
	return _BindingSkateGateway.Contract.GetMsg(&_BindingSkateGateway.CallOpts, taskId)
}

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewayCallerSession) GetMsg(taskId *big.Int) (string, error) {
	return _BindingSkateGateway.Contract.GetMsg(&_BindingSkateGateway.CallOpts, taskId)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewayCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	var out []interface{}
	err := _BindingSkateGateway.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		Message string
		Signer  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Message = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewaySession) Messages(arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	return _BindingSkateGateway.Contract.Messages(&_BindingSkateGateway.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewayCallerSession) Messages(arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	return _BindingSkateGateway.Contract.Messages(&_BindingSkateGateway.CallOpts, arg0)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_BindingSkateGateway *BindingSkateGatewayCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateGateway.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_BindingSkateGateway *BindingSkateGatewaySession) Relayer() (common.Address, error) {
	return _BindingSkateGateway.Contract.Relayer(&_BindingSkateGateway.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_BindingSkateGateway *BindingSkateGatewayCallerSession) Relayer() (common.Address, error) {
	return _BindingSkateGateway.Contract.Relayer(&_BindingSkateGateway.CallOpts)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactor) PostMsg(opts *bind.TransactOpts, taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.contract.Transact(opts, "postMsg", taskId, message, signer)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewaySession) PostMsg(taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.PostMsg(&_BindingSkateGateway.TransactOpts, taskId, message, signer)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactorSession) PostMsg(taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.PostMsg(&_BindingSkateGateway.TransactOpts, taskId, message, signer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x56f1612f.
//
// Solidity: function registerRelayer(address newRelayer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactor) RegisterRelayer(opts *bind.TransactOpts, newRelayer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.contract.Transact(opts, "registerRelayer", newRelayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x56f1612f.
//
// Solidity: function registerRelayer(address newRelayer) returns()
func (_BindingSkateGateway *BindingSkateGatewaySession) RegisterRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.RegisterRelayer(&_BindingSkateGateway.TransactOpts, newRelayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x56f1612f.
//
// Solidity: function registerRelayer(address newRelayer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactorSession) RegisterRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.RegisterRelayer(&_BindingSkateGateway.TransactOpts, newRelayer)
}
