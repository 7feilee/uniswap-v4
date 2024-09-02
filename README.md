
# Uniswap V4 RPC

This project provides a Go-based JSON-RPC server/client that interacts with the Uniswap V4 protocol to approve tokens, provide liquidity, and perform swaps. It includes functionality for managing Ethereum wallets, signing transactions, and sending them to the Ethereum network.

## Features

- **Approve Tokens**: Approve tokens for Uniswap V4 contracts.
- **Provide Liquidity**: Add liquidity to a Uniswap V4 pool.
- **Perform Swaps**: Swap tokens within a Uniswap V4 pool.

## Requirements

- **Go**: Ensure that Go is installed on your system. You can download it [here](https://golang.org/dl/).
- **Anvil (Foundry)**: A local Ethereum node for testing. You can install Foundry [here](https://book.getfoundry.sh/getting-started/installation.html).
- **Solidity Compiler (solc)**: Required to compile the smart contracts. Install using apt:
  ```bash
  sudo apt install solc
  ```
- **abigen**: Part of the Go-Ethereum toolkit to generate Go bindings for the smart contracts. You can install it with:
  ```bash
  go install github.com/ethereum/go-ethereum/cmd/abigen@latest
  ```

## Setup

### 1. Deploy the v4 contract on local Anvil

Following the readme in [github](https://github.com/7feilee/v4-template/blob/main/README.md#local-development-anvil).



### 2. Clone the RPC Repository

```bash
git clone https://github.com/7feilee/uniswap-v4.git
cd uniswap-v4
```

### 3. Install Go Dependencies

Navigate to the project directory and run:

```bash
go mod tidy
```

This will download and install the necessary Go dependencies.


### 4. Create a Keystore Directory

Create a directory to store Ethereum keystore files (private keys):

```bash
mkdir keystore
```

### 5. Run the JSON-RPC Server

Start the JSON-RPC server:

```bash
go run main.go
```

This will start the server on `http://localhost:8080`.

### 6. Use the Client CLI

Use the provided client CLI to interact with the JSON-RPC server. Here are a few example commands:

- **Approve Tokens**:

Please first run Approve operation so that you could do provide liquidity and swap operations.

```bash
go run client-cli/client.go -cmd approve -from 0xYourAddress -password YourPassword -token0 0xToken0Address -token1 0xToken1Address -lpRouter 0xLpRouterAddress -swapRouter 0xSwapRouterAddress
```

- **Provide Liquidity**:

```bash
go run client-cli/client.go -cmd provide -from 0xYourAddress -password YourPassword -token0 0xToken0Address -token1 0xToken1Address -lpRouter 0xLpRouterAddress -tickLower -600 -tickUpper 600 -liquidity 1000000000000000000
```

short example in the anvil test environment (Please change the contract address in command-line flags).
bash
```
go run client-cli/client.go -cmd provide -from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -password TEST_abc123
```

- **Perform Swap**:

```bash
go run client-cli/client.go -cmd swap -from 0xYourAddress -password YourPassword -token0 0xToken0Address -token1 0xToken1Address -swapRouter 0xSwapRouterAddress -amount 1000000000000000000 -zeroForOne true
```

short example in the anvil test environment (Please change the contract address in command-line flags).

```bash
go run client-cli/client.go -cmd swap -from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -password TEST_abc123 -amount  1000000
```

### 7. Running Tests

To run unit tests:

```bash
go test -v
```

This will execute all tests within the project.

## Note

### Generate the [Go Contract Bindings](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings)

```bash
solc --abi --bin lib/v4-core/lib/solmate/src/test/utils/mocks/MockERC20.sol -o output/ --overwrite
solc --abi --bin lib/v4-core/src/test/PoolModifyLiquidityTest.sol  -o output/ --overwrite
solc --abi --bin lib/v4-core/src/interfaces/IPoolManager.sol -o output/ --overwrite
solc --abi --bin lib/v4-core/src/test/PoolSwapTest.sol -o output/ --overwrite
abigen --abi output/MockERC20.abi --pkg MockERC20 --out contracts/MockERC20.go
abigen --abi output/PoolModifyLiquidityTest.abi --pkg LpRouter --out contracts/LpRouter.go
abigen --abi output/IPoolManager.abi --pkg PoolManager --out contracts/PoolManager.go
abigen --abi output/PoolSwapTest.abi --pkg SwapRouter --out contracts/SwapRouter.go
```

## Contributing

Feel free to contribute to this project by submitting issues or pull requests.

## License

This project is licensed under the MIT License.

