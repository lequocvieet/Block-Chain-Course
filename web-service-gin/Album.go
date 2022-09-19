// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// AlbumAlbumInfo is an auto generated low-level Go binding around an user-defined struct.
type AlbumAlbumInfo struct {
	Price  *big.Int
	Id     string
	Artist string
	Title  string
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"albumInfo\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artist\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artist\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"internalType\":\"structAlbum.AlbumInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_artist\",\"type\":\"string\"}],\"name\":\"setInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610bde806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635a9b0b891461004657806390fe63ac14610064578063d49f016a14610080575b600080fd5b61004e6100a1565b60405161005b91906105bb565b60405180910390f35b61007e60048036038101906100799190610752565b61027b565b005b6100886102c3565b6040516100989493929190610866565b60405180910390f35b6100a9610479565b6000604051806080016040529081600082015481526020016001820180546100d0906108ef565b80601f01602080910402602001604051908101604052809291908181526020018280546100fc906108ef565b80156101495780601f1061011e57610100808354040283529160200191610149565b820191906000526020600020905b81548152906001019060200180831161012c57829003601f168201915b50505050508152602001600282018054610162906108ef565b80601f016020809104026020016040519081016040528092919081815260200182805461018e906108ef565b80156101db5780601f106101b0576101008083540402835291602001916101db565b820191906000526020600020905b8154815290600101906020018083116101be57829003601f168201915b505050505081526020016003820180546101f4906108ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610220906108ef565b801561026d5780601f106102425761010080835404028352916020019161026d565b820191906000526020600020905b81548152906001019060200180831161025057829003601f168201915b505050505081525050905090565b83600080018190555082600060010190816102969190610ad6565b5081600060030190816102a99190610ad6565b5080600060020190816102bc9190610ad6565b5050505050565b60008060000154908060010180546102da906108ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610306906108ef565b80156103535780601f1061032857610100808354040283529160200191610353565b820191906000526020600020905b81548152906001019060200180831161033657829003601f168201915b505050505090806002018054610368906108ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610394906108ef565b80156103e15780601f106103b6576101008083540402835291602001916103e1565b820191906000526020600020905b8154815290600101906020018083116103c457829003601f168201915b5050505050908060030180546103f6906108ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610422906108ef565b801561046f5780601f106104445761010080835404028352916020019161046f565b820191906000526020600020905b81548152906001019060200180831161045257829003601f168201915b5050505050905084565b6040518060800160405280600081526020016060815260200160608152602001606081525090565b6000819050919050565b6104b4816104a1565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104f45780820151818401526020810190506104d9565b60008484015250505050565b6000601f19601f8301169050919050565b600061051c826104ba565b61052681856104c5565b93506105368185602086016104d6565b61053f81610500565b840191505092915050565b600060808301600083015161056260008601826104ab565b506020830151848203602086015261057a8282610511565b915050604083015184820360408601526105948282610511565b915050606083015184820360608601526105ae8282610511565b9150508091505092915050565b600060208201905081810360008301526105d5818461054a565b905092915050565b6000604051905090565b600080fd5b600080fd5b6105fa816104a1565b811461060557600080fd5b50565b600081359050610617816105f1565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61065f82610500565b810181811067ffffffffffffffff8211171561067e5761067d610627565b5b80604052505050565b60006106916105dd565b905061069d8282610656565b919050565b600067ffffffffffffffff8211156106bd576106bc610627565b5b6106c682610500565b9050602081019050919050565b82818337600083830152505050565b60006106f56106f0846106a2565b610687565b90508281526020810184848401111561071157610710610622565b5b61071c8482856106d3565b509392505050565b600082601f8301126107395761073861061d565b5b81356107498482602086016106e2565b91505092915050565b6000806000806080858703121561076c5761076b6105e7565b5b600061077a87828801610608565b945050602085013567ffffffffffffffff81111561079b5761079a6105ec565b5b6107a787828801610724565b935050604085013567ffffffffffffffff8111156107c8576107c76105ec565b5b6107d487828801610724565b925050606085013567ffffffffffffffff8111156107f5576107f46105ec565b5b61080187828801610724565b91505092959194509250565b610816816104a1565b82525050565b600082825260208201905092915050565b6000610838826104ba565b610842818561081c565b93506108528185602086016104d6565b61085b81610500565b840191505092915050565b600060808201905061087b600083018761080d565b818103602083015261088d818661082d565b905081810360408301526108a1818561082d565b905081810360608301526108b5818461082d565b905095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061090757607f821691505b60208210810361091a576109196108c0565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610945565b61098c8683610945565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006109d36109ce6109c9846109a4565b6109ae565b6109a4565b9050919050565b6000819050919050565b6109ed836109b8565b610a016109f9826109da565b848454610952565b825550505050565b600090565b610a16610a09565b610a218184846109e4565b505050565b5b81811015610a4557610a3a600082610a0e565b600181019050610a27565b5050565b601f821115610a8a57610a5b81610920565b610a6484610935565b81016020851015610a73578190505b610a87610a7f85610935565b830182610a26565b50505b505050565b600082821c905092915050565b6000610aad60001984600802610a8f565b1980831691505092915050565b6000610ac68383610a9c565b9150826002028217905092915050565b610adf826104ba565b67ffffffffffffffff811115610af857610af7610627565b5b610b0282546108ef565b610b0d828285610a49565b600060209050601f831160018114610b405760008415610b2e578287015190505b610b388582610aba565b865550610ba0565b601f198416610b4e86610920565b60005b82811015610b7657848901518255600182019150602085019450602081019050610b51565b86831015610b935784890151610b8f601f891682610a9c565b8355505b6001600288020188555050505b50505050505056fea26469706673582212200be61659b8fa497cd367a85645a913435f6747a729540e4cea1164017b91c7a764736f6c63430008110033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// AlbumInfo is a free data retrieval call binding the contract method 0xd49f016a.
//
// Solidity: function albumInfo() view returns(int256 price, string id, string artist, string title)
func (_Storage *StorageCaller) AlbumInfo(opts *bind.CallOpts) (struct {
	Price  *big.Int
	Id     string
	Artist string
	Title  string
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "albumInfo")

	outstruct := new(struct {
		Price  *big.Int
		Id     string
		Artist string
		Title  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Artist = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// AlbumInfo is a free data retrieval call binding the contract method 0xd49f016a.
//
// Solidity: function albumInfo() view returns(int256 price, string id, string artist, string title)
func (_Storage *StorageSession) AlbumInfo() (struct {
	Price  *big.Int
	Id     string
	Artist string
	Title  string
}, error) {
	return _Storage.Contract.AlbumInfo(&_Storage.CallOpts)
}

// AlbumInfo is a free data retrieval call binding the contract method 0xd49f016a.
//
// Solidity: function albumInfo() view returns(int256 price, string id, string artist, string title)
func (_Storage *StorageCallerSession) AlbumInfo() (struct {
	Price  *big.Int
	Id     string
	Artist string
	Title  string
}, error) {
	return _Storage.Contract.AlbumInfo(&_Storage.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((int256,string,string,string))
func (_Storage *StorageCaller) GetInfo(opts *bind.CallOpts) (AlbumAlbumInfo, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getInfo")

	if err != nil {
		return *new(AlbumAlbumInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AlbumAlbumInfo)).(*AlbumAlbumInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((int256,string,string,string))
func (_Storage *StorageSession) GetInfo() (AlbumAlbumInfo, error) {
	return _Storage.Contract.GetInfo(&_Storage.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((int256,string,string,string))
func (_Storage *StorageCallerSession) GetInfo() (AlbumAlbumInfo, error) {
	return _Storage.Contract.GetInfo(&_Storage.CallOpts)
}

// SetInfo is a paid mutator transaction binding the contract method 0x90fe63ac.
//
// Solidity: function setInfo(int256 _price, string _id, string _title, string _artist) returns()
func (_Storage *StorageTransactor) SetInfo(opts *bind.TransactOpts, _price *big.Int, _id string, _title string, _artist string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setInfo", _price, _id, _title, _artist)
}

// SetInfo is a paid mutator transaction binding the contract method 0x90fe63ac.
//
// Solidity: function setInfo(int256 _price, string _id, string _title, string _artist) returns()
func (_Storage *StorageSession) SetInfo(_price *big.Int, _id string, _title string, _artist string) (*types.Transaction, error) {
	return _Storage.Contract.SetInfo(&_Storage.TransactOpts, _price, _id, _title, _artist)
}

// SetInfo is a paid mutator transaction binding the contract method 0x90fe63ac.
//
// Solidity: function setInfo(int256 _price, string _id, string _title, string _artist) returns()
func (_Storage *StorageTransactorSession) SetInfo(_price *big.Int, _id string, _title string, _artist string) (*types.Transaction, error) {
	return _Storage.Contract.SetInfo(&_Storage.TransactOpts, _price, _id, _title, _artist)
}
