// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// L1CustomERC20GatewayMetaData contains all meta data concerning the L1CustomERC20Gateway contract.
var (
	L1CustomERC20GatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DepositERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeWithdrawERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"RefundERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"oldL2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newL2Token\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20AndCall\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeWithdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"onDropMessage\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateTokenMapping\"}]",
	}
	// L1CustomERC20GatewayABI is the input ABI used to generate the binding from.
	L1CustomERC20GatewayABI, _ = L1CustomERC20GatewayMetaData.GetAbi()
)

// L1CustomERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L1CustomERC20Gateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1CustomERC20GatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1CustomERC20GatewayCaller     // Read-only binding to the contract
	L1CustomERC20GatewayTransactor // Write-only binding to the contract
}

// GetName return L1CustomERC20Gateway's Name.
func (o *L1CustomERC20Gateway) GetName() string {
	return o.Name
}

// GetAddress return L1CustomERC20Gateway's contract address.
func (o *L1CustomERC20Gateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1CustomERC20Gateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1CustomERC20Gateway) GetSigHashes() []common.Hash {
	if len(o.parsers) == 0 {
		return nil
	}
	var sigHashes = make([]common.Hash, 0, len(o.parsers))
	for id := range o.parsers {
		sigHashes = append(sigHashes, id)
	}
	return sigHashes
}

// ParseLog parse the log if parse func is exist.
func (o *L1CustomERC20Gateway) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterDepositERC20, the DepositERC20 event ID is 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
func (o *L1CustomERC20Gateway) RegisterDepositERC20(handler func(vLog *types.Log, data *L1CustomERC20GatewayDepositERC20Event) error) {
	_id := o.ABI.Events["DepositERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayDepositERC20Event)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "DepositERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DepositERC20"
}

// RegisterFinalizeWithdrawERC20, the FinalizeWithdrawERC20 event ID is 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
func (o *L1CustomERC20Gateway) RegisterFinalizeWithdrawERC20(handler func(vLog *types.Log, data *L1CustomERC20GatewayFinalizeWithdrawERC20Event) error) {
	_id := o.ABI.Events["FinalizeWithdrawERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayFinalizeWithdrawERC20Event)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "FinalizeWithdrawERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeWithdrawERC20"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1CustomERC20Gateway) RegisterInitialized(handler func(vLog *types.Log, data *L1CustomERC20GatewayInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayInitializedEvent)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1CustomERC20Gateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1CustomERC20GatewayOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayOwnershipTransferredEvent)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterRefundERC20, the RefundERC20 event ID is 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
func (o *L1CustomERC20Gateway) RegisterRefundERC20(handler func(vLog *types.Log, data *L1CustomERC20GatewayRefundERC20Event) error) {
	_id := o.ABI.Events["RefundERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayRefundERC20Event)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "RefundERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RefundERC20"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L1CustomERC20Gateway) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L1CustomERC20GatewayUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayUpdateRateLimiterEvent)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// RegisterUpdateTokenMapping, the UpdateTokenMapping event ID is 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
func (o *L1CustomERC20Gateway) RegisterUpdateTokenMapping(handler func(vLog *types.Log, data *L1CustomERC20GatewayUpdateTokenMappingEvent) error) {
	_id := o.ABI.Events["UpdateTokenMapping"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1CustomERC20GatewayUpdateTokenMappingEvent)
		if err := o.L1CustomERC20GatewayCaller.contract.UnpackLog(event, "UpdateTokenMapping", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateTokenMapping"
}

// L1CustomERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1CustomERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1CustomERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1CustomERC20Gateway creates a new instance of L1CustomERC20Gateway, bound to a specific deployed contract.
func NewL1CustomERC20Gateway(address common.Address, backend bind.ContractBackend) (*L1CustomERC20Gateway, error) {
	contract, err := bindL1CustomERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1CustomERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1CustomERC20Gateway{
		Name:                           "L1CustomERC20Gateway",
		ABI:                            sigAbi,
		Address:                        address,
		topics:                         topics,
		parsers:                        map[common.Hash]func(log *types.Log) error{},
		L1CustomERC20GatewayCaller:     L1CustomERC20GatewayCaller{contract: contract},
		L1CustomERC20GatewayTransactor: L1CustomERC20GatewayTransactor{contract: contract},
	}, nil
}

// NewL1CustomERC20GatewayCaller creates a new read-only instance of L1CustomERC20Gateway, bound to a specific deployed contract.
func NewL1CustomERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1CustomERC20GatewayCaller, error) {
	contract, err := bindL1CustomERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20GatewayCaller{contract: contract}, nil
}

// NewL1CustomERC20GatewayTransactor creates a new write-only instance of L1CustomERC20Gateway, bound to a specific deployed contract.
func NewL1CustomERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1CustomERC20GatewayTransactor, error) {
	contract, err := bindL1CustomERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20GatewayTransactor{contract: contract}, nil
}

// bindL1CustomERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL1CustomERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1CustomERC20GatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1CustomERC20Gateway *L1CustomERC20GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "onDropMessage", _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1CustomERC20Gateway *L1CustomERC20GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1CustomERC20Gateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// L1CustomERC20GatewayDepositERC20 represents a DepositERC20 event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayDepositERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L1CustomERC20GatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayFinalizeWithdrawERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L1CustomERC20GatewayInitialized represents a Initialized event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1CustomERC20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1CustomERC20GatewayRefundERC20 represents a RefundERC20 event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayRefundERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
}

// L1CustomERC20GatewayUpdateRateLimiter represents a UpdateRateLimiter event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}

// L1CustomERC20GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1CustomERC20Gateway contract.
type L1CustomERC20GatewayUpdateTokenMappingEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
}
