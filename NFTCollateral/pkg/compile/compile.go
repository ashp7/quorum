package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

const (
	nodeURL    = "http://localhost:22000" // Replace with your Quorum node URL
	gasLimit = 700000
	gasPrice = 0
)

var (
 privateKey	string  = "0x692a322c4eaacafc4ea25c8513270c9d62009f88a7c2866d51b820d5e4bd9696"
)

func displayHelp() {
	fmt.Println("Usage: myContractBinary [options]")
	fmt.Println("Options:")
	fmt.Println("  -nodeURL            URL of the Quorum node. Required.")
	fmt.Println("  -privateKey         Private key for the transaction. Required.")
	fmt.Println("  -contractFile       Path to the Solidity contract file. Required.")
	fmt.Println("  -contractName       Name of the contract. Required.")
	fmt.Println("  -contractAddress    Address of the deployed contract. Optional.")
	fmt.Println("  -methodName         Method name to call on the contract. Optional.")
	fmt.Println("  -methodArgs         Method arguments, comma-separated. Optional.")
	fmt.Println("  -help               Display this help message.")
	fmt.Println("\nExample: ./myContractBinary -nodeURL http://localhost:22000 -privateKey 0x123... -contractFile /path/to/MyContract.sol -contractName MyContract")
	fmt.Println("Example with method call: ./myContractBinary -nodeURL http://localhost:22000 -privateKey 0x123... -contractAddress 0xContractAddr -methodName myMethod -methodArgs arg1,arg2,arg3")
}




func main() {
	// Command-line arguments
	nodeURL := flag.String("nodeURL", "http://localhost:22000", "URL of the Quorum node")
	privateKey := flag.String("privateKey", "", "Private key for the transaction")
	contractFile := flag.String("contractFile", "", "Path to the Solidity contract file")
	contractName := flag.String("contractName", "", "Name of the contract")
	contractAddressString := flag.String("contractAddress", "", "Address of the deployed contract (optional)")
	methodName := flag.String("methodName", "", "Method name to call on the contract (optional)")
	methodArgs := flag.String("methodArgs", "", "Method arguments, comma-separated (optional)")
	help := flag.Bool("help", false, "Display help message")

	flag.Parse()

	if *help {
		displayHelp()
		return
	}


	// Validate required arguments
	if *nodeURL == "" || *privateKey == "" || *contractFile == "" || *contractName == "" {
		log.Fatal("nodeURL, privateKey, contractFile, and contractName are required")
	}

	contract, err := CompileContract(*contractFile, "istanbul", *contractName)
	if err != nil {
		log.Fatalf("Failed to compile contract: %v", err)
	}


	// Perform actions based on the presence of contractAddress
	if *contractAddressString == "" {
		// Deploy the contract
		address, err := DeployContract(contract, *nodeURL, *privateKey, *contractFile, *contractName)
		if err != nil {
			log.Fatalf("Failed to deploy contract: %v", err)
		}
		fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	} else {
		// Call a method on the contract
		if *methodName == "" {
			log.Fatal("methodName is required when contractAddress is provided")
		}

		nonce, err := GetNonce(*nodeURL, *privateKey)// Replace with the nonce of the account
		amount := big.NewInt(0)                        // Amount of Ether to send, set to 0 if not sending Ether

		contractAddress := common.HexToAddress(*contractAddressString)

		_, err = CallContractMethod(*nodeURL, *privateKey, contract, *methodName, *methodArgs, nonce,  contractAddress,amount, gasLimit, big.NewInt(gasPrice))
		if err != nil {
			log.Fatalf("Failed to call contract method: %v", err)
		}
	}
}

func DeployContract(contract *compiler.Contract, nodeURL string, privateKey string, contractFile string, contractName string) (common.Address, interface{}) {


	log.Info("abi definition is %v", contract.Info.AbiDefinition)

	abiBytes, err := json.Marshal(contract.Info.AbiDefinition)
	if err != nil {
		log.Fatalf("Failed to marshal ABI: %v", err)
	}


	// Filter out custom error types
	filteredABIBytes, err := filterABIToRemoveCustomErrors(abiBytes)
	if err != nil {
		log.Fatalf("Failed to filter ABI: %v", err)
	}


	parsedABI, err := abi.JSON(bytes.NewReader(filteredABIBytes))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	log.Info("methods are", parsedABI.Methods)

	// Decode the private key from hex string
	privateKeyBytes, err := hex.DecodeString(privateKey[2:]) // Assuming privateKey is a hex string with "0x" prefix
	if err != nil {
		log.Fatal("Failed to decode hex string of private key: ", err)
	}

	// Parse the ECDSA private key
	privKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse ECDSA private key: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(1337))
	if err != nil {
		log.Fatal(err)
	}
	// Set the gas price and gas limit
	auth.GasPrice = big.NewInt(gasPrice)
	auth.GasLimit = gasLimit

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, _, err := bind.DeployContract(auth, parsedABI, common.FromHex(contract.Code), client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Printf("Contract deployed at address: %s\n", address.Hex())

	// Wait for the transaction to be mined

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Transaction mining exceeded 30 seconds, attempting to replace transaction")

			// Replace the transaction
			nonce := tx.Nonce()
			var price int64 = 10000
			newGasPrice := new(big.Int).Mul(big.NewInt(price), big.NewInt(2)) // Increase the gas price

			// Create a new transaction with the same nonce and higher gas price
			newTx := types.NewTransaction(nonce, auth.From, big.NewInt(0), tx.Gas(), newGasPrice, nil)

			signer := types.NewEIP155Signer(big.NewInt(1337))
			// Sign the new transaction
			signedTx, err := types.SignTx(newTx, signer, privKey)
			if err != nil {
				log.Fatalf("Failed to sign new transaction: %v", err)
			}

			opts := bind.PrivateTxArgs{}
			err = client.SendTransaction(context.Background(), signedTx, opts)
			if err != nil {
				log.Fatalf("Failed to send new transaction: %v", err)
			}

			log.Printf("New transaction sent to replace original: %s\n", signedTx.Hash().Hex())
		} else {
			log.Fatalf("Failed to wait for transaction mining: %v", err)
		}
	}

	fmt.Printf("Transaction hash: %s\n", receipt.TxHash.Hex())
	return address, nil
}


const (
solcVer = "0.8.23"
solcPath = "/usr/local/bin/solc"
)

// CompileContract uses solc to compile the Solidity source and
func CompileContract(solidityFile, evmVersion, contractName string) (*compiler.Contract, error) {
	//var c CompiledSolidity

	solcArgs := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata",
		"--optimize",
		"--evm-version", evmVersion,
		"--allow-paths", ".",
		"--include-path", "../../node_modules",
		"--base-path", "../../contracts",
		solidityFile,
	}
	solOptionsString := strings.Join(append([]string{solcPath}, solcArgs...), " ")
	log.Infof("Compiling: %s", solOptionsString)
	cmd := exec.Command(solcPath, solcArgs...)

	// Compile the solidity
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to compile [%s]: %s", err, stderr.String())
	}

	compiled, err := compiler.ParseCombinedJSON(stdout.Bytes(), "", solcVer, solcVer, solOptionsString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solc output: %s", err)
	}

	// Check we only have one conract and grab the code/info
	var contract *compiler.Contract
	contractNames := reflect.ValueOf(compiled).MapKeys()
	if contractName != "" {
		if _, ok := compiled[contractName]; !ok {
			return nil, fmt.Errorf("contract %s not found in Solidity file: %s", contractName, contractNames)
		}
		contract = compiled[contractName]
	} else if len(contractNames) != 1 {
		return nil, fmt.Errorf("more than one contract in Solidity file, please set one to call: %s", contractNames)
	} else {
		contractName = contractNames[0].String()
		contract = compiled[contractName]
	}

	return contract, nil

}

// CallContractMethod creates, signs, and sends a transaction to call a method on a smart contract
func CallContractMethod(nodeURL, privateKeyHex string, contract *compiler.Contract, methodName string, args string,
	nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int) (*types.Transaction, error) {

	// Connect to the Ethereum client
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}


	abiBytes, err := json.Marshal(contract.Info.AbiDefinition)
	if err != nil {
		log.Fatalf("Failed to marshal ABI: %v", err)
	}

	// Filter out custom error types
	filteredABIBytes, err := filterABIToRemoveCustomErrors(abiBytes)
	if err != nil {
		log.Fatalf("Failed to filter ABI: %v", err)
	}


	parsedABI, err := abi.JSON(bytes.NewReader(filteredABIBytes))
	if err != nil {
		log.Errorf("Failed to parse ABI: %v", err)
	}

	log.Info("methods are", parsedABI.Methods)

	method, exists := parsedABI.Methods[methodName]
	if !exists {
		return nil, fmt.Errorf("method '%s' not found in ABI", methodName)
	}

	// Split the string into individual arguments
	strArgs := strings.Split(args, ",")

	if len(method.Inputs) != len(strArgs) {
		return nil, fmt.Errorf("method '%s' requires %d args but got %d", methodName, len(method.Inputs), len(strArgs))
	}


	typedArgs, err := GenerateTypedArgs(parsedABI, methodName, strArgs)
	if err != nil {
		log.Fatalf("Failed to generate typed arguments: %v", err)
	}


	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	// Decode the private key from hex string
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %v", err)
	}

	// Pack the arguments for the contract call
	packedArgs, err := parsedABI.Pack(methodName, typedArgs...)
	if err != nil {
		return nil, fmt.Errorf("error packing arguments for method %s: %v", methodName, err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, packedArgs)

	// Sign the transaction with the private key
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	privateArgs := bind.PrivateTxArgs{}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx, privateArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx, nil
}

// GenerateTypedArgs converts string arguments into their corresponding types based on the ABI and method name
func GenerateTypedArgs(contractAbi abi.ABI, methodName string, strArgs []string) ([]interface{}, error) {
	method, exists := contractAbi.Methods[methodName]
	if !exists {
		return nil, fmt.Errorf("method '%s' not found in ABI", methodName)
	}

	if len(method.Inputs) != len(strArgs) {
		return nil, fmt.Errorf("method '%s' requires %d args but got %d", methodName, len(method.Inputs), len(strArgs))
	}

	var typedArgs []interface{}
	for i, input := range method.Inputs {
		arg := strArgs[i]
		switch input.Type.T {
		case abi.StringTy: // string
			typedArgs = append(typedArgs, arg)

		case abi.IntTy, abi.UintTy: // int, uint
			argInt, ok := new(big.Int).SetString(arg, 10)
			if !ok {
				return nil, fmt.Errorf("failed to convert argument %d (%s) to big.Int for type %s", i, arg, input.Type)
			}
			typedArgs = append(typedArgs, argInt)

		case abi.BoolTy: // bool
			argBool, err := strconv.ParseBool(arg)
			if err != nil {
				return nil, fmt.Errorf("failed to convert argument %d (%s) to bool: %v", i, arg, err)
			}
			typedArgs = append(typedArgs, argBool)

		case abi.AddressTy: // address
			if !common.IsHexAddress(arg) {
				return nil, fmt.Errorf("argument %d (%s) is not a valid hex address", i, arg)
			}
			typedArgs = append(typedArgs, common.HexToAddress(arg))

		default:
			// Add support for additional types as needed
			return nil, fmt.Errorf("unsupported argument type %s for argument %d", input.Type, i)
		}
	}

	return typedArgs, nil
}

func GetNonce(nodeURL, privateKeyHex string) (uint64, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return 0, err
	}
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return 0, err
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}


func filterABIToRemoveCustomErrors(abiJSON []byte) ([]byte, error) {
	var abiElements []map[string]interface{}
	err := json.Unmarshal(abiJSON, &abiElements)
	if err != nil {
		return nil, err
	}

	filteredABIElements := make([]map[string]interface{}, 0)
	for _, element := range abiElements {
		if typ, ok := element["type"]; ok && typ != "error" {
			filteredABIElements = append(filteredABIElements, element)
		}
	}

	return json.Marshal(filteredABIElements)
}

