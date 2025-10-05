2、说明各层关键模块：
les（轻节点协议）
trie（默克尔树实现）
core/types（区块数据结构）

一、les（轻节点协议）：
轻量以太坊子协议是“轻量”客户端使用的协议，这些客户端仅在区块头出现时下载，并按需获取区块链的其他部分。他们提供安全访问区块链的完整功能，但不参与挖矿，因此不参与共识过程。
使用该协议通过Merkle Patricia Trie 证明验证账户状态的正确性，不下载完整的区块，仅验证根哈希是否匹配
对比ETH全节点协议优点为：带宽消耗极低、节点存储为MB级
通信流程：
1、Handshake
（1）LES 客户端通过 p2p 层建立连接。
（2）交换链状态（head、genesis、network ID）。

2、Header Synchronization
（1）轻节点请求区块头（GetBlockHeaders）。
（2）全节点返回区块头及部分证明。

3、State Proof Request
（1）客户端发送 GetProofsV2 请求。
（2）服务端提供账户状态的 MPT 证明（可验证，不需下载整树）。

4、Tx & Receipt Retrieval
（1）客户端可请求交易或收据，验证其哈希一致性。

5、流量控制机制
（1）LES 协议中有 “request credits” 系统，防止滥用全节点资源。


二、trie(默克尔树)：处于状态存储层，trie的结构主要由hash组成，均为偶数叶子节点，每个非叶子节点是由相邻的叶子节点的hash值拼接的哈希值。同时继续向上哈希，计算根节点
如：H1 H2 H3 H4 为叶子节点
P1 = keccak256(H1 || H2)
P2 = keccak256(H3 || H4)

        P1           P2
       /  \         /  \
     H1    H2     H3    H4

Root = keccak256(P1 || P2)

                 Root
                 /  \
               P1    P2
              / \    / \
            H1  H2 H3  H4

核心作用：负责以太坊中所有数据（账户、存储、交易收据）的高效存储和验证
注意：
若节点数为奇数，复制最后一个节点

如果叶子节点数为奇数（例如 5 个），为了保证能两两组合，
则复制最后一个节点，使节点数变为偶数：

H1, H2, H3, H4, H5 → H1, H2, H3, H4, H5, H5
P3 = keccak256(P5 || P5)


core/types (区块数据结构)：

三、core/types（区块数据结构）
定义了区块、交易、收据、签名 链上数据格式的基础结构
包括：
区块（Block）
区块头 （Header）
交易 （Transaction）
交易收据 （Receipt）
签名与哈希 （Signer）
Bloom过滤器 （日志检索）

数据结构
type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`  //父区块哈希
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`  //叔块哈希列表哈希
	Coinbase    common.Address `json:"miner"`                                 //矿工地址 
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`  //状态树根哈希
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`  //交易叔根哈希
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`  //收据树根哈希
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`  //日志过滤器
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`  //难度值，POS后为0x0
	Number      *big.Int       `json:"number"           gencodec:"required"`  //区块高度
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`  //附加信息（共识算法课使用）
	MixDigest   common.Hash    `json:"mixHash"`                    
	Nonce       BlockNonce     `json:"nonce"`                                 //工作量证明用

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// ExcessDataGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessDataGas *big.Int `json:"excessDataGas" rlp:"optional"`

	/*
		TODO (MariusVanDerWijden) Add this field once needed
		// Random was added during the merge and contains the BeaconState randomness
		Random common.Hash `json:"random" rlp:"optional"`
	*/
}