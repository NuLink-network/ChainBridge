// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package policy

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PolicyABI is the input ABI used to generate the binding from.
const PolicyABI = "[{\"inputs\":[{\"internalType\":\"contractIStaking\",\"name\":\"_staking\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"cancelPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_rewardBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"createPolicy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"giveOutReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policyIDList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policyInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"policyID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policyRewardedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicySession) Owner() (common.Address, error) {
	return _Policy.Contract.Owner(&_Policy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicyCallerSession) Owner() (common.Address, error) {
	return _Policy.Contract.Owner(&_Policy.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x717ab112.
//
// Solidity: function pID() view returns(uint256)
func (_Policy *PolicyCaller) PID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "pID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PID is a free data retrieval call binding the contract method 0x717ab112.
//
// Solidity: function pID() view returns(uint256)
func (_Policy *PolicySession) PID() (*big.Int, error) {
	return _Policy.Contract.PID(&_Policy.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x717ab112.
//
// Solidity: function pID() view returns(uint256)
func (_Policy *PolicyCallerSession) PID() (*big.Int, error) {
	return _Policy.Contract.PID(&_Policy.CallOpts)
}

// PolicyIDList is a free data retrieval call binding the contract method 0x922fdd18.
//
// Solidity: function policyIDList(uint256 ) view returns(uint256)
func (_Policy *PolicyCaller) PolicyIDList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "policyIDList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolicyIDList is a free data retrieval call binding the contract method 0x922fdd18.
//
// Solidity: function policyIDList(uint256 ) view returns(uint256)
func (_Policy *PolicySession) PolicyIDList(arg0 *big.Int) (*big.Int, error) {
	return _Policy.Contract.PolicyIDList(&_Policy.CallOpts, arg0)
}

// PolicyIDList is a free data retrieval call binding the contract method 0x922fdd18.
//
// Solidity: function policyIDList(uint256 ) view returns(uint256)
func (_Policy *PolicyCallerSession) PolicyIDList(arg0 *big.Int) (*big.Int, error) {
	return _Policy.Contract.PolicyIDList(&_Policy.CallOpts, arg0)
}

// PolicyInfo is a free data retrieval call binding the contract method 0x576be6c5.
//
// Solidity: function policyInfo(uint256 ) view returns(uint256 policyID, address creater, uint256 rewardBalance, uint256 begin, uint256 end, uint256 cancelBlock)
func (_Policy *PolicyCaller) PolicyInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PolicyID      *big.Int
	Creater       common.Address
	RewardBalance *big.Int
	Begin         *big.Int
	End           *big.Int
	CancelBlock   *big.Int
}, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "policyInfo", arg0)

	outstruct := new(struct {
		PolicyID      *big.Int
		Creater       common.Address
		RewardBalance *big.Int
		Begin         *big.Int
		End           *big.Int
		CancelBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PolicyID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creater = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RewardBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Begin = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CancelBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PolicyInfo is a free data retrieval call binding the contract method 0x576be6c5.
//
// Solidity: function policyInfo(uint256 ) view returns(uint256 policyID, address creater, uint256 rewardBalance, uint256 begin, uint256 end, uint256 cancelBlock)
func (_Policy *PolicySession) PolicyInfo(arg0 *big.Int) (struct {
	PolicyID      *big.Int
	Creater       common.Address
	RewardBalance *big.Int
	Begin         *big.Int
	End           *big.Int
	CancelBlock   *big.Int
}, error) {
	return _Policy.Contract.PolicyInfo(&_Policy.CallOpts, arg0)
}

// PolicyInfo is a free data retrieval call binding the contract method 0x576be6c5.
//
// Solidity: function policyInfo(uint256 ) view returns(uint256 policyID, address creater, uint256 rewardBalance, uint256 begin, uint256 end, uint256 cancelBlock)
func (_Policy *PolicyCallerSession) PolicyInfo(arg0 *big.Int) (struct {
	PolicyID      *big.Int
	Creater       common.Address
	RewardBalance *big.Int
	Begin         *big.Int
	End           *big.Int
	CancelBlock   *big.Int
}, error) {
	return _Policy.Contract.PolicyInfo(&_Policy.CallOpts, arg0)
}

// PolicyRewardedBlock is a free data retrieval call binding the contract method 0x2fd7c791.
//
// Solidity: function policyRewardedBlock(uint256 ) view returns(uint256)
func (_Policy *PolicyCaller) PolicyRewardedBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "policyRewardedBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolicyRewardedBlock is a free data retrieval call binding the contract method 0x2fd7c791.
//
// Solidity: function policyRewardedBlock(uint256 ) view returns(uint256)
func (_Policy *PolicySession) PolicyRewardedBlock(arg0 *big.Int) (*big.Int, error) {
	return _Policy.Contract.PolicyRewardedBlock(&_Policy.CallOpts, arg0)
}

// PolicyRewardedBlock is a free data retrieval call binding the contract method 0x2fd7c791.
//
// Solidity: function policyRewardedBlock(uint256 ) view returns(uint256)
func (_Policy *PolicyCallerSession) PolicyRewardedBlock(arg0 *big.Int) (*big.Int, error) {
	return _Policy.Contract.PolicyRewardedBlock(&_Policy.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Policy *PolicyCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Policy *PolicySession) RewardToken() (common.Address, error) {
	return _Policy.Contract.RewardToken(&_Policy.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Policy *PolicyCallerSession) RewardToken() (common.Address, error) {
	return _Policy.Contract.RewardToken(&_Policy.CallOpts)
}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Policy *PolicyCaller) StakerReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "stakerReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Policy *PolicySession) StakerReward(arg0 common.Address) (*big.Int, error) {
	return _Policy.Contract.StakerReward(&_Policy.CallOpts, arg0)
}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Policy *PolicyCallerSession) StakerReward(arg0 common.Address) (*big.Int, error) {
	return _Policy.Contract.StakerReward(&_Policy.CallOpts, arg0)
}

// CancelPolicy is a paid mutator transaction binding the contract method 0xb20e8bcc.
//
// Solidity: function cancelPolicy(uint256 _pid) returns()
func (_Policy *PolicyTransactor) CancelPolicy(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "cancelPolicy", _pid)
}

// CancelPolicy is a paid mutator transaction binding the contract method 0xb20e8bcc.
//
// Solidity: function cancelPolicy(uint256 _pid) returns()
func (_Policy *PolicySession) CancelPolicy(_pid *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.CancelPolicy(&_Policy.TransactOpts, _pid)
}

// CancelPolicy is a paid mutator transaction binding the contract method 0xb20e8bcc.
//
// Solidity: function cancelPolicy(uint256 _pid) returns()
func (_Policy *PolicyTransactorSession) CancelPolicy(_pid *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.CancelPolicy(&_Policy.TransactOpts, _pid)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _amount) returns()
func (_Policy *PolicyTransactor) Claim(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "claim", _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _amount) returns()
func (_Policy *PolicySession) Claim(_amount *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.Claim(&_Policy.TransactOpts, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _amount) returns()
func (_Policy *PolicyTransactorSession) Claim(_amount *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.Claim(&_Policy.TransactOpts, _amount)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x1ef34dc9.
//
// Solidity: function createPolicy(address[] _stakerAddresses, uint256 _rewardBalance, uint256 _begin, uint256 _end) returns(uint256)
func (_Policy *PolicyTransactor) CreatePolicy(opts *bind.TransactOpts, _stakerAddresses []common.Address, _rewardBalance *big.Int, _begin *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "createPolicy", _stakerAddresses, _rewardBalance, _begin, _end)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x1ef34dc9.
//
// Solidity: function createPolicy(address[] _stakerAddresses, uint256 _rewardBalance, uint256 _begin, uint256 _end) returns(uint256)
func (_Policy *PolicySession) CreatePolicy(_stakerAddresses []common.Address, _rewardBalance *big.Int, _begin *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _stakerAddresses, _rewardBalance, _begin, _end)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x1ef34dc9.
//
// Solidity: function createPolicy(address[] _stakerAddresses, uint256 _rewardBalance, uint256 _begin, uint256 _end) returns(uint256)
func (_Policy *PolicyTransactorSession) CreatePolicy(_stakerAddresses []common.Address, _rewardBalance *big.Int, _begin *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _stakerAddresses, _rewardBalance, _begin, _end)
}

// GiveOutReward is a paid mutator transaction binding the contract method 0xedc95130.
//
// Solidity: function giveOutReward() returns()
func (_Policy *PolicyTransactor) GiveOutReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "giveOutReward")
}

// GiveOutReward is a paid mutator transaction binding the contract method 0xedc95130.
//
// Solidity: function giveOutReward() returns()
func (_Policy *PolicySession) GiveOutReward() (*types.Transaction, error) {
	return _Policy.Contract.GiveOutReward(&_Policy.TransactOpts)
}

// GiveOutReward is a paid mutator transaction binding the contract method 0xedc95130.
//
// Solidity: function giveOutReward() returns()
func (_Policy *PolicyTransactorSession) GiveOutReward() (*types.Transaction, error) {
	return _Policy.Contract.GiveOutReward(&_Policy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicySession) RenounceOwnership() (*types.Transaction, error) {
	return _Policy.Contract.RenounceOwnership(&_Policy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Policy.Contract.RenounceOwnership(&_Policy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policy.Contract.TransferOwnership(&_Policy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policy.Contract.TransferOwnership(&_Policy.TransactOpts, newOwner)
}

// PolicyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Policy contract.
type PolicyOwnershipTransferredIterator struct {
	Event *PolicyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolicyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyOwnershipTransferred)
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
		it.Event = new(PolicyOwnershipTransferred)
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
func (it *PolicyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyOwnershipTransferred represents a OwnershipTransferred event raised by the Policy contract.
type PolicyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policy *PolicyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolicyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolicyOwnershipTransferredIterator{contract: _Policy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policy *PolicyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolicyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyOwnershipTransferred)
				if err := _Policy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policy *PolicyFilterer) ParseOwnershipTransferred(log types.Log) (*PolicyOwnershipTransferred, error) {
	event := new(PolicyOwnershipTransferred)
	if err := _Policy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
