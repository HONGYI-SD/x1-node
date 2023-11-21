# Setup Production zkNode
X1 is now available on the Testnet for developers to launch smart contracts, execute transactions, and experiment with the network. This tutorial extends the exploration by allowing developers to launch their own node on the Public Testnet.
Before we begin, this document is fairly technical and requires prior exposure to Docker and CLI. Post spinning up your zkNode instance, you will be able to run the Synchronizer and utilize the JSON-RPC interface.

## Prerequisites
This tutorial assumes that you have docker-compose already installed. If you need any help with the installation, please check the [official docker-compose installation guide](https://docs.docker.com/compose/install/).

### Minimum System Requirements
<Tip title="CAUTION">zkProver does not work on ARM-based Macs yet, and using WSL/WSL2 on Windows is not advisable. Currently, zkProver optimizations require CPUs that support the AVX2 instruction, which means some non-M1 computers, such as AMD, won't work with the software regardless of the OS.</Tip> 
 
- 16GB RAM
- 4-core CPU
- 20GB Storage (This will increase over time) 

### Network Components
Here is a list of crucial network components that are required before you can run the zkNode:
- Ethereum Node - Use geth or any service providing a JSON RPC interface for accessing the L1 network
- X1-Node (or zkNode)  - L2 Network
  - Synchronizer - Responsible for synchronizing data between L1 and L2
  - JSON RPC Server - Interface to L2 network 
  - State DB - Save the L2 account, block and tx data.

Let's set up each of the above components!

## Ethereum Node Setup
The Ethereum RPC Node is the first component to be deployed because zkNode needs to synchronize blocks and transactions on L1. You can invoke the ETH RPC (Testnet: Sepolia) service through any of the following methods:
- Third-party RPC services, such as [Infura](https://www.infura.io/) or [Ankr](https://www.ankr.com/).
- Set up your own Ethereum node. Follow the instructions provided in this [guide to set up and install Geth](https://geth.ethereum.org/docs/getting-started/installing-geth).

## Installing
Once the L1 RPC component is complete, we can start the zkNode setup. This is the most straightforward way to run a zkNode and it's fine for most use cases. 
Furthermore, this method is purely subjective and feel free to run this software in a different manner. For example, Docker is not required, you could simply use the Go binaries directly.
Let's start setting up our zkNode:

1. Download the installation scrip
``` bash
mkdir -p ./x1-node && cd ./x1-node

wget https://static.okex.org/cdn/chain/x1/snapshot/run_x1_testnet.sh && chmod +x run_x1_testnet.sh && ./run_x1_testnet.sh init && cp ./testnet/example.env ./testnet/.env
```
2. The example.env file must be modified according to your configurations. Edit the .env file with your favourite editor (we'll use vim in this guide): ```vim ./testnet/.env.```
``` bash
# URL of a JSON RPC for Sepolia
X1_NODE_ETHERMAN_URL = "http://your.L1node.url"

# PATH WHERE THE STATEDB POSTGRES CONTAINER WILL STORE PERSISTENT DATA
X1_NODE_STATEDB_DATA_DIR = "./x1_testnet_data/statedb"

# PATH WHERE THE POOLDB POSTGRES CONTAINER WILL STORE PERSISTENT DATA
X1_NODE_POOLDB_DATA_DIR = "/x1_testnet_data/pooldb"
```
3. Restore the latest L2 snapshot  locally database for synchronizing  L2 data quickly.
``` bash
./run_x1_testnet.sh restore 
```

## Starting
Use the below command to start the zkNode instance:
``` bash
./run_x1_testnet.sh start

docker ps -a
```

You will see a list of the following containers :
  - x1-rpc
  - x1-sync
  - x1-state-db
  - x1-pool-db
  - x1-prover

You should now be able to run queries to the JSON-RPC endpoint at http://localhost:8545.
Run the following query to get the most recently synchronized L2 block; if you call it every few seconds, you should see the number grow:
``` bash
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}' http://localhost:8545
```

## Stopping
Use the below command to stop the zkNode instance:
``` bash
./run_x1_testnet.sh stop
```

## Restarting
Use the below command to stop the zkNode instance:
``` bash
./run_x1_testnet.sh restart
```
## Updating
To update the zkNode software, run the below command, and the file ```./testnet/.env``` will be retained, the other config will be deleted.
``` bash
./run_x1_testnet.sh update
```

## Troubleshooting
- It's possible that the machine you're using already uses some of the necessary ports. In this case, you can change them directly ```./testnet/docker-compose.yml```.
- If one or more containers are crashing, please check the logs using the command below:
``` bash
docker ps -a

docker logs <cointainer_name>
```

## Advanced setup

> DISCLAIMER: right now this part of the documentation attempts to give ideas on how to improve the setup for better performance, but is far from being a detailed guide on how to achieve this. Please open issues requesting more details if you don't understand how to achieve something. We will keep improving this doc for sure!

There are some fundamental changes that can be done towards the basic setup, in order to get better performance and scale better:

### DB

In the basic setup, there are Postgres being instanciated as Docker containers. For better performance is recommended to:

- Run dedicated instances for Postgres. To achieve this you will need to:
  - Remove the Postgres services (`x1-pool-db` and `x1-state-db`) from the `docker-compose.yml`
  - Instantiate Postgres elsewhere (note that you will have to create credentials and run some queries to make this work, following the config files and docker-compose should give a clear idea of what to do)
  - Update the `node.config.toml` to use the correct URI for both DBs
  - Update `prover.config.json` to use the correct URI for the state DB
- Use a setup of Postgres that allows to have separated endpoints for read / write replicas

### JSON RPC

Unlike the synchronizer, that needs to have only one instance running (having more than one synchronizer running at the same time connected to the same DB can be fatal), the JSON RPC can scale horizontally.

There can be as many instances of it as needed, but in order to not introduce other bottlenecks, it's important to consider the following:

- Read replicas of the State DB should be used
- Synchronizer should have an exclusive instance of `x1-prover`
- JSON RPCs should scale in correlation with instances of `x1-prover`. The most obvious way to do so is by having a dedicated `x1-prover` for each `x1-rpc`. But depending on the payload of your solution it could be worth to have `1 x1-rpc : many x1-prover` or `many x1-rpc : 1 x1-prover`, ... For reference, the `x1-prover` implements the EVM, and therefore will be heavily used when calling endpoints such as `eth_call`. On the other hand, there are other endpoints that relay on the `x1-state-db`
