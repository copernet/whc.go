whc.go
---

whc.go提供了wormhole客户端全部的RPC请求API(包含bitcoin-cash原始的RPC API). 使用当前wormhole sdk需要最新版本的wormhole客户端以及Go环境。

#### 安装依赖:

- 安装最新版本的wormhole客户端，安装文档请参考[https://github.com/copernet/wormhole/tree/master/doc](https://github.com/copernet/wormhole/tree/master/doc)， 由于不同的操作系统，安装方法有些差异，请参考与自己操作系统相同的安装文档，例如ubuntu安装文档需要参考:[https://github.com/copernet/wormhole/blob/master/doc/build-unix.md](https://github.com/copernet/wormhole/blob/master/doc/build-unix.md)

-  [Go](http://golang.org/) 语言版本需要在1.8以上.

- 安装:

  1. 克隆本仓库，推荐使用`go get github.com/copernet/whc.go`命令安装本仓库；这样在克隆本仓库同时也已安装完依赖。

     ```
     git clone https://github.com/copernet/whc.go.git $GOPATH/src/github.com/copernet/whc.go
     ```

  2. 安装依赖，如果使用`go get github.com/copernet/whc.go `安装完成本仓库，下面的操作是没有必要的。

     ```
     go get github.com/btcsuite/go-socks
     go get github.com/btcsuite/websocket
     go get github.com/btcsuite/btclog
     go get github.com/bcext/gcash
     go get github.com/bcext/cashutil
     ```

  3. 在你的工程中引入这个sdk，`import github.com/copernet/whc.go/rpcclient`. 下面是具体的使用示例。

#### 示例:

1. 现在需要一个安装并运行wormhole客户端的服务器，并且需要配置`rpcuser`和`rpcpassword`两个参数。如果工程项目和wormhole客户端不在同一台机器，需要配置`rpcallowip`选项。下面是一个简单的wormhole配置示例，通常位于bitcoin.conf文件中。

   ```
   rpcuser=D313FF53C1
   rpcpassword=E51DD672DED3FA7D3FB00E946C10C681EF164
   # server ip of your project using this sdk
   rpcallowip=54.78.09.134
   startclean=1
   ```

2. 在真实的项目中，API相关的配置项推荐的做法是使用配置文件。处于简单演示的目的，下面的示例直接将配置项写在了代码中。

   ```
   func main() {
   	// Connect to local wormhole RPC server using HTTP POST mode.
   	connCfg := &rpcclient.ConnConfig{
   		Host:         "127.0.0.1:18332",
   		User:         "D313FF53C1",
   		Pass:         "E51DD672DED3FA7D3FB00E946C10C681EF164",
   		HTTPPostMode: true, 		// wormhole only supports HTTP POST mode
   		DisableTLS:   true, 		// wormhole does not provide TLS by default
   	}
   	// Notice the notification parameter is nil since notifications are
   	// not supported in HTTP POST mode.
   	client, err := rpcclient.New(connCfg, nil)
   	if err != nil {
   		log.Fatal(err)
   	}
   	defer client.Shutdown()
   
   	// Get the current wormhole relative information.
   	info, err := client.WhcGetInfo()
   	if err != nil {
   		log.Fatal(err)
   	}
   	log.Printf("Block count: %d", info.Block)
   }
   
   // possible output:
   1251782
   ```

   >  注意wormhole客户端当前只支持HTTP POST请求模式。完整的示例代码移至[examples/whcinfo.go](https://github.com/copernet/whc.go/blob/master/examples/whcinfo.go). 更多的API请参考以下内容. 如果发现任何bug和问题请移至 [issue](https://github.com/copernet/whc.go/issues/new)。

#### 支持的API:

当前以下的API涵盖所有wormhole相关的功能，并于最新版本的wormhole保持同步。该SDK有意的被设计成单个package形式，可以被任何相关的工程所引用，方便查询和创建wormhole相关的交易。

|   #   |  Method    |   Description   |
| ---- | ---- | ---- |
| 1 |WhcSetAutoCommit| 设置时候自动提交一个交易 |
| 2 |WhcGetActiveCrowd| 返回指定地址的活跃众筹 |
| 3 |WhcGetAllBalancesForAddress| 返回一个地址所有的token余额信息 |
| 4 |WhcGetAllBalancesForID| 返回一个token的余额信息列表 |
| 5 |WhcGetBalances| 返回一个给定地址和token的余额信息 |
| 6 |WhcGetBalanceHash| 返回指定token的余额信息hash |
| 7 |WhcGetCrowdSale| 返回一个众筹详情 |
| 8 |WhcGetCurrentConsensusHash| 返回当前区块的余额信息hash |
| 9 |WhcSendSto| 创建并广播一个空投交易 |
| 10 |WhcGetGrants| 返回可管理资产的增发、销毁信息 |
| 11 |WhcGetInfo| 返回当前客户端和协议相关信息 |
| 12 |WhcGetPayload| 返回wormhole交易的payload信息 |
| 13 |WhcGetProperty| 返回token信息 |
| 14 |WhcGetSeedBlocks| 返回指定区块区间的含有wormhole交易的区块 |
| 15 |WhcGetSto| 返回空投交易的详情 |
| 16 |WhcGetTransaction| 返回wormhole交易的详情(不同交易类型返回字段不同) |
| 17 |WhcListBlockTransactions| 返回给定区块所含有的wormhole交易列表 |
| 18 |WhcListPendingTransactions| 返回内存池中未被确认的wormhole交易列表 |
| 19 |WhcListProperties| 返回当前系统中所有wormhole相关的资产列表 |
| 20 |WhcListTransactions| 返回指定条件的wormhole交易列表 |
| 21 |WhcCreatePayloadBurnBCH| 创建获取基础货币WHC的交易payload |
| 22 |WhcCreatePayloadChangeIssuer| 创建改变发行者交易的payload |
| 23 |WhcCreatePayloadCloseCrowdSale| 创建手动关闭众筹的交易payload |
| 24 |WhcCreatePayloadGrant| 创建增发交易的payload |
| 25 |WhcCreatePayloadIssuanceCrowdSale| 创建众筹交易的payload |
| 26 |WhcCreatePayloadIssuanceFixed| 创建固定资产的交易payload |
| 27 |WhcCreatePayloadIssuanceManaged| 创建可管理资产的交易payload |
| 28 |WhcCreatePayloadPartiCrowdSale| 创建参与众筹的交易payload |
| 29 |WhcCreatePayloadRevoke| 创建销毁token的交易payload |
| 30 |WhcCreatePayloadSendAll| 创建发送所有token到指定地址交易的payload |
| 31 |WhcCreatePayloadSimpleSend| 创建转账交易的payload |
| 32 |WhcCreatePayloadSto| 创建空投交易的payload |
| 33 |WhcCreateRawTxChange| 添加一个找零输出 |
| 34 |WhcCreateRawTxInput| 添加一个交易输入 |
| 35 |WhcCreateRawTxOpReturn| 添加含有payload的输出 |
| 36 |WhcCreateRawTxReference| 添加一个找零输出 |
| 37 |WhcDecodeTransaction| 解码wormhole交易 |
| 38 |WhcBurnBCHGetWhc| 获取wormhole基础货币WHC |
| 39 |WhcPartiCrowdSale| 创建并广播参与众筹的交易 |
| 40 |WhcSend| 创建并广播转账交易 |
| 41 |WhcSendAll| 创建并广播发送所有token到指定地址的交易 |
| 42 |WhcSendChangeIssuer| 创建并广播改变发行者交易 |
| 43 |WhcSendCloseCrowdSale| 创建并广播关闭众筹交易 |
| 44 |WhcSendGrant| 创建并广播增发交易 |
| 45 |WhcSendIssuanceCrowdSale| 创建并广播众筹交易 |
| 46 |WhcSendIssuanceFixed| 创建并广播固定数量token交易 |
| 47 |WhcSendIssuanceManaged| 创建并广播管理资产交易 |
| 48 |WhcSendRawTx| 广播wormhole交易 |
| 49 |WhcSendRevoke| 创建并广播销毁管理资产交易 |

当前所有的API均支持同步请求和异步请求方式，并且所有异步请求API的方法名均以 `Async`结尾。在许多使用场景中，使用异步模式是非常高效的。下面是一个使用异步模式请求API的例子:

```
// Notice the notification parameter is nil since notifications are
// not supported in HTTP POST mode.
client, err := rpcclient.New(connCfg, nil)
if err != nil {
	log.Fatal(err)
}
defer client.Shutdown()

r := client.WhcGetCrowdSaleAsync(34, nil)
// do some time-consuming code
time.Sleep(3 * time.Second)

result, err := r.Receive()
```

异步模式之所以更加高效，是因为在执行API请求的过程中，程序不会被阻塞，同时程序可以执行一个比较耗时的操作，在完成这个操作的之后，再接收API请求的结果。当然异步模式并不适合于所有使用场景。