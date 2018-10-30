### SDK 依赖:

- Go 安装步骤请参考：https://golang.org/dl/

- wormhole client：安装步骤请参考：https://github.com/copernet/wormhole/tree/master/doc#building

  > 需要注意的是在安装完wormhole客户端之后，需要同步wormhole全部区块链数据，其同步时间视电脑的硬件配置和网络情况而定，一般在1~3天时间不等。

### 安装:

使用`go get`命令安装：

 ```
go get github.com/copernet/whc.go
 ```

### 使用:

引入whc.go到您的Go项目中：

```
import "github.com/copernet/whc.go/rpcclient"
```

