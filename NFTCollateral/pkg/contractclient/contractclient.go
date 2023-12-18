package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const (
	gasLimit = 700000
	gasPrice = 0
	solcVer  = "0.8.23"
	solcPath = "/usr/local/bin/solc"
)

var (
	verboseOutput *bool
)

// EthClient is an interface that wraps the necessary methods from ethclient.Client.
type EthClient interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, args bind.PrivateTxArgs) error
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
}

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
	fmt.Println(" -verbose              Verbose output")
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

	// If the contract Address is present, it means we are trying to call a method on a contract
	contractAddressString := flag.String("contractAddress", "", "Address of the deployed contract (optional)")
	methodName := flag.String("methodName", "", "Method name to call on the contract (optional)")
	eventName := flag.String("eventName", "", "Event name to view logs(optional)")
	methodArgs := flag.String("methodArgs", "", "Method arguments, comma-separated (optional)")

	// call a view methid
	isViewMethod := flag.Bool("isViewMethod", false, "Set to true if calling a view method")

	help := flag.Bool("help", false, "Display help message")
	verboseOutput = flag.Bool("verbose", false, "Display help message")

	flag.Parse()

	if *help {
		displayHelp()
		return
	}



	// Validate required arguments
	if *nodeURL == "" || *privateKey == "" || *contractFile == "" || *contractName == "" {
		log.Fatal("nodeURL, privateKey, contractFile, and contractName are required")
	}

	// Compile the contract using solc
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

		// connect to the node
		client, err := ethclient.Dial(*nodeURL)
		if err != nil {
			log.Fatal(err)
		}

		nonce, err := GetNonce(client, *privateKey)
		contractAddress := common.HexToAddress(*contractAddressString)

		if *isViewMethod {
			err := CallViewMethod(client, contract, common.HexToAddress(*contractAddressString), *methodName, *methodArgs)
			if err != nil {
				log.Fatalf("Failed to call view method: %v", err)
			}
		} else {

			txn, err := CallContractMethod(client, *privateKey, contract, *methodName, *methodArgs, nonce, contractAddress, *eventName)
			if err != nil {
				log.Fatalf("Failed to call contract method: %v", err)
			}

			log.Info("transaction created ", " txn From: ", txn.From(), " txn To: ", txn.To(), " type: ", txn.Type())
		}
	}
}

func DeployContract(contract *compiler.Contract, nodeURL string, privateKey string, contractFile string, contractName string) (common.Address, interface{}) {

	if *verboseOutput {
		log.Infof("abi definition is %v", contract.Info.AbiDefinition)
	}

	abiBytes, err := json.Marshal(contract.Info.AbiDefinition)
	if err != nil {
		log.Fatalf("Failed to marshal ABI: %v", err)
	}

	// Filter out custom error types because there was an issue with ERC721 related error while converting to json
	filteredABIBytes, err := filterABIToRemoveCustomErrors(abiBytes)
	if err != nil {
		log.Fatalf("Failed to filter ABI: %v", err)
	}

	parsedABI, err := abi.JSON(bytes.NewReader(filteredABIBytes))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	if *verboseOutput {
		log.Info("methods are ", parsedABI.Methods)
		log.Info("events are ", parsedABI.Events)
	}

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

	// connect to the node
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// deploy the contract
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
			errorStr := "Transaction mining exceeded 30 seconds"
			log.Println(errorStr)

			return address, errors.New(errorStr)
		}
	}

	fmt.Printf("Transaction hash: %s\n", receipt.TxHash.Hex())
	return address, nil
}

// CompileContract uses solc to compile the Solidity source and
func CompileContract(solidityFile, evmVersion, contractName string) (*compiler.Contract, error) {

	solcArgs := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata",
		"--optimize",
		"--evm-version", evmVersion,
		"--allow-paths", ".",
		"--include-path", "../../node_modules", // Needed this for ERC721
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
func CallContractMethod(client EthClient,
	privateKeyHex string,
	contract *compiler.Contract,
	methodName string,
	args string,
	nonce uint64,
	to common.Address,
	eventName string,
) (*types.Transaction, error) {

	ctx := context.Background()

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

	if *verboseOutput {
		log.Info("methods are", parsedABI.Methods)
		log.Info("events are ", parsedABI.Events)
	}

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

	// Decode the private key from hex string
	privateKeyBytes, err := hex.DecodeString(privateKeyHex[2:]) // Assuming privateKey is a hex string with "0x" prefix
	if err != nil {
		log.Fatal("Failed to decode hex string of private key: ", err)
	}

	// Parse the ECDSA private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %v", err)
	}

	// Pack the arguments for the contract call
	packedArgs, err := parsedABI.Pack(methodName, typedArgs...)
	if err != nil {
		return nil, fmt.Errorf("error packing arguments for method %s: %v", methodName, err)
	}

	amount := big.NewInt(0)
	// Create the transaction
	tx := types.NewTransaction(nonce, to, amount, gasLimit, big.NewInt(gasPrice), packedArgs)

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
	err = client.SendTransaction(ctx, signedTx, privateArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		return signedTx, err
	}

	log.Info("Transaction receipt ", " Block ", receipt.BlockNumber, " Status ", receipt.Status)
	if eventName != "" {
		if event, ok := parsedABI.Events[eventName]; ok {
			for _, vLog := range receipt.Logs {
				fmt.Printf("receipt log %v", vLog)
				if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == event.ID.Hex() {
					processEvent(event, vLog.Data)
				}
			}
		}
	}

	return signedTx, nil
}

func CallViewMethod(client *ethclient.Client,
	contract *compiler.Contract,
	contractAddress common.Address,
	methodName string,
	args string,
) error {

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

	if *verboseOutput {
		log.Info("methods are", parsedABI.Methods)
		log.Info("events are ", parsedABI.Events)
	}

	method, exists := parsedABI.Methods[methodName]
	if !exists {
		return fmt.Errorf("method '%s' not found in ABI", methodName)
	}

	// Split the string into individual arguments
	strArgs := strings.Split(args, ",")

	if len(method.Inputs) != len(strArgs) {
		return fmt.Errorf("method '%s' requires %d args but got %d", methodName, len(method.Inputs), len(strArgs))
	}

	typedArgs, err := GenerateTypedArgs(parsedABI, methodName, strArgs)
	if err != nil {
		log.Fatalf("Failed to generate typed arguments: %v", err)
	}

	// Pack the arguments for the contract call
	packedArgs, err := parsedABI.Pack(methodName, typedArgs...)
	if err != nil {
		return fmt.Errorf("error packing arguments for method %s: %v", methodName, err)
	}

	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: append(method.ID, packedArgs...),
	}

	// Perform the call
	rawResult, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return fmt.Errorf("failed to call contract method: %v", err)
	}

	// Unpack the raw data into the expected outputs
	results := make(map[string]interface{})
	err = method.Outputs.UnpackIntoMap(results, rawResult)
	if err != nil {
		return fmt.Errorf("failed to unpack returned values: %v", err)
	}

	for key, value := range results {
		fmt.Printf("Key: %s Value: %v", key, value)
	}

	return nil
}

func processEvent(event abi.Event, data []byte) {

	results, err := event.Inputs.Unpack(data)
	if err != nil {
		log.Errorf("Failed to parse event logs Error: %s", err.Error())
		return
	}

	eventParams := make(map[string]interface{})
	for i, input := range event.Inputs {
		fmt.Printf("Parameter %s: Value %v\n", eventParams[input.Name], results[i])
	}

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

func GetNonce(client EthClient, privateKeyHex string) (uint64, error) {

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

// filterABIToRemoveCustomErrors removes custom errors from the input abiJSON and returns the filtered byteArray
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
