package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/local/go-eth-demo/erc20"
	store "github.com/local/go-eth-demo/store"
	"golang.org/x/crypto/sha3"
)

func main_block1() {
	// 查询区块
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//blockNumber := big.NewInt(5671744)
	//// 块头 使用 HeaderByNumber
	//header, err := client.HeaderByNumber(context.Background(), blockNumber)
	//fmt.Println(header.Number.Uint64())     // 5671744 //区块号，即区块高度
	//fmt.Println(header.Time)                // 1712798400 //区块计时器
	//fmt.Println(header.Difficulty.Uint64()) // 0   //出块难度，因为是测试网，默认为0
	//fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 使用 BodyByNumber
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//fmt.Println(block.Number().Uint64())
	//fmt.Println(block.Time())
	//fmt.Println(block.Difficulty().Uint64())
	//fmt.Println(block.Hash().Hex()) //区块哈希
	///**
	//block.Hash().Hex()
	//
	//返回的是 字符串类型 (string)。
	//
	//底层调用了 common.Hash.Hex() 方法，把 common.Hash 转成 0x 开头的十六进制字符串
	//*/
	//fmt.Println(len(block.Transactions())) //获取交易总书数
	//fmt.Println(block.Transactions())
	//fmt.Println(block.Hash()) //区块哈希
	//
	///**
	//block.Hash()
	//
	//返回的是 common.Hash 类型（本质上是 [32]byte 的数组封装）。
	//
	//它代表区块的哈希值，二进制存储形式。
	//
	//在 Go 代码中如果直接 fmt.Println(block.Hash())，会看到类似：
	//
	//0x3f1e3c4f5f9c8bb80e4a5f497d5a5c7e6b2cbe9d84f65a246f8c3a1c7d3e2c4f
	//*/
	//fmt.Println(client)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	////调用TransactionCount仅返回一个区块的交易数量
	//count, err := client.TransactionCount(context.Background(), block.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(count)

	// 查询交易
	// 最重要的是Data，会调用合约的方法签名信息和合约的入参
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//方式一：知道块高度
	//需要拿到发送方的地址,使用当前事务进行，context.Background()
	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//blockNumber := big.NewInt(5671744)
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, tx := range block.Transactions() {
	//	fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	fmt.Println(tx.Value().String())    // 100000000000000000
	//	fmt.Println(tx.Gas())               // 21000
	//	fmt.Println(tx.GasPrice().Uint64()) // 100000000000
	//	fmt.Println(tx.Nonce())             // 245132 重要
	//	fmt.Println(tx.Data())              // []  重要，记录=
	//	fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587
	//	if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
	//		fmt.Println("sender", sender.Hex())
	//	} else {
	//		log.Fatal(err)
	//	}
	//	//每笔交易都有一个收据，其中包含执行交易的结果，例如所有的返回值和日志，以及“1”（成功）或“0”（失败）的交易结果状态。
	//	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//获取交易状态
	//	fmt.Println(receipt.Status)
	//	fmt.Println(receipt.Logs)
	//
	//	//终止
	//	break
	//}

	//方式二：知道块的哈希
	//字符串转成common.Hash
	//blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	//count, err := client.TransactionCount(context.Background(), blockHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for idx := uint(0); idx < count; idx++ {
	//	tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	fmt.Println(tx.Value().String())    // 100000000000000000  交易转账的金额。单位为wei 数据类型为 big.int
	//	fmt.Println(tx.Gas())               // 21000
	//	fmt.Println(tx.GasPrice().Uint64()) // 100000000000
	//	fmt.Println(tx.Nonce())             // 245132 重要
	//	fmt.Println(tx.Data())              // []  重要，记录=
	//	fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587
	//	break
	//}
	////获取txHash对应的isPending状态
	//txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(isPending)
	//fmt.Println(tx.Hash().Hex())

	//查询收据
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	////client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////两种方式查询，一区块高度，二块hash
	//blockNumber := big.NewInt(5671744)
	//blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	///**
	//1. BlockNumberOrHashWithHash(hash, requireCanonical bool)
	//
	//这个函数的作用是构造一个参数，可以让 RPC 方法用 区块哈希 来定位区块。
	//
	//第二个布尔值叫 requireCanonical，字面意思是：
	//
	//是否强制要求区块属于 规范链（canonical chain）。
	//
	//2. requireCanonical = false
	//
	//表示：不强制要求该区块在规范链上。
	//
	//只要找到这个哈希对应的区块，就会返回结果。
	//
	//好处是：即使该区块后来变成了孤块（uncle/orphan），你也能查到它的交易收据。
	//
	//3. requireCanonical = true
	//
	//表示：必须保证区块在规范链上。
	//
	//如果你传入的哈希对应的区块 不是主链的一部分（比如被链重组替换掉了），那么 RPC 会返回 not found。
	//
	//这种用法适合严格需要主链数据的场景。
	//*/
	//receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(receiptByHash[0] == receiptsByNum[0])
	//
	//for _, receipt := range receiptByHash {
	//	fmt.Println(receipt.Status)                // 1
	//	fmt.Println(receipt.Logs)                  // []
	//	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	fmt.Println(receipt.TransactionIndex)      // 0
	//	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
	//	break
	//}
	//
	//txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//receipt, err := client.TransactionReceipt(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(receipt.Status)                // 1
	//fmt.Println(receipt.Logs)                  // []
	//fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//fmt.Println(receipt.TransactionIndex)      // 0
	//fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000

	//创建钱包
	//两种方式（1.私钥随机生成，2.指定私钥生成）
	//1.私钥随机生成
	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}
	/**
	随机结果1：
	1372a42dc6a042fd047deac1dbb66bcc83d3631d4a9a76fdf9988dc1371029ec

	公钥full: 0xeb1c9274258695bb95e45e3e7ab24dfd7fceda48a044d5bea1c7e3ada6227f08
	0x7ab24dfd7fceda48a044d5bea1c7e3ada6227f08
	*/

	//2.如果有16进制的私钥，可以使用指定私钥生成
	// Hex 将 16进制字符串，恢复成ECDSA的原生私钥
	//privateKey, err := crypto.HexToECDSA("1372a42dc6a042fd047deac1dbb66bcc83d3631d4a9a76fdf9988dc1371029ec")
	//if err != nil {
	//	log.Fatal(err)
	//}
	///**
	//指定结果2： 与随机中一样的1372a42dc6a042fd047deac1dbb66bcc83d3631d4a9a76fdf9988dc1371029ec密钥，生成的公钥一致
	//1372a42dc6a042fd047deac1dbb66bcc83d3631d4a9a76fdf9988dc1371029ec
	//公钥full: 0xeb1c9274258695bb95e45e3e7ab24dfd7fceda48a044d5bea1c7e3ada6227f08
	//0x7ab24dfd7fceda48a044d5bea1c7e3ada6227f08
	//*/
	////3.crypto/ecdsa包并使用FromECDSA方法将其转换为Bytes字节
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//
	////4.使用步骤3的私钥的Bytes字节，生成签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为拥有它的人可以访问你的所有资产
	//// 称为钱包密码
	//// hexutil转换为16进制字符串
	//// Encode方法转成切片
	//// [2:],去掉头部0x,即为钱包密码
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) //钱包密码 （256位）：1372a42dc6a042fd047deac1dbb66bcc83d3631d4a9a76fdf9988dc1371029ec
	//
	////5.（私钥初始化生成时会有一个公钥方法publicPub）使用私钥派生公钥，作为钱包地址
	//publicKey := privateKey.Public()
	////6. 将其转成16进制字符串的步骤，6.1先转成指针，6.2生成bytes字节 6.3 再hash 6.4HexUtil转成16进制字符串 6.5截取 || 或直接6.1 publicKey.(*ecdsa.PublickKey)转成指针，6.2 PublicKeyToAddress转成公钥
	////6.1先转成指针
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	////6.2FromECDSAPub生成bytes字节
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	////6.3 再hash
	////1.创建一个 Keccak-256 哈希函数实例。
	//hash := sha3.NewLegacyKeccak256()
	//////2.使用会把这 64 字节数据输入到 Keccak-256 哈希函数中
	//hash.Write(publicKeyBytes[1:]) //去掉0x40 再bytes中表示1个字节
	////6.4 hexutil.Encode
	//fmt.Println("公钥full:", hexutil.Encode(hash.Sum(nil)[:]))
	////6.5 使用hexutil.Encode成字符串切片截取12位
	//fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位  公钥地址：0x7ab24dfd7fceda48a044d5bea1c7e3ada6227f08
	//// 上面的所有步骤，可以用PubkeyToAddress封装方法一步到位生成公钥地址
	////fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //0x7AB24DfD7Fceda48A044D5BEA1C7E3ada6227f08
	////fmt.Println(fromAddress)

}

// ETH 转账
func main_transETH() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")

	if err != nil {
		log.Fatal(err)
	}

	//通过私钥获取到钱包公钥地址
	privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()                   //
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // 默认生成Account1账户地址:0x52C49173a0a5C824cD6504275EAD5E602425be02
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//交易前需要拿到当前上下文fromAddress钱包的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	//设置转账ETH
	value := big.NewInt(10000000000000000) //0.01ETH
	gasLimit := uint64(21000)              //纯ETH转账消耗gas 为：210000
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	//转入账户
	toAddress := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")
	var data []byte
	//纯ETH data为nil
	//1、发起转账
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	// 2、使用转账账户的私钥对未签名的交易事务进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}
	//3、过在client实例调用SendTransaction来将已签名的事务广播到整个网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent:%s", signedTx.Hash().Hex())

}

func main_transERC20() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")

	if err != nil {
		log.Fatal(err)
	}
	//1.钱包私钥派生账户公钥地址
	publicKey := privateKey.Public()
	publicKeyECDA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDA)
	//2.拿到当前公钥地址的交易nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//3.ETH设置为0
	value := big.NewInt(0)

	//4.网上获取gasPrice最新价格，和gasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前区块的gas限制
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置合理的gas限制（通常使用区块限制的50-80%）
	gasLimit := uint64(float64(header.GasLimit) * 0.8)

	// 5.发送地址，已经合约地址
	toAddress := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")
	tokenAddress := common.HexToAddress("0x038aA69737f6F284f82E43654D71BA5DaB9911c3")

	//5.生成函数签名
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	menthodID := hash.Sum(nil)[:4]

	//6.ERC20的data不能为空，需要拼接data
	//参数toAddress
	paddedtoAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//传送的
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10) //1000tokens

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, menthodID...)
	data = append(data, paddedtoAddress...)
	data = append(data, paddedAmount...)

	//7.发起ERC20转账
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	//获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//发送广播
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent:%s", signedTx.Hash().Hex())
}

func main_Test() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	blockNumber := big.NewInt(9289667)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}

func main_setString() {
	f := new(big.Float)

	// 支持的各种格式
	formats := []string{
		"3.14159",    // 常规小数
		"-123.456",   // 负数
		"1.23e10",    // 科学计数法（大写E）
		"1.23e-5",    // 负指数
		"0.000123",   // 小数
		"1234567890", // 大整数
		"0",          // 零
		"-0",         // 负零
	}

	for _, str := range formats {
		if _, success := f.SetString(str); success {
			fmt.Printf("成功解析: %-15s -> %s\n", str, f.String())
		} else {
			fmt.Printf("解析失败: %s\n", str)
		}
	}
}

func main_查询代币余额() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}
	// Golem (GNT) Address
	// 合约地址
	tokenAddress := common.HexToAddress("0x038aA69737f6F284f82E43654D71BA5DaB9911c3")
	// 生成代币合约实例，可以调用合约的方法
	instance, err := token.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	// 账户地址
	address := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")
	/**
	&bind.CallOpts{} 是 go-ethereum 库中用于只读合约调用的配置选项。

	1. 基本概念
	CallOpts vs TransactOpts
	类型	用途	Gas费用	区块链状态
	CallOpts	只读调用	免费	不修改状态
	TransactOpts	写入交易	需要Gas	修改状态
	*/
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	// 进行decimal单位换算
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals)))) // fbal / 10**18

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

	approveInt := new(big.Int)
	approveInt.SetString("100", 10)
	result, err := instance.Approve(&bind.TransactOpts{}, tokenAddress, approveInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func main_subcribe() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			//heardHas := common.HexToHash("0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f")
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}

// 部署合约
func main_depoly() {
	//两种方式部署合约
	//1.abigen工具，abi文件+bin文件部署
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}
	//私钥派生公钥钱包地址
	privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取&bin.Transport()进行交易操作，消耗gas
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	//填写部署配置
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
	//2.二进制文件部署
	//const (
	//	// store合约的字节码
	//	contractBytecode = "608060405234801561000f575f5ffd5b5060405161087838038061087883398181016040528101906100319190610193565b805f908161003f91906103ea565b50506104b9565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f5f905090565b61032f610320565b61033a8184846102fb565b505050565b5b8181101561035d576103525f82610327565b600181019050610340565b5050565b601f8211156103a25761037381610241565b61037c84610253565b8101602085101561038b578190505b61039f61039785610253565b83018261033f565b50505b505050565b5f82821c905092915050565b5f6103c25f19846008026103a7565b1980831691505092915050565b5f6103da83836103b3565b9150826002028217905092915050565b6103f3826101da565b67ffffffffffffffff81111561040c5761040b61006f565b5b6104168254610211565b610421828285610361565b5f60209050601f831160018114610452575f8415610440578287015190505b61044a85826103cf565b8655506104b1565b601f19841661046086610241565b5f5b8281101561048757848901518255600182019150602085019450602081019050610462565b868310156104a457848901516104a0601f8916826103b3565b8355505b6001600288020188555050505b505050505050565b6103b2806104c65f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f5ffd5b61005d600480360381019061005891906101d7565b6100ad565b60405161006a9190610211565b60405180910390f35b61007b6100c2565b604051610088919061029a565b60405180910390f35b6100ab60048036038101906100a691906102ba565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610325565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051610194929190610355565b60405180910390a15050565b5f5ffd5b5f819050919050565b6101b6816101a4565b81146101c0575f5ffd5b50565b5f813590506101d1816101ad565b92915050565b5f602082840312156101ec576101eb6101a0565b5b5f6101f9848285016101c3565b91505092915050565b61020b816101a4565b82525050565b5f6020820190506102245f830184610202565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026c8261022a565b6102768185610234565b9350610286818560208601610244565b61028f81610252565b840191505092915050565b5f6020820190508181035f8301526102b28184610262565b905092915050565b5f5f604083850312156102d0576102cf6101a0565b5b5f6102dd858286016101c3565b92505060206102ee858286016101c3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033c57607f821691505b60208210810361034f5761034e6102f8565b5b50919050565b5f6040820190506103685f830185610202565b6103756020830184610202565b939250505056fea26469706673582212209ed396d79b52f8a99904c38c2f0aafe8f8863de501c2afead75e24bf8084c95064736f6c634300081e0033"
	//)
	//// 连接到以太坊网络（这里使用 Goerli 测试网络作为示例）
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//////私钥派生公钥钱包地址
	//privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	////获取nonce
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////获取网上建议的gas价格
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice = big.NewInt(int64(64579546054))
	////设置gasLimit
	//header, err := client.HeaderByNumber(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 设置合理的gas限制（通常使用区块限制的50-80%）
	//gasLimit := uint64(float64(header.GasLimit) * 0.8)
	//
	////解码合约字节码
	//data, err := hex.DecodeString(contractBytecode)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////创建交易
	//tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, data)
	//
	////获取当前链ID
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	////交易签名
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//
	////发送交易
	//err = client.SendTransaction(context.Background(), signedTx)
	//
	//fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	//
	////广播，等待被挖矿，生成块上链
	//receipt, err := waitForReceipt(client, signedTx.Hash())
	//fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex())

}

//func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
//	for {
//		receipt, err := client.TransactionReceipt(context.Background(), txHash)
//		if err == nil {
//			return receipt, nil
//		}
//		if err != ethereum.NotFound {
//			return nil, err
//		}
//		// 等待一段时间后再次查询
//		time.Sleep(1 * time.Second)
//	}
//}

// // 加载合约
// 0x1b45344Ab499bBeeEbEbEE8CA2c52f5B393d4DC6
// 0x889161d35d93f685e5ba4596365fbacab2fa3c2bd3f5e6c92c14025006ed88a1
func main_loadContract() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}
	const (
		contractAddr = "0x1b45344Ab499bBeeEbEbEE8CA2c52f5B393d4DC6"
	)
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	_ = storeContract
	version, err := storeContract.Version(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("version: %d\n", version)
}

// 执行合约
/**
有4种方式
1、使用工具生成Go合约代码
2、使用ethclient库
2.1、使用abi文件签约合约
2.2、不使用abi文件签约合约
3、remix手动执行
*/
const (
	contractAddr = "0x1b45344Ab499bBeeEbEbEE8CA2c52f5B393d4DC6"
)

func main_执行合约() {
	////1、使用工具生成Go合约代码
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////拿到私钥实例签约，初始化交易opt实例
	//var key [32]byte
	//var value [32]byte
	//
	//copy(key[:], []byte("demo_save_key"))
	//copy(value[:], []byte("demo_save_value11111"))
	//
	//chainID, err := client.NetworkID(context.Background())
	//fmt.Println(chainID)
	//opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) //消耗gas费用
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tx, err := storeContract.SetItem(opt, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(tx.Hash().Hex())
	//
	////生成只读调用callOpts 不消耗gas
	//callOpt := &bind.CallOpts{Context: context.Background()}
	//valueInContract, err := storeContract.Items(callOpt, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////fmt.Println(string(valueInContract[:]))
	//fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)

	//2.1、使用abi文件签约合约
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 获取公钥地址
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	//// 获取 nonce
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 估算 gas 价格
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	////准备交易数据
	//contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	//if err != nil {
	//	log.Fatal(err)
	//}
	////交易数据配置
	//methodName := "setItem"
	//var key [32]byte
	//var value [32]byte
	//copy(key[:], []byte("demo_save_key_use_abi"))
	//copy(value[:], []byte("demo_save_value_use_abi_11111"))
	////生成带参数的选择器
	//input, err := contractABI.Pack(methodName, key, value)
	////发起交易前，先获取链ID
	//chain, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//toAddress := common.HexToAddress(contractAddr)
	////发起交易，未签名阶段，需调用签名
	//tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), 300000, gasPrice, input)
	//
	////调用签名
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chain), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 发送交易，广播
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(signedTx.Hash().Hex())
	////等待收据
	//_, err = waitForReceipt(client, signedTx.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	////接收到收据后，再发起查询，查询刚刚设置的值
	//methodName2 := "items"
	//callInput, err := contractABI.Pack(methodName2, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//callMsg := ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: callInput}
	//
	////解析返回值
	//result, err := client.CallContract(context.Background(), callMsg, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var unpacked [32]byte
	//contractABI.UnpackIntoInterface(&unpacked, "items", result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("is value saving in contract equals to origin value:", unpacked == value)

	//2.2不使用abi文件签约合约
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 估算 gas 价格 ，EIP155使用兼容所有网络，但是gas费用不明确，没有burn
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//准备交易数据
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)
	//发起交易前，先获取链ID
	chain, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(contractAddr)
	//发起交易，未签名阶段，需调用签名
	//1.使用固定发起交易
	//tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), 300000, gasPrice, input)

	//2.使用动态发起交易
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 2.1设置GasFeeCap（总费用上限）
	header, err := client.HeaderByNumber(context.Background(), nil)
	gasFeeCap := new(big.Int).Add(gasTipCap,
		new(big.Int).Mul(header.BaseFee, big.NewInt(1))) //2倍费用，如果增加多次倍数后，还是发现交易仍在pending状态，则需要考虑是否nonce连续。

	// 验证费用关系
	if gasTipCap.Cmp(gasFeeCap) > 0 {
		fmt.Errorf("小费(%s)不能高于总费用上限(%s)", gasTipCap.String(), gasFeeCap.String())
		return
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chain,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       uint64(300000),
		To:        &toAddress,
		Value:     big.NewInt(0),
		Data:      input,
	})
	fmt.Println(tx.Hash().Hex())
	//出现transaction type not supported，异常的原因是调用签名时NewEIP155Signer不支持，需要换成NewLondonSigner
	//调用签名
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chain), privateKey)
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chain), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易，广播
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signedTx.Hash().Hex())
	//等待收据
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	//接收到收据后，再发起查询，查询刚刚设置的值
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	//解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)

	/**
	signedTx 已签名交易查询
	交易哈希：
	0x0f1f1d471cdf0e5e113bcc5eb2c7473879cf786c521e6388dc42c92ad8668585
	地位：
	成功
	堵塞：
	9297287
	1 区块确认
	时间戳：
	16 秒前 ( 2025 年 9 月 28 日 09:16:36 AM UTC )
	从：
	0x52C49173a0a5C824cD6504275EAD5E602425be02
	到：
	0x1b45344Ab499bBeeEbEbEE8CA2c52f5B393d4DC6
	价值：
	0 ETH
	交易费：
	0.000000026421317052以太币
	汽油价格：
	0 . 001000012 Gwei (0 . 000000000001000012 ETH)
	*/

}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}

var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func main_事件() {
	//0xD813af45a7b6807A7077BB61021EA80844686a3C
	//0x19b81d4bc7d47920b8b4d43543c1f5690d6769ae8f70287dcf2e10d7216b2b68
	//方法1：
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xD813af45a7b6807A7077BB61021EA80844686a3C")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(common.Bytes2Hex(event.Key[:]))
		fmt.Println(common.Bytes2Hex(event.Value[:]))
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}
		fmt.Println("topics[0]=", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}

	//方法2：通过订阅合约事件方式，获取监听事件的值
	//client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/4803e056901a43e39b0c47c7c2602648")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contractAddress := common.HexToAddress("0xD813af45a7b6807A7077BB61021EA80844686a3C")
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{contractAddress},
	//}
	//logs := make(chan types.Log)
	//sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	select {
	//	case err := <-sub.Err():
	//		log.Fatal(err)
	//	case vLog := <-logs:
	//		fmt.Println(vLog.BlockHash.Hex())
	//		fmt.Println(vLog.BlockNumber)
	//		fmt.Println(vLog.TxHash.Hex())
	//		event := struct {
	//			Key   [32]byte
	//			Value [32]byte
	//		}{}
	//		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		fmt.Println(common.Bytes2Hex(event.Key[:]))
	//		fmt.Println(common.Bytes2Hex(event.Value[:]))
	//		var topics []string
	//		for i := range vLog.Topics {
	//			topics = append(topics, vLog.Topics[i].Hex())
	//		}
	//		fmt.Println("topics[0]=", topics[0])
	//		if len(topics) > 1 {
	//			fmt.Println("index topic:", topics[1:])
	//		}
	//	}
	//}
}

type ethClient struct {
	client *ethclient.Client
}

// 获取balance
func (c *ethClient) mainEth_getBalance(fromAddress [20]byte) {
	//获取A账户的余额
	fromBalance, err := c.client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fromBalance=", fromBalance)
	toAddress := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")

	toBalance, err := c.client.BalanceAt(context.Background(), toAddress, nil)
	fmt.Println("toBalance=", toBalance)
}

// 获取store存储的值，通过key
func (c *ethClient) mainEth_getStorageAt(account common.Address, key [32]byte) []byte {
	sotreValue, err := c.client.StorageAt(context.Background(), account, key, nil)
	if err != nil {
		log.Fatal(err)
	}
	return sotreValue
}

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4803e056901a43e39b0c47c7c2602648")
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.HexToECDSA("121a0bb66540275df6c3e5cc315e4049df250bd589633bb5c607bda418f4a848")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 获取公钥地址
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//ec := &ethClient{client: client}
	////查询余额
	//ec.mainEth_getBalance(fromAddress)
	//从给定地址的存储位置返回值
	//contractAddrs := common.HexToAddress("0xDbf53929f1091deC15AE2a3092C9623C9BC93Eb2")
	//key := common.HexToHash("name")
	//
	//itemsSlot := big.NewInt(0) // 假设 items 映射在存储槽 1
	//
	//// 计算映射项的存储位置
	//mapPosition := crypto.Keccak256(key.Bytes(), common.LeftPadBytes(itemsSlot.Bytes(), 32))
	//mapStorageKey := common.BytesToHash(mapPosition)
	//itemData, err := client.StorageAt(context.Background(), contractAddrs, mapStorageKey, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Item data: %x\n", itemData)
	//
	//count, err := client.PendingNonceAt(context.Background(), contractAddrs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("PendingNonceAt(address): %d\n", count)
	//
	//code, err := client.CodeAt(context.Background(), contractAddrs, nil)
	//fmt.Printf("Item data: %x\n", code)

	//返回匹配给定区块哈希的区块中的交易数量
	//block, err := client.BlockByNumber(context.Background(), nil)
	//blockCount, err := client.TransactionCount(context.Background(), block.Hash())
	//fmt.Println(blockCount)

	//返回哈希中处于pedding的交易数量
	blockCount, err := client.PendingTransactionCount(context.Background())
	fmt.Print(blockCount)
	//返回匹配给定区块哈希的区块中的叔块数量。

	//拼装数据
	//gasTipCap, err := client.SuggestGasTipCap(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//header, err := client.HeaderByNumber(context.Background(), nil)
	//gasFeeCap := new(big.Int).Add(gasTipCap,
	//	new(big.Int).Mul(header.BaseFee, big.NewInt(1))) //2倍费用，如果增加多次倍数后，还是发现交易仍在pending状态，则需要考虑是否nonce连续。
	//
	//// 验证费用关系
	//if gasTipCap.Cmp(gasFeeCap) > 0 {
	//	fmt.Errorf("小费(%s)不能高于总费用上限(%s)", gasTipCap.String(), gasFeeCap.String())
	//	return
	//}
	//chain, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////转账不需要拼装数据，直接input的data为nil
	////合约地址 或对应的EOA账户地址
	//toAddress := common.HexToAddress("0xbC4166AeC87Ef009245BBa2E8Ac4848e79a565Cf")
	//tx := types.NewTx(&types.DynamicFeeTx{
	//	ChainID:   chain,
	//	Nonce:     nonce,
	//	GasTipCap: gasTipCap,
	//	GasFeeCap: gasFeeCap,
	//	Gas:       uint64(21000),
	//	To:        &toAddress,
	//	Value:     big.NewInt(10000000000000000),
	//	Data:      nil,
	//})
	//signedTx, err := types.SignTx(tx, types.NewLondonSigner(chain), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(signedTx.Hash().Hex())

	//获取B账户的余额
}
