# 部署合约
## 部署步骤（todo）
产生4个文件  
`aggregator.keystore`  
`sequencer.keystore`  
`genesis.json`  
`deploy_output.json`  
# 部署zkNode
进入`x1-node`的`test`目录下  
## 1. 修改config/test.genesis.json
将genesis.json内容复制到test.genesis.json中，并且在root字段前添加如下内容：
``` json
  "l1Config": {
    "chainId": 11155111,
    // 以下字段内容可在deploy_out.json中找到
    "polygonZkEVMAddress": "", 
    "maticTokenAddress": "",
    "polygonZkEVMGlobalExitRootAddress": "",
    "dataCommitteeContract": ""
  },
  "genesisBlockNumber":  123
```
## 2. 替换sequencer.keystore, aggregator.keystore
## 3. 修改config/test.node.config.toml
```
FreeGasAddress = "0xdc5711dd1b5feb2af756f5264734127592e25df4"

[SequenceSender]
WaitPeriodSendSequence = "30m"
LastBatchVirtualizationTimeMaxWaitPeriod = "5m"
SenderAddress = "0x19536e2a814f041662b44850eb2a045b837ca548"
L2Coinbase = "0x19536e2a814f041662b44850eb2a045b837ca548"

[Aggregator]
SenderAddress = "0x4e2dd2a9349dd39f65fd6b4a3a79916fc43bcb11"
```
## 4. 修改config/test.bridge.config.toml
```
[NetworkConfig]
GenBlockNumber = 5367840
PolygonBridgeAddress = "0xaEB5133aDD4A529d6783BC81bDEADF2630c4507a"
PolygonZkEVMGlobalExitRootAddress = "0x4D149fA2264f5E1F55690d9adCBDADe9F4c3b0d1"
L2PolygonBridgeAddresses = ["0xaEB5133aDD4A529d6783BC81bDEADF2630c4507a"] //为genesis.json中PolygonZkEVMBridge proxy的address
```
# cdk-bridge-ui配置
.env中
```
// 见deploy_output.json
VITE_ETHEREUM_BRIDGE_CONTRACT_ADDRESS=0x7cc8f4b4357Cff5897f8b50Ea115a2d11798DCD4  // polygonZkEVMBridgeAddress
VITE_ETHEREUM_PROOF_OF_EFFICIENCY_CONTRACT_ADDRESS=0x0d543Fb5aEA6cc82c1F91837ea47866e236FCD53 // polygonZkEVMAddress
VITE_POLYGON_ZK_EVM_BRIDGE_CONTRACT_ADDRESS=0x7cc8f4b4357Cff5897f8b50Ea115a2d11798DCD4 // genesis.json 中PolygonZkEVMBridge proxy的地址。
```
constants.ts中
```
export const L1GASTOKEN_ADDRESS    = "0xcf41Fd1317CE3b7B04628548fCE038D7669cd521";// 为第一步部署erc20合约时产生的地址(详细说明)
export const L2WRAPPERETH_ADDRESS  = "0x96C32B6250A191DD79D79b86A1EebaC0ABc7aDb9"; // 为genesis.json中WETHzkEVM的address
```
# prover 分离模式下，node的启动命令
1. make run1
2. 等待prover启动
3. make run3

# 部署节点加入zkNode(todo)
