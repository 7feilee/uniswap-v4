package main

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"uniswap-v4/contracts/PoolManager"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func initWalletManager(t *testing.T) *WalletManager {
	// Connect to the actual Ethereum client
	walletManager, err := NewWalletManager("http://localhost:8545", "./keystore")
	if err != nil {
		t.Fatalf("Failed to create wallet manager: %v", err)
	}

	_, err = walletManager.ImportAccount("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "TEST_abc123")
	if err != nil {
		t.Fatalf("Failed to import account: %v", err)
	}
	return walletManager
}

// Test the ApproveTokens method
func TestApproveTokens(t *testing.T) {
	walletManager := initWalletManager(t)

	// Mock data
	params := ApproveTokensArgs{
		From:              common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Password:          "TEST_abc123",
		Token0Address:     common.HexToAddress("0xC7f2Cf4845C6db0e1a1e91ED41Bcd0FcC1b0E141"),
		Token1Address:     common.HexToAddress("0xdaE97900D4B184c5D2012dcdB658c008966466DD"),
		LpRouterAddress:   common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		SwapRouterAddress: common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
	}

	// Convert params to json.RawMessage
	paramsBytes, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsBytes)

	// Prepare the request body
	reqBody, _ := json.Marshal(JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "approveTokens",
		Params:  paramsRaw,
		ID:      1,
	})

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(walletManager.ServeHTTP)

	// Call the ServeHTTP method with our mock request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Parse the response
	var resp JSONRPCResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the result is correct
	assert.NotNil(t, resp.Result, "Result should not be nil")
	assert.Nil(t, resp.Error, "Error should be nil")
}

// Test the ProvideLiquidity method
func TestProvideLiquidity(t *testing.T) {
	walletManager := initWalletManager(t)

	// Mock data
	params := ProvideLiquidityArgs{
		From:            common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Password:        "TEST_abc123",
		LpRouterAddress: common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		PoolKey: PoolManager.PoolKey{
			Currency0:   common.HexToAddress("0xC7f2Cf4845C6db0e1a1e91ED41Bcd0FcC1b0E141"),
			Currency1:   common.HexToAddress("0xdaE97900D4B184c5D2012dcdB658c008966466DD"),
			Fee:         big.NewInt(3000),
			TickSpacing: big.NewInt(60),
		},
		TickLower: -600,
		TickUpper: 600,
		Liquidity: big.NewInt(1000000000000000000),
	}

	// Convert params to json.RawMessage
	paramsBytes, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsBytes)

	// Prepare the request body
	reqBody, _ := json.Marshal(JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "provideLiquidity",
		Params:  paramsRaw,
		ID:      1,
	})

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(walletManager.ServeHTTP)

	// Call the ServeHTTP method with our mock request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Parse the response
	var resp JSONRPCResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the result is correct
	assert.NotNil(t, resp.Result, "Result should not be nil")
	assert.Nil(t, resp.Error, "Error should be nil")
}

// Test the PerformSwap method
func TestPerformSwap(t *testing.T) {
	walletManager := initWalletManager(t)

	// Mock data
	params := PerformSwapArgs{
		From:              common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Password:          "TEST_abc123",
		SwapRouterAddress: common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
		PoolKey: PoolManager.PoolKey{
			Currency0:   common.HexToAddress("0xC7f2Cf4845C6db0e1a1e91ED41Bcd0FcC1b0E141"),
			Currency1:   common.HexToAddress("0xdaE97900D4B184c5D2012dcdB658c008966466DD"),
			Fee:         big.NewInt(3000),
			TickSpacing: big.NewInt(60),
		},
		AmountSpecified: big.NewInt(1000000000000000000),
		ZeroForOne:      true,
	}

	// Convert params to json.RawMessage
	paramsBytes, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsBytes)

	// Prepare the request body
	reqBody, _ := json.Marshal(JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "performSwap",
		Params:  paramsRaw,
		ID:      1,
	})

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(walletManager.ServeHTTP)

	// Call the ServeHTTP method with our mock request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Parse the response
	var resp JSONRPCResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the result is correct
	assert.NotNil(t, resp.Result, "Result should not be nil")
	assert.Nil(t, resp.Error, "Error should be nil")
}
