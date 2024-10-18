// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"u256\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b322\",\"type\":\"bytes32\"}],\"name\":\"Output\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"SetBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetUint256\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"requestedBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"requestedBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"requestedUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610a898061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c806368895979116100645780636889597914610131578063c2b12a731461014f578063d2282dc51461016b578063da359dc814610187578063ed53e511146101a35761009b565b80626d6cae1461009f5780630bcd3b33146100bd5780631f903037146100db5780633345b4d0146100f957806346ddd1ff14610115575b5f80fd5b6100a76101bf565b6040516100b491906103ac565b60405180910390f35b6100c56101c5565b6040516100d29190610435565b60405180910390f35b6100e3610251565b6040516100f091906103ac565b60405180910390f35b610113600480360381019061010e91906104c3565b610256565b005b61012f600480360381019061012a919061062d565b61026a565b005b61013961027e565b6040516101469190610696565b60405180910390f35b610169600480360381019061016491906106af565b610284565b005b610185600480360381019061018091906106da565b6102d1565b005b6101a1600480360381019061019c9190610705565b61031f565b005b6101bd60048036038101906101b8919061074c565b610380565b005b60025481565b600380546101d2906107b7565b80601f01602080910402602001604051908101604052809291908181526020018280546101fe906107b7565b80156102495780601f1061022057610100808354040283529160200191610249565b820191905f5260205f20905b81548152906001019060200180831161022c57829003601f168201915b505050505081565b5f5481565b81600281905550610266816102d1565b5050565b8160028190555061027a8161031f565b5050565b60015481565b805f81905550803373ffffffffffffffffffffffffffffffffffffffff167fdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d60405160405180910390a350565b80600181905550803373ffffffffffffffffffffffffffffffffffffffff167fd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d60405160405180910390a350565b806003908161032e9190610984565b503373ffffffffffffffffffffffffffffffffffffffff167ff22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf826040516103759190610435565b60405180910390a250565b8160028190555061039081610284565b5050565b5f819050919050565b6103a681610394565b82525050565b5f6020820190506103bf5f83018461039d565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610407826103c5565b61041181856103cf565b93506104218185602086016103df565b61042a816103ed565b840191505092915050565b5f6020820190508181035f83015261044d81846103fd565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b61046f81610394565b8114610479575f80fd5b50565b5f8135905061048a81610466565b92915050565b5f819050919050565b6104a281610490565b81146104ac575f80fd5b50565b5f813590506104bd81610499565b92915050565b5f80604083850312156104d9576104d861045e565b5b5f6104e68582860161047c565b92505060206104f7858286016104af565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61053f826103ed565b810181811067ffffffffffffffff8211171561055e5761055d610509565b5b80604052505050565b5f610570610455565b905061057c8282610536565b919050565b5f67ffffffffffffffff82111561059b5761059a610509565b5b6105a4826103ed565b9050602081019050919050565b828183375f83830152505050565b5f6105d16105cc84610581565b610567565b9050828152602081018484840111156105ed576105ec610505565b5b6105f88482856105b1565b509392505050565b5f82601f83011261061457610613610501565b5b81356106248482602086016105bf565b91505092915050565b5f80604083850312156106435761064261045e565b5b5f6106508582860161047c565b925050602083013567ffffffffffffffff81111561067157610670610462565b5b61067d85828601610600565b9150509250929050565b61069081610490565b82525050565b5f6020820190506106a95f830184610687565b92915050565b5f602082840312156106c4576106c361045e565b5b5f6106d18482850161047c565b91505092915050565b5f602082840312156106ef576106ee61045e565b5b5f6106fc848285016104af565b91505092915050565b5f6020828403121561071a5761071961045e565b5b5f82013567ffffffffffffffff81111561073757610736610462565b5b61074384828501610600565b91505092915050565b5f80604083850312156107625761076161045e565b5b5f61076f8582860161047c565b92505060206107808582860161047c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806107ce57607f821691505b6020821081036107e1576107e061078a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108437fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610808565b61084d8683610808565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61088861088361087e84610490565b610865565b610490565b9050919050565b5f819050919050565b6108a18361086e565b6108b56108ad8261088f565b848454610814565b825550505050565b5f90565b6108c96108bd565b6108d4818484610898565b505050565b5b818110156108f7576108ec5f826108c1565b6001810190506108da565b5050565b601f82111561093c5761090d816107e7565b610916846107f9565b81016020851015610925578190505b610939610931856107f9565b8301826108d9565b50505b505050565b5f82821c905092915050565b5f61095c5f1984600802610941565b1980831691505092915050565b5f610974838361094d565b9150826002028217905092915050565b61098d826103c5565b67ffffffffffffffff8111156109a6576109a5610509565b5b6109b082546107b7565b6109bb8282856108fb565b5f60209050601f8311600181146109ec575f84156109da578287015190505b6109e48582610969565b865550610a4b565b601f1984166109fa866107e7565b5f5b82811015610a21578489015182556001820191506020850194506020810190506109fc565b86831015610a3e5784890151610a3a601f89168261094d565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220d3f409b92b9b94290cc632fbaf65a21476bb097e83619fef33e68c30843d9d4c64736f6c63430008190033",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// BindingsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingsMetaData.Bin instead.
var BindingsBin = BindingsMetaData.Bin

// DeployBindings deploys a new Ethereum contract, binding an instance of Bindings to it.
func DeployBindings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bindings, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Bindings *BindingsCaller) GetBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Bindings *BindingsSession) GetBytes() ([]byte, error) {
	return _Bindings.Contract.GetBytes(&_Bindings.CallOpts)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Bindings *BindingsCallerSession) GetBytes() ([]byte, error) {
	return _Bindings.Contract.GetBytes(&_Bindings.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Bindings *BindingsCaller) GetBytes32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBytes32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Bindings *BindingsSession) GetBytes32() ([32]byte, error) {
	return _Bindings.Contract.GetBytes32(&_Bindings.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Bindings *BindingsCallerSession) GetBytes32() ([32]byte, error) {
	return _Bindings.Contract.GetBytes32(&_Bindings.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Bindings *BindingsCaller) GetUint256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getUint256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Bindings *BindingsSession) GetUint256() (*big.Int, error) {
	return _Bindings.Contract.GetUint256(&_Bindings.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetUint256() (*big.Int, error) {
	return _Bindings.Contract.GetUint256(&_Bindings.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Bindings *BindingsCaller) RequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Bindings *BindingsSession) RequestId() ([32]byte, error) {
	return _Bindings.Contract.RequestId(&_Bindings.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Bindings *BindingsCallerSession) RequestId() ([32]byte, error) {
	return _Bindings.Contract.RequestId(&_Bindings.CallOpts)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Bindings *BindingsTransactor) RequestedBytes(opts *bind.TransactOpts, _requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "requestedBytes", _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Bindings *BindingsSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedBytes(&_Bindings.TransactOpts, _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Bindings *BindingsTransactorSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedBytes(&_Bindings.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Bindings *BindingsTransactor) RequestedBytes32(opts *bind.TransactOpts, _requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "requestedBytes32", _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Bindings *BindingsSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedBytes32(&_Bindings.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Bindings *BindingsTransactorSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedBytes32(&_Bindings.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Bindings *BindingsTransactor) RequestedUint256(opts *bind.TransactOpts, _requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "requestedUint256", _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Bindings *BindingsSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedUint256(&_Bindings.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Bindings *BindingsTransactorSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RequestedUint256(&_Bindings.TransactOpts, _requestId, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Bindings *BindingsTransactor) SetBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setBytes", _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Bindings *BindingsSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SetBytes(&_Bindings.TransactOpts, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Bindings *BindingsTransactorSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SetBytes(&_Bindings.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Bindings *BindingsTransactor) SetBytes32(opts *bind.TransactOpts, _value [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setBytes32", _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Bindings *BindingsSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SetBytes32(&_Bindings.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Bindings *BindingsTransactorSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SetBytes32(&_Bindings.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Bindings *BindingsTransactor) SetUint256(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setUint256", _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Bindings *BindingsSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetUint256(&_Bindings.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Bindings *BindingsTransactorSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetUint256(&_Bindings.TransactOpts, _value)
}

// BindingsOutputIterator is returned from FilterOutput and is used to iterate over the raw logs and unpacked data for Output events raised by the Bindings contract.
type BindingsOutputIterator struct {
	Event *BindingsOutput // Event containing the contract specifics and raw log

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
func (it *BindingsOutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOutput)
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
		it.Event = new(BindingsOutput)
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
func (it *BindingsOutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOutput represents a Output event raised by the Bindings contract.
type BindingsOutput struct {
	B32  [32]byte
	U256 *big.Int
	B322 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOutput is a free log retrieval operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Bindings *BindingsFilterer) FilterOutput(opts *bind.FilterOpts) (*BindingsOutputIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return &BindingsOutputIterator{contract: _Bindings.contract, event: "Output", logs: logs, sub: sub}, nil
}

// WatchOutput is a free log subscription operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Bindings *BindingsFilterer) WatchOutput(opts *bind.WatchOpts, sink chan<- *BindingsOutput) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOutput)
				if err := _Bindings.contract.UnpackLog(event, "Output", log); err != nil {
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

// ParseOutput is a log parse operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Bindings *BindingsFilterer) ParseOutput(log types.Log) (*BindingsOutput, error) {
	event := new(BindingsOutput)
	if err := _Bindings.contract.UnpackLog(event, "Output", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSetBytesIterator is returned from FilterSetBytes and is used to iterate over the raw logs and unpacked data for SetBytes events raised by the Bindings contract.
type BindingsSetBytesIterator struct {
	Event *BindingsSetBytes // Event containing the contract specifics and raw log

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
func (it *BindingsSetBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSetBytes)
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
		it.Event = new(BindingsSetBytes)
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
func (it *BindingsSetBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSetBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSetBytes represents a SetBytes event raised by the Bindings contract.
type BindingsSetBytes struct {
	From  common.Address
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes is a free log retrieval operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Bindings *BindingsFilterer) FilterSetBytes(opts *bind.FilterOpts, from []common.Address) (*BindingsSetBytesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSetBytesIterator{contract: _Bindings.contract, event: "SetBytes", logs: logs, sub: sub}, nil
}

// WatchSetBytes is a free log subscription operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Bindings *BindingsFilterer) WatchSetBytes(opts *bind.WatchOpts, sink chan<- *BindingsSetBytes, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSetBytes)
				if err := _Bindings.contract.UnpackLog(event, "SetBytes", log); err != nil {
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

// ParseSetBytes is a log parse operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Bindings *BindingsFilterer) ParseSetBytes(log types.Log) (*BindingsSetBytes, error) {
	event := new(BindingsSetBytes)
	if err := _Bindings.contract.UnpackLog(event, "SetBytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSetBytes32Iterator is returned from FilterSetBytes32 and is used to iterate over the raw logs and unpacked data for SetBytes32 events raised by the Bindings contract.
type BindingsSetBytes32Iterator struct {
	Event *BindingsSetBytes32 // Event containing the contract specifics and raw log

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
func (it *BindingsSetBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSetBytes32)
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
		it.Event = new(BindingsSetBytes32)
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
func (it *BindingsSetBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSetBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSetBytes32 represents a SetBytes32 event raised by the Bindings contract.
type BindingsSetBytes32 struct {
	From  common.Address
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes32 is a free log retrieval operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Bindings *BindingsFilterer) FilterSetBytes32(opts *bind.FilterOpts, from []common.Address, value [][32]byte) (*BindingsSetBytes32Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSetBytes32Iterator{contract: _Bindings.contract, event: "SetBytes32", logs: logs, sub: sub}, nil
}

// WatchSetBytes32 is a free log subscription operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Bindings *BindingsFilterer) WatchSetBytes32(opts *bind.WatchOpts, sink chan<- *BindingsSetBytes32, from []common.Address, value [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSetBytes32)
				if err := _Bindings.contract.UnpackLog(event, "SetBytes32", log); err != nil {
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

// ParseSetBytes32 is a log parse operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Bindings *BindingsFilterer) ParseSetBytes32(log types.Log) (*BindingsSetBytes32, error) {
	event := new(BindingsSetBytes32)
	if err := _Bindings.contract.UnpackLog(event, "SetBytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSetUint256Iterator is returned from FilterSetUint256 and is used to iterate over the raw logs and unpacked data for SetUint256 events raised by the Bindings contract.
type BindingsSetUint256Iterator struct {
	Event *BindingsSetUint256 // Event containing the contract specifics and raw log

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
func (it *BindingsSetUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSetUint256)
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
		it.Event = new(BindingsSetUint256)
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
func (it *BindingsSetUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSetUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSetUint256 represents a SetUint256 event raised by the Bindings contract.
type BindingsSetUint256 struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetUint256 is a free log retrieval operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Bindings *BindingsFilterer) FilterSetUint256(opts *bind.FilterOpts, from []common.Address, value []*big.Int) (*BindingsSetUint256Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSetUint256Iterator{contract: _Bindings.contract, event: "SetUint256", logs: logs, sub: sub}, nil
}

// WatchSetUint256 is a free log subscription operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Bindings *BindingsFilterer) WatchSetUint256(opts *bind.WatchOpts, sink chan<- *BindingsSetUint256, from []common.Address, value []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSetUint256)
				if err := _Bindings.contract.UnpackLog(event, "SetUint256", log); err != nil {
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

// ParseSetUint256 is a log parse operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Bindings *BindingsFilterer) ParseSetUint256(log types.Log) (*BindingsSetUint256, error) {
	event := new(BindingsSetUint256)
	if err := _Bindings.contract.UnpackLog(event, "SetUint256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
