// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platon

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

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"setTotalReword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakaAmount\",\"type\":\"uint256\"}],\"name\":\"updateStakeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_workCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isWork\",\"type\":\"bool\"}],\"name\":\"updateStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_workCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isWorks\",\"type\":\"bool[]\"}],\"name\":\"UpdateStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workcount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWork\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _sender) view returns(bool)
func (_Staking *StakingCaller) IsManager(opts *bind.CallOpts, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isManager", _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _sender) view returns(bool)
func (_Staking *StakingSession) IsManager(_sender common.Address) (bool, error) {
	return _Staking.Contract.IsManager(&_Staking.CallOpts, _sender)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _sender) view returns(bool)
func (_Staking *StakingCallerSession) IsManager(_sender common.Address) (bool, error) {
	return _Staking.Contract.IsManager(&_Staking.CallOpts, _sender)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address _user) view returns(bool)
func (_Staking *StakingCaller) IsStaker(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isStaker", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address _user) view returns(bool)
func (_Staking *StakingSession) IsStaker(_user common.Address) (bool, error) {
	return _Staking.Contract.IsStaker(&_Staking.CallOpts, _user)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address _user) view returns(bool)
func (_Staking *StakingCallerSession) IsStaker(_user common.Address) (bool, error) {
	return _Staking.Contract.IsStaker(&_Staking.CallOpts, _user)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_Staking *StakingCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "managers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_Staking *StakingSession) Managers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Managers(&_Staking.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_Staking *StakingCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Managers(&_Staking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(address owner, address stakeToken, uint256 stakeBalance, uint256 workcount, bool isWork)
func (_Staking *StakingCaller) StakerInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner        common.Address
	StakeToken   common.Address
	StakeBalance *big.Int
	Workcount    *big.Int
	IsWork       bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakerInfo", arg0)

	outstruct := new(struct {
		Owner        common.Address
		StakeToken   common.Address
		StakeBalance *big.Int
		Workcount    *big.Int
		IsWork       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakeToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Workcount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsWork = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(address owner, address stakeToken, uint256 stakeBalance, uint256 workcount, bool isWork)
func (_Staking *StakingSession) StakerInfo(arg0 common.Address) (struct {
	Owner        common.Address
	StakeToken   common.Address
	StakeBalance *big.Int
	Workcount    *big.Int
	IsWork       bool
}, error) {
	return _Staking.Contract.StakerInfo(&_Staking.CallOpts, arg0)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(address owner, address stakeToken, uint256 stakeBalance, uint256 workcount, bool isWork)
func (_Staking *StakingCallerSession) StakerInfo(arg0 common.Address) (struct {
	Owner        common.Address
	StakeToken   common.Address
	StakeBalance *big.Int
	Workcount    *big.Int
	IsWork       bool
}, error) {
	return _Staking.Contract.StakerInfo(&_Staking.CallOpts, arg0)
}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Staking *StakingCaller) StakerReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakerReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Staking *StakingSession) StakerReward(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakerReward(&_Staking.CallOpts, arg0)
}

// StakerReward is a free data retrieval call binding the contract method 0x0ea474c7.
//
// Solidity: function stakerReward(address ) view returns(uint256)
func (_Staking *StakingCallerSession) StakerReward(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakerReward(&_Staking.CallOpts, arg0)
}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_Staking *StakingCaller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_Staking *StakingSession) TotalReward() (*big.Int, error) {
	return _Staking.Contract.TotalReward(&_Staking.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_Staking *StakingCallerSession) TotalReward() (*big.Int, error) {
	return _Staking.Contract.TotalReward(&_Staking.CallOpts)
}

// TotalStakeBalance is a free data retrieval call binding the contract method 0x34949cc1.
//
// Solidity: function totalStakeBalance() view returns(uint256)
func (_Staking *StakingCaller) TotalStakeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStakeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakeBalance is a free data retrieval call binding the contract method 0x34949cc1.
//
// Solidity: function totalStakeBalance() view returns(uint256)
func (_Staking *StakingSession) TotalStakeBalance() (*big.Int, error) {
	return _Staking.Contract.TotalStakeBalance(&_Staking.CallOpts)
}

// TotalStakeBalance is a free data retrieval call binding the contract method 0x34949cc1.
//
// Solidity: function totalStakeBalance() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStakeBalance() (*big.Int, error) {
	return _Staking.Contract.TotalStakeBalance(&_Staking.CallOpts)
}

// UpdateStakers is a paid mutator transaction binding the contract method 0x57df2864.
//
// Solidity: function UpdateStakers(address[] _owners, uint256[] _balances, uint256[] _workCounts, bool[] _isWorks) returns()
func (_Staking *StakingTransactor) UpdateStakers(opts *bind.TransactOpts, _owners []common.Address, _balances []*big.Int, _workCounts []*big.Int, _isWorks []bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "UpdateStakers", _owners, _balances, _workCounts, _isWorks)
}

// UpdateStakers is a paid mutator transaction binding the contract method 0x57df2864.
//
// Solidity: function UpdateStakers(address[] _owners, uint256[] _balances, uint256[] _workCounts, bool[] _isWorks) returns()
func (_Staking *StakingSession) UpdateStakers(_owners []common.Address, _balances []*big.Int, _workCounts []*big.Int, _isWorks []bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStakers(&_Staking.TransactOpts, _owners, _balances, _workCounts, _isWorks)
}

// UpdateStakers is a paid mutator transaction binding the contract method 0x57df2864.
//
// Solidity: function UpdateStakers(address[] _owners, uint256[] _balances, uint256[] _workCounts, bool[] _isWorks) returns()
func (_Staking *StakingTransactorSession) UpdateStakers(_owners []common.Address, _balances []*big.Int, _workCounts []*big.Int, _isWorks []bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStakers(&_Staking.TransactOpts, _owners, _balances, _workCounts, _isWorks)
}

// Claim is a paid mutator transaction binding the contract method 0xaab8ab0c.
//
// Solidity: function claim(uint64 _amount) returns()
func (_Staking *StakingTransactor) Claim(opts *bind.TransactOpts, _amount uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim", _amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaab8ab0c.
//
// Solidity: function claim(uint64 _amount) returns()
func (_Staking *StakingSession) Claim(_amount uint64) (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaab8ab0c.
//
// Solidity: function claim(uint64 _amount) returns()
func (_Staking *StakingTransactorSession) Claim(_amount uint64) (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts, _amount)
}

// LeaveStaker is a paid mutator transaction binding the contract method 0xd89085de.
//
// Solidity: function leaveStaker() returns()
func (_Staking *StakingTransactor) LeaveStaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "leaveStaker")
}

// LeaveStaker is a paid mutator transaction binding the contract method 0xd89085de.
//
// Solidity: function leaveStaker() returns()
func (_Staking *StakingSession) LeaveStaker() (*types.Transaction, error) {
	return _Staking.Contract.LeaveStaker(&_Staking.TransactOpts)
}

// LeaveStaker is a paid mutator transaction binding the contract method 0xd89085de.
//
// Solidity: function leaveStaker() returns()
func (_Staking *StakingTransactorSession) LeaveStaker() (*types.Transaction, error) {
	return _Staking.Contract.LeaveStaker(&_Staking.TransactOpts)
}

// NewStaker is a paid mutator transaction binding the contract method 0x7c59a910.
//
// Solidity: function newStaker() returns()
func (_Staking *StakingTransactor) NewStaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "newStaker")
}

// NewStaker is a paid mutator transaction binding the contract method 0x7c59a910.
//
// Solidity: function newStaker() returns()
func (_Staking *StakingSession) NewStaker() (*types.Transaction, error) {
	return _Staking.Contract.NewStaker(&_Staking.TransactOpts)
}

// NewStaker is a paid mutator transaction binding the contract method 0x7c59a910.
//
// Solidity: function newStaker() returns()
func (_Staking *StakingTransactorSession) NewStaker() (*types.Transaction, error) {
	return _Staking.Contract.NewStaker(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetTotalReword is a paid mutator transaction binding the contract method 0xc7162ea6.
//
// Solidity: function setTotalReword(uint256 _totalReward) returns()
func (_Staking *StakingTransactor) SetTotalReword(opts *bind.TransactOpts, _totalReward *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setTotalReword", _totalReward)
}

// SetTotalReword is a paid mutator transaction binding the contract method 0xc7162ea6.
//
// Solidity: function setTotalReword(uint256 _totalReward) returns()
func (_Staking *StakingSession) SetTotalReword(_totalReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetTotalReword(&_Staking.TransactOpts, _totalReward)
}

// SetTotalReword is a paid mutator transaction binding the contract method 0xc7162ea6.
//
// Solidity: function setTotalReword(uint256 _totalReward) returns()
func (_Staking *StakingTransactorSession) SetTotalReword(_totalReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetTotalReword(&_Staking.TransactOpts, _totalReward)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// UpdateStakeInfo is a paid mutator transaction binding the contract method 0x30527b66.
//
// Solidity: function updateStakeInfo(address _stakeToken, uint256 _stakaAmount) returns()
func (_Staking *StakingTransactor) UpdateStakeInfo(opts *bind.TransactOpts, _stakeToken common.Address, _stakaAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateStakeInfo", _stakeToken, _stakaAmount)
}

// UpdateStakeInfo is a paid mutator transaction binding the contract method 0x30527b66.
//
// Solidity: function updateStakeInfo(address _stakeToken, uint256 _stakaAmount) returns()
func (_Staking *StakingSession) UpdateStakeInfo(_stakeToken common.Address, _stakaAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStakeInfo(&_Staking.TransactOpts, _stakeToken, _stakaAmount)
}

// UpdateStakeInfo is a paid mutator transaction binding the contract method 0x30527b66.
//
// Solidity: function updateStakeInfo(address _stakeToken, uint256 _stakaAmount) returns()
func (_Staking *StakingTransactorSession) UpdateStakeInfo(_stakeToken common.Address, _stakaAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStakeInfo(&_Staking.TransactOpts, _stakeToken, _stakaAmount)
}

// UpdateStaker is a paid mutator transaction binding the contract method 0x3b5036d7.
//
// Solidity: function updateStaker(address _owner, uint256 _balance, uint256 _workCount, bool _isWork) returns()
func (_Staking *StakingTransactor) UpdateStaker(opts *bind.TransactOpts, _owner common.Address, _balance *big.Int, _workCount *big.Int, _isWork bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateStaker", _owner, _balance, _workCount, _isWork)
}

// UpdateStaker is a paid mutator transaction binding the contract method 0x3b5036d7.
//
// Solidity: function updateStaker(address _owner, uint256 _balance, uint256 _workCount, bool _isWork) returns()
func (_Staking *StakingSession) UpdateStaker(_owner common.Address, _balance *big.Int, _workCount *big.Int, _isWork bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStaker(&_Staking.TransactOpts, _owner, _balance, _workCount, _isWork)
}

// UpdateStaker is a paid mutator transaction binding the contract method 0x3b5036d7.
//
// Solidity: function updateStaker(address _owner, uint256 _balance, uint256 _workCount, bool _isWork) returns()
func (_Staking *StakingTransactorSession) UpdateStaker(_owner common.Address, _balance *big.Int, _workCount *big.Int, _isWork bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateStaker(&_Staking.TransactOpts, _owner, _balance, _workCount, _isWork)
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
