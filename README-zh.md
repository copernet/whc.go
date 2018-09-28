whc.go
---

whc.go提供了请求wormhole客户端RPC的Golang版SDK（包含bitcoin-cash原始的RPC SDK）。

#### 安装:

- 安装wormhole客户端（版本: [Earth-0.1.0](https://github.com/copernet/wormhole/releases)），安装文档请参考：[https://github.com/copernet/wormhole/tree/master/doc](https://github.com/copernet/wormhole/tree/master/doc)。

- [Go](http://golang.org/) 版本需要在1.8以上。

- 安装:

  - 安装方法一：

    ```
    go get github.com/copernet/whc.go
    ```

  - 安装方法二：
    1. 克隆本仓库：

      ```
      git clone https://github.com/copernet/whc.go.git $GOPATH/src/github.com/copernet/whc.go
      ```

    2. 安装依赖：

      ```
      go get github.com/btcsuite/go-socks
      go get github.com/btcsuite/websocket
      go get github.com/btcsuite/btclog
      go get github.com/bcext/gcash
      go get github.com/bcext/cashutil
      ```


#### 使用:

1. 安装并运行wormhole客户端，配置`rpcuser`、`rpcpassword`和`rpcallowip`选项。下面是一个简单的wormhole配置示例，通常位于bitcoin.conf文件中。

   ```
   rpcuser=D313FF53C1
   rpcpassword=E51DD672DED3FA7D3FB00E946C10C681EF164
   # server ip of your project using this sdk
   rpcallowip=54.78.09.134
   rpcport=18332
   startclean=1
   ```

2. 创建与wormhole客户端的连接：

   ```
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
   ```

3. 使用示例:

   - 获取wormhole客户端区块高度：
        ```
        // Get the current wormhole relative information.
        info, err := client.WhcGetInfo()
        if err != nil {
        	log.Fatal(err)
        }
        
        log.Printf("Block count: %d", info.Block)
        
        ```

   - 获取wormhole系统中的资产列表：

        ```
        // Fetch all wormhole properties.
        list, err := client.WhcListProperties()
        if err != nil {
        	log.Fatal(err)
        }
        
        log.Println(list)
        ```

   - 发起众筹：

     ```
     // necessary field
     addressFrom := "bchtest:qqg2fwfzd4xeywf8h2zajqy77357gk0v7yvsvhd4xu"
     ecosystem := int64(1) // must be 1
     precision := int64(8)
     previousID := int64(0) // must be 0
     category := "Blockchain research"
     subCatetory := "Bitcoin cash"
     name := "wormhole"
     url := "https://www.wormhole.cash"
     data := "working for the future"
     desiredID := int64(1) // must be 1
     tokensPerUnit := "100"
     deadline := int64(1582772366)
     earlyBonus := int64(2)
     issuerPercentage := int64(0) // must be 0
     amount := "10000000.987"
     
     // Create a crowdsalle transaction
     txHash, err := client.WhcSendIssuanceCrowdSale(addressFrom, ecosystem, precision, previousID, category,subCatetory, name, url, data, desiredID, tokensPerUnit, deadline, earlyBonus, issuerPercentage, amount)
     
     log.Println(txHash)
     ```

    >  注意wormhole客户端当前只支持HTTP POST请求方式。完整的示例代码移至[examples](https://github.com/copernet/whc.go/blob/master/examples). 更多的API请参考以下内容. 如果发现任何bug和问题请提交 [issue](https://github.com/copernet/whc.go/issues/new)，开发者将会定期答复。

#### 接口列表:

以下接口涵盖所有wormhole相关功能。该SDK有意的被设计成单个package形式，方便导入使用。详细的wormhole json-rpc文档，请参考[https://github.com/copernet/spec/blob/master/wormhole-RPC.md](https://github.com/copernet/spec/blob/master/wormhole-RPC.md)

| #    | Method                            | Description                                      |
| ---- | --------------------------------- | ------------------------------------------------ |
| 1    | WhcSetAutoCommit                  | 设置是否自动提交一个交易                         |
| 2    | WhcGetActiveCrowd                 | 返回指定地址的活跃众筹                           |
| 3    | WhcGetAllBalancesForAddress       | 返回一个地址所有的token余额信息                  |
| 4    | WhcGetAllBalancesForID            | 返回一个token的余额信息列表                      |
| 5    | WhcGetBalances                    | 返回一个给定地址和token的余额信息                |
| 6    | WhcGetBalanceHash                 | 返回指定token的余额信息hash                      |
| 7    | WhcGetCrowdSale                   | 返回指定众筹详情                                 |
| 8    | WhcGetCurrentConsensusHash        | 返回当前区块的余额信息hashs                      |
| 9    | WhcSendSto                        | 创建并广播空投交易                               |
| 10   | WhcGetGrants                      | 返回可管理资产的增发、销毁信息                   |
| 11   | WhcGetInfo                        | 返回当前客户端和协议相关信息                     |
| 12   | WhcGetPayload                     | 返回wormhole交易的payload信息                    |
| 13   | WhcGetProperty                    | 返回token信息                                    |
| 14   | WhcGetSeedBlocks                  | 返回指定区块区间的含有wormhole交易的区块         |
| 15   | WhcGetSto                         | 返回空投交易的详情                               |
| 16   | WhcGetTransaction                 | 返回wormhole交易的详情(不同交易类型返回字段不同) |
| 17   | WhcListBlockTransactions          | 返回给定区块所含有的wormhole交易列表             |
| 18   | WhcListPendingTransactions        | 返回内存池中未被确认的wormhole交易列表           |
| 19   | WhcListProperties                 | 返回当前系统中所有wormhole相关的资产列表         |
| 20   | WhcListTransactions               | 返回指定条件的wormhole交易列表                   |
| 21   | WhcCreatePayloadBurnBCH           | 创建获取基础货币WHC的交易payload                 |
| 22   | WhcCreatePayloadChangeIssuer      | 创建改变发行者交易的payload                      |
| 23   | WhcCreatePayloadCloseCrowdSale    | 创建手动关闭众筹交易的payload                    |
| 24   | WhcCreatePayloadGrant             | 创建增发交易的payload                            |
| 25   | WhcCreatePayloadIssuanceCrowdSale | 创建众筹交易的payload                            |
| 26   | WhcCreatePayloadIssuanceFixed     | 创建固定资产交易的payload                        |
| 27   | WhcCreatePayloadIssuanceManaged   | 创建可管理资产交易的payload                      |
| 28   | WhcCreatePayloadPartiCrowdSale    | 创建参与众筹交易的payload                        |
| 29   | WhcCreatePayloadRevoke            | 创建销毁token交易的payload                       |
| 30   | WhcCreatePayloadSendAll           | 创建发送所有token到指定地址交易的payload         |
| 31   | WhcCreatePayloadSimpleSend        | 创建转账交易的payload                            |
| 32   | WhcCreatePayloadSto               | 创建空投交易的payload                            |
| 33   | WhcCreateRawTxChange              | 添加一个找零输出                                 |
| 34   | WhcCreateRawTxInput               | 添加一个交易输入                                 |
| 35   | WhcCreateRawTxOpReturn            | 添加含有payload的输出                            |
| 36   | WhcCreateRawTxReference           | 添加一个找零输出                                 |
| 37   | WhcDecodeTransaction              | 解码wormhole交易                                 |
| 38   | WhcBurnBCHGetWhc                  | 获取wormhole基础货币WHC                          |
| 39   | WhcPartiCrowdSale                 | 创建并广播参与众筹的交易                         |
| 40   | WhcSend                           | 创建并广播转账交易                               |
| 41   | WhcSendAll                        | 创建并广播发送所有token到指定地址的交易          |
| 42   | WhcSendChangeIssuer               | 创建并广播改变发行者交易                         |
| 43   | WhcSendCloseCrowdSale             | 创建并广播关闭众筹交易                           |
| 44   | WhcSendGrant                      | 创建并广播增发交易                               |
| 45   | WhcSendIssuanceCrowdSale          | 创建并广播众筹交易                               |
| 46   | WhcSendIssuanceFixed              | 创建并广播固定数量token交易                      |
| 47   | WhcSendIssuanceManaged            | 创建并广播管理资产交易                           |
| 48   | WhcSendRawTx                      | 广播wormhole交易                                 |
| 49   | WhcSendRevoke                     | 创建并广播销毁管理资产交易                       |

当前所有的API均支持同步请求和异步请求方式，并且所有异步请求API的方法名均以 `Async`结尾。在许多使用场景中，使用异步模式是非常高效的。下面是一个使用异步模式请求API的例子:

```
// Notice the notification parameter is nil since notifications are
// not supported in HTTP POST mode.
client, err := rpcclient.New(connCfg, nil)
if err != nil {
	log.Fatal(err)
}
defer client.Shutdown()

// 异步模式
r := client.WhcGetCrowdSaleAsync(34, nil)
// 这里执行一个耗时的任务
// ...
// 接收结果
result, err := r.Receive()

// 同步模式
r := client.WhcGetCrowdSale(34, nil)
```
