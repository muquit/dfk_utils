// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profile

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
)

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"created\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"picId\",\"type\":\"uint8\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"created\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"picId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"}],\"name\":\"ProfileUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"name\":\"changeHeroPic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"changeName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_picId\",\"type\":\"uint8\"}],\"name\":\"changePic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_picId\",\"type\":\"uint8\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAddressByName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"profileAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileAddress\",\"type\":\"address\"}],\"name\":\"getProfileByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_created\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_picId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getProfileByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_created\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_picId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProfileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"heroesNftContract\",\"outputs\":[{\"internalType\":\"contractIHeroCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"nameTaken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"taken\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileAddress\",\"type\":\"address\"}],\"name\":\"profileExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"created\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"picId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setHeroes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_min\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_max\",\"type\":\"uint8\"}],\"name\":\"setNameLengths\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_max\",\"type\":\"uint8\"}],\"name\":\"setPicMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// Profile is an auto generated Go binding around an Ethereum contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProfileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// AddressToIndex is a free data retrieval call binding the contract method 0xf667e0aa.
//
// Solidity: function addressToIndex(address ) view returns(uint256)
func (_Profile *ProfileCaller) AddressToIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "addressToIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToIndex is a free data retrieval call binding the contract method 0xf667e0aa.
//
// Solidity: function addressToIndex(address ) view returns(uint256)
func (_Profile *ProfileSession) AddressToIndex(arg0 common.Address) (*big.Int, error) {
	return _Profile.Contract.AddressToIndex(&_Profile.CallOpts, arg0)
}

// AddressToIndex is a free data retrieval call binding the contract method 0xf667e0aa.
//
// Solidity: function addressToIndex(address ) view returns(uint256)
func (_Profile *ProfileCallerSession) AddressToIndex(arg0 common.Address) (*big.Int, error) {
	return _Profile.Contract.AddressToIndex(&_Profile.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Profile *ProfileCaller) Addresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Profile *ProfileSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Profile.Contract.Addresses(&_Profile.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Profile *ProfileCallerSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Profile.Contract.Addresses(&_Profile.CallOpts, arg0)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string name) view returns(address profileAddress)
func (_Profile *ProfileCaller) GetAddressByName(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getAddressByName", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string name) view returns(address profileAddress)
func (_Profile *ProfileSession) GetAddressByName(name string) (common.Address, error) {
	return _Profile.Contract.GetAddressByName(&_Profile.CallOpts, name)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string name) view returns(address profileAddress)
func (_Profile *ProfileCallerSession) GetAddressByName(name string) (common.Address, error) {
	return _Profile.Contract.GetAddressByName(&_Profile.CallOpts, name)
}

// GetProfileByAddress is a free data retrieval call binding the contract method 0x4cd6959c.
//
// Solidity: function getProfileByAddress(address profileAddress) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileCaller) GetProfileByAddress(opts *bind.CallOpts, profileAddress common.Address) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfileByAddress", profileAddress)

	outstruct := new(struct {
		Id      *big.Int
		Owner   common.Address
		Name    string
		Created uint64
		PicId   uint8
		HeroId  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Created = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.PicId = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.HeroId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProfileByAddress is a free data retrieval call binding the contract method 0x4cd6959c.
//
// Solidity: function getProfileByAddress(address profileAddress) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileSession) GetProfileByAddress(profileAddress common.Address) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	return _Profile.Contract.GetProfileByAddress(&_Profile.CallOpts, profileAddress)
}

// GetProfileByAddress is a free data retrieval call binding the contract method 0x4cd6959c.
//
// Solidity: function getProfileByAddress(address profileAddress) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileCallerSession) GetProfileByAddress(profileAddress common.Address) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	return _Profile.Contract.GetProfileByAddress(&_Profile.CallOpts, profileAddress)
}

// GetProfileByName is a free data retrieval call binding the contract method 0x7d6df699.
//
// Solidity: function getProfileByName(string name) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileCaller) GetProfileByName(opts *bind.CallOpts, name string) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfileByName", name)

	outstruct := new(struct {
		Id      *big.Int
		Owner   common.Address
		Name    string
		Created uint64
		PicId   uint8
		HeroId  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Created = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.PicId = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.HeroId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProfileByName is a free data retrieval call binding the contract method 0x7d6df699.
//
// Solidity: function getProfileByName(string name) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileSession) GetProfileByName(name string) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	return _Profile.Contract.GetProfileByName(&_Profile.CallOpts, name)
}

// GetProfileByName is a free data retrieval call binding the contract method 0x7d6df699.
//
// Solidity: function getProfileByName(string name) view returns(uint256 _id, address _owner, string _name, uint64 _created, uint8 _picId, uint256 _heroId)
func (_Profile *ProfileCallerSession) GetProfileByName(name string) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
}, error) {
	return _Profile.Contract.GetProfileByName(&_Profile.CallOpts, name)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() view returns(uint256 count)
func (_Profile *ProfileCaller) GetProfileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() view returns(uint256 count)
func (_Profile *ProfileSession) GetProfileCount() (*big.Int, error) {
	return _Profile.Contract.GetProfileCount(&_Profile.CallOpts)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() view returns(uint256 count)
func (_Profile *ProfileCallerSession) GetProfileCount() (*big.Int, error) {
	return _Profile.Contract.GetProfileCount(&_Profile.CallOpts)
}

// HeroesNftContract is a free data retrieval call binding the contract method 0xb5708ccc.
//
// Solidity: function heroesNftContract() view returns(address)
func (_Profile *ProfileCaller) HeroesNftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "heroesNftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HeroesNftContract is a free data retrieval call binding the contract method 0xb5708ccc.
//
// Solidity: function heroesNftContract() view returns(address)
func (_Profile *ProfileSession) HeroesNftContract() (common.Address, error) {
	return _Profile.Contract.HeroesNftContract(&_Profile.CallOpts)
}

// HeroesNftContract is a free data retrieval call binding the contract method 0xb5708ccc.
//
// Solidity: function heroesNftContract() view returns(address)
func (_Profile *ProfileCallerSession) HeroesNftContract() (common.Address, error) {
	return _Profile.Contract.HeroesNftContract(&_Profile.CallOpts)
}

// NameTaken is a free data retrieval call binding the contract method 0x91fc437e.
//
// Solidity: function nameTaken(string name) view returns(bool taken)
func (_Profile *ProfileCaller) NameTaken(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "nameTaken", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NameTaken is a free data retrieval call binding the contract method 0x91fc437e.
//
// Solidity: function nameTaken(string name) view returns(bool taken)
func (_Profile *ProfileSession) NameTaken(name string) (bool, error) {
	return _Profile.Contract.NameTaken(&_Profile.CallOpts, name)
}

// NameTaken is a free data retrieval call binding the contract method 0x91fc437e.
//
// Solidity: function nameTaken(string name) view returns(bool taken)
func (_Profile *ProfileCallerSession) NameTaken(name string) (bool, error) {
	return _Profile.Contract.NameTaken(&_Profile.CallOpts, name)
}

// NameToIndex is a free data retrieval call binding the contract method 0xc1481626.
//
// Solidity: function nameToIndex(string ) view returns(uint256)
func (_Profile *ProfileCaller) NameToIndex(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "nameToIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToIndex is a free data retrieval call binding the contract method 0xc1481626.
//
// Solidity: function nameToIndex(string ) view returns(uint256)
func (_Profile *ProfileSession) NameToIndex(arg0 string) (*big.Int, error) {
	return _Profile.Contract.NameToIndex(&_Profile.CallOpts, arg0)
}

// NameToIndex is a free data retrieval call binding the contract method 0xc1481626.
//
// Solidity: function nameToIndex(string ) view returns(uint256)
func (_Profile *ProfileCallerSession) NameToIndex(arg0 string) (*big.Int, error) {
	return _Profile.Contract.NameToIndex(&_Profile.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCallerSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// ProfileExists is a free data retrieval call binding the contract method 0x05cfa06f.
//
// Solidity: function profileExists(address profileAddress) view returns(bool exists)
func (_Profile *ProfileCaller) ProfileExists(opts *bind.CallOpts, profileAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "profileExists", profileAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProfileExists is a free data retrieval call binding the contract method 0x05cfa06f.
//
// Solidity: function profileExists(address profileAddress) view returns(bool exists)
func (_Profile *ProfileSession) ProfileExists(profileAddress common.Address) (bool, error) {
	return _Profile.Contract.ProfileExists(&_Profile.CallOpts, profileAddress)
}

// ProfileExists is a free data retrieval call binding the contract method 0x05cfa06f.
//
// Solidity: function profileExists(address profileAddress) view returns(bool exists)
func (_Profile *ProfileCallerSession) ProfileExists(profileAddress common.Address) (bool, error) {
	return _Profile.Contract.ProfileExists(&_Profile.CallOpts, profileAddress)
}

// Profiles is a free data retrieval call binding the contract method 0xc36fe3d6.
//
// Solidity: function profiles(uint256 ) view returns(uint256 id, address owner, string name, uint64 created, uint8 picId, uint256 heroId, bool set)
func (_Profile *ProfileCaller) Profiles(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
	Set     bool
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "profiles", arg0)

	outstruct := new(struct {
		Id      *big.Int
		Owner   common.Address
		Name    string
		Created uint64
		PicId   uint8
		HeroId  *big.Int
		Set     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Created = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.PicId = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.HeroId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Set = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Profiles is a free data retrieval call binding the contract method 0xc36fe3d6.
//
// Solidity: function profiles(uint256 ) view returns(uint256 id, address owner, string name, uint64 created, uint8 picId, uint256 heroId, bool set)
func (_Profile *ProfileSession) Profiles(arg0 *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
	Set     bool
}, error) {
	return _Profile.Contract.Profiles(&_Profile.CallOpts, arg0)
}

// Profiles is a free data retrieval call binding the contract method 0xc36fe3d6.
//
// Solidity: function profiles(uint256 ) view returns(uint256 id, address owner, string name, uint64 created, uint8 picId, uint256 heroId, bool set)
func (_Profile *ProfileCallerSession) Profiles(arg0 *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	Name    string
	Created uint64
	PicId   uint8
	HeroId  *big.Int
	Set     bool
}, error) {
	return _Profile.Contract.Profiles(&_Profile.CallOpts, arg0)
}

// ChangeHeroPic is a paid mutator transaction binding the contract method 0xf6ebd47e.
//
// Solidity: function changeHeroPic(uint256 profileId, uint256 _heroId) returns(bool success)
func (_Profile *ProfileTransactor) ChangeHeroPic(opts *bind.TransactOpts, profileId *big.Int, _heroId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "changeHeroPic", profileId, _heroId)
}

// ChangeHeroPic is a paid mutator transaction binding the contract method 0xf6ebd47e.
//
// Solidity: function changeHeroPic(uint256 profileId, uint256 _heroId) returns(bool success)
func (_Profile *ProfileSession) ChangeHeroPic(profileId *big.Int, _heroId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.ChangeHeroPic(&_Profile.TransactOpts, profileId, _heroId)
}

// ChangeHeroPic is a paid mutator transaction binding the contract method 0xf6ebd47e.
//
// Solidity: function changeHeroPic(uint256 profileId, uint256 _heroId) returns(bool success)
func (_Profile *ProfileTransactorSession) ChangeHeroPic(profileId *big.Int, _heroId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.ChangeHeroPic(&_Profile.TransactOpts, profileId, _heroId)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 profileId, string _name) returns(bool success)
func (_Profile *ProfileTransactor) ChangeName(opts *bind.TransactOpts, profileId *big.Int, _name string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "changeName", profileId, _name)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 profileId, string _name) returns(bool success)
func (_Profile *ProfileSession) ChangeName(profileId *big.Int, _name string) (*types.Transaction, error) {
	return _Profile.Contract.ChangeName(&_Profile.TransactOpts, profileId, _name)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 profileId, string _name) returns(bool success)
func (_Profile *ProfileTransactorSession) ChangeName(profileId *big.Int, _name string) (*types.Transaction, error) {
	return _Profile.Contract.ChangeName(&_Profile.TransactOpts, profileId, _name)
}

// ChangePic is a paid mutator transaction binding the contract method 0x3f5df246.
//
// Solidity: function changePic(uint256 profileId, uint8 _picId) returns(bool success)
func (_Profile *ProfileTransactor) ChangePic(opts *bind.TransactOpts, profileId *big.Int, _picId uint8) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "changePic", profileId, _picId)
}

// ChangePic is a paid mutator transaction binding the contract method 0x3f5df246.
//
// Solidity: function changePic(uint256 profileId, uint8 _picId) returns(bool success)
func (_Profile *ProfileSession) ChangePic(profileId *big.Int, _picId uint8) (*types.Transaction, error) {
	return _Profile.Contract.ChangePic(&_Profile.TransactOpts, profileId, _picId)
}

// ChangePic is a paid mutator transaction binding the contract method 0x3f5df246.
//
// Solidity: function changePic(uint256 profileId, uint8 _picId) returns(bool success)
func (_Profile *ProfileTransactorSession) ChangePic(profileId *big.Int, _picId uint8) (*types.Transaction, error) {
	return _Profile.Contract.ChangePic(&_Profile.TransactOpts, profileId, _picId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd3eaa6f.
//
// Solidity: function createProfile(string _name, uint8 _picId) returns(bool success)
func (_Profile *ProfileTransactor) CreateProfile(opts *bind.TransactOpts, _name string, _picId uint8) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "createProfile", _name, _picId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd3eaa6f.
//
// Solidity: function createProfile(string _name, uint8 _picId) returns(bool success)
func (_Profile *ProfileSession) CreateProfile(_name string, _picId uint8) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, _name, _picId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd3eaa6f.
//
// Solidity: function createProfile(string _name, uint8 _picId) returns(bool success)
func (_Profile *ProfileTransactorSession) CreateProfile(_name string, _picId uint8) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, _name, _picId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Profile *ProfileTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Profile *ProfileSession) RenounceOwnership() (*types.Transaction, error) {
	return _Profile.Contract.RenounceOwnership(&_Profile.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Profile *ProfileTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Profile.Contract.RenounceOwnership(&_Profile.TransactOpts)
}

// SetHeroes is a paid mutator transaction binding the contract method 0x2fa32eb9.
//
// Solidity: function setHeroes(address _address) returns(bool success)
func (_Profile *ProfileTransactor) SetHeroes(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setHeroes", _address)
}

// SetHeroes is a paid mutator transaction binding the contract method 0x2fa32eb9.
//
// Solidity: function setHeroes(address _address) returns(bool success)
func (_Profile *ProfileSession) SetHeroes(_address common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetHeroes(&_Profile.TransactOpts, _address)
}

// SetHeroes is a paid mutator transaction binding the contract method 0x2fa32eb9.
//
// Solidity: function setHeroes(address _address) returns(bool success)
func (_Profile *ProfileTransactorSession) SetHeroes(_address common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetHeroes(&_Profile.TransactOpts, _address)
}

// SetNameLengths is a paid mutator transaction binding the contract method 0x393bca2d.
//
// Solidity: function setNameLengths(uint8 _min, uint8 _max) returns(bool success)
func (_Profile *ProfileTransactor) SetNameLengths(opts *bind.TransactOpts, _min uint8, _max uint8) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setNameLengths", _min, _max)
}

// SetNameLengths is a paid mutator transaction binding the contract method 0x393bca2d.
//
// Solidity: function setNameLengths(uint8 _min, uint8 _max) returns(bool success)
func (_Profile *ProfileSession) SetNameLengths(_min uint8, _max uint8) (*types.Transaction, error) {
	return _Profile.Contract.SetNameLengths(&_Profile.TransactOpts, _min, _max)
}

// SetNameLengths is a paid mutator transaction binding the contract method 0x393bca2d.
//
// Solidity: function setNameLengths(uint8 _min, uint8 _max) returns(bool success)
func (_Profile *ProfileTransactorSession) SetNameLengths(_min uint8, _max uint8) (*types.Transaction, error) {
	return _Profile.Contract.SetNameLengths(&_Profile.TransactOpts, _min, _max)
}

// SetPicMax is a paid mutator transaction binding the contract method 0x21319c49.
//
// Solidity: function setPicMax(uint8 _max) returns(bool success)
func (_Profile *ProfileTransactor) SetPicMax(opts *bind.TransactOpts, _max uint8) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setPicMax", _max)
}

// SetPicMax is a paid mutator transaction binding the contract method 0x21319c49.
//
// Solidity: function setPicMax(uint8 _max) returns(bool success)
func (_Profile *ProfileSession) SetPicMax(_max uint8) (*types.Transaction, error) {
	return _Profile.Contract.SetPicMax(&_Profile.TransactOpts, _max)
}

// SetPicMax is a paid mutator transaction binding the contract method 0x21319c49.
//
// Solidity: function setPicMax(uint8 _max) returns(bool success)
func (_Profile *ProfileTransactorSession) SetPicMax(_max uint8) (*types.Transaction, error) {
	return _Profile.Contract.SetPicMax(&_Profile.TransactOpts, _max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Profile *ProfileTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Profile *ProfileSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Profile.Contract.TransferOwnership(&_Profile.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Profile *ProfileTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Profile.Contract.TransferOwnership(&_Profile.TransactOpts, newOwner)
}

// ProfileOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Profile contract.
type ProfileOwnershipTransferredIterator struct {
	Event *ProfileOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProfileOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileOwnershipTransferred)
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
		it.Event = new(ProfileOwnershipTransferred)
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
func (it *ProfileOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileOwnershipTransferred represents a OwnershipTransferred event raised by the Profile contract.
type ProfileOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Profile *ProfileFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProfileOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProfileOwnershipTransferredIterator{contract: _Profile.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Profile *ProfileFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProfileOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileOwnershipTransferred)
				if err := _Profile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseOwnershipTransferred(log types.Log) (*ProfileOwnershipTransferred, error) {
	event := new(ProfileOwnershipTransferred)
	if err := _Profile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Profile contract.
type ProfileProfileCreatedIterator struct {
	Event *ProfileProfileCreated // Event containing the contract specifics and raw log

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
func (it *ProfileProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileProfileCreated)
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
		it.Event = new(ProfileProfileCreated)
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
func (it *ProfileProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileProfileCreated represents a ProfileCreated event raised by the Profile contract.
type ProfileProfileCreated struct {
	ProfileId *big.Int
	Owner     common.Address
	Name      string
	Created   uint64
	PicId     uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0x72a68484472f11fcd73602564e346c18c3d22b4869842397c7347330035c8a38.
//
// Solidity: event ProfileCreated(uint256 profileId, address owner, string name, uint64 created, uint8 picId)
func (_Profile *ProfileFilterer) FilterProfileCreated(opts *bind.FilterOpts) (*ProfileProfileCreatedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "ProfileCreated")
	if err != nil {
		return nil, err
	}
	return &ProfileProfileCreatedIterator{contract: _Profile.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0x72a68484472f11fcd73602564e346c18c3d22b4869842397c7347330035c8a38.
//
// Solidity: event ProfileCreated(uint256 profileId, address owner, string name, uint64 created, uint8 picId)
func (_Profile *ProfileFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *ProfileProfileCreated) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "ProfileCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileProfileCreated)
				if err := _Profile.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0x72a68484472f11fcd73602564e346c18c3d22b4869842397c7347330035c8a38.
//
// Solidity: event ProfileCreated(uint256 profileId, address owner, string name, uint64 created, uint8 picId)
func (_Profile *ProfileFilterer) ParseProfileCreated(log types.Log) (*ProfileProfileCreated, error) {
	event := new(ProfileProfileCreated)
	if err := _Profile.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileProfileUpdatedIterator is returned from FilterProfileUpdated and is used to iterate over the raw logs and unpacked data for ProfileUpdated events raised by the Profile contract.
type ProfileProfileUpdatedIterator struct {
	Event *ProfileProfileUpdated // Event containing the contract specifics and raw log

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
func (it *ProfileProfileUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileProfileUpdated)
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
		it.Event = new(ProfileProfileUpdated)
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
func (it *ProfileProfileUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileProfileUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileProfileUpdated represents a ProfileUpdated event raised by the Profile contract.
type ProfileProfileUpdated struct {
	ProfileId *big.Int
	Owner     common.Address
	Name      string
	Created   uint64
	PicId     uint8
	HeroId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileUpdated is a free log retrieval operation binding the contract event 0x69e29765eee9ce41e5c9e024d867f612cb1a6690a5c5f905201d4c591c423f1e.
//
// Solidity: event ProfileUpdated(uint256 profileId, address owner, string name, uint64 created, uint8 picId, uint256 heroId)
func (_Profile *ProfileFilterer) FilterProfileUpdated(opts *bind.FilterOpts) (*ProfileProfileUpdatedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "ProfileUpdated")
	if err != nil {
		return nil, err
	}
	return &ProfileProfileUpdatedIterator{contract: _Profile.contract, event: "ProfileUpdated", logs: logs, sub: sub}, nil
}

// WatchProfileUpdated is a free log subscription operation binding the contract event 0x69e29765eee9ce41e5c9e024d867f612cb1a6690a5c5f905201d4c591c423f1e.
//
// Solidity: event ProfileUpdated(uint256 profileId, address owner, string name, uint64 created, uint8 picId, uint256 heroId)
func (_Profile *ProfileFilterer) WatchProfileUpdated(opts *bind.WatchOpts, sink chan<- *ProfileProfileUpdated) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "ProfileUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileProfileUpdated)
				if err := _Profile.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
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

// ParseProfileUpdated is a log parse operation binding the contract event 0x69e29765eee9ce41e5c9e024d867f612cb1a6690a5c5f905201d4c591c423f1e.
//
// Solidity: event ProfileUpdated(uint256 profileId, address owner, string name, uint64 created, uint8 picId, uint256 heroId)
func (_Profile *ProfileFilterer) ParseProfileUpdated(log types.Log) (*ProfileProfileUpdated, error) {
	event := new(ProfileProfileUpdated)
	if err := _Profile.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
