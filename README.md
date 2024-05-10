# Skate AVS services

Built with target:

+ eigenlayer-contracts [v0.2.3-mainnet-m2](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v0.2.3-mainnet-m2)
+ eigenlayer-middleware [v0.1.3-mainnet-m2+pragma-change](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v0.1.3-mainnet-m2%2Bpragma-change)
+ eigensdk-go [v0.1.6](https://github.com/Layr-Labs/eigensdk-go/tree/cff810715271da986a7e594f7967a86fc4299834)

---

# Pre-built packages

## Operator:

Images: [https://github.com/orgs/Skate-Org/packages/container/package/skate-operator](https://github.com/orgs/Skate-Org/packages/container/package/skate-operator)

# From source

## CLI tools

All CLI tools come with help, if you are curious about how things work underneath. Try

`go run <service>/main.go <action> -h`

for example:

```bash
go run kms/main.go store -h
```

**@For dev** - entry point for each service is invoked from `<service>/cmd/cmd.go`, traceback logic from there

### I. Configuring accounts

Key management system (kms) is an utility tool following the [geth guideline](https://geth.ethereum.org/docs/developers/dapp-developer/native-accounts)

To run, start with:

```bash
go run kms/main.go store -p <passphrase> -k <account_private_key> [OPTIONAL] -s <save_path>
```

Key information will be dumped in `./keystore` folders and used for subsequent services. To load a signer to be used with respective services, 
specify **ACCOUNT ADDRESS** and corresponding **passphrase**. Example config in `./configs/signer/operator/1.yaml`

__WARNING ⚠️ - NEVER LEAK THE FILES GENERATED IN `./keystore`. Otherwise, private key can potentially be exposed.__

### II. Register Operator to Skate AVS

Command for Operator onboarding

#### Step 1: Register with EigenLayer to become an operator

```bash
go run operator/main.go register-el -h
```

see `./operator/cmd/register.go:registerAvsCmd(..)` for full logic and options. 

Default run with:

+ environment: `configs/environment/testnet.yaml` 
+ signer: `configs/signer/operator/1.yaml` 

populate signer config for auto account import, else manually specify `--signer-address` and `--passphrase`

#### Step 2: Deposit into token strategy

1. AVS required minimum shares/stakes allowance to opt-in. Skate AVS (holesky testnet) uses [stETH](https://holesky.etherscan.io/token/0x3f1c547b21f65e10480de3ad8e19faac46c95034)
as the underlying token strategy. Get stETH buy calling submit function (recommend amount: 0.01 ETH)

2. Once stETH is obtained, deposit into EigenLayer StrategyManager.

```bash
go run operator/main.go deposit -h
```

Default value from config is the minimum required to participate in Skate AVS

Follow the same step to customize environments as in step 1


#### Step 3: Register with Skate AVS

```bash
go run operator/main.go register-avs -h
```

Default options is good, same environments configuration as in step 1.

_**NOTE**: Default Operator details don't include any metadata URI. Other initialize params specified in `operator/cmd/register.go:L126`_

### III. Running Operator services

Operator will monitor Skate App (on Nollie testnet) activity and sign confirmation message then publish to relayer.

Cache is maintained locally on operator environments, db config in `operator/db/**`. 
By default an SQLite db will be created under `data/operator/skateapp.db`, files logs for schemas maintained in the same folder.

To participate (sign and publish task), run:

```bash
go run operator/monitor.go
```

_NOTE: Follow default signer configs. if no signer provided, run with watch only-mode (no sign and publish to relayer)_

### IV. Running Relayer services

Relayer will collect operator signatures, aggregates and send over to AVS for confirmation (see `SkateAVS.sol:submitData(...)`)

#### 1. Listen for Operator

To collect operator signatures for a task and store in database

```bash
go run relayer/main.go retrieve
```


#### 2. Publish to AVS and settle on Gateway contract

To publish quorums that reached threshold to the avs

```bash
go run relayer/main.go publish
```

_TODO: publish verified message to SkateGateway on destination chains as well_

## Quick start

Make files `./Makefile` contains shortcut to bootstrap pilot avs services.
Please explore and config respective nodes using designated CLI tools

__Prerequisites: 3 operator accounts registered with avs and 1 relayer__

1. Start all 3 operators:

```bash
make start-operators
```

2. Start all relayer service (retrieve and publish):

```bash
make start-relayer
```


--- 
## Deployment info

### Testnet

See `configs/environment/testnet.yaml`

Currently 3 operator (registered with Skate AVS):

+ `configs/signer/operator/1.yaml`: 0x786775c9ecB916bd7f5a59c150491871fCfCEe86
+ `configs/signer/operator/2.yaml`: 0x72b3793B2A476c055A88dfd5e38D1E032a27e038
+ `configs/signer/operator/3.yaml`: 0xCD6DB57894AfE39AC5Db4B62E17971Ae07c5EC91

**@For dev** - 3 are burner accounts, contacts for private keys

---

# To Be Completed

## Skate gateway

Improve configuration format + finish logic in `relayer/publish/skateapp.go:submitTasksToAvs(..)`

## Docker images


## Metrics server

WIP
