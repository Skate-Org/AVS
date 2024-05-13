// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingSkateApp

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

// BindingSkateAppMetaData contains all meta data concerning the BindingSkateApp contract.
var BindingSkateAppMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMsg\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainType\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProof\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainType\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161049b38038061049b83398101604081905261002f91610054565b600280546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610408806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806311149ada1461004657806320d41de9146100785780638d9776721461008d575b600080fd5b61006661005436600461019c565b60009081526020819052604090205490565b60405190815260200160405180910390f35b61008b6100863660046101e4565b6100ad565b005b61006661009b36600461019c565b60006020819052908152604090205481565b600060405180606001604052808581526020018363ffffffff1681526020018463ffffffff168152509050806040516020016100e991906102fc565b60405160208183030381529060405280519060200120600080600160008154809291906101159061035a565b909155508152602081019190915260400160002055600180546101389190610373565b7f900c1bba72823a7d83b1db8670af9c73250734013566f7308581ee5044fffec66000806001805461016a9190610373565b8152602001908152602001600020548633878760405161018e95949392919061038c565b60405180910390a250505050565b6000602082840312156101ae57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b803563ffffffff811681146101df57600080fd5b919050565b6000806000606084860312156101f957600080fd5b833567ffffffffffffffff8082111561021157600080fd5b818601915086601f83011261022557600080fd5b813581811115610237576102376101b5565b604051601f8201601f19908116603f0116810190838211818310171561025f5761025f6101b5565b8160405282815289602084870101111561027857600080fd5b82602086016020830137600060208483010152809750505050505061029f602085016101cb565b91506102ad604085016101cb565b90509250925092565b6000815180845260005b818110156102dc576020818501810151868301820152016102c0565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600082516060602084015261031860808401826102b6565b9050602084015163ffffffff808216604086015280604087015116606086015250508091505092915050565b634e487b7160e01b600052601160045260246000fd5b60006001820161036c5761036c610344565b5060010190565b8181038181111561038657610386610344565b92915050565b85815260a0602082015260006103a560a08301876102b6565b6001600160a01b039590951660408301525063ffffffff928316606082015291166080909101529291505056fea2646970667358221220a2f753f59a7fcbefed901a007d956200a18e1de8414f09794496212c85aa10e364736f6c63430008140033",
}

// BindingSkateAppABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateAppMetaData.ABI instead.
var BindingSkateAppABI = BindingSkateAppMetaData.ABI

// BindingSkateAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateAppMetaData.Bin instead.
var BindingSkateAppBin = BindingSkateAppMetaData.Bin

// DeployBindingSkateApp deploys a new Ethereum contract, binding an instance of BindingSkateApp to it.
func DeployBindingSkateApp(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *BindingSkateApp, error) {
	parsed, err := BindingSkateAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateAppBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingSkateApp{BindingSkateAppCaller: BindingSkateAppCaller{contract: contract}, BindingSkateAppTransactor: BindingSkateAppTransactor{contract: contract}, BindingSkateAppFilterer: BindingSkateAppFilterer{contract: contract}}, nil
}

// BindingSkateApp is an auto generated Go binding around an Ethereum contract.
type BindingSkateApp struct {
	BindingSkateAppCaller     // Read-only binding to the contract
	BindingSkateAppTransactor // Write-only binding to the contract
	BindingSkateAppFilterer   // Log filterer for contract events
}

// BindingSkateAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingSkateAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingSkateAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingSkateAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSkateAppSession struct {
	Contract     *BindingSkateApp  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingSkateAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingSkateAppCallerSession struct {
	Contract *BindingSkateAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BindingSkateAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingSkateAppTransactorSession struct {
	Contract     *BindingSkateAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BindingSkateAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingSkateAppRaw struct {
	Contract *BindingSkateApp // Generic contract binding to access the raw methods on
}

// BindingSkateAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingSkateAppCallerRaw struct {
	Contract *BindingSkateAppCaller // Generic read-only contract binding to access the raw methods on
}

// BindingSkateAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingSkateAppTransactorRaw struct {
	Contract *BindingSkateAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingSkateApp creates a new instance of BindingSkateApp, bound to a specific deployed contract.
func NewBindingSkateApp(address common.Address, backend bind.ContractBackend) (*BindingSkateApp, error) {
	contract, err := bindBindingSkateApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingSkateApp{BindingSkateAppCaller: BindingSkateAppCaller{contract: contract}, BindingSkateAppTransactor: BindingSkateAppTransactor{contract: contract}, BindingSkateAppFilterer: BindingSkateAppFilterer{contract: contract}}, nil
}

// NewBindingSkateAppCaller creates a new read-only instance of BindingSkateApp, bound to a specific deployed contract.
func NewBindingSkateAppCaller(address common.Address, caller bind.ContractCaller) (*BindingSkateAppCaller, error) {
	contract, err := bindBindingSkateApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAppCaller{contract: contract}, nil
}

// NewBindingSkateAppTransactor creates a new write-only instance of BindingSkateApp, bound to a specific deployed contract.
func NewBindingSkateAppTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingSkateAppTransactor, error) {
	contract, err := bindBindingSkateApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAppTransactor{contract: contract}, nil
}

// NewBindingSkateAppFilterer creates a new log filterer instance of BindingSkateApp, bound to a specific deployed contract.
func NewBindingSkateAppFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingSkateAppFilterer, error) {
	contract, err := bindBindingSkateApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAppFilterer{contract: contract}, nil
}

// bindBindingSkateApp binds a generic wrapper to an already deployed contract.
func bindBindingSkateApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingSkateAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateApp *BindingSkateAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateApp.Contract.BindingSkateAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateApp *BindingSkateAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.BindingSkateAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateApp *BindingSkateAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.BindingSkateAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateApp *BindingSkateAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateApp *BindingSkateAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateApp *BindingSkateAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.contract.Transact(opts, method, params...)
}

// GetProof is a free data retrieval call binding the contract method 0x11149ada.
//
// Solidity: function getProof(uint256 taskId) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppCaller) GetProof(opts *bind.CallOpts, taskId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BindingSkateApp.contract.Call(opts, &out, "getProof", taskId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x11149ada.
//
// Solidity: function getProof(uint256 taskId) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppSession) GetProof(taskId *big.Int) ([32]byte, error) {
	return _BindingSkateApp.Contract.GetProof(&_BindingSkateApp.CallOpts, taskId)
}

// GetProof is a free data retrieval call binding the contract method 0x11149ada.
//
// Solidity: function getProof(uint256 taskId) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppCallerSession) GetProof(taskId *big.Int) ([32]byte, error) {
	return _BindingSkateApp.Contract.GetProof(&_BindingSkateApp.CallOpts, taskId)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BindingSkateApp.contract.Call(opts, &out, "tasks", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppSession) Tasks(arg0 *big.Int) ([32]byte, error) {
	return _BindingSkateApp.Contract.Tasks(&_BindingSkateApp.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(bytes32)
func (_BindingSkateApp *BindingSkateAppCallerSession) Tasks(arg0 *big.Int) ([32]byte, error) {
	return _BindingSkateApp.Contract.Tasks(&_BindingSkateApp.CallOpts, arg0)
}

// CreateMsg is a paid mutator transaction binding the contract method 0x20d41de9.
//
// Solidity: function createMsg(string message, uint32 chainType, uint32 chainId) returns()
func (_BindingSkateApp *BindingSkateAppTransactor) CreateMsg(opts *bind.TransactOpts, message string, chainType uint32, chainId uint32) (*types.Transaction, error) {
	return _BindingSkateApp.contract.Transact(opts, "createMsg", message, chainType, chainId)
}

// CreateMsg is a paid mutator transaction binding the contract method 0x20d41de9.
//
// Solidity: function createMsg(string message, uint32 chainType, uint32 chainId) returns()
func (_BindingSkateApp *BindingSkateAppSession) CreateMsg(message string, chainType uint32, chainId uint32) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.CreateMsg(&_BindingSkateApp.TransactOpts, message, chainType, chainId)
}

// CreateMsg is a paid mutator transaction binding the contract method 0x20d41de9.
//
// Solidity: function createMsg(string message, uint32 chainType, uint32 chainId) returns()
func (_BindingSkateApp *BindingSkateAppTransactorSession) CreateMsg(message string, chainType uint32, chainId uint32) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.CreateMsg(&_BindingSkateApp.TransactOpts, message, chainType, chainId)
}

// BindingSkateAppTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the BindingSkateApp contract.
type BindingSkateAppTaskCreatedIterator struct {
	Event *BindingSkateAppTaskCreated // Event containing the contract specifics and raw log

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
func (it *BindingSkateAppTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateAppTaskCreated)
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
		it.Event = new(BindingSkateAppTaskCreated)
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
func (it *BindingSkateAppTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateAppTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateAppTaskCreated represents a TaskCreated event raised by the BindingSkateApp contract.
type BindingSkateAppTaskCreated struct {
	TaskId    *big.Int
	TaskHash  [32]byte
	Message   string
	Signer    common.Address
	ChainType uint32
	ChainId   uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x900c1bba72823a7d83b1db8670af9c73250734013566f7308581ee5044fffec6.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chainType, uint32 chainId)
func (_BindingSkateApp *BindingSkateAppFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskId []*big.Int) (*BindingSkateAppTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BindingSkateApp.contract.FilterLogs(opts, "TaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAppTaskCreatedIterator{contract: _BindingSkateApp.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x900c1bba72823a7d83b1db8670af9c73250734013566f7308581ee5044fffec6.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chainType, uint32 chainId)
func (_BindingSkateApp *BindingSkateAppFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *BindingSkateAppTaskCreated, taskId []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BindingSkateApp.contract.WatchLogs(opts, "TaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateAppTaskCreated)
				if err := _BindingSkateApp.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x900c1bba72823a7d83b1db8670af9c73250734013566f7308581ee5044fffec6.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chainType, uint32 chainId)
func (_BindingSkateApp *BindingSkateAppFilterer) ParseTaskCreated(log types.Log) (*BindingSkateAppTaskCreated, error) {
	event := new(BindingSkateAppTaskCreated)
	if err := _BindingSkateApp.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
