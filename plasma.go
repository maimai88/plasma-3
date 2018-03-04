// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plasma

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PlasmaABI is the input ABI used to generate the binding from.
const PlasmaABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"depositor\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"__init__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"submitBlock\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"root\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":81689},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"tx\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":103504},{\"name\":\"authority\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":543},{\"name\":\"last_child_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":573},{\"name\":\"last_parent_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":603},{\"name\":\"child_chain__root\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":900},{\"name\":\"child_chain__created_at\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":924}]"

// PlasmaBin is the compiled bytecode used for deploying new contracts.
const PlasmaBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b3360005560016001554360025561090556600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263baa47694600051141561013857602060046101403734156100b457600080fd5b60005433146100c257600080fd5b606051600660015401806040519013585780919012156100e157600080fd5b43146100ec57600080fd5b600154600360c052602060c0200160c052602060c02042815561014051600182015550600160605160018254018060405190135857809190121561012f57600080fd5b81555043600255005b6398b1e06a600051141561071f5760206004610140376104206004356004016101603761040060043560040135111561017057600080fd5b610a206101606104406105a08251602084016000735185d17c44699cecc3133114f8df70753b856709611720f150506101806105a051146101b057600080fd5b6105a0516105a0018060200151600082518060209013585780919012156101d657600080fd5b601f6101000a82048115176101ea57600080fd5b606051816020036101000a83048060405190135857809190121561020d57600080fd5b90509050905081526105c0516105a00180602001516000825180602090135857809190121561023b57600080fd5b601f6101000a820481151761024f57600080fd5b606051816020036101000a83048060405190135857809190121561027257600080fd5b90509050905081602001526105e0516105a0018060200151600082518060209013585780919012156102a357600080fd5b601f6101000a82048115176102b757600080fd5b606051816020036101000a8304806040519013585780919012156102da57600080fd5b9050905090508160400152610600516105a00180602001516000825180602090135857809190121561030b57600080fd5b601f6101000a820481151761031f57600080fd5b606051816020036101000a83048060405190135857809190121561034257600080fd5b9050905090508160600152610620516105a00180602001516000825180602090135857809190121561037357600080fd5b601f6101000a820481151761038757600080fd5b606051816020036101000a8304806040519013585780919012156103aa57600080fd5b9050905090508160800152610640516105a0018060200151600082518060209013585780919012156103db57600080fd5b601f6101000a82048115176103ef57600080fd5b606051816020036101000a83048060405190135857809190121561041257600080fd5b9050905090508160a001526014610660516105a001511461043257600080fd5b602051610660516105b40151068160c00152610680516105a00180602001516000825180602090135857809190121561046a57600080fd5b601f6101000a820481151761047e57600080fd5b606051816020036101000a8304806040519013585780919012156104a157600080fd5b9050905090508160e0015260146106a0516105a00151146104c157600080fd5b6020516106a0516105b40151068161010001526106c0516105a0018060200151600082518060209013585780919012156104fa57600080fd5b601f6101000a820481151761050e57600080fd5b606051816020036101000a83048060405190135857809190121561053157600080fd5b9050905090508161012001526106e0516105a00180602001516000825180602090135857809190121561056357600080fd5b601f6101000a820481151761057757600080fd5b606051816020036101000a83048060405190135857809190121561059a57600080fd5b90509050905081610140015250600061016061040080602084610bc001018260208501600060046078f1505080518201915050610b8051602082610bc001015260208101905080610bc052610bc09050805160208201209050610ba05261102060006010818352015b6000610ba051602082611040010152602081019050610b805160208261104001015260208101905080611040526110409050805160208201209050610ba0526000610b80516020826110c0010152602081019050610b80516020826110c0010152602081019050806110c0526110c09050805160208201209050610b80525b8151600101808352811415610603575b5050600154600360c052602060c0200160c052602060c020428155610b805160018201555060016060516001825401806040519013585780919012156106d757600080fd5b81555043600255610ae05161118052602061114052610ae051611180527f8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef123461114051611180a1005b63bf7e214f600051141561074557341561073857600080fd5b60005460005260206000f3005b6324c5ad9a600051141561076b57341561075e57600080fd5b60015460005260206000f3005b635258b093600051141561079157341561078457600080fd5b60025460005260206000f3005b6324f00f5b60005114156107f457602060046101403734156107b257600080fd5b606051600435806040519013585780919012156107ce57600080fd5b50600161014051600360c052602060c0200160c052602060c020015460005260206000f3005b6388c73d2d6000511415610854576020600461014037341561081557600080fd5b6060516004358060405190135857809190121561083157600080fd5b5061014051600360c052602060c0200160c052602060c0205460005260206000f3005b5b6100b0610905036100b06000396100b0610905036000f3`

// DeployPlasma deploys a new Ethereum contract, binding an instance of Plasma to it.
func DeployPlasma(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Plasma, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Plasma{PlasmaCaller: PlasmaCaller{contract: contract}, PlasmaTransactor: PlasmaTransactor{contract: contract}, PlasmaFilterer: PlasmaFilterer{contract: contract}}, nil
}

// Plasma is an auto generated Go binding around an Ethereum contract.
type Plasma struct {
	PlasmaCaller     // Read-only binding to the contract
	PlasmaTransactor // Write-only binding to the contract
	PlasmaFilterer   // Log filterer for contract events
}

// PlasmaCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaSession struct {
	Contract     *Plasma           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaCallerSession struct {
	Contract *PlasmaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlasmaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaTransactorSession struct {
	Contract     *PlasmaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaRaw struct {
	Contract *Plasma // Generic contract binding to access the raw methods on
}

// PlasmaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaCallerRaw struct {
	Contract *PlasmaCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaTransactorRaw struct {
	Contract *PlasmaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasma creates a new instance of Plasma, bound to a specific deployed contract.
func NewPlasma(address common.Address, backend bind.ContractBackend) (*Plasma, error) {
	contract, err := bindPlasma(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Plasma{PlasmaCaller: PlasmaCaller{contract: contract}, PlasmaTransactor: PlasmaTransactor{contract: contract}, PlasmaFilterer: PlasmaFilterer{contract: contract}}, nil
}

// NewPlasmaCaller creates a new read-only instance of Plasma, bound to a specific deployed contract.
func NewPlasmaCaller(address common.Address, caller bind.ContractCaller) (*PlasmaCaller, error) {
	contract, err := bindPlasma(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaCaller{contract: contract}, nil
}

// NewPlasmaTransactor creates a new write-only instance of Plasma, bound to a specific deployed contract.
func NewPlasmaTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaTransactor, error) {
	contract, err := bindPlasma(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaTransactor{contract: contract}, nil
}

// NewPlasmaFilterer creates a new log filterer instance of Plasma, bound to a specific deployed contract.
func NewPlasmaFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaFilterer, error) {
	contract, err := bindPlasma(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaFilterer{contract: contract}, nil
}

// bindPlasma binds a generic wrapper to an already deployed contract.
func bindPlasma(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plasma *PlasmaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Plasma.Contract.PlasmaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plasma *PlasmaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plasma.Contract.PlasmaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plasma *PlasmaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plasma.Contract.PlasmaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plasma *PlasmaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Plasma.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plasma *PlasmaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plasma.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plasma *PlasmaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plasma.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaSession) Authority() (common.Address, error) {
	return _Plasma.Contract.Authority(&_Plasma.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaCallerSession) Authority() (common.Address, error) {
	return _Plasma.Contract.Authority(&_Plasma.CallOpts)
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCaller) Child_chain__created_at(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "child_chain__created_at", arg0)
	return *ret0, err
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaSession) Child_chain__created_at(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Child_chain__created_at(&_Plasma.CallOpts, arg0)
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Child_chain__created_at(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Child_chain__created_at(&_Plasma.CallOpts, arg0)
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaCaller) Child_chain__root(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "child_chain__root", arg0)
	return *ret0, err
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaSession) Child_chain__root(arg0 *big.Int) ([32]byte, error) {
	return _Plasma.Contract.Child_chain__root(&_Plasma.CallOpts, arg0)
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaCallerSession) Child_chain__root(arg0 *big.Int) ([32]byte, error) {
	return _Plasma.Contract.Child_chain__root(&_Plasma.CallOpts, arg0)
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaCaller) Last_child_block(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "last_child_block")
	return *ret0, err
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaSession) Last_child_block() (*big.Int, error) {
	return _Plasma.Contract.Last_child_block(&_Plasma.CallOpts)
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Last_child_block() (*big.Int, error) {
	return _Plasma.Contract.Last_child_block(&_Plasma.CallOpts)
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaCaller) Last_parent_block(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "last_parent_block")
	return *ret0, err
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaSession) Last_parent_block() (*big.Int, error) {
	return _Plasma.Contract.Last_parent_block(&_Plasma.CallOpts)
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Last_parent_block() (*big.Int, error) {
	return _Plasma.Contract.Last_parent_block(&_Plasma.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaTransactor) Deposit(opts *bind.TransactOpts, tx []byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "deposit", tx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaSession) Deposit(tx []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Deposit(&_Plasma.TransactOpts, tx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaTransactorSession) Deposit(tx []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Deposit(&_Plasma.TransactOpts, tx)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "submitBlock", root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _Plasma.Contract.SubmitBlock(&_Plasma.TransactOpts, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaTransactorSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _Plasma.Contract.SubmitBlock(&_Plasma.TransactOpts, root)
}

// PlasmaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Plasma contract.
type PlasmaDepositIterator struct {
	Event *PlasmaDeposit // Event containing the contract specifics and raw log

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
func (it *PlasmaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaDeposit)
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
		it.Event = new(PlasmaDeposit)
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
func (it *PlasmaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaDeposit represents a Deposit event raised by the Plasma contract.
type PlasmaDeposit struct {
	Depositor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(depositor address)
func (_Plasma *PlasmaFilterer) FilterDeposit(opts *bind.FilterOpts) (*PlasmaDepositIterator, error) {

	logs, sub, err := _Plasma.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PlasmaDepositIterator{contract: _Plasma.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(depositor address)
func (_Plasma *PlasmaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PlasmaDeposit) (event.Subscription, error) {

	logs, sub, err := _Plasma.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaDeposit)
				if err := _Plasma.contract.UnpackLog(event, "Deposit", log); err != nil {
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
