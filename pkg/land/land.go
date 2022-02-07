// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package land

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

// LandCoreLandMeta is an auto generated low-level Go binding around an user-defined struct.
type LandCoreLandMeta struct {
	LandId  *big.Int
	Name    string
	Owner   common.Address
	Region  *big.Int
	Level   uint8
	Steward *big.Int
	Score   uint64
}

// LandMetaData contains all meta data concerning the Land contract.
var LandMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"}],\"name\":\"LandClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRegion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRegion\",\"type\":\"uint256\"}],\"name\":\"LandMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIMER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountLands\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"steward\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structLandCore.LandMeta[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllLands\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"steward\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structLandCore.LandMeta[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"}],\"name\":\"getLand\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"steward\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structLandCore.LandMeta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_region\",\"type\":\"uint32\"}],\"name\":\"getLandsByRegion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"steward\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structLandCore.LandMeta[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"landIdToMeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"region\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"steward\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"regionToLandCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"regionToLands\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_region\",\"type\":\"uint32\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_region\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldLandIndex\",\"type\":\"uint256\"}],\"name\":\"updateLandRegion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LandABI is the input ABI used to generate the binding from.
// Deprecated: Use LandMetaData.ABI instead.
var LandABI = LandMetaData.ABI

// Land is an auto generated Go binding around an Ethereum contract.
type Land struct {
	LandCaller     // Read-only binding to the contract
	LandTransactor // Write-only binding to the contract
	LandFilterer   // Log filterer for contract events
}

// LandCaller is an auto generated read-only Go binding around an Ethereum contract.
type LandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LandSession struct {
	Contract     *Land             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LandCallerSession struct {
	Contract *LandCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LandTransactorSession struct {
	Contract     *LandTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandRaw is an auto generated low-level Go binding around an Ethereum contract.
type LandRaw struct {
	Contract *Land // Generic contract binding to access the raw methods on
}

// LandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LandCallerRaw struct {
	Contract *LandCaller // Generic read-only contract binding to access the raw methods on
}

// LandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LandTransactorRaw struct {
	Contract *LandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLand creates a new instance of Land, bound to a specific deployed contract.
func NewLand(address common.Address, backend bind.ContractBackend) (*Land, error) {
	contract, err := bindLand(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Land{LandCaller: LandCaller{contract: contract}, LandTransactor: LandTransactor{contract: contract}, LandFilterer: LandFilterer{contract: contract}}, nil
}

// NewLandCaller creates a new read-only instance of Land, bound to a specific deployed contract.
func NewLandCaller(address common.Address, caller bind.ContractCaller) (*LandCaller, error) {
	contract, err := bindLand(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LandCaller{contract: contract}, nil
}

// NewLandTransactor creates a new write-only instance of Land, bound to a specific deployed contract.
func NewLandTransactor(address common.Address, transactor bind.ContractTransactor) (*LandTransactor, error) {
	contract, err := bindLand(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LandTransactor{contract: contract}, nil
}

// NewLandFilterer creates a new log filterer instance of Land, bound to a specific deployed contract.
func NewLandFilterer(address common.Address, filterer bind.ContractFilterer) (*LandFilterer, error) {
	contract, err := bindLand(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LandFilterer{contract: contract}, nil
}

// bindLand binds a generic wrapper to an already deployed contract.
func bindLand(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LandABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Land.Contract.LandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Land.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.contract.Transact(opts, method, params...)
}

// CLAIMERROLE is a free data retrieval call binding the contract method 0xfc5f18d3.
//
// Solidity: function CLAIMER_ROLE() view returns(bytes32)
func (_Land *LandCaller) CLAIMERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "CLAIMER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMERROLE is a free data retrieval call binding the contract method 0xfc5f18d3.
//
// Solidity: function CLAIMER_ROLE() view returns(bytes32)
func (_Land *LandSession) CLAIMERROLE() ([32]byte, error) {
	return _Land.Contract.CLAIMERROLE(&_Land.CallOpts)
}

// CLAIMERROLE is a free data retrieval call binding the contract method 0xfc5f18d3.
//
// Solidity: function CLAIMER_ROLE() view returns(bytes32)
func (_Land *LandCallerSession) CLAIMERROLE() ([32]byte, error) {
	return _Land.Contract.CLAIMERROLE(&_Land.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Land *LandCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Land *LandSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Land.Contract.DEFAULTADMINROLE(&_Land.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Land *LandCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Land.Contract.DEFAULTADMINROLE(&_Land.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Land *LandCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Land *LandSession) MINTERROLE() ([32]byte, error) {
	return _Land.Contract.MINTERROLE(&_Land.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Land *LandCallerSession) MINTERROLE() ([32]byte, error) {
	return _Land.Contract.MINTERROLE(&_Land.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Land *LandCaller) MODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Land *LandSession) MODERATORROLE() ([32]byte, error) {
	return _Land.Contract.MODERATORROLE(&_Land.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Land *LandCallerSession) MODERATORROLE() ([32]byte, error) {
	return _Land.Contract.MODERATORROLE(&_Land.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Land *LandCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Land *LandSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Land.Contract.BalanceOf(&_Land.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Land *LandCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Land.Contract.BalanceOf(&_Land.CallOpts, owner)
}

// GetAccountLands is a free data retrieval call binding the contract method 0xe4c14f22.
//
// Solidity: function getAccountLands(address _account) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCaller) GetAccountLands(opts *bind.CallOpts, _account common.Address) ([]LandCoreLandMeta, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getAccountLands", _account)

	if err != nil {
		return *new([]LandCoreLandMeta), err
	}

	out0 := *abi.ConvertType(out[0], new([]LandCoreLandMeta)).(*[]LandCoreLandMeta)

	return out0, err

}

// GetAccountLands is a free data retrieval call binding the contract method 0xe4c14f22.
//
// Solidity: function getAccountLands(address _account) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandSession) GetAccountLands(_account common.Address) ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetAccountLands(&_Land.CallOpts, _account)
}

// GetAccountLands is a free data retrieval call binding the contract method 0xe4c14f22.
//
// Solidity: function getAccountLands(address _account) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCallerSession) GetAccountLands(_account common.Address) ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetAccountLands(&_Land.CallOpts, _account)
}

// GetAllLands is a free data retrieval call binding the contract method 0xb7b6e543.
//
// Solidity: function getAllLands() view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCaller) GetAllLands(opts *bind.CallOpts) ([]LandCoreLandMeta, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getAllLands")

	if err != nil {
		return *new([]LandCoreLandMeta), err
	}

	out0 := *abi.ConvertType(out[0], new([]LandCoreLandMeta)).(*[]LandCoreLandMeta)

	return out0, err

}

// GetAllLands is a free data retrieval call binding the contract method 0xb7b6e543.
//
// Solidity: function getAllLands() view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandSession) GetAllLands() ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetAllLands(&_Land.CallOpts)
}

// GetAllLands is a free data retrieval call binding the contract method 0xb7b6e543.
//
// Solidity: function getAllLands() view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCallerSession) GetAllLands() ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetAllLands(&_Land.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Land *LandCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Land *LandSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.GetApproved(&_Land.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Land *LandCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.GetApproved(&_Land.CallOpts, tokenId)
}

// GetLand is a free data retrieval call binding the contract method 0xf02dd53f.
//
// Solidity: function getLand(uint256 _landId) view returns((uint256,string,address,uint256,uint8,uint256,uint64))
func (_Land *LandCaller) GetLand(opts *bind.CallOpts, _landId *big.Int) (LandCoreLandMeta, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getLand", _landId)

	if err != nil {
		return *new(LandCoreLandMeta), err
	}

	out0 := *abi.ConvertType(out[0], new(LandCoreLandMeta)).(*LandCoreLandMeta)

	return out0, err

}

// GetLand is a free data retrieval call binding the contract method 0xf02dd53f.
//
// Solidity: function getLand(uint256 _landId) view returns((uint256,string,address,uint256,uint8,uint256,uint64))
func (_Land *LandSession) GetLand(_landId *big.Int) (LandCoreLandMeta, error) {
	return _Land.Contract.GetLand(&_Land.CallOpts, _landId)
}

// GetLand is a free data retrieval call binding the contract method 0xf02dd53f.
//
// Solidity: function getLand(uint256 _landId) view returns((uint256,string,address,uint256,uint8,uint256,uint64))
func (_Land *LandCallerSession) GetLand(_landId *big.Int) (LandCoreLandMeta, error) {
	return _Land.Contract.GetLand(&_Land.CallOpts, _landId)
}

// GetLandsByRegion is a free data retrieval call binding the contract method 0xa24b5af1.
//
// Solidity: function getLandsByRegion(uint32 _region) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCaller) GetLandsByRegion(opts *bind.CallOpts, _region uint32) ([]LandCoreLandMeta, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getLandsByRegion", _region)

	if err != nil {
		return *new([]LandCoreLandMeta), err
	}

	out0 := *abi.ConvertType(out[0], new([]LandCoreLandMeta)).(*[]LandCoreLandMeta)

	return out0, err

}

// GetLandsByRegion is a free data retrieval call binding the contract method 0xa24b5af1.
//
// Solidity: function getLandsByRegion(uint32 _region) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandSession) GetLandsByRegion(_region uint32) ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetLandsByRegion(&_Land.CallOpts, _region)
}

// GetLandsByRegion is a free data retrieval call binding the contract method 0xa24b5af1.
//
// Solidity: function getLandsByRegion(uint32 _region) view returns((uint256,string,address,uint256,uint8,uint256,uint64)[])
func (_Land *LandCallerSession) GetLandsByRegion(_region uint32) ([]LandCoreLandMeta, error) {
	return _Land.Contract.GetLandsByRegion(&_Land.CallOpts, _region)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Land *LandCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Land *LandSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Land.Contract.GetRoleAdmin(&_Land.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Land *LandCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Land.Contract.GetRoleAdmin(&_Land.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Land *LandCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Land *LandSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Land.Contract.HasRole(&_Land.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Land *LandCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Land.Contract.HasRole(&_Land.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Land *LandCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Land *LandSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Land.Contract.IsApprovedForAll(&_Land.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Land *LandCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Land.Contract.IsApprovedForAll(&_Land.CallOpts, owner, operator)
}

// LandIdToMeta is a free data retrieval call binding the contract method 0x6d6dcc52.
//
// Solidity: function landIdToMeta(uint256 ) view returns(uint256 landId, string name, address owner, uint256 region, uint8 level, uint256 steward, uint64 score)
func (_Land *LandCaller) LandIdToMeta(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LandId  *big.Int
	Name    string
	Owner   common.Address
	Region  *big.Int
	Level   uint8
	Steward *big.Int
	Score   uint64
}, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "landIdToMeta", arg0)

	outstruct := new(struct {
		LandId  *big.Int
		Name    string
		Owner   common.Address
		Region  *big.Int
		Level   uint8
		Steward *big.Int
		Score   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LandId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Region = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Steward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Score = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// LandIdToMeta is a free data retrieval call binding the contract method 0x6d6dcc52.
//
// Solidity: function landIdToMeta(uint256 ) view returns(uint256 landId, string name, address owner, uint256 region, uint8 level, uint256 steward, uint64 score)
func (_Land *LandSession) LandIdToMeta(arg0 *big.Int) (struct {
	LandId  *big.Int
	Name    string
	Owner   common.Address
	Region  *big.Int
	Level   uint8
	Steward *big.Int
	Score   uint64
}, error) {
	return _Land.Contract.LandIdToMeta(&_Land.CallOpts, arg0)
}

// LandIdToMeta is a free data retrieval call binding the contract method 0x6d6dcc52.
//
// Solidity: function landIdToMeta(uint256 ) view returns(uint256 landId, string name, address owner, uint256 region, uint8 level, uint256 steward, uint64 score)
func (_Land *LandCallerSession) LandIdToMeta(arg0 *big.Int) (struct {
	LandId  *big.Int
	Name    string
	Owner   common.Address
	Region  *big.Int
	Level   uint8
	Steward *big.Int
	Score   uint64
}, error) {
	return _Land.Contract.LandIdToMeta(&_Land.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandSession) Name() (string, error) {
	return _Land.Contract.Name(&_Land.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandCallerSession) Name() (string, error) {
	return _Land.Contract.Name(&_Land.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Land *LandCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Land *LandSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Land.Contract.OnERC721Received(&_Land.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Land *LandCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Land.Contract.OnERC721Received(&_Land.CallOpts, arg0, arg1, arg2, arg3)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Land *LandCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Land *LandSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOf(&_Land.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Land *LandCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOf(&_Land.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandSession) Paused() (bool, error) {
	return _Land.Contract.Paused(&_Land.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandCallerSession) Paused() (bool, error) {
	return _Land.Contract.Paused(&_Land.CallOpts)
}

// RegionToLandCount is a free data retrieval call binding the contract method 0xd27fb5b7.
//
// Solidity: function regionToLandCount(uint256 ) view returns(uint256)
func (_Land *LandCaller) RegionToLandCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "regionToLandCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegionToLandCount is a free data retrieval call binding the contract method 0xd27fb5b7.
//
// Solidity: function regionToLandCount(uint256 ) view returns(uint256)
func (_Land *LandSession) RegionToLandCount(arg0 *big.Int) (*big.Int, error) {
	return _Land.Contract.RegionToLandCount(&_Land.CallOpts, arg0)
}

// RegionToLandCount is a free data retrieval call binding the contract method 0xd27fb5b7.
//
// Solidity: function regionToLandCount(uint256 ) view returns(uint256)
func (_Land *LandCallerSession) RegionToLandCount(arg0 *big.Int) (*big.Int, error) {
	return _Land.Contract.RegionToLandCount(&_Land.CallOpts, arg0)
}

// RegionToLands is a free data retrieval call binding the contract method 0xa1140396.
//
// Solidity: function regionToLands(uint256 , uint256 ) view returns(uint256)
func (_Land *LandCaller) RegionToLands(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "regionToLands", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegionToLands is a free data retrieval call binding the contract method 0xa1140396.
//
// Solidity: function regionToLands(uint256 , uint256 ) view returns(uint256)
func (_Land *LandSession) RegionToLands(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Land.Contract.RegionToLands(&_Land.CallOpts, arg0, arg1)
}

// RegionToLands is a free data retrieval call binding the contract method 0xa1140396.
//
// Solidity: function regionToLands(uint256 , uint256 ) view returns(uint256)
func (_Land *LandCallerSession) RegionToLands(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Land.Contract.RegionToLands(&_Land.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Land *LandCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Land *LandSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Land.Contract.SupportsInterface(&_Land.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Land *LandCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Land.Contract.SupportsInterface(&_Land.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandSession) Symbol() (string, error) {
	return _Land.Contract.Symbol(&_Land.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandCallerSession) Symbol() (string, error) {
	return _Land.Contract.Symbol(&_Land.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Land *LandCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Land *LandSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenByIndex(&_Land.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Land *LandCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenByIndex(&_Land.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Land *LandCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Land *LandSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenOfOwnerByIndex(&_Land.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Land *LandCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenOfOwnerByIndex(&_Land.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Land *LandCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Land *LandSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Land.Contract.TokenURI(&_Land.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Land *LandCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Land.Contract.TokenURI(&_Land.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Land *LandCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Land *LandSession) TotalSupply() (*big.Int, error) {
	return _Land.Contract.TotalSupply(&_Land.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Land *LandCallerSession) TotalSupply() (*big.Int, error) {
	return _Land.Contract.TotalSupply(&_Land.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Land *LandTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Land *LandSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Approve(&_Land.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Land *LandTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Approve(&_Land.TransactOpts, to, tokenId)
}

// ClaimLand is a paid mutator transaction binding the contract method 0x9705c2d6.
//
// Solidity: function claimLand(address _to, uint256 _tokenId) returns()
func (_Land *LandTransactor) ClaimLand(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "claimLand", _to, _tokenId)
}

// ClaimLand is a paid mutator transaction binding the contract method 0x9705c2d6.
//
// Solidity: function claimLand(address _to, uint256 _tokenId) returns()
func (_Land *LandSession) ClaimLand(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.ClaimLand(&_Land.TransactOpts, _to, _tokenId)
}

// ClaimLand is a paid mutator transaction binding the contract method 0x9705c2d6.
//
// Solidity: function claimLand(address _to, uint256 _tokenId) returns()
func (_Land *LandTransactorSession) ClaimLand(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.ClaimLand(&_Land.TransactOpts, _to, _tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Land *LandTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Land *LandSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.GrantRole(&_Land.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Land *LandTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.GrantRole(&_Land.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Land *LandTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Land *LandSession) Initialize() (*types.Transaction, error) {
	return _Land.Contract.Initialize(&_Land.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Land *LandTransactorSession) Initialize() (*types.Transaction, error) {
	return _Land.Contract.Initialize(&_Land.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandSession) Pause() (*types.Transaction, error) {
	return _Land.Contract.Pause(&_Land.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandTransactorSession) Pause() (*types.Transaction, error) {
	return _Land.Contract.Pause(&_Land.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Land *LandTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Land *LandSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.RenounceRole(&_Land.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Land *LandTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.RenounceRole(&_Land.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Land *LandTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Land *LandSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.RevokeRole(&_Land.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Land *LandTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Land.Contract.RevokeRole(&_Land.TransactOpts, role, account)
}

// SafeMint is a paid mutator transaction binding the contract method 0x402b0436.
//
// Solidity: function safeMint(address _owner, uint256 _landId, string _name, uint32 _region) returns()
func (_Land *LandTransactor) SafeMint(opts *bind.TransactOpts, _owner common.Address, _landId *big.Int, _name string, _region uint32) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "safeMint", _owner, _landId, _name, _region)
}

// SafeMint is a paid mutator transaction binding the contract method 0x402b0436.
//
// Solidity: function safeMint(address _owner, uint256 _landId, string _name, uint32 _region) returns()
func (_Land *LandSession) SafeMint(_owner common.Address, _landId *big.Int, _name string, _region uint32) (*types.Transaction, error) {
	return _Land.Contract.SafeMint(&_Land.TransactOpts, _owner, _landId, _name, _region)
}

// SafeMint is a paid mutator transaction binding the contract method 0x402b0436.
//
// Solidity: function safeMint(address _owner, uint256 _landId, string _name, uint32 _region) returns()
func (_Land *LandTransactorSession) SafeMint(_owner common.Address, _landId *big.Int, _name string, _region uint32) (*types.Transaction, error) {
	return _Land.Contract.SafeMint(&_Land.TransactOpts, _owner, _landId, _name, _region)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom(&_Land.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom(&_Land.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Land *LandTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Land *LandSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom0(&_Land.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Land *LandTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom0(&_Land.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Land *LandTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Land *LandSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Land.Contract.SetApprovalForAll(&_Land.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Land *LandTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Land.Contract.SetApprovalForAll(&_Land.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.TransferFrom(&_Land.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Land *LandTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.TransferFrom(&_Land.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandSession) Unpause() (*types.Transaction, error) {
	return _Land.Contract.Unpause(&_Land.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandTransactorSession) Unpause() (*types.Transaction, error) {
	return _Land.Contract.Unpause(&_Land.TransactOpts)
}

// UpdateLandRegion is a paid mutator transaction binding the contract method 0x81ea4eb9.
//
// Solidity: function updateLandRegion(uint256 _landId, uint256 _region, uint256 _oldLandIndex) returns()
func (_Land *LandTransactor) UpdateLandRegion(opts *bind.TransactOpts, _landId *big.Int, _region *big.Int, _oldLandIndex *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "updateLandRegion", _landId, _region, _oldLandIndex)
}

// UpdateLandRegion is a paid mutator transaction binding the contract method 0x81ea4eb9.
//
// Solidity: function updateLandRegion(uint256 _landId, uint256 _region, uint256 _oldLandIndex) returns()
func (_Land *LandSession) UpdateLandRegion(_landId *big.Int, _region *big.Int, _oldLandIndex *big.Int) (*types.Transaction, error) {
	return _Land.Contract.UpdateLandRegion(&_Land.TransactOpts, _landId, _region, _oldLandIndex)
}

// UpdateLandRegion is a paid mutator transaction binding the contract method 0x81ea4eb9.
//
// Solidity: function updateLandRegion(uint256 _landId, uint256 _region, uint256 _oldLandIndex) returns()
func (_Land *LandTransactorSession) UpdateLandRegion(_landId *big.Int, _region *big.Int, _oldLandIndex *big.Int) (*types.Transaction, error) {
	return _Land.Contract.UpdateLandRegion(&_Land.TransactOpts, _landId, _region, _oldLandIndex)
}

// LandApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Land contract.
type LandApprovalIterator struct {
	Event *LandApproval // Event containing the contract specifics and raw log

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
func (it *LandApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandApproval)
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
		it.Event = new(LandApproval)
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
func (it *LandApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandApproval represents a Approval event raised by the Land contract.
type LandApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Land *LandFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LandApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LandApprovalIterator{contract: _Land.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Land *LandFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LandApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandApproval)
				if err := _Land.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Land *LandFilterer) ParseApproval(log types.Log) (*LandApproval, error) {
	event := new(LandApproval)
	if err := _Land.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Land contract.
type LandApprovalForAllIterator struct {
	Event *LandApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LandApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandApprovalForAll)
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
		it.Event = new(LandApprovalForAll)
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
func (it *LandApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandApprovalForAll represents a ApprovalForAll event raised by the Land contract.
type LandApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Land *LandFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LandApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LandApprovalForAllIterator{contract: _Land.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Land *LandFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LandApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandApprovalForAll)
				if err := _Land.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Land *LandFilterer) ParseApprovalForAll(log types.Log) (*LandApprovalForAll, error) {
	event := new(LandApprovalForAll)
	if err := _Land.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandLandClaimedIterator is returned from FilterLandClaimed and is used to iterate over the raw logs and unpacked data for LandClaimed events raised by the Land contract.
type LandLandClaimedIterator struct {
	Event *LandLandClaimed // Event containing the contract specifics and raw log

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
func (it *LandLandClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandLandClaimed)
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
		it.Event = new(LandLandClaimed)
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
func (it *LandLandClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandLandClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandLandClaimed represents a LandClaimed event raised by the Land contract.
type LandLandClaimed struct {
	Owner  common.Address
	LandId *big.Int
	Region *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLandClaimed is a free log retrieval operation binding the contract event 0xe9d066f43c8a7c765b15b359553a8d0058f345fc38d18799ff4e431ea4f9ea9e.
//
// Solidity: event LandClaimed(address indexed owner, uint256 landId, uint256 indexed region)
func (_Land *LandFilterer) FilterLandClaimed(opts *bind.FilterOpts, owner []common.Address, region []*big.Int) (*LandLandClaimedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var regionRule []interface{}
	for _, regionItem := range region {
		regionRule = append(regionRule, regionItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "LandClaimed", ownerRule, regionRule)
	if err != nil {
		return nil, err
	}
	return &LandLandClaimedIterator{contract: _Land.contract, event: "LandClaimed", logs: logs, sub: sub}, nil
}

// WatchLandClaimed is a free log subscription operation binding the contract event 0xe9d066f43c8a7c765b15b359553a8d0058f345fc38d18799ff4e431ea4f9ea9e.
//
// Solidity: event LandClaimed(address indexed owner, uint256 landId, uint256 indexed region)
func (_Land *LandFilterer) WatchLandClaimed(opts *bind.WatchOpts, sink chan<- *LandLandClaimed, owner []common.Address, region []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var regionRule []interface{}
	for _, regionItem := range region {
		regionRule = append(regionRule, regionItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "LandClaimed", ownerRule, regionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandLandClaimed)
				if err := _Land.contract.UnpackLog(event, "LandClaimed", log); err != nil {
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

// ParseLandClaimed is a log parse operation binding the contract event 0xe9d066f43c8a7c765b15b359553a8d0058f345fc38d18799ff4e431ea4f9ea9e.
//
// Solidity: event LandClaimed(address indexed owner, uint256 landId, uint256 indexed region)
func (_Land *LandFilterer) ParseLandClaimed(log types.Log) (*LandLandClaimed, error) {
	event := new(LandLandClaimed)
	if err := _Land.contract.UnpackLog(event, "LandClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandLandMovedIterator is returned from FilterLandMoved and is used to iterate over the raw logs and unpacked data for LandMoved events raised by the Land contract.
type LandLandMovedIterator struct {
	Event *LandLandMoved // Event containing the contract specifics and raw log

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
func (it *LandLandMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandLandMoved)
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
		it.Event = new(LandLandMoved)
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
func (it *LandLandMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandLandMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandLandMoved represents a LandMoved event raised by the Land contract.
type LandLandMoved struct {
	LandId    *big.Int
	OldRegion *big.Int
	NewRegion *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLandMoved is a free log retrieval operation binding the contract event 0xa2fb11ef6ad5a913e97a22811dfeacf74a4ac38a4f4591b83be1861f934f8b36.
//
// Solidity: event LandMoved(uint256 landId, uint256 oldRegion, uint256 newRegion)
func (_Land *LandFilterer) FilterLandMoved(opts *bind.FilterOpts) (*LandLandMovedIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "LandMoved")
	if err != nil {
		return nil, err
	}
	return &LandLandMovedIterator{contract: _Land.contract, event: "LandMoved", logs: logs, sub: sub}, nil
}

// WatchLandMoved is a free log subscription operation binding the contract event 0xa2fb11ef6ad5a913e97a22811dfeacf74a4ac38a4f4591b83be1861f934f8b36.
//
// Solidity: event LandMoved(uint256 landId, uint256 oldRegion, uint256 newRegion)
func (_Land *LandFilterer) WatchLandMoved(opts *bind.WatchOpts, sink chan<- *LandLandMoved) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "LandMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandLandMoved)
				if err := _Land.contract.UnpackLog(event, "LandMoved", log); err != nil {
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

// ParseLandMoved is a log parse operation binding the contract event 0xa2fb11ef6ad5a913e97a22811dfeacf74a4ac38a4f4591b83be1861f934f8b36.
//
// Solidity: event LandMoved(uint256 landId, uint256 oldRegion, uint256 newRegion)
func (_Land *LandFilterer) ParseLandMoved(log types.Log) (*LandLandMoved, error) {
	event := new(LandLandMoved)
	if err := _Land.contract.UnpackLog(event, "LandMoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Land contract.
type LandPausedIterator struct {
	Event *LandPaused // Event containing the contract specifics and raw log

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
func (it *LandPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandPaused)
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
		it.Event = new(LandPaused)
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
func (it *LandPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandPaused represents a Paused event raised by the Land contract.
type LandPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Land *LandFilterer) FilterPaused(opts *bind.FilterOpts) (*LandPausedIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LandPausedIterator{contract: _Land.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Land *LandFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LandPaused) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandPaused)
				if err := _Land.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Land *LandFilterer) ParsePaused(log types.Log) (*LandPaused, error) {
	event := new(LandPaused)
	if err := _Land.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Land contract.
type LandRoleAdminChangedIterator struct {
	Event *LandRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LandRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandRoleAdminChanged)
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
		it.Event = new(LandRoleAdminChanged)
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
func (it *LandRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandRoleAdminChanged represents a RoleAdminChanged event raised by the Land contract.
type LandRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Land *LandFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LandRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LandRoleAdminChangedIterator{contract: _Land.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Land *LandFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LandRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandRoleAdminChanged)
				if err := _Land.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Land *LandFilterer) ParseRoleAdminChanged(log types.Log) (*LandRoleAdminChanged, error) {
	event := new(LandRoleAdminChanged)
	if err := _Land.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Land contract.
type LandRoleGrantedIterator struct {
	Event *LandRoleGranted // Event containing the contract specifics and raw log

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
func (it *LandRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandRoleGranted)
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
		it.Event = new(LandRoleGranted)
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
func (it *LandRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandRoleGranted represents a RoleGranted event raised by the Land contract.
type LandRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LandRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LandRoleGrantedIterator{contract: _Land.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LandRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandRoleGranted)
				if err := _Land.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) ParseRoleGranted(log types.Log) (*LandRoleGranted, error) {
	event := new(LandRoleGranted)
	if err := _Land.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Land contract.
type LandRoleRevokedIterator struct {
	Event *LandRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LandRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandRoleRevoked)
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
		it.Event = new(LandRoleRevoked)
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
func (it *LandRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandRoleRevoked represents a RoleRevoked event raised by the Land contract.
type LandRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LandRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LandRoleRevokedIterator{contract: _Land.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LandRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandRoleRevoked)
				if err := _Land.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Land *LandFilterer) ParseRoleRevoked(log types.Log) (*LandRoleRevoked, error) {
	event := new(LandRoleRevoked)
	if err := _Land.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Land contract.
type LandTransferIterator struct {
	Event *LandTransfer // Event containing the contract specifics and raw log

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
func (it *LandTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandTransfer)
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
		it.Event = new(LandTransfer)
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
func (it *LandTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandTransfer represents a Transfer event raised by the Land contract.
type LandTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Land *LandFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LandTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LandTransferIterator{contract: _Land.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Land *LandFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LandTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandTransfer)
				if err := _Land.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Land *LandFilterer) ParseTransfer(log types.Log) (*LandTransfer, error) {
	event := new(LandTransfer)
	if err := _Land.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Land contract.
type LandUnpausedIterator struct {
	Event *LandUnpaused // Event containing the contract specifics and raw log

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
func (it *LandUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandUnpaused)
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
		it.Event = new(LandUnpaused)
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
func (it *LandUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandUnpaused represents a Unpaused event raised by the Land contract.
type LandUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Land *LandFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LandUnpausedIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LandUnpausedIterator{contract: _Land.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Land *LandFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LandUnpaused) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandUnpaused)
				if err := _Land.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Land *LandFilterer) ParseUnpaused(log types.Log) (*LandUnpaused, error) {
	event := new(LandUnpaused)
	if err := _Land.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
