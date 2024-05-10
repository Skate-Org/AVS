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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMsg\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chain\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProof\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161046638038061046683398101604081905261002f91610054565b600280546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6103d3806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806311149ada146100465780638d97767214610078578063a7fcc6ab14610098575b600080fd5b61006661005436600461018d565b60009081526020819052604090205490565b60405190815260200160405180910390f35b61006661008636600461018d565b60006020819052908152604090205481565b6100ab6100a63660046101d5565b6100ad565b005b600060405180604001604052808481526020018363ffffffff168152509050806040516020016100dd91906102dd565b60405160208183030381529060405280519060200120600080600160008154809291906101099061032b565b9091555081526020810191909152604001600020556001805461012c9190610344565b7ff4334a6b82c42b3da48195ff8f637cc42e08d795f4df38d4313dbed7b00f201a6000806001805461015e9190610344565b815260200190815260200160002054853386604051610180949392919061035d565b60405180910390a2505050565b60006020828403121561019f57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b803563ffffffff811681146101d057600080fd5b919050565b600080604083850312156101e857600080fd5b823567ffffffffffffffff8082111561020057600080fd5b818501915085601f83011261021457600080fd5b813581811115610226576102266101a6565b604051601f8201601f19908116603f0116810190838211818310171561024e5761024e6101a6565b8160405282815288602084870101111561026757600080fd5b82602086016020830137600060208483010152809650505050505061028e602084016101bc565b90509250929050565b6000815180845260005b818110156102bd576020818501810151868301820152016102a1565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260008251604060208401526102f96060840182610297565b905063ffffffff60208501511660408401528091505092915050565b634e487b7160e01b600052601160045260246000fd5b60006001820161033d5761033d610315565b5060010190565b8181038181111561035757610357610315565b92915050565b8481526080602082015260006103766080830186610297565b6001600160a01b039490941660408301525063ffffffff919091166060909101529291505056fea2646970667358221220e0d047844d3f0e4da26488178e75f5f68ddf2f8f99c01168a7ac3bcaf24aca3464736f6c63430008140033",
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

// CreateMsg is a paid mutator transaction binding the contract method 0xa7fcc6ab.
//
// Solidity: function createMsg(string message, uint32 chain) returns()
func (_BindingSkateApp *BindingSkateAppTransactor) CreateMsg(opts *bind.TransactOpts, message string, chain uint32) (*types.Transaction, error) {
	return _BindingSkateApp.contract.Transact(opts, "createMsg", message, chain)
}

// CreateMsg is a paid mutator transaction binding the contract method 0xa7fcc6ab.
//
// Solidity: function createMsg(string message, uint32 chain) returns()
func (_BindingSkateApp *BindingSkateAppSession) CreateMsg(message string, chain uint32) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.CreateMsg(&_BindingSkateApp.TransactOpts, message, chain)
}

// CreateMsg is a paid mutator transaction binding the contract method 0xa7fcc6ab.
//
// Solidity: function createMsg(string message, uint32 chain) returns()
func (_BindingSkateApp *BindingSkateAppTransactorSession) CreateMsg(message string, chain uint32) (*types.Transaction, error) {
	return _BindingSkateApp.Contract.CreateMsg(&_BindingSkateApp.TransactOpts, message, chain)
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
	TaskId   *big.Int
	TaskHash [32]byte
	Message  string
	Signer   common.Address
	Chain    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xf4334a6b82c42b3da48195ff8f637cc42e08d795f4df38d4313dbed7b00f201a.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chain)
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

// WatchTaskCreated is a free log subscription operation binding the contract event 0xf4334a6b82c42b3da48195ff8f637cc42e08d795f4df38d4313dbed7b00f201a.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chain)
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

// ParseTaskCreated is a log parse operation binding the contract event 0xf4334a6b82c42b3da48195ff8f637cc42e08d795f4df38d4313dbed7b00f201a.
//
// Solidity: event TaskCreated(uint256 indexed taskId, bytes32 taskHash, string message, address signer, uint32 chain)
func (_BindingSkateApp *BindingSkateAppFilterer) ParseTaskCreated(log types.Log) (*BindingSkateAppTaskCreated, error) {
	event := new(BindingSkateAppTaskCreated)
	if err := _BindingSkateApp.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
