# Quorum Network Go Client

This repository contains a Go client for interacting with a Quorum network. The client allows you to deploy a contract and call methods on it.

## Usage

To use the client, follow these steps:

1. **Clone the repository**

   ```shell
   git clone https://github.com/ashp7/quorum

2. Build the binary using the provided Makefile:
make build

3. Run the binary with the desired command-line arguments. For example, to deploy a contract, use the following command:
   
    ```` 
    To deploy a contract
   
    ./contractclient --privateKey=<> --contractFile=../../contracts/NFTCollateral.sol --contractName NFTCollateral.sol:NFTCollateralLoan --nodeURL=http://localhost:22000 --verbose
    
   This would output the contract address
     
   To call a method on the contract

   ./contractclient --privateKey=<> --contractFile=../../contracts/NFTCollateral.sol  --contractName NFTCollateral.sol:NFTCollateralLoan  \
   --contractAddress 0x26FDd2B8ce6e8B73F7C4725bC6FCE9A2703686BA --nodeURL http://localhost:22001 --methodName acceptProposal --methodArgs 0


## Makefile

The provided Makefile simplifies the build process. Here are the available commands:

    make build: Builds the Go client binary.
    make clean: Removes the built binary and any temporary files.
    make test: Runs the tests for the Go client.

## Code Flow

The Go client follows the following code flow:

    Import required packages and libraries.
    Connect to the Quorum network using an Ethereum client.
    Load and compile the contract using solc.
    Deploy the contract to the network.
    Call methods on the deployed contract.

The code is organized into separate functions and structs to handle different tasks, making it easy to understand and modify as needed.

For more information on the available functions and structs, refer to the code comments and the GoDoc documentation.

## NFTCollateral Directory

The `NFTCollateral` folder contains the code for the Go client and Solidity contracts.

### Contracts

The Solidity contracts are located in the `contracts` directory. These contracts define the NFTToken and the NFTCollateralLoan contracts

### Client Code

The Go client code is located in the `pkg/contractclient` directory. This code provides an interface to interact with the Quorum network and perform various operations.

### Makefile
The repository includes a `Makefile` that provides convenient commands for building, testing, and deploying the Quorum network.

### Test Files
The test files in the repository have the suffix `_test.go` and are located in the `pkg/contractclient` directory. These are currently very basic and cover the usage of a mock ethereum client

### Mocks Directory
The `mocks` directory contains mock implementations used for testing purposes. 


