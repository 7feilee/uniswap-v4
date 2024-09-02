package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"uniswap-v4/contracts/PoolManager"

	"github.com/ethereum/go-ethereum/common"
)

type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	ID      int         `json:"id"`
}

func main() {
	rpcURL := "http://localhost:8080/rpc"

	// Define command-line flags
	cmd := flag.String("cmd", "", "Command to execute: approve, provide, swap")
	from := flag.String("from", "", "Sender address")
	password := flag.String("password", "", "Password for keystore")
	token0 := flag.String("token0", "0xC7f2Cf4845C6db0e1a1e91ED41Bcd0FcC1b0E141", "Token0 address")
	token1 := flag.String("token1", "0xdaE97900D4B184c5D2012dcdB658c008966466DD", "Token1 address")
	lpRouter := flag.String("lpRouter", "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "Liquidity pool router address")
	swapRouter := flag.String("swapRouter", "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9", "Swap router address")
	tickLower := flag.Int64("tickLower", -600, "Lower tick for liquidity")
	tickUpper := flag.Int64("tickUpper", 600, "Upper tick for liquidity")
	liquidity := flag.String("liquidity", "0", "Amount of liquidity")
	amountSpecified := flag.String("amount", "0", "Amount for swap")
	zeroForOne := flag.Bool("zeroForOne", true, "Direction of swap: token0 -> token1 if true")

	flag.Parse()

	if *cmd == "" || *from == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	fromAddress := common.HexToAddress(*from)

	switch *cmd {
	case "approve":
		approveParams := &ApproveTokensArgs{
			From:              fromAddress,
			Password:          *password,
			Token0Address:     common.HexToAddress(*token0),
			Token1Address:     common.HexToAddress(*token1),
			LpRouterAddress:   common.HexToAddress(*lpRouter),
			SwapRouterAddress: common.HexToAddress(*swapRouter),
		}
		response := makeRPCRequest(rpcURL, "approveTokens", approveParams)
		fmt.Println(response)

	case "provide":
		liquidityAmount, _ := new(big.Int).SetString(*liquidity, 10)
		provideParams := &ProvideLiquidityArgs{
			From:            fromAddress,
			Password:        *password,
			LpRouterAddress: common.HexToAddress(*lpRouter),
			PoolKey: PoolManager.PoolKey{
				Currency0:   common.HexToAddress(*token0),
				Currency1:   common.HexToAddress(*token1),
				Fee:         big.NewInt(3000),
				TickSpacing: big.NewInt(60),
			},
			TickLower: *tickLower,
			TickUpper: *tickUpper,
			Liquidity: liquidityAmount,
		}
		response := makeRPCRequest(rpcURL, "provideLiquidity", provideParams)
		fmt.Println(response)

	case "swap":
		swapAmount, _ := new(big.Int).SetString(*amountSpecified, 10)
		swapParams := &PerformSwapArgs{
			From:              fromAddress,
			Password:          *password,
			SwapRouterAddress: common.HexToAddress(*swapRouter),
			PoolKey: PoolManager.PoolKey{
				Currency0:   common.HexToAddress(*token0),
				Currency1:   common.HexToAddress(*token1),
				Fee:         big.NewInt(3000),
				TickSpacing: big.NewInt(60),
			},
			AmountSpecified: swapAmount,
			ZeroForOne:      *zeroForOne,
		}
		response := makeRPCRequest(rpcURL, "performSwap", swapParams)
		fmt.Println(response)

	default:
		fmt.Println("Unknown command:", *cmd)
		flag.Usage()
		os.Exit(1)
	}
}

func makeRPCRequest(rpcURL, method string, params interface{}) string {
	reqBody, err := json.Marshal(JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	})
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("Failed to make RPC request: %v", err)
	}
	defer resp.Body.Close()

	// Read the entire response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var rpcResp JSONRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		log.Printf("Raw response: %s", body) // Print the raw response for debugging
		log.Fatalf("Failed to decode RPC response: %v", err)
	}

	if rpcResp.Error != nil {
		return fmt.Sprintf("RPC Error: %v", rpcResp.Error)
	}

	return fmt.Sprintf("RPC Result: %v", rpcResp.Result)
}

type ApproveTokensArgs struct {
	From              common.Address
	Password          string
	Token0Address     common.Address
	Token1Address     common.Address
	LpRouterAddress   common.Address
	SwapRouterAddress common.Address
}

type ProvideLiquidityArgs struct {
	From            common.Address
	Password        string
	LpRouterAddress common.Address
	PoolKey         PoolManager.PoolKey
	TickLower       int64
	TickUpper       int64
	Liquidity       *big.Int
}

type PerformSwapArgs struct {
	From              common.Address
	Password          string
	SwapRouterAddress common.Address
	PoolKey           PoolManager.PoolKey
	AmountSpecified   *big.Int
	ZeroForOne        bool
}
