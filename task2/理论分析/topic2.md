
2、解析核心模块交互关系：
区块链同步协议（eth/62,eth/63）
; eth/63 (2016)
; 版本 63 添加了 GetNodeData，NodeData，GetReceipts 和 Receipts 消息，这些消息允许同步交易执行结果。

; eth/62 (2015)
; 在版本 62 中，扩展了 NewBlockHashes 消息，以包含区块哈希以及区块编号。Status 中的区块编号已删除。消息 GetBlockHashes (0x03)，BlockHashes (0x04)，GetBlocks (0x05) 和 Blocks (0x06) 由提取区块标头和主体的消息替换。BlockHashesFromNumber (0x08) 消息已删除。
交易池管理与Gas机制
EVM执行环境构建
共识算法实现（Ethash/POS）

答题：
（1）
用于节点之间同步区块头、区块体、收据和状态数据
服务于新区块传播和状态恢复

（2）
交易通过JSON-RPC创建交易，P2p接收上层交易
验证签名、nonce、余额、gasLimit
维护pending和queued队列
根据gasprice进行排序和丢弃，管理交易池容量和价格高的交易先处理原则
广播到其他几点

(3)
从TxPool取出交易，并在当前状态下用于执行智能合约代码进行交易转账等操作
维护账户状态、存储、日志
根据交易计算消耗gas
生成receipts给共识层验证

（4）
出块由共识层验证者完成
对于正式的主网，miner不再进行区块打包，而是仍然会使用rpc调用，构建执行层部分区块，并使用EngineAPI与共识层通信，协助共识层完成签名、验证和广播。