// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nucypher

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

// NucypherABI is the input ABI used to generate the binding from.
const NucypherABI = "[{\"inputs\":[{\"internalType\":\"contractNuCypherToken\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractWorkLockInterface\",\"name\":\"_workLock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investigator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"testTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"StateVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UpgradeFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceHistory\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMintingPeriod\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodSupply\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_unlockingDuration\",\"type\":\"uint16\"}],\"name\":\"depositFromWorkLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"finishUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getAllTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getCompletedWork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPastDowntime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"startPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endPeriod\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getPastDowntimeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSubStakeInfo\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"firstPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lastPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"unlockingDuration\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"lockedValue\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getSubStakesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgrade\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"lockedPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousPeriodSupply\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_measureWork\",\"type\":\"bool\"}],\"name\":\"setWorkMeasurement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerFromWorker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"currentCommittedPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"nextCommittedPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lastCommittedPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stub1\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"completedWork\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"workerStartPeriod\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedSlot1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedSlot2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedSlot3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedSlot4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedSlot5\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportsHistory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractNuCypherToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalStakedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalStakedForAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_testTarget\",\"type\":\"address\"}],\"name\":\"verifyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workLock\",\"outputs\":[{\"internalType\":\"contractWorkLockInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Nucypher is an auto generated Go binding around an Ethereum contract.
type Nucypher struct {
	NucypherCaller     // Read-only binding to the contract
	NucypherTransactor // Write-only binding to the contract
	NucypherFilterer   // Log filterer for contract events
}

// NucypherCaller is an auto generated read-only Go binding around an Ethereum contract.
type NucypherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NucypherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NucypherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NucypherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NucypherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NucypherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NucypherSession struct {
	Contract     *Nucypher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NucypherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NucypherCallerSession struct {
	Contract *NucypherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NucypherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NucypherTransactorSession struct {
	Contract     *NucypherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NucypherRaw is an auto generated low-level Go binding around an Ethereum contract.
type NucypherRaw struct {
	Contract *Nucypher // Generic contract binding to access the raw methods on
}

// NucypherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NucypherCallerRaw struct {
	Contract *NucypherCaller // Generic read-only contract binding to access the raw methods on
}

// NucypherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NucypherTransactorRaw struct {
	Contract *NucypherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNucypher creates a new instance of Nucypher, bound to a specific deployed contract.
func NewNucypher(address common.Address, backend bind.ContractBackend) (*Nucypher, error) {
	contract, err := bindNucypher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nucypher{NucypherCaller: NucypherCaller{contract: contract}, NucypherTransactor: NucypherTransactor{contract: contract}, NucypherFilterer: NucypherFilterer{contract: contract}}, nil
}

// NewNucypherCaller creates a new read-only instance of Nucypher, bound to a specific deployed contract.
func NewNucypherCaller(address common.Address, caller bind.ContractCaller) (*NucypherCaller, error) {
	contract, err := bindNucypher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NucypherCaller{contract: contract}, nil
}

// NewNucypherTransactor creates a new write-only instance of Nucypher, bound to a specific deployed contract.
func NewNucypherTransactor(address common.Address, transactor bind.ContractTransactor) (*NucypherTransactor, error) {
	contract, err := bindNucypher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NucypherTransactor{contract: contract}, nil
}

// NewNucypherFilterer creates a new log filterer instance of Nucypher, bound to a specific deployed contract.
func NewNucypherFilterer(address common.Address, filterer bind.ContractFilterer) (*NucypherFilterer, error) {
	contract, err := bindNucypher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NucypherFilterer{contract: contract}, nil
}

// bindNucypher binds a generic wrapper to an already deployed contract.
func bindNucypher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NucypherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nucypher *NucypherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nucypher.Contract.NucypherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nucypher *NucypherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nucypher.Contract.NucypherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nucypher *NucypherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nucypher.Contract.NucypherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nucypher *NucypherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nucypher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nucypher *NucypherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nucypher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nucypher *NucypherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nucypher.Contract.contract.Transact(opts, method, params...)
}

// BalanceHistory is a free data retrieval call binding the contract method 0xdfee9e0d.
//
// Solidity: function balanceHistory(uint256 ) view returns(uint128)
func (_Nucypher *NucypherCaller) BalanceHistory(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "balanceHistory", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceHistory is a free data retrieval call binding the contract method 0xdfee9e0d.
//
// Solidity: function balanceHistory(uint256 ) view returns(uint128)
func (_Nucypher *NucypherSession) BalanceHistory(arg0 *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.BalanceHistory(&_Nucypher.CallOpts, arg0)
}

// BalanceHistory is a free data retrieval call binding the contract method 0xdfee9e0d.
//
// Solidity: function balanceHistory(uint256 ) view returns(uint128)
func (_Nucypher *NucypherCallerSession) BalanceHistory(arg0 *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.BalanceHistory(&_Nucypher.CallOpts, arg0)
}

// CurrentMintingPeriod is a free data retrieval call binding the contract method 0x0f64ab14.
//
// Solidity: function currentMintingPeriod() view returns(uint16)
func (_Nucypher *NucypherCaller) CurrentMintingPeriod(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "currentMintingPeriod")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CurrentMintingPeriod is a free data retrieval call binding the contract method 0x0f64ab14.
//
// Solidity: function currentMintingPeriod() view returns(uint16)
func (_Nucypher *NucypherSession) CurrentMintingPeriod() (uint16, error) {
	return _Nucypher.Contract.CurrentMintingPeriod(&_Nucypher.CallOpts)
}

// CurrentMintingPeriod is a free data retrieval call binding the contract method 0x0f64ab14.
//
// Solidity: function currentMintingPeriod() view returns(uint16)
func (_Nucypher *NucypherCallerSession) CurrentMintingPeriod() (uint16, error) {
	return _Nucypher.Contract.CurrentMintingPeriod(&_Nucypher.CallOpts)
}

// CurrentPeriodSupply is a free data retrieval call binding the contract method 0x204612a8.
//
// Solidity: function currentPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherCaller) CurrentPeriodSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "currentPeriodSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodSupply is a free data retrieval call binding the contract method 0x204612a8.
//
// Solidity: function currentPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherSession) CurrentPeriodSupply() (*big.Int, error) {
	return _Nucypher.Contract.CurrentPeriodSupply(&_Nucypher.CallOpts)
}

// CurrentPeriodSupply is a free data retrieval call binding the contract method 0x204612a8.
//
// Solidity: function currentPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherCallerSession) CurrentPeriodSupply() (*big.Int, error) {
	return _Nucypher.Contract.CurrentPeriodSupply(&_Nucypher.CallOpts)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x178b6de6.
//
// Solidity: function getAllTokens(address _staker) view returns(uint256)
func (_Nucypher *NucypherCaller) GetAllTokens(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getAllTokens", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllTokens is a free data retrieval call binding the contract method 0x178b6de6.
//
// Solidity: function getAllTokens(address _staker) view returns(uint256)
func (_Nucypher *NucypherSession) GetAllTokens(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetAllTokens(&_Nucypher.CallOpts, _staker)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x178b6de6.
//
// Solidity: function getAllTokens(address _staker) view returns(uint256)
func (_Nucypher *NucypherCallerSession) GetAllTokens(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetAllTokens(&_Nucypher.CallOpts, _staker)
}

// GetCompletedWork is a free data retrieval call binding the contract method 0xd094adbf.
//
// Solidity: function getCompletedWork(address _staker) view returns(uint256)
func (_Nucypher *NucypherCaller) GetCompletedWork(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getCompletedWork", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompletedWork is a free data retrieval call binding the contract method 0xd094adbf.
//
// Solidity: function getCompletedWork(address _staker) view returns(uint256)
func (_Nucypher *NucypherSession) GetCompletedWork(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetCompletedWork(&_Nucypher.CallOpts, _staker)
}

// GetCompletedWork is a free data retrieval call binding the contract method 0xd094adbf.
//
// Solidity: function getCompletedWork(address _staker) view returns(uint256)
func (_Nucypher *NucypherCallerSession) GetCompletedWork(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetCompletedWork(&_Nucypher.CallOpts, _staker)
}

// GetPastDowntime is a free data retrieval call binding the contract method 0x8fa95a15.
//
// Solidity: function getPastDowntime(address _staker, uint256 _index) view returns(uint16 startPeriod, uint16 endPeriod)
func (_Nucypher *NucypherCaller) GetPastDowntime(opts *bind.CallOpts, _staker common.Address, _index *big.Int) (struct {
	StartPeriod uint16
	EndPeriod   uint16
}, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getPastDowntime", _staker, _index)

	outstruct := new(struct {
		StartPeriod uint16
		EndPeriod   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartPeriod = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.EndPeriod = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetPastDowntime is a free data retrieval call binding the contract method 0x8fa95a15.
//
// Solidity: function getPastDowntime(address _staker, uint256 _index) view returns(uint16 startPeriod, uint16 endPeriod)
func (_Nucypher *NucypherSession) GetPastDowntime(_staker common.Address, _index *big.Int) (struct {
	StartPeriod uint16
	EndPeriod   uint16
}, error) {
	return _Nucypher.Contract.GetPastDowntime(&_Nucypher.CallOpts, _staker, _index)
}

// GetPastDowntime is a free data retrieval call binding the contract method 0x8fa95a15.
//
// Solidity: function getPastDowntime(address _staker, uint256 _index) view returns(uint16 startPeriod, uint16 endPeriod)
func (_Nucypher *NucypherCallerSession) GetPastDowntime(_staker common.Address, _index *big.Int) (struct {
	StartPeriod uint16
	EndPeriod   uint16
}, error) {
	return _Nucypher.Contract.GetPastDowntime(&_Nucypher.CallOpts, _staker, _index)
}

// GetPastDowntimeLength is a free data retrieval call binding the contract method 0x90dcb51f.
//
// Solidity: function getPastDowntimeLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherCaller) GetPastDowntimeLength(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getPastDowntimeLength", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastDowntimeLength is a free data retrieval call binding the contract method 0x90dcb51f.
//
// Solidity: function getPastDowntimeLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherSession) GetPastDowntimeLength(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetPastDowntimeLength(&_Nucypher.CallOpts, _staker)
}

// GetPastDowntimeLength is a free data retrieval call binding the contract method 0x90dcb51f.
//
// Solidity: function getPastDowntimeLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherCallerSession) GetPastDowntimeLength(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetPastDowntimeLength(&_Nucypher.CallOpts, _staker)
}

// GetStakersLength is a free data retrieval call binding the contract method 0x9fabeb00.
//
// Solidity: function getStakersLength() view returns(uint256)
func (_Nucypher *NucypherCaller) GetStakersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getStakersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakersLength is a free data retrieval call binding the contract method 0x9fabeb00.
//
// Solidity: function getStakersLength() view returns(uint256)
func (_Nucypher *NucypherSession) GetStakersLength() (*big.Int, error) {
	return _Nucypher.Contract.GetStakersLength(&_Nucypher.CallOpts)
}

// GetStakersLength is a free data retrieval call binding the contract method 0x9fabeb00.
//
// Solidity: function getStakersLength() view returns(uint256)
func (_Nucypher *NucypherCallerSession) GetStakersLength() (*big.Int, error) {
	return _Nucypher.Contract.GetStakersLength(&_Nucypher.CallOpts)
}

// GetSubStakeInfo is a free data retrieval call binding the contract method 0x521ca4da.
//
// Solidity: function getSubStakeInfo(address _staker, uint256 _index) view returns(uint16 firstPeriod, uint16 lastPeriod, uint16 unlockingDuration, uint128 lockedValue)
func (_Nucypher *NucypherCaller) GetSubStakeInfo(opts *bind.CallOpts, _staker common.Address, _index *big.Int) (struct {
	FirstPeriod       uint16
	LastPeriod        uint16
	UnlockingDuration uint16
	LockedValue       *big.Int
}, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getSubStakeInfo", _staker, _index)

	outstruct := new(struct {
		FirstPeriod       uint16
		LastPeriod        uint16
		UnlockingDuration uint16
		LockedValue       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstPeriod = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.LastPeriod = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.UnlockingDuration = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.LockedValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSubStakeInfo is a free data retrieval call binding the contract method 0x521ca4da.
//
// Solidity: function getSubStakeInfo(address _staker, uint256 _index) view returns(uint16 firstPeriod, uint16 lastPeriod, uint16 unlockingDuration, uint128 lockedValue)
func (_Nucypher *NucypherSession) GetSubStakeInfo(_staker common.Address, _index *big.Int) (struct {
	FirstPeriod       uint16
	LastPeriod        uint16
	UnlockingDuration uint16
	LockedValue       *big.Int
}, error) {
	return _Nucypher.Contract.GetSubStakeInfo(&_Nucypher.CallOpts, _staker, _index)
}

// GetSubStakeInfo is a free data retrieval call binding the contract method 0x521ca4da.
//
// Solidity: function getSubStakeInfo(address _staker, uint256 _index) view returns(uint16 firstPeriod, uint16 lastPeriod, uint16 unlockingDuration, uint128 lockedValue)
func (_Nucypher *NucypherCallerSession) GetSubStakeInfo(_staker common.Address, _index *big.Int) (struct {
	FirstPeriod       uint16
	LastPeriod        uint16
	UnlockingDuration uint16
	LockedValue       *big.Int
}, error) {
	return _Nucypher.Contract.GetSubStakeInfo(&_Nucypher.CallOpts, _staker, _index)
}

// GetSubStakesLength is a free data retrieval call binding the contract method 0x04bcafe5.
//
// Solidity: function getSubStakesLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherCaller) GetSubStakesLength(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "getSubStakesLength", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubStakesLength is a free data retrieval call binding the contract method 0x04bcafe5.
//
// Solidity: function getSubStakesLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherSession) GetSubStakesLength(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetSubStakesLength(&_Nucypher.CallOpts, _staker)
}

// GetSubStakesLength is a free data retrieval call binding the contract method 0x04bcafe5.
//
// Solidity: function getSubStakesLength(address _staker) view returns(uint256)
func (_Nucypher *NucypherCallerSession) GetSubStakesLength(_staker common.Address) (*big.Int, error) {
	return _Nucypher.Contract.GetSubStakesLength(&_Nucypher.CallOpts, _staker)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Nucypher *NucypherCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Nucypher *NucypherSession) IsOwner() (bool, error) {
	return _Nucypher.Contract.IsOwner(&_Nucypher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Nucypher *NucypherCallerSession) IsOwner() (bool, error) {
	return _Nucypher.Contract.IsOwner(&_Nucypher.CallOpts)
}

// IsUpgrade is a free data retrieval call binding the contract method 0xe38a303b.
//
// Solidity: function isUpgrade() view returns(uint8)
func (_Nucypher *NucypherCaller) IsUpgrade(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "isUpgrade")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// IsUpgrade is a free data retrieval call binding the contract method 0xe38a303b.
//
// Solidity: function isUpgrade() view returns(uint8)
func (_Nucypher *NucypherSession) IsUpgrade() (uint8, error) {
	return _Nucypher.Contract.IsUpgrade(&_Nucypher.CallOpts)
}

// IsUpgrade is a free data retrieval call binding the contract method 0xe38a303b.
//
// Solidity: function isUpgrade() view returns(uint8)
func (_Nucypher *NucypherCallerSession) IsUpgrade() (uint8, error) {
	return _Nucypher.Contract.IsUpgrade(&_Nucypher.CallOpts)
}

// LockedPerPeriod is a free data retrieval call binding the contract method 0xef95e0ba.
//
// Solidity: function lockedPerPeriod(uint16 ) view returns(uint256)
func (_Nucypher *NucypherCaller) LockedPerPeriod(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "lockedPerPeriod", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedPerPeriod is a free data retrieval call binding the contract method 0xef95e0ba.
//
// Solidity: function lockedPerPeriod(uint16 ) view returns(uint256)
func (_Nucypher *NucypherSession) LockedPerPeriod(arg0 uint16) (*big.Int, error) {
	return _Nucypher.Contract.LockedPerPeriod(&_Nucypher.CallOpts, arg0)
}

// LockedPerPeriod is a free data retrieval call binding the contract method 0xef95e0ba.
//
// Solidity: function lockedPerPeriod(uint16 ) view returns(uint256)
func (_Nucypher *NucypherCallerSession) LockedPerPeriod(arg0 uint16) (*big.Int, error) {
	return _Nucypher.Contract.LockedPerPeriod(&_Nucypher.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nucypher *NucypherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nucypher *NucypherSession) Owner() (common.Address, error) {
	return _Nucypher.Contract.Owner(&_Nucypher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nucypher *NucypherCallerSession) Owner() (common.Address, error) {
	return _Nucypher.Contract.Owner(&_Nucypher.CallOpts)
}

// PreviousPeriodSupply is a free data retrieval call binding the contract method 0x3250843d.
//
// Solidity: function previousPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherCaller) PreviousPeriodSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "previousPeriodSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousPeriodSupply is a free data retrieval call binding the contract method 0x3250843d.
//
// Solidity: function previousPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherSession) PreviousPeriodSupply() (*big.Int, error) {
	return _Nucypher.Contract.PreviousPeriodSupply(&_Nucypher.CallOpts)
}

// PreviousPeriodSupply is a free data retrieval call binding the contract method 0x3250843d.
//
// Solidity: function previousPeriodSupply() view returns(uint128)
func (_Nucypher *NucypherCallerSession) PreviousPeriodSupply() (*big.Int, error) {
	return _Nucypher.Contract.PreviousPeriodSupply(&_Nucypher.CallOpts)
}

// PreviousTarget is a free data retrieval call binding the contract method 0x4b2cd118.
//
// Solidity: function previousTarget() view returns(address)
func (_Nucypher *NucypherCaller) PreviousTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "previousTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousTarget is a free data retrieval call binding the contract method 0x4b2cd118.
//
// Solidity: function previousTarget() view returns(address)
func (_Nucypher *NucypherSession) PreviousTarget() (common.Address, error) {
	return _Nucypher.Contract.PreviousTarget(&_Nucypher.CallOpts)
}

// PreviousTarget is a free data retrieval call binding the contract method 0x4b2cd118.
//
// Solidity: function previousTarget() view returns(address)
func (_Nucypher *NucypherCallerSession) PreviousTarget() (common.Address, error) {
	return _Nucypher.Contract.PreviousTarget(&_Nucypher.CallOpts)
}

// StakerFromWorker is a free data retrieval call binding the contract method 0xb2eeb36e.
//
// Solidity: function stakerFromWorker(address ) view returns(address)
func (_Nucypher *NucypherCaller) StakerFromWorker(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "stakerFromWorker", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerFromWorker is a free data retrieval call binding the contract method 0xb2eeb36e.
//
// Solidity: function stakerFromWorker(address ) view returns(address)
func (_Nucypher *NucypherSession) StakerFromWorker(arg0 common.Address) (common.Address, error) {
	return _Nucypher.Contract.StakerFromWorker(&_Nucypher.CallOpts, arg0)
}

// StakerFromWorker is a free data retrieval call binding the contract method 0xb2eeb36e.
//
// Solidity: function stakerFromWorker(address ) view returns(address)
func (_Nucypher *NucypherCallerSession) StakerFromWorker(arg0 common.Address) (common.Address, error) {
	return _Nucypher.Contract.StakerFromWorker(&_Nucypher.CallOpts, arg0)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(uint256 value, uint16 currentCommittedPeriod, uint16 nextCommittedPeriod, uint16 lastCommittedPeriod, uint16 stub1, uint256 completedWork, uint16 workerStartPeriod, address worker, uint256 flags, uint256 reservedSlot1, uint256 reservedSlot2, uint256 reservedSlot3, uint256 reservedSlot4, uint256 reservedSlot5)
func (_Nucypher *NucypherCaller) StakerInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Value                  *big.Int
	CurrentCommittedPeriod uint16
	NextCommittedPeriod    uint16
	LastCommittedPeriod    uint16
	Stub1                  uint16
	CompletedWork          *big.Int
	WorkerStartPeriod      uint16
	Worker                 common.Address
	Flags                  *big.Int
	ReservedSlot1          *big.Int
	ReservedSlot2          *big.Int
	ReservedSlot3          *big.Int
	ReservedSlot4          *big.Int
	ReservedSlot5          *big.Int
}, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "stakerInfo", arg0)

	outstruct := new(struct {
		Value                  *big.Int
		CurrentCommittedPeriod uint16
		NextCommittedPeriod    uint16
		LastCommittedPeriod    uint16
		Stub1                  uint16
		CompletedWork          *big.Int
		WorkerStartPeriod      uint16
		Worker                 common.Address
		Flags                  *big.Int
		ReservedSlot1          *big.Int
		ReservedSlot2          *big.Int
		ReservedSlot3          *big.Int
		ReservedSlot4          *big.Int
		ReservedSlot5          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentCommittedPeriod = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.NextCommittedPeriod = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.LastCommittedPeriod = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Stub1 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CompletedWork = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.WorkerStartPeriod = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.Worker = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Flags = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ReservedSlot1 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.ReservedSlot2 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.ReservedSlot3 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.ReservedSlot4 = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.ReservedSlot5 = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(uint256 value, uint16 currentCommittedPeriod, uint16 nextCommittedPeriod, uint16 lastCommittedPeriod, uint16 stub1, uint256 completedWork, uint16 workerStartPeriod, address worker, uint256 flags, uint256 reservedSlot1, uint256 reservedSlot2, uint256 reservedSlot3, uint256 reservedSlot4, uint256 reservedSlot5)
func (_Nucypher *NucypherSession) StakerInfo(arg0 common.Address) (struct {
	Value                  *big.Int
	CurrentCommittedPeriod uint16
	NextCommittedPeriod    uint16
	LastCommittedPeriod    uint16
	Stub1                  uint16
	CompletedWork          *big.Int
	WorkerStartPeriod      uint16
	Worker                 common.Address
	Flags                  *big.Int
	ReservedSlot1          *big.Int
	ReservedSlot2          *big.Int
	ReservedSlot3          *big.Int
	ReservedSlot4          *big.Int
	ReservedSlot5          *big.Int
}, error) {
	return _Nucypher.Contract.StakerInfo(&_Nucypher.CallOpts, arg0)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address ) view returns(uint256 value, uint16 currentCommittedPeriod, uint16 nextCommittedPeriod, uint16 lastCommittedPeriod, uint16 stub1, uint256 completedWork, uint16 workerStartPeriod, address worker, uint256 flags, uint256 reservedSlot1, uint256 reservedSlot2, uint256 reservedSlot3, uint256 reservedSlot4, uint256 reservedSlot5)
func (_Nucypher *NucypherCallerSession) StakerInfo(arg0 common.Address) (struct {
	Value                  *big.Int
	CurrentCommittedPeriod uint16
	NextCommittedPeriod    uint16
	LastCommittedPeriod    uint16
	Stub1                  uint16
	CompletedWork          *big.Int
	WorkerStartPeriod      uint16
	Worker                 common.Address
	Flags                  *big.Int
	ReservedSlot1          *big.Int
	ReservedSlot2          *big.Int
	ReservedSlot3          *big.Int
	ReservedSlot4          *big.Int
	ReservedSlot5          *big.Int
}, error) {
	return _Nucypher.Contract.StakerInfo(&_Nucypher.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Nucypher *NucypherCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "stakers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Nucypher *NucypherSession) Stakers(arg0 *big.Int) (common.Address, error) {
	return _Nucypher.Contract.Stakers(&_Nucypher.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Nucypher *NucypherCallerSession) Stakers(arg0 *big.Int) (common.Address, error) {
	return _Nucypher.Contract.Stakers(&_Nucypher.CallOpts, arg0)
}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() pure returns(bool)
func (_Nucypher *NucypherCaller) SupportsHistory(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "supportsHistory")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() pure returns(bool)
func (_Nucypher *NucypherSession) SupportsHistory() (bool, error) {
	return _Nucypher.Contract.SupportsHistory(&_Nucypher.CallOpts)
}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() pure returns(bool)
func (_Nucypher *NucypherCallerSession) SupportsHistory() (bool, error) {
	return _Nucypher.Contract.SupportsHistory(&_Nucypher.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Nucypher *NucypherCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Nucypher *NucypherSession) Target() (common.Address, error) {
	return _Nucypher.Contract.Target(&_Nucypher.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Nucypher *NucypherCallerSession) Target() (common.Address, error) {
	return _Nucypher.Contract.Target(&_Nucypher.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nucypher *NucypherCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nucypher *NucypherSession) Token() (common.Address, error) {
	return _Nucypher.Contract.Token(&_Nucypher.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nucypher *NucypherCallerSession) Token() (common.Address, error) {
	return _Nucypher.Contract.Token(&_Nucypher.CallOpts)
}

// TotalStakedAt is a free data retrieval call binding the contract method 0xc9c53232.
//
// Solidity: function totalStakedAt(uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherCaller) TotalStakedAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "totalStakedAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedAt is a free data retrieval call binding the contract method 0xc9c53232.
//
// Solidity: function totalStakedAt(uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherSession) TotalStakedAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.TotalStakedAt(&_Nucypher.CallOpts, _blockNumber)
}

// TotalStakedAt is a free data retrieval call binding the contract method 0xc9c53232.
//
// Solidity: function totalStakedAt(uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherCallerSession) TotalStakedAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.TotalStakedAt(&_Nucypher.CallOpts, _blockNumber)
}

// TotalStakedForAt is a free data retrieval call binding the contract method 0xede38421.
//
// Solidity: function totalStakedForAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherCaller) TotalStakedForAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "totalStakedForAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedForAt is a free data retrieval call binding the contract method 0xede38421.
//
// Solidity: function totalStakedForAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherSession) TotalStakedForAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.TotalStakedForAt(&_Nucypher.CallOpts, _owner, _blockNumber)
}

// TotalStakedForAt is a free data retrieval call binding the contract method 0xede38421.
//
// Solidity: function totalStakedForAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Nucypher *NucypherCallerSession) TotalStakedForAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Nucypher.Contract.TotalStakedForAt(&_Nucypher.CallOpts, _owner, _blockNumber)
}

// WorkLock is a free data retrieval call binding the contract method 0x1c7fd532.
//
// Solidity: function workLock() view returns(address)
func (_Nucypher *NucypherCaller) WorkLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nucypher.contract.Call(opts, &out, "workLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkLock is a free data retrieval call binding the contract method 0x1c7fd532.
//
// Solidity: function workLock() view returns(address)
func (_Nucypher *NucypherSession) WorkLock() (common.Address, error) {
	return _Nucypher.Contract.WorkLock(&_Nucypher.CallOpts)
}

// WorkLock is a free data retrieval call binding the contract method 0x1c7fd532.
//
// Solidity: function workLock() view returns(address)
func (_Nucypher *NucypherCallerSession) WorkLock() (common.Address, error) {
	return _Nucypher.Contract.WorkLock(&_Nucypher.CallOpts)
}

// DepositFromWorkLock is a paid mutator transaction binding the contract method 0x171f03c2.
//
// Solidity: function depositFromWorkLock(address _staker, uint256 _value, uint16 _unlockingDuration) returns()
func (_Nucypher *NucypherTransactor) DepositFromWorkLock(opts *bind.TransactOpts, _staker common.Address, _value *big.Int, _unlockingDuration uint16) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "depositFromWorkLock", _staker, _value, _unlockingDuration)
}

// DepositFromWorkLock is a paid mutator transaction binding the contract method 0x171f03c2.
//
// Solidity: function depositFromWorkLock(address _staker, uint256 _value, uint16 _unlockingDuration) returns()
func (_Nucypher *NucypherSession) DepositFromWorkLock(_staker common.Address, _value *big.Int, _unlockingDuration uint16) (*types.Transaction, error) {
	return _Nucypher.Contract.DepositFromWorkLock(&_Nucypher.TransactOpts, _staker, _value, _unlockingDuration)
}

// DepositFromWorkLock is a paid mutator transaction binding the contract method 0x171f03c2.
//
// Solidity: function depositFromWorkLock(address _staker, uint256 _value, uint16 _unlockingDuration) returns()
func (_Nucypher *NucypherTransactorSession) DepositFromWorkLock(_staker common.Address, _value *big.Int, _unlockingDuration uint16) (*types.Transaction, error) {
	return _Nucypher.Contract.DepositFromWorkLock(&_Nucypher.TransactOpts, _staker, _value, _unlockingDuration)
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0xe8dccd06.
//
// Solidity: function finishUpgrade(address _target) returns()
func (_Nucypher *NucypherTransactor) FinishUpgrade(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "finishUpgrade", _target)
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0xe8dccd06.
//
// Solidity: function finishUpgrade(address _target) returns()
func (_Nucypher *NucypherSession) FinishUpgrade(_target common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.FinishUpgrade(&_Nucypher.TransactOpts, _target)
}

// FinishUpgrade is a paid mutator transaction binding the contract method 0xe8dccd06.
//
// Solidity: function finishUpgrade(address _target) returns()
func (_Nucypher *NucypherTransactorSession) FinishUpgrade(_target common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.FinishUpgrade(&_Nucypher.TransactOpts, _target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nucypher *NucypherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nucypher *NucypherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nucypher.Contract.RenounceOwnership(&_Nucypher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nucypher *NucypherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nucypher.Contract.RenounceOwnership(&_Nucypher.TransactOpts)
}

// SetWorkMeasurement is a paid mutator transaction binding the contract method 0x49e5add4.
//
// Solidity: function setWorkMeasurement(address _staker, bool _measureWork) returns(uint256)
func (_Nucypher *NucypherTransactor) SetWorkMeasurement(opts *bind.TransactOpts, _staker common.Address, _measureWork bool) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "setWorkMeasurement", _staker, _measureWork)
}

// SetWorkMeasurement is a paid mutator transaction binding the contract method 0x49e5add4.
//
// Solidity: function setWorkMeasurement(address _staker, bool _measureWork) returns(uint256)
func (_Nucypher *NucypherSession) SetWorkMeasurement(_staker common.Address, _measureWork bool) (*types.Transaction, error) {
	return _Nucypher.Contract.SetWorkMeasurement(&_Nucypher.TransactOpts, _staker, _measureWork)
}

// SetWorkMeasurement is a paid mutator transaction binding the contract method 0x49e5add4.
//
// Solidity: function setWorkMeasurement(address _staker, bool _measureWork) returns(uint256)
func (_Nucypher *NucypherTransactorSession) SetWorkMeasurement(_staker common.Address, _measureWork bool) (*types.Transaction, error) {
	return _Nucypher.Contract.SetWorkMeasurement(&_Nucypher.TransactOpts, _staker, _measureWork)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nucypher *NucypherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nucypher *NucypherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.TransferOwnership(&_Nucypher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nucypher *NucypherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.TransferOwnership(&_Nucypher.TransactOpts, newOwner)
}

// VerifyState is a paid mutator transaction binding the contract method 0xb9626d21.
//
// Solidity: function verifyState(address _testTarget) returns()
func (_Nucypher *NucypherTransactor) VerifyState(opts *bind.TransactOpts, _testTarget common.Address) (*types.Transaction, error) {
	return _Nucypher.contract.Transact(opts, "verifyState", _testTarget)
}

// VerifyState is a paid mutator transaction binding the contract method 0xb9626d21.
//
// Solidity: function verifyState(address _testTarget) returns()
func (_Nucypher *NucypherSession) VerifyState(_testTarget common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.VerifyState(&_Nucypher.TransactOpts, _testTarget)
}

// VerifyState is a paid mutator transaction binding the contract method 0xb9626d21.
//
// Solidity: function verifyState(address _testTarget) returns()
func (_Nucypher *NucypherTransactorSession) VerifyState(_testTarget common.Address) (*types.Transaction, error) {
	return _Nucypher.Contract.VerifyState(&_Nucypher.TransactOpts, _testTarget)
}

// NucypherDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Nucypher contract.
type NucypherDepositedIterator struct {
	Event *NucypherDeposited // Event containing the contract specifics and raw log

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
func (it *NucypherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherDeposited)
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
		it.Event = new(NucypherDeposited)
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
func (it *NucypherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherDeposited represents a Deposited event raised by the Nucypher contract.
type NucypherDeposited struct {
	Staker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) FilterDeposited(opts *bind.FilterOpts, staker []common.Address) (*NucypherDepositedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "Deposited", stakerRule)
	if err != nil {
		return nil, err
	}
	return &NucypherDepositedIterator{contract: _Nucypher.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NucypherDeposited, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "Deposited", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherDeposited)
				if err := _Nucypher.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) ParseDeposited(log types.Log) (*NucypherDeposited, error) {
	event := new(NucypherDeposited)
	if err := _Nucypher.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NucypherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nucypher contract.
type NucypherOwnershipTransferredIterator struct {
	Event *NucypherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NucypherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherOwnershipTransferred)
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
		it.Event = new(NucypherOwnershipTransferred)
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
func (it *NucypherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherOwnershipTransferred represents a OwnershipTransferred event raised by the Nucypher contract.
type NucypherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nucypher *NucypherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NucypherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NucypherOwnershipTransferredIterator{contract: _Nucypher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nucypher *NucypherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NucypherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherOwnershipTransferred)
				if err := _Nucypher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nucypher *NucypherFilterer) ParseOwnershipTransferred(log types.Log) (*NucypherOwnershipTransferred, error) {
	event := new(NucypherOwnershipTransferred)
	if err := _Nucypher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NucypherSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Nucypher contract.
type NucypherSlashedIterator struct {
	Event *NucypherSlashed // Event containing the contract specifics and raw log

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
func (it *NucypherSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherSlashed)
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
		it.Event = new(NucypherSlashed)
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
func (it *NucypherSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherSlashed represents a Slashed event raised by the Nucypher contract.
type NucypherSlashed struct {
	Staker       common.Address
	Penalty      *big.Int
	Investigator common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x34040fa1c0549c8572589c39cb87ff7638286467446a4790971d58d931c654bb.
//
// Solidity: event Slashed(address indexed staker, uint256 penalty, address indexed investigator, uint256 reward)
func (_Nucypher *NucypherFilterer) FilterSlashed(opts *bind.FilterOpts, staker []common.Address, investigator []common.Address) (*NucypherSlashedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	var investigatorRule []interface{}
	for _, investigatorItem := range investigator {
		investigatorRule = append(investigatorRule, investigatorItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "Slashed", stakerRule, investigatorRule)
	if err != nil {
		return nil, err
	}
	return &NucypherSlashedIterator{contract: _Nucypher.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x34040fa1c0549c8572589c39cb87ff7638286467446a4790971d58d931c654bb.
//
// Solidity: event Slashed(address indexed staker, uint256 penalty, address indexed investigator, uint256 reward)
func (_Nucypher *NucypherFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *NucypherSlashed, staker []common.Address, investigator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	var investigatorRule []interface{}
	for _, investigatorItem := range investigator {
		investigatorRule = append(investigatorRule, investigatorItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "Slashed", stakerRule, investigatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherSlashed)
				if err := _Nucypher.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x34040fa1c0549c8572589c39cb87ff7638286467446a4790971d58d931c654bb.
//
// Solidity: event Slashed(address indexed staker, uint256 penalty, address indexed investigator, uint256 reward)
func (_Nucypher *NucypherFilterer) ParseSlashed(log types.Log) (*NucypherSlashed, error) {
	event := new(NucypherSlashed)
	if err := _Nucypher.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NucypherStateVerifiedIterator is returned from FilterStateVerified and is used to iterate over the raw logs and unpacked data for StateVerified events raised by the Nucypher contract.
type NucypherStateVerifiedIterator struct {
	Event *NucypherStateVerified // Event containing the contract specifics and raw log

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
func (it *NucypherStateVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherStateVerified)
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
		it.Event = new(NucypherStateVerified)
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
func (it *NucypherStateVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherStateVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherStateVerified represents a StateVerified event raised by the Nucypher contract.
type NucypherStateVerified struct {
	TestTarget common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateVerified is a free log retrieval operation binding the contract event 0x1e8d98c1b4a0d9bd2e2371026b632eb2773fcce41742e41f02f574ab69868d4c.
//
// Solidity: event StateVerified(address indexed testTarget, address sender)
func (_Nucypher *NucypherFilterer) FilterStateVerified(opts *bind.FilterOpts, testTarget []common.Address) (*NucypherStateVerifiedIterator, error) {

	var testTargetRule []interface{}
	for _, testTargetItem := range testTarget {
		testTargetRule = append(testTargetRule, testTargetItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "StateVerified", testTargetRule)
	if err != nil {
		return nil, err
	}
	return &NucypherStateVerifiedIterator{contract: _Nucypher.contract, event: "StateVerified", logs: logs, sub: sub}, nil
}

// WatchStateVerified is a free log subscription operation binding the contract event 0x1e8d98c1b4a0d9bd2e2371026b632eb2773fcce41742e41f02f574ab69868d4c.
//
// Solidity: event StateVerified(address indexed testTarget, address sender)
func (_Nucypher *NucypherFilterer) WatchStateVerified(opts *bind.WatchOpts, sink chan<- *NucypherStateVerified, testTarget []common.Address) (event.Subscription, error) {

	var testTargetRule []interface{}
	for _, testTargetItem := range testTarget {
		testTargetRule = append(testTargetRule, testTargetItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "StateVerified", testTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherStateVerified)
				if err := _Nucypher.contract.UnpackLog(event, "StateVerified", log); err != nil {
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

// ParseStateVerified is a log parse operation binding the contract event 0x1e8d98c1b4a0d9bd2e2371026b632eb2773fcce41742e41f02f574ab69868d4c.
//
// Solidity: event StateVerified(address indexed testTarget, address sender)
func (_Nucypher *NucypherFilterer) ParseStateVerified(log types.Log) (*NucypherStateVerified, error) {
	event := new(NucypherStateVerified)
	if err := _Nucypher.contract.UnpackLog(event, "StateVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NucypherUpgradeFinishedIterator is returned from FilterUpgradeFinished and is used to iterate over the raw logs and unpacked data for UpgradeFinished events raised by the Nucypher contract.
type NucypherUpgradeFinishedIterator struct {
	Event *NucypherUpgradeFinished // Event containing the contract specifics and raw log

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
func (it *NucypherUpgradeFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherUpgradeFinished)
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
		it.Event = new(NucypherUpgradeFinished)
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
func (it *NucypherUpgradeFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherUpgradeFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherUpgradeFinished represents a UpgradeFinished event raised by the Nucypher contract.
type NucypherUpgradeFinished struct {
	Target common.Address
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpgradeFinished is a free log retrieval operation binding the contract event 0xd55ec27c5c6316913ed8803c18cfd1bfefea953db909dcba6140744a9d8b0d1f.
//
// Solidity: event UpgradeFinished(address indexed target, address sender)
func (_Nucypher *NucypherFilterer) FilterUpgradeFinished(opts *bind.FilterOpts, target []common.Address) (*NucypherUpgradeFinishedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "UpgradeFinished", targetRule)
	if err != nil {
		return nil, err
	}
	return &NucypherUpgradeFinishedIterator{contract: _Nucypher.contract, event: "UpgradeFinished", logs: logs, sub: sub}, nil
}

// WatchUpgradeFinished is a free log subscription operation binding the contract event 0xd55ec27c5c6316913ed8803c18cfd1bfefea953db909dcba6140744a9d8b0d1f.
//
// Solidity: event UpgradeFinished(address indexed target, address sender)
func (_Nucypher *NucypherFilterer) WatchUpgradeFinished(opts *bind.WatchOpts, sink chan<- *NucypherUpgradeFinished, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "UpgradeFinished", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherUpgradeFinished)
				if err := _Nucypher.contract.UnpackLog(event, "UpgradeFinished", log); err != nil {
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

// ParseUpgradeFinished is a log parse operation binding the contract event 0xd55ec27c5c6316913ed8803c18cfd1bfefea953db909dcba6140744a9d8b0d1f.
//
// Solidity: event UpgradeFinished(address indexed target, address sender)
func (_Nucypher *NucypherFilterer) ParseUpgradeFinished(log types.Log) (*NucypherUpgradeFinished, error) {
	event := new(NucypherUpgradeFinished)
	if err := _Nucypher.contract.UnpackLog(event, "UpgradeFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NucypherWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Nucypher contract.
type NucypherWithdrawnIterator struct {
	Event *NucypherWithdrawn // Event containing the contract specifics and raw log

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
func (it *NucypherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NucypherWithdrawn)
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
		it.Event = new(NucypherWithdrawn)
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
func (it *NucypherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NucypherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NucypherWithdrawn represents a Withdrawn event raised by the Nucypher contract.
type NucypherWithdrawn struct {
	Staker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) FilterWithdrawn(opts *bind.FilterOpts, staker []common.Address) (*NucypherWithdrawnIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Nucypher.contract.FilterLogs(opts, "Withdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return &NucypherWithdrawnIterator{contract: _Nucypher.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NucypherWithdrawn, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Nucypher.contract.WatchLogs(opts, "Withdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NucypherWithdrawn)
				if err := _Nucypher.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 value)
func (_Nucypher *NucypherFilterer) ParseWithdrawn(log types.Log) (*NucypherWithdrawn, error) {
	event := new(NucypherWithdrawn)
	if err := _Nucypher.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
