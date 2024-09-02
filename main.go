package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"uniswap-v4/contracts/LpRouter"
	"uniswap-v4/contracts/MockERC20"
	"uniswap-v4/contracts/PoolManager"
	"uniswap-v4/contracts/SwapRouter"
)

// Global variable for hook address
var hookAddress common.Address

type WalletManager struct {
	ethClient *ethclient.Client
	keystore  *keystore.KeyStore
}

func NewWalletManager(nodeURL string, keystoreDir string) (*WalletManager, error) {
	ethClient, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	return &WalletManager{
		ethClient: ethClient,
		keystore:  ks,
	}, nil
}

func (wm *WalletManager) ImportAccount(privateKeyHex string, password string) (common.Address, error) {
	// Convert the private key hex string to an ECDSA private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid private key: %v", err)
	}

	// Derive the public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	// Get the address corresponding to the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Check if the account already exists in the keystore
	for _, account := range wm.keystore.Accounts() {
		if account.Address == address {
			return address, nil
		}
	}

	// Create a new key using the private key and store it in the keystore
	ksAccount, err := wm.keystore.ImportECDSA(privateKey, password)
	if err != nil {
		if err == keystore.ErrAccountAlreadyExists {
			log.Printf("Account already exists in keystore: %s", address.Hex())
			return address, nil
		}
		return common.Address{}, fmt.Errorf("error importing account: %v", err)
	}

	// Return the address of the newly imported account
	return ksAccount.Address, nil
}

func findAccountByAddress(ks *keystore.KeyStore, address common.Address) (*accounts.Account, error) {
	for _, account := range ks.Accounts() {
		if account.Address == address {
			return &account, nil
		}
	}
	return nil, fmt.Errorf("account with address %s not found", address.Hex())
}

func (wm *WalletManager) SignTransaction(from common.Address, password string, tx *bind.TransactOpts) error {
	// Find the account in the keystore
	account, err := findAccountByAddress(wm.keystore, from)
	if err != nil {
		return fmt.Errorf("account not found: %v", err)
	}

	// Unlock the account using the provided password
	err = wm.keystore.Unlock(*account, password)
	if err != nil {
		return fmt.Errorf("failed to unlock account: %v", err)
	}

	// Get the nonce for the account
	nonce, err := wm.ethClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}
	tx.Nonce = big.NewInt(int64(nonce))

	// Suggest a gas price
	gasPrice, err := wm.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	tx.GasPrice = gasPrice
	tx.From = from
	tx.Context = context.Background()

	// Get the chain ID for the network
	chainID, err := wm.ethClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Decrypt the private key directly from the keystore
	privateKey, err := wm.keystore.Export(*account, password, password)
	if err != nil {
		return fmt.Errorf("failed to export private key: %v", err)
	}

	// Convert the key into a usable format
	key, err := keystore.DecryptKey(privateKey, password)
	if err != nil {
		return fmt.Errorf("failed to decrypt private key: %v", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	tx.Signer = signer.Signer

	return nil
}

type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

func (wm *WalletManager) ApproveTokens(params *ApproveTokensArgs) (string, error) {
	token0, err := MockERC20.NewMockERC20(params.Token0Address, wm.ethClient)
	if err != nil {
		return "", err
	}
	token1, err := MockERC20.NewMockERC20(params.Token1Address, wm.ethClient)
	if err != nil {
		return "", err
	}

	tx := &bind.TransactOpts{
		From:     params.From,
		GasLimit: 300000,
		Value:    big.NewInt(0),
		Context:  context.Background(),
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}

	_, err = token0.Approve(tx, params.LpRouterAddress, big.NewInt(0).SetUint64(^uint64(0)))
	if err != nil {
		return "", err
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}

	_, err = token0.Approve(tx, params.SwapRouterAddress, big.NewInt(0).SetUint64(^uint64(0)))
	if err != nil {
		return "", err
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}

	_, err = token1.Approve(tx, params.LpRouterAddress, big.NewInt(0).SetUint64(^uint64(0)))
	if err != nil {
		return "", err
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}
	_, err = token1.Approve(tx, params.SwapRouterAddress, big.NewInt(0).SetUint64(^uint64(0)))
	if err != nil {
		return "", err
	}

	return "Tokens approved successfully!", nil
}

func (wm *WalletManager) ProvideLiquidity(params *ProvideLiquidityArgs) (string, error) {
	lpRouter, err := LpRouter.NewLpRouter(params.LpRouterAddress, wm.ethClient)
	if err != nil {
		return "", err
	}

	// Add hook address from the global variable
	params.PoolKey.Hooks = hookAddress

	tx := &bind.TransactOpts{
		From:     params.From,
		GasLimit: 300000,
		Value:    big.NewInt(0),
		Context:  context.Background(),
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}

	_, err = lpRouter.ModifyLiquidity(tx, LpRouter.PoolKey(params.PoolKey), LpRouter.IPoolManagerModifyLiquidityParams{
		TickLower:      big.NewInt(params.TickLower),
		TickUpper:      big.NewInt(params.TickUpper),
		LiquidityDelta: params.Liquidity,
		Salt:           [32]byte{},
	}, nil, false, false)

	if err != nil {
		return "", err
	}

	return "Liquidity provided successfully!", nil
}

func (wm *WalletManager) PerformSwap(params *PerformSwapArgs) (string, error) {
	swapRouter, err := SwapRouter.NewSwapRouter(params.SwapRouterAddress, wm.ethClient)
	if err != nil {
		return "", err
	}

	// Add hook address from the global variable
	params.PoolKey.Hooks = hookAddress

	tx := &bind.TransactOpts{
		From:     params.From,
		GasLimit: 300000,
		Value:    big.NewInt(0),
		Context:  context.Background(),
	}

	err = wm.SignTransaction(params.From, params.Password, tx)
	if err != nil {
		return "", err
	}

	minSqrtPrice := big.NewInt(4295128739 + 1)
	maxSqrtPrice, _ := big.NewInt(0).SetString("1461446703485210103287273052203988822378723970341", 10)
	sqrtPriceLimit := minSqrtPrice
	if !params.ZeroForOne {
		sqrtPriceLimit = maxSqrtPrice
	}

	_, err = swapRouter.Swap(tx, SwapRouter.PoolKey(params.PoolKey), SwapRouter.IPoolManagerSwapParams{
		ZeroForOne:        params.ZeroForOne,
		AmountSpecified:   params.AmountSpecified,
		SqrtPriceLimitX96: sqrtPriceLimit,
	}, SwapRouter.PoolSwapTestTestSettings{
		TakeClaims:      false,
		SettleUsingBurn: false,
	}, nil)

	if err != nil {
		return "", err
	}
	return "Swap executed successfully!", nil
}

func (wm *WalletManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res JSONRPCResponse
	res.JSONRPC = "2.0"
	res.ID = req.ID

	switch req.Method {
	case "approveTokens":
		var params ApproveTokensArgs
		if err := json.Unmarshal(req.Params, &params); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
		result, err := wm.ApproveTokens(&params)
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Result = result
		}

	case "provideLiquidity":
		var params ProvideLiquidityArgs
		if err := json.Unmarshal(req.Params, &params); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		result, err := wm.ProvideLiquidity(&params)
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Result = result
		}

	case "performSwap":
		var params PerformSwapArgs
		if err := json.Unmarshal(req.Params, &params); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
		result, err := wm.PerformSwap(&params)
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Result = result
		}

	default:
		res.Error = "Method not found"
	}

	json.NewEncoder(w).Encode(res)
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

func main() {
	// Initialize the global hook address
	hookAddress = common.HexToAddress("0x85Fdc1D643EE6152765CD0c72c2434be2C238aC0")

	walletManager, err := NewWalletManager("http://localhost:8545", "./keystore")
	if err != nil {
		log.Fatalf("Failed to create wallet manager: %v", err)
	}

	_, err = walletManager.ImportAccount("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "TEST_abc123")
	if err != nil {
		log.Fatalf("Failed to import account: %v", err)
	}

	http.Handle("/rpc", walletManager)
	fmt.Println("JSON-RPC server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
