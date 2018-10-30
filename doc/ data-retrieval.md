##### `WhcGetAllBalancesForAddress`

获取指定地址所有种类的token金额

**Arguments**

1. address `string` required

##### 返回值

##### `array`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetAllBalancesForAddress("bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5")

// result:
[
    {
      "propertyid": 1,
      "balance": "127.88577709",
      "reserved": "0.00000000"
    },
   ...
]
```

------

##### `WhcGetAllBalancesForID`

获取wormhole系统中含有指定token的所有地址及金额信息

**Arguments**

1. property id `uint64`

##### 返回值

##### `array`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetAllBalancesForID(34)

// result:
[
    {
      "address": "bchtest:qqwrj8yer7us830ca4y7fw89q24gh3cu8u8a7z489j",
      "balance": "99999999999999.9",
      "reserved": "0.0"
    }
]
```

------

##### `WhcGetBalances`

获取指定地址上指定token的金额信息

**Arguments**

1. address `uint64`
2. property id `uint64`

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetBalances("bchtest:qqwrj8yer7us830ca4y7fw89q24gh3cu8u8a7z489j", 34)

// result:
Result: 
{
    "balance": "99999999999999.9",
    "reserved": "0.0"
}
```

------

##### `WhcGetBalanceHash`

获取本节点当前高度下指定token的状态哈希

**Arguments**

1. property id `uint64`

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetBalanceHash(34)

// result:
{
    "block": 1265214,
    "blockhash": "00000000f519e18cf3c9bfdd80629bfc32c3b953f6abc1376de9766e1eb2825e",
    "propertyid": 34,
    "balanceshash": "03fd20f7b9292e88ac6f5d57d167189e1576aa1cbeeca8ad68d4303eb22e0e24"
}
```

------

#####   `WhcGetCrowdSale`

获取众筹token的详细信息

**Arguments**

1. property id `uint64`
2. verbose `pointer`

##### 返回值

##### `pointer`

##### 示例

```Go
var verbose = true
rpc.NewRPCInstance().WhcGetCrowdSale(34, &verbose)

// result:
{
    "propertyid": 30,
    "name": "qshuaifinalcrowsale",
    "active": true,
    "issuer": "bchtest:qque8sgqke0w94sjmkqdhs9zwvdlgepgnvu2nfkguy",
    "propertyiddesired": 1,
    "precision": "2",
    "tokensperunit": "10.00000000",
    "earlybonus": 10,
    "starttime": 1533018995,
    "deadline": 1659199632,
    "amountraised": "34.99981568",
    "tokensissued": "12345.12",
    "addedissuertokens": "0.00",
    "closetx": null,
    "participanttransactions": [
      {
        "txid": "f71b2b538592a965a2b20655af5d44ebdf3ba8b93671ca77a61045dec83de6e4",
        "amountsent": "0.99997384",
        "participanttokens": "217.34"
      },
      ...
    ]
}
```

------

##### `WhcGetCurrentConsensusHash`

获取本节点当前高度下wormhole系统的状态哈希

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetCurrentConsensusHash()

// result:
{
    "block": 1265214,
    "blockhash": "00000000f519e18cf3c9bfdd80629bfc32c3b953f6abc1376de9766e1eb2825e",
    "consensushash": "16dfec22f81abcaf3824a9e8fcb08044c287e597233a18645b03e57d60bd8ce5"
}
```

------

##### `WhcGetGrants`

获取指定管理token的增发，销毁信息

**Arguments**

1. property id `uint64`

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetGrants(3)

// result:
{
    "propertyid": 3,
    "name": "test_token1",
    "issuer": "bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5",
    "creationtxid": "1c3f95acbd6eb38e2a7c26b12dc9138b4523c355a20944874bdc3c82f4c5e4e1",
    "totaltokens": "90",
    "issuances": [
      {
        "txid": "a1f2d0126a04296aad6f492a0ef8c1c1afb781efc6c5f37de105790a7debcf87",
        "revoke": "10"
      },
      {
        "txid": "6afffd7d14060b6e79c504c5f17596616eb99356b71d301b5b37df1df065b9a0",
        "grant": "100"
      }
    ]
}
```

------

#####  `WhcGetInfo`

获取wormhole节点的基础信息

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetInfo()

// result:
Result: 
{
    "wormholeversion_int": 10000000,
    "wormholeversion": "0.1.0",
    "bitcoincoreversion": "0.17.2",
    "block": 1265220,
    "blocktime": 1540877466,
    "blocktransactions": 0,
    "totaltrades": 0,
    "totaltransactions": 2704,
    "alerts": []
}
```

------

##### `WhcGetPayload`

获取指定交易的wormhole载荷数据

**Arguments**

1. tx hash `string`

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetPayload("6579dea76c3d0b4463671c5476f90f20c746992c300a4b8ec4ce6748c0960836")

// result:
{
    "payload": "000000360100020000000074657374206d616e6167656420746f6b656e2032007465737400746573745f746f6b656e31007777772e746573746d616e61676564746f6b656e2e636f6d006d79206461746100",
    "payloadsize": 82
}
```

------

##### `WhcGetProperty`

获取指定token的信息

**Arguments**

1. property id `uint64`

##### 返回值

##### `pointer`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetProperty(123)

// result:
{
    "propertyid": 123,
    "name": "212",
    "category": "",
    "subcategory": "",
    "data": "1",
    "url": "1",
    "precision": 1,
    "issuer": "bchtest:qq6qag6mv2fzuq73qanm6k60wppy23djnv7ddk3lpk",
    "creationtxid": "5b757ee98840b97ef14445fe25ace4dc5f2aba9eb2d447e7bc886db20bfa65b7",
    "fixedissuance": true,
    "managedissuance": false,
    "freezingenabled": null,
    "totaltokens": "1.0"
}
```

------

##### `WhcListTransactions`

列出与节点钱包中的wormhole交易

**Arguments**

1. address `address` optinal
2. count `int64` optional
3. skip `int64` optional
4. startblock `int64` optional
5. endblock `int64` optional

##### 返回值

##### `array`

##### 示例

```json
height := 1437719290
rpc.NewRPCInstance().WhcListTransactions(nil, nil, nil, &height, nil)

// result:
[
    {
      "txid": "35f3210fd2963f901f7d49a4bff2a0f537ed9f27698c5372253a2e1d1d9f9a4c",
      "fee": "520",
      "sendingaddress": "bchtest:qqu9lh4jpc05p59pfhu9amyv9uvder8j3sa2up95vs",
      "referenceaddress": "bchtest:qqgrkkpyj050wxzj8vaca36nvyy9qgxm0s93z4fqac",
      "ismine": true,
      "version": 0,
      "type_int": 0,
      "type": "Simple Send",
      "block": 1263586,
      "confirmations": 1625,
      "valid": true,
      "blockhash": "00000000000003562e87cb521904cfe3022d8ee2d7caae111b6dcbd134cc51d3",
      "blocktime": 1540037998,
      "propertyid": 1,
      "precision": "8",
      "amount": "100.00000000",
      "subsends": null
    }
]
```

------

#####  `WhcGetTransaction`

获取指定交易的wormhole协议信息

**Arguments**

1. tx hash `string`

##### 返回值

##### `pointer`

##### 示例

```json
filter := "*"
rpc.NewRPCInstance().WhcGetTransaction("b319fecd3998dd689044c59d33d09b28654f31b8ae0123492b8a6e295467b6dc")

// result:
{
    "txid": "b319fecd3998dd689044c59d33d09b28654f31b8ae0123492b8a6e295467b6dc",
    "fee": "4555",
    "sendingaddress": "bchtest:qqwxfed2405mk3k4l82ejdt85s4ng8fmdyxye4rtwl",
    "ismine": false,
    "version": 0,
    "type_int": 3,
    "type": "Send To Owners",
    "block": 1249821,
    "confirmations": 15390,
    "valid": true,
    "blockhash": "000000000000323f30f14da5c5f15efa24bbc9d567b0b34070595398dc09c5f9",
    "blocktime": 1533205768,
    "propertyid": 1,
    "precision": "8",
    "amount": "10.00000000",
    "subsends": null
}
```

------

##### `WhcListPendingTransactions`

获取节点待确认的wormhole交易列表

**Arguments**

1. address `pointer`

##### 返回值

##### `array`

##### 示例

```json
filter := "*"
rpc.NewRPCInstance().WhcListPendingTransactions(1249129)

// result:
[
    {
      "txid": "e50c03aa0691c9a55872fa40da0acbab7b545015bed74e0e3231e80597fe27d6",
      "fee": "282",
      "sendingaddress": "bchtest:qrpdz05x8n4fxs80w07gsdxp5qe208kg6us3r2l27l",
      "referenceaddress": "bchtest:qrengre4vk4h5dd2uarnssmkqg43y0ep5cn7lf2w0w",
      "ismine": false,
      "version": 0,
      "type_int": 0,
      "type": "Simple Send",
      "propertyid": 1,
      "precision": "8",
      "amount": "0.00010000",
      "subsends": null
    }
]
```

------

##### `WhcListProperties`

列出wormhole系统中的所有token

##### 返回值

##### `array`

##### 示例

```json
filter := "*"
rpc.NewRPCInstance().WhcListProperties(1249129)

// result:
[
    {
      "propertyid": 1,
      "name": "WHC",
      "category": "N/A",
      "subcategory": "N/A",
      "data": "WHC serve as the binding between Bitcoin cash, smart properties and contracts created on the Wormhole.",
      "url": "http://www.wormhole.cash",
      "precision": 8
    },
    ...
]
```

------

##### 