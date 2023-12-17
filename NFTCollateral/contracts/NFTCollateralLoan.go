// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftCollateralLoan

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

// NFTCollateralLoanMetaData contains all meta data concerning the NFTCollateralLoan contract.
var NFTCollateralLoanMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"acceptProposal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"actOnLoanDateExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dueDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAccepted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPaidBack\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"repayLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"retractProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_loanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061164c806100206000396000f3fe6080604052600436106100705760003560e01c806360c5cc3a1161004e57806360c5cc3a1461011057806365c0322d1461012c578063ab7b1c8914610155578063ccdd27eb1461017157610070565b8063013cf08b146100755780632ab09d14146100bc5780633c32219f146100e7575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610e80565b61019a565b6040516100b39b9a99989796959493929190610f18565b60405180910390f35b3480156100c857600080fd5b506100d161026e565b6040516100de9190610fc3565b60405180910390f35b3480156100f357600080fd5b5061010e60048036038101906101099190610e80565b610274565b005b61012a60048036038101906101259190610e80565b610511565b005b34801561013857600080fd5b50610153600480360381019061014e9190610e80565b6106d1565b005b61016f600480360381019061016a9190610e80565b61082f565b005b34801561017d57600080fd5b506101986004803603810190610193919061100a565b610a59565b005b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154908060050154908060060154908060070154908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060080160149054906101000a900460ff16908060080160159054906101000a900460ff1690508b565b60005481565b60006001600083815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461031d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610314906110e2565b60405180910390fd5b8060080160149054906101000a900460ff161561036f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103669061114e565b60405180910390fd5b8060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd308360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600201546040518463ffffffff1660e01b81526004016103f89392919061116e565b600060405180830381600087803b15801561041257600080fd5b505af1158015610426573d6000803e3d6000fd5b5050505060016000838152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160009055600382016000905560048201600090556005820160009055600682016000905560078201600090556008820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556008820160146101000a81549060ff02191690556008820160156101000a81549060ff021916905550505050565b60006001600083815260200190815260200160002090508060030154341461056e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610565906111f1565b60405180910390fd5b600015158160080160149054906101000a900460ff161515146105c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bd9061114e565b60405180910390fd5b338160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018160080160146101000a81548160ff02191690831515021790555060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050509050806106c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ba9061125d565b60405180910390fd5b428260070181905550505050565b60006001600083815260200190815260200160002090508060080160159054906101000a900460ff161561073a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610731906112c9565b60405180910390fd5b806006015442111561074c575061082c565b60008160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166323b872dd308460080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600201546040518463ffffffff1660e01b81526004016107da9392919061116e565b600060405180830381600087803b1580156107f457600080fd5b505af1158015610808573d6000803e3d6000fd5b5050505060018260080160156101000a81548160ff02191690831515021790555050505b50565b6000600160008381526020019081526020016000209050600061085182610de2565b82600301546108609190611318565b9050803410156108a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089c90611398565b60405180910390fd5b8160080160149054906101000a900460ff1680156108d257508160080160159054906101000a900460ff16155b610911576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109089061142a565b60405180910390fd5b8160080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015801561097b573d6000803e3d6000fd5b508160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd308460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600201546040518463ffffffff1660e01b8152600401610a059392919061116e565b600060405180830381600087803b158015610a1f57600080fd5b505af1158015610a33573d6000803e3d6000fd5b5050505060018260080160156101000a81548160ff021916908315150217905550505050565b60008590503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636352211e876040518263ffffffff1660e01b8152600401610aae9190610fc3565b602060405180830381865afa158015610acb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aef919061145f565b73ffffffffffffffffffffffffffffffffffffffff1614610b45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3c906114d8565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b8152600401610b829392919061116e565b600060405180830381600087803b158015610b9c57600080fd5b505af1158015610bb0573d6000803e3d6000fd5b5050505060006040518061016001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020018581526020018481526020018442610c1e9190611318565b815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016000151581525090508060016000806000815480929190610c70906114f8565b91905055815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101208201518160080160146101000a81548160ff0219169083151502179055506101408201518160080160156101000a81548160ff02191690831515021790555090505050505050505050565b600080826007015442610df59190611540565b905060006301e1338090506000606482610e0f9190611574565b8386600401548760030154610e249190611574565b610e2e9190611574565b610e3891906115e5565b9050809350505050919050565b600080fd5b6000819050919050565b610e5d81610e4a565b8114610e6857600080fd5b50565b600081359050610e7a81610e54565b92915050565b600060208284031215610e9657610e95610e45565b5b6000610ea484828501610e6b565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ed882610ead565b9050919050565b610ee881610ecd565b82525050565b610ef781610e4a565b82525050565b60008115159050919050565b610f1281610efd565b82525050565b600061016082019050610f2e600083018e610edf565b610f3b602083018d610edf565b610f48604083018c610eee565b610f55606083018b610eee565b610f62608083018a610eee565b610f6f60a0830189610eee565b610f7c60c0830188610eee565b610f8960e0830187610eee565b610f97610100830186610edf565b610fa5610120830185610f09565b610fb3610140830184610f09565b9c9b505050505050505050505050565b6000602082019050610fd86000830184610eee565b92915050565b610fe781610ecd565b8114610ff257600080fd5b50565b60008135905061100481610fde565b92915050565b600080600080600060a0868803121561102657611025610e45565b5b600061103488828901610ff5565b955050602061104588828901610e6b565b945050604061105688828901610e6b565b935050606061106788828901610e6b565b925050608061107888828901610e6b565b9150509295509295909350565b600082825260208201905092915050565b7f4e6f742074686520626f72726f77657200000000000000000000000000000000600082015250565b60006110cc601083611085565b91506110d782611096565b602082019050919050565b600060208201905081810360008301526110fb816110bf565b9050919050565b7f50726f706f73616c20616c726561647920616363657074656400000000000000600082015250565b6000611138601983611085565b915061114382611102565b602082019050919050565b600060208201905081810360008301526111678161112b565b9050919050565b60006060820190506111836000830186610edf565b6111906020830185610edf565b61119d6040830184610eee565b949350505050565b7f496e636f7272656374206c6f616e20616d6f756e740000000000000000000000600082015250565b60006111db601583611085565b91506111e6826111a5565b602082019050919050565b6000602082019050818103600083015261120a816111ce565b9050919050565b7f4661696c656420746f206c656e64000000000000000000000000000000000000600082015250565b6000611247600e83611085565b915061125282611211565b602082019050919050565b600060208201905081810360008301526112768161123a565b9050919050565b7f4c6f616e20697320616c72656164792070616964206261636b00000000000000600082015250565b60006112b3601983611085565b91506112be8261127d565b602082019050919050565b600060208201905081810360008301526112e2816112a6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061132382610e4a565b915061132e83610e4a565b9250828201905080821115611346576113456112e9565b5b92915050565b7f4e6f7420656e6f7567682066756e647300000000000000000000000000000000600082015250565b6000611382601083611085565b915061138d8261134c565b602082019050919050565b600060208201905081810360008301526113b181611375565b9050919050565b7f4c6f616e206e6f74206163636570746564206f7220616c72656164792070616960008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000611414602183611085565b915061141f826113b8565b604082019050919050565b6000602082019050818103600083015261144381611407565b9050919050565b60008151905061145981610fde565b92915050565b60006020828403121561147557611474610e45565b5b60006114838482850161144a565b91505092915050565b7f4e6f7420746865204e4654206f776e6572000000000000000000000000000000600082015250565b60006114c2601183611085565b91506114cd8261148c565b602082019050919050565b600060208201905081810360008301526114f1816114b5565b9050919050565b600061150382610e4a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611535576115346112e9565b5b600182019050919050565b600061154b82610e4a565b915061155683610e4a565b925082820390508181111561156e5761156d6112e9565b5b92915050565b600061157f82610e4a565b915061158a83610e4a565b925082820261159881610e4a565b915082820484148315176115af576115ae6112e9565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006115f082610e4a565b91506115fb83610e4a565b92508261160b5761160a6115b6565b5b82820490509291505056fea26469706673582212200808a9fc0d3b8701ed8eb9e23eb64245f67a3c68796e3d59cced394835ab9eac64736f6c63430008170033",
}

// NFTCollateralLoanABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTCollateralLoanMetaData.ABI instead.
var NFTCollateralLoanABI = NFTCollateralLoanMetaData.ABI

// NFTCollateralLoanBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NFTCollateralLoanMetaData.Bin instead.
var NFTCollateralLoanBin = NFTCollateralLoanMetaData.Bin

// DeployNFTCollateralLoan deploys a new Ethereum contract, binding an instance of NFTCollateralLoan to it.
func DeployNFTCollateralLoan(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NFTCollateralLoan, error) {
	parsed, err := NFTCollateralLoanMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NFTCollateralLoanBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NFTCollateralLoan{NFTCollateralLoanCaller: NFTCollateralLoanCaller{contract: contract}, NFTCollateralLoanTransactor: NFTCollateralLoanTransactor{contract: contract}, NFTCollateralLoanFilterer: NFTCollateralLoanFilterer{contract: contract}}, nil
}

// NFTCollateralLoan is an auto generated Go binding around an Ethereum contract.
type NFTCollateralLoan struct {
	NFTCollateralLoanCaller     // Read-only binding to the contract
	NFTCollateralLoanTransactor // Write-only binding to the contract
	NFTCollateralLoanFilterer   // Log filterer for contract events
}

// NFTCollateralLoanCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCollateralLoanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollateralLoanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTCollateralLoanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollateralLoanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTCollateralLoanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollateralLoanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTCollateralLoanSession struct {
	Contract     *NFTCollateralLoan // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NFTCollateralLoanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCollateralLoanCallerSession struct {
	Contract *NFTCollateralLoanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NFTCollateralLoanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTCollateralLoanTransactorSession struct {
	Contract     *NFTCollateralLoanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NFTCollateralLoanRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTCollateralLoanRaw struct {
	Contract *NFTCollateralLoan // Generic contract binding to access the raw methods on
}

// NFTCollateralLoanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCollateralLoanCallerRaw struct {
	Contract *NFTCollateralLoanCaller // Generic read-only contract binding to access the raw methods on
}

// NFTCollateralLoanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTCollateralLoanTransactorRaw struct {
	Contract *NFTCollateralLoanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTCollateralLoan creates a new instance of NFTCollateralLoan, bound to a specific deployed contract.
func NewNFTCollateralLoan(address common.Address, backend bind.ContractBackend) (*NFTCollateralLoan, error) {
	contract, err := bindNFTCollateralLoan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTCollateralLoan{NFTCollateralLoanCaller: NFTCollateralLoanCaller{contract: contract}, NFTCollateralLoanTransactor: NFTCollateralLoanTransactor{contract: contract}, NFTCollateralLoanFilterer: NFTCollateralLoanFilterer{contract: contract}}, nil
}

// NewNFTCollateralLoanCaller creates a new read-only instance of NFTCollateralLoan, bound to a specific deployed contract.
func NewNFTCollateralLoanCaller(address common.Address, caller bind.ContractCaller) (*NFTCollateralLoanCaller, error) {
	contract, err := bindNFTCollateralLoan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCollateralLoanCaller{contract: contract}, nil
}

// NewNFTCollateralLoanTransactor creates a new write-only instance of NFTCollateralLoan, bound to a specific deployed contract.
func NewNFTCollateralLoanTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTCollateralLoanTransactor, error) {
	contract, err := bindNFTCollateralLoan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCollateralLoanTransactor{contract: contract}, nil
}

// NewNFTCollateralLoanFilterer creates a new log filterer instance of NFTCollateralLoan, bound to a specific deployed contract.
func NewNFTCollateralLoanFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTCollateralLoanFilterer, error) {
	contract, err := bindNFTCollateralLoan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTCollateralLoanFilterer{contract: contract}, nil
}

// bindNFTCollateralLoan binds a generic wrapper to an already deployed contract.
func bindNFTCollateralLoan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTCollateralLoanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTCollateralLoan *NFTCollateralLoanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTCollateralLoan.Contract.NFTCollateralLoanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTCollateralLoan *NFTCollateralLoanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.NFTCollateralLoanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTCollateralLoan *NFTCollateralLoanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.NFTCollateralLoanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTCollateralLoan *NFTCollateralLoanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTCollateralLoan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTCollateralLoan *NFTCollateralLoanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTCollateralLoan *NFTCollateralLoanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.contract.Transact(opts, method, params...)
}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_NFTCollateralLoan *NFTCollateralLoanCaller) NextProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollateralLoan.contract.Call(opts, &out, "nextProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_NFTCollateralLoan *NFTCollateralLoanSession) NextProposalId() (*big.Int, error) {
	return _NFTCollateralLoan.Contract.NextProposalId(&_NFTCollateralLoan.CallOpts)
}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_NFTCollateralLoan *NFTCollateralLoanCallerSession) NextProposalId() (*big.Int, error) {
	return _NFTCollateralLoan.Contract.NextProposalId(&_NFTCollateralLoan.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address borrower, address nftContractAddress, uint256 nftTokenId, uint256 loanAmount, uint256 interestRate, uint256 duration, uint256 dueDate, uint256 startTime, address lender, bool isAccepted, bool isPaidBack)
func (_NFTCollateralLoan *NFTCollateralLoanCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Borrower           common.Address
	NftContractAddress common.Address
	NftTokenId         *big.Int
	LoanAmount         *big.Int
	InterestRate       *big.Int
	Duration           *big.Int
	DueDate            *big.Int
	StartTime          *big.Int
	Lender             common.Address
	IsAccepted         bool
	IsPaidBack         bool
}, error) {
	var out []interface{}
	err := _NFTCollateralLoan.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Borrower           common.Address
		NftContractAddress common.Address
		NftTokenId         *big.Int
		LoanAmount         *big.Int
		InterestRate       *big.Int
		Duration           *big.Int
		DueDate            *big.Int
		StartTime          *big.Int
		Lender             common.Address
		IsAccepted         bool
		IsPaidBack         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Borrower = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftTokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LoanAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.InterestRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DueDate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Lender = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.IsAccepted = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.IsPaidBack = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address borrower, address nftContractAddress, uint256 nftTokenId, uint256 loanAmount, uint256 interestRate, uint256 duration, uint256 dueDate, uint256 startTime, address lender, bool isAccepted, bool isPaidBack)
func (_NFTCollateralLoan *NFTCollateralLoanSession) Proposals(arg0 *big.Int) (struct {
	Borrower           common.Address
	NftContractAddress common.Address
	NftTokenId         *big.Int
	LoanAmount         *big.Int
	InterestRate       *big.Int
	Duration           *big.Int
	DueDate            *big.Int
	StartTime          *big.Int
	Lender             common.Address
	IsAccepted         bool
	IsPaidBack         bool
}, error) {
	return _NFTCollateralLoan.Contract.Proposals(&_NFTCollateralLoan.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address borrower, address nftContractAddress, uint256 nftTokenId, uint256 loanAmount, uint256 interestRate, uint256 duration, uint256 dueDate, uint256 startTime, address lender, bool isAccepted, bool isPaidBack)
func (_NFTCollateralLoan *NFTCollateralLoanCallerSession) Proposals(arg0 *big.Int) (struct {
	Borrower           common.Address
	NftContractAddress common.Address
	NftTokenId         *big.Int
	LoanAmount         *big.Int
	InterestRate       *big.Int
	Duration           *big.Int
	DueDate            *big.Int
	StartTime          *big.Int
	Lender             common.Address
	IsAccepted         bool
	IsPaidBack         bool
}, error) {
	return _NFTCollateralLoan.Contract.Proposals(&_NFTCollateralLoan.CallOpts, arg0)
}

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactor) AcceptProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.contract.Transact(opts, "acceptProposal", _proposalId)
}

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanSession) AcceptProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.AcceptProposal(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactorSession) AcceptProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.AcceptProposal(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// ActOnLoanDateExpiry is a paid mutator transaction binding the contract method 0x65c0322d.
//
// Solidity: function actOnLoanDateExpiry(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactor) ActOnLoanDateExpiry(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.contract.Transact(opts, "actOnLoanDateExpiry", _proposalId)
}

// ActOnLoanDateExpiry is a paid mutator transaction binding the contract method 0x65c0322d.
//
// Solidity: function actOnLoanDateExpiry(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanSession) ActOnLoanDateExpiry(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.ActOnLoanDateExpiry(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// ActOnLoanDateExpiry is a paid mutator transaction binding the contract method 0x65c0322d.
//
// Solidity: function actOnLoanDateExpiry(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactorSession) ActOnLoanDateExpiry(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.ActOnLoanDateExpiry(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// RepayLoan is a paid mutator transaction binding the contract method 0xab7b1c89.
//
// Solidity: function repayLoan(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactor) RepayLoan(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.contract.Transact(opts, "repayLoan", _proposalId)
}

// RepayLoan is a paid mutator transaction binding the contract method 0xab7b1c89.
//
// Solidity: function repayLoan(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanSession) RepayLoan(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.RepayLoan(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// RepayLoan is a paid mutator transaction binding the contract method 0xab7b1c89.
//
// Solidity: function repayLoan(uint256 _proposalId) payable returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactorSession) RepayLoan(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.RepayLoan(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// RetractProposal is a paid mutator transaction binding the contract method 0x3c32219f.
//
// Solidity: function retractProposal(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactor) RetractProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.contract.Transact(opts, "retractProposal", _proposalId)
}

// RetractProposal is a paid mutator transaction binding the contract method 0x3c32219f.
//
// Solidity: function retractProposal(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanSession) RetractProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.RetractProposal(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// RetractProposal is a paid mutator transaction binding the contract method 0x3c32219f.
//
// Solidity: function retractProposal(uint256 _proposalId) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactorSession) RetractProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.RetractProposal(&_NFTCollateralLoan.TransactOpts, _proposalId)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xccdd27eb.
//
// Solidity: function submitProposal(address _nftContractAddress, uint256 _nftTokenId, uint256 _loanAmount, uint256 _interestRate, uint256 _duration) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactor) SubmitProposal(opts *bind.TransactOpts, _nftContractAddress common.Address, _nftTokenId *big.Int, _loanAmount *big.Int, _interestRate *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.contract.Transact(opts, "submitProposal", _nftContractAddress, _nftTokenId, _loanAmount, _interestRate, _duration)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xccdd27eb.
//
// Solidity: function submitProposal(address _nftContractAddress, uint256 _nftTokenId, uint256 _loanAmount, uint256 _interestRate, uint256 _duration) returns()
func (_NFTCollateralLoan *NFTCollateralLoanSession) SubmitProposal(_nftContractAddress common.Address, _nftTokenId *big.Int, _loanAmount *big.Int, _interestRate *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.SubmitProposal(&_NFTCollateralLoan.TransactOpts, _nftContractAddress, _nftTokenId, _loanAmount, _interestRate, _duration)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xccdd27eb.
//
// Solidity: function submitProposal(address _nftContractAddress, uint256 _nftTokenId, uint256 _loanAmount, uint256 _interestRate, uint256 _duration) returns()
func (_NFTCollateralLoan *NFTCollateralLoanTransactorSession) SubmitProposal(_nftContractAddress common.Address, _nftTokenId *big.Int, _loanAmount *big.Int, _interestRate *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _NFTCollateralLoan.Contract.SubmitProposal(&_NFTCollateralLoan.TransactOpts, _nftContractAddress, _nftTokenId, _loanAmount, _interestRate, _duration)
}
