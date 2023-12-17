package main

import (
	"context"
	"github.com/ethereum/go-ethereum/NFTCollateral/pkg/contractclient/mocks"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

const (
	nodeURL    = "http://localhost:23000"
	privateKey = "0x692a322c4eaacafc4ea25c8513270c9d62119f88a7c2866d51b820d5e4bd9696"
)

func TestFilterABIToRemoveCustomErrors(t *testing.T) {
	// Mock the necessary components and dependencies
	abiJSON := []byte(`[
		{
			"type": "function",
			"name": "myFunction",
			"inputs": [],
			"outputs": []
		},
		{
			"type": "error",
			"name": "MyCustomError"
		}
	]`)

	// Call the filterABIToRemoveCustomErrors function
	filteredABI, err := filterABIToRemoveCustomErrors(abiJSON)

	// Assert the expected results
	assert.NoError(t, err)
	assert.NotNil(t, filteredABI)
	// Assert that the filteredABI does not contain the custom error type
	assert.Contains(t, string(filteredABI), `"type":"function"`)
	assert.NotContains(t, string(filteredABI), `"type":"error"`)
}
func TestGenerateTypedArgs(t *testing.T) {
	// Mock the necessary components and dependencies
	contractAbi := abi.ABI{
		Methods: map[string]abi.Method{
			"testMethod": {
				Inputs: []abi.Argument{
					{
						Name: "arg1",
						Type: abi.Type{T: abi.StringTy},
					},
					{
						Name: "arg2",
						Type: abi.Type{T: abi.IntTy},
					},
					{
						Name: "arg3",
						Type: abi.Type{T: abi.BoolTy},
					},
				},
			},
		},
	}

	methodName := "testMethod"
	strArgs := []string{"test", "4567", "true"}

	// Call the GenerateTypedArgs function
	typedArgs, err := GenerateTypedArgs(contractAbi, methodName, strArgs)

	// Assert the expected results
	assert.NoError(t, err)
	assert.NotNil(t, typedArgs)
	// Assert the types of the generated arguments
	assert.IsType(t, "", typedArgs[0])
	assert.IsType(t, &big.Int{}, typedArgs[1])
	assert.IsType(t, false, typedArgs[2])
}

// MockEthClient is a mock implementation of the Ethereum client
type MockEthClient struct{}

func (m *MockEthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}

func (m *MockEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// Mock implementation
	return nil
}

func (m *MockEthClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	// Mock implementation
	return 0, nil
}

func TestGetNonce(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockEthClient := mocks.NewMockEthClient(mockCtrl)
	// Convert the private key to the expected format and derive the address
	privateKeyHex := privateKey
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privKey, _ := crypto.HexToECDSA(privateKeyHex) // Assume no error for simplicity
	expectedAddress := crypto.PubkeyToAddress(privKey.PublicKey)

	// Expected nonce value
	expectedNonce := uint64(12345)

	// Mock setup
	mockEthClient.EXPECT().
		PendingNonceAt(gomock.Any(), gomock.Eq(expectedAddress)).
		Return(expectedNonce, nil)

	// Call GetNonce
	nonce, err := GetNonce(mockEthClient, privateKeyHex)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedNonce, nonce)
}
