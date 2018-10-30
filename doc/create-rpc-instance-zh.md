### RPC客户端实例化:

文档中使用的所有 `client` 对象都是经过如下的初始化之后才能使用的，推荐使用单例模式获取RPC客户端实例化对象。

RPC客户端对象实例化示例：

```Go
package rpc

import (
	"log"

	"github.com/copernet/whc.go/rpcclient"
)

var client *rpcclient.Client

func NewRPCInstance() *rpcclient.Client {
	if client != nil {
		return client
	}

	// Connect to local wormhole RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig {
		Host:         "39.104.152.31",
		User:         "5xX57X3Ad7KdjUv/lYFR",
		Pass:         "cz8S5tM/BUdQv6tyQdET",
		HTTPPostMode: true, 	// wormhole only supports HTTP POST mode
		DisableTLS:   true, 	// wormhole does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
```

> 您的项目中应该包含上面的RPC客户端实例化代码。SDK仅仅提供实例化的方法: `rpcclient.New(connCfg, nil)`。

