// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package btcrelay

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
	//_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BTCProofMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BTCProofMerkleProof struct {
	ProofHash [32]byte
	IsLeft    bool
}

// BtcrelayABI is the input ABI used to generate the binding from.
const BtcrelayABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isLeft\",\"type\":\"bool\"}],\"internalType\":\"structBTCProof.MerkleProof[]\",\"name\":\"merkleProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BtcrelayBin is the compiled bytecode used for deploying new contracts.
var BtcrelayBin = "0x608060405234801561001057600080fd5b50610583806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a724681114610030575b600080fd5b61004a6004803603810190610045919061032a565b610060565b6040516100579190610446565b60405180910390f35b60008082905060008090505b84518110156100eb5784818151811061008157fe5b602002602001015160200151156100ba576100b38582815181106100a157fe5b602002602001015160000151836100f9565b91506100de565b6100db828683815181106100ca57fe5b6020026020010151600001516100f9565b91505b808060010191505061006c565b508481149150509392505050565b60006101258383604051602001610111929190610403565b60405160208183030381529060405261012d565b905092915050565b600060028083604051610140919061042f565b602060405180830381855afa15801561015d573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101809190610301565b60405160200161019091906103e8565b6040516020818303038152906040526040516101ac919061042f565b602060405180830381855afa1580156101c9573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101ec9190610301565b9050919050565b600082601f83011261020457600080fd5b81356102176102128261048e565b610461565b9150818183526020840193506020810190508385604084028201111561023c57600080fd5b60005b8381101561026c578161025288826102b5565b84526020840193506040830192505060018101905061023f565b5050505092915050565b6000813590506102858161051f565b92915050565b60008135905061029a81610536565b92915050565b6000815190506102af81610536565b92915050565b6000604082840312156102c757600080fd5b6102d16040610461565b905060006102e18482850161028b565b60008301525060206102f584828501610276565b60208301525092915050565b60006020828403121561031357600080fd5b6000610321848285016102a0565b91505092915050565b60008060006060848603121561033f57600080fd5b600061034d8682870161028b565b935050602084013567ffffffffffffffff81111561036a57600080fd5b610376868287016101f3565b92505060406103878682870161028b565b9150509250925092565b61039a816104cc565b82525050565b6103b16103ac826104d8565b610515565b82525050565b60006103c2826104b6565b6103cc81856104c1565b93506103dc8185602086016104e2565b80840191505092915050565b60006103f482846103a0565b60208201915081905092915050565b600061040f82856103a0565b60208201915061041f82846103a0565b6020820191508190509392505050565b600061043b82846103b7565b915081905092915050565b600060208201905061045b6000830184610391565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561048457600080fd5b8060405250919050565b600067ffffffffffffffff8211156104a557600080fd5b602082029050602081019050919050565b600081519050919050565b600081905092915050565b60008115159050919050565b6000819050919050565b60005b838110156105005780820151818401526020810190506104e5565b8381111561050f576000848401525b50505050565b6000819050919050565b610528816104cc565b811461053357600080fd5b50565b61053f816104d8565b811461054a57600080fd5b5056fea2646970667358221220d36fbe1f5c06e9df96ee56afabf418d546d4f15e5021ca8a070a09a673020fc564736f6c63430006060033"

// DeployBtcrelay deploys a new Ethereum contract, binding an instance of Btcrelay to it.
func DeployBtcrelay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Btcrelay, error) {
	parsed, err := abi.JSON(strings.NewReader(BtcrelayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BtcrelayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Btcrelay{BtcrelayCaller: BtcrelayCaller{contract: contract}, BtcrelayTransactor: BtcrelayTransactor{contract: contract}, BtcrelayFilterer: BtcrelayFilterer{contract: contract}}, nil
}

// Btcrelay is an auto generated Go binding around an Ethereum contract.
type Btcrelay struct {
	BtcrelayCaller     // Read-only binding to the contract
	BtcrelayTransactor // Write-only binding to the contract
	BtcrelayFilterer   // Log filterer for contract events
}

// BtcrelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtcrelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcrelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtcrelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcrelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtcrelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcrelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtcrelaySession struct {
	Contract     *Btcrelay         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcrelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtcrelayCallerSession struct {
	Contract *BtcrelayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BtcrelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtcrelayTransactorSession struct {
	Contract     *BtcrelayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BtcrelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtcrelayRaw struct {
	Contract *Btcrelay // Generic contract binding to access the raw methods on
}

// BtcrelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtcrelayCallerRaw struct {
	Contract *BtcrelayCaller // Generic read-only contract binding to access the raw methods on
}

// BtcrelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtcrelayTransactorRaw struct {
	Contract *BtcrelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtcrelay creates a new instance of Btcrelay, bound to a specific deployed contract.
func NewBtcrelay(address common.Address, backend bind.ContractBackend) (*Btcrelay, error) {
	contract, err := bindBtcrelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Btcrelay{BtcrelayCaller: BtcrelayCaller{contract: contract}, BtcrelayTransactor: BtcrelayTransactor{contract: contract}, BtcrelayFilterer: BtcrelayFilterer{contract: contract}}, nil
}

// NewBtcrelayCaller creates a new read-only instance of Btcrelay, bound to a specific deployed contract.
func NewBtcrelayCaller(address common.Address, caller bind.ContractCaller) (*BtcrelayCaller, error) {
	contract, err := bindBtcrelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtcrelayCaller{contract: contract}, nil
}

// NewBtcrelayTransactor creates a new write-only instance of Btcrelay, bound to a specific deployed contract.
func NewBtcrelayTransactor(address common.Address, transactor bind.ContractTransactor) (*BtcrelayTransactor, error) {
	contract, err := bindBtcrelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtcrelayTransactor{contract: contract}, nil
}

// NewBtcrelayFilterer creates a new log filterer instance of Btcrelay, bound to a specific deployed contract.
func NewBtcrelayFilterer(address common.Address, filterer bind.ContractFilterer) (*BtcrelayFilterer, error) {
	contract, err := bindBtcrelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtcrelayFilterer{contract: contract}, nil
}

// bindBtcrelay binds a generic wrapper to an already deployed contract.
func bindBtcrelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BtcrelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcrelay *BtcrelayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Btcrelay.Contract.BtcrelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcrelay *BtcrelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcrelay.Contract.BtcrelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcrelay *BtcrelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcrelay.Contract.BtcrelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcrelay *BtcrelayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Btcrelay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcrelay *BtcrelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcrelay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcrelay *BtcrelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcrelay.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xa7246811.
//
// Solidity: function verify(bytes32 merkleRoot, []BTCProofMerkleProof merkleProofs, bytes32 txHash) pure returns(bool)
func (_Btcrelay *BtcrelayCaller) Verify(opts *bind.CallOpts, merkleRoot [32]byte, merkleProofs []BTCProofMerkleProof, txHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Btcrelay.contract.Call(opts, out, "verify", merkleRoot, merkleProofs, txHash)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0xa7246811.
//
// Solidity: function verify(bytes32 merkleRoot, []BTCProofMerkleProof merkleProofs, bytes32 txHash) pure returns(bool)
func (_Btcrelay *BtcrelaySession) Verify(merkleRoot [32]byte, merkleProofs []BTCProofMerkleProof, txHash [32]byte) (bool, error) {
	return _Btcrelay.Contract.Verify(&_Btcrelay.CallOpts, merkleRoot, merkleProofs, txHash)
}

// Verify is a free data retrieval call binding the contract method 0xa7246811.
//
// Solidity: function verify(bytes32 merkleRoot, []BTCProofMerkleProof merkleProofs, bytes32 txHash) pure returns(bool)
func (_Btcrelay *BtcrelayCallerSession) Verify(merkleRoot [32]byte, merkleProofs []BTCProofMerkleProof, txHash [32]byte) (bool, error) {
	return _Btcrelay.Contract.Verify(&_Btcrelay.CallOpts, merkleRoot, merkleProofs, txHash)
}
