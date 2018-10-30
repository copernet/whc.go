`WhcGetAllBalancesForAddress`

获取指定地址所有种类的token金额

**Arguments**

1. address `string` required

##### 返回值

##### `array`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetAllBalancesForAddress("bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5")

// result:
Result: 
[
    {PropertyID:1, Balance:127.88577709, Reserved:0.00000000},
    {PropertyID:3, Balance:80, Reserved:0}
]
```

------

`WhcGetAllBalancesForID`

获取wormhole系统中含有指定token的所有地址及金额信息

**Arguments**

1. property id `uint64`

##### 返回值

##### `array`

##### 示例

```Go
rpc.NewRPCInstance().WhcGetAllBalancesForID(34)

// result:
Result: 
[
    {Address:"bchtest:qqwrj8yer7us830ca4y7fw89q24gh3cu8u8a7z489j", Balance:99999999999999.9, Reserved:0.0}
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
&{Balance:99999999999999.9, Reserved:0.0}
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
Result: 
&{
    Block:1265108, 		     BlockHash:000000000000022adc8dc94b42bca3d1fd19ed7ff5ee4788fc8417871bba9b98,
    PropertyID:34, BalanceHash:03fd20f7b9292e88ac6f5d57d167189e1576aa1cbeeca8ad68d4303eb22e0e24
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
Result: 
&{PropertyId:30,
Name:qshuaifinalcrowsale,
Active:true,
Issuer:bchtest:qque8sgqke0w94sjmkqdhs9zwvdlgepgnvu2nfkguy,
PropertyIdDesired:1,
Precision:2,
TokensPerUnit:10.00000000,
EarlyBonus:10,
StartTime:1533018995,
DeadLine:1659199632,
AmountRaised:34.99981568,
TokensIssued:12345.12,
AddedIssuerTokens:0.00,
CloseDearly:<nil>,
MaxTokens:<nil>,
EndedTime:0,
TxidClosed:<nil>,
ParticipantTxs:[
{TxID:f71b2b538592a965a2b20655af5d44ebdf3ba8b93671ca77a61045dec83de6e4, Amount:0.99997384, ParticipantTokens:217.34} 		{TxID:89b435caef90f07afb4d1376dd66d2980931ac136eff4a226993de4fad5c1462, Amount:9.99998423, ParticipantTokens:2173.42} {TxID:6c94e33ab5f03788b2a9ad8573ca007339a8578658c886f5b034165589fd66ba, Amount:19.99996988, ParticipantTokens:4346.75} {TxID:f94f0863ba03ef252fcc821538e825665a598f0798f4d03a26380b9a732cdf10, Amount:0.99997398, ParticipantTokens:216.65} {TxID:c9650ca8eb4e5186eb56ead7f2b1b7ca32cee635af9ef087c7f7f9fe816cb058, Amount:0.99997387, ParticipantTokens:216.30} {TxID:ae7ce16077efc1e7241a8a35e699c87cb284761ff644c4a7d9bdc6594be31fc1, Amount:0.99997305, ParticipantTokens:216.18} {TxID:ba7df1bd956cd7c751b7abb7ed83f3a87c783226ca6d80d7020481b8e2772288, Amount:0.99996683, ParticipantTokens:214.21}
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
Result: 
&{Block:1265109, BlockHash:00000000000001e9f846c2818ca644e9148e52cae3198d05bc11c567f0740acb, ConsensusHash:739bc69e4743a1ff51c57df7ae7b75b8f2340c1724c4eb7dc70eb6cc3c9d8144}

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
Result: 
&btcjson.WhcGetGrantsResult{PropertyID:0x3, Name:"test_token1", Issuer:"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5", CreateTxID:"1c3f95acbd6eb38e2a7c26b12dc9138b4523c355a20944874bdc3c82f4c5e4e1", TotalTokens:"90", Issuance:[]btcjson.IssuanceTxs{btcjson.IssuanceTxs{TxID:"a1f2d0126a04296aad6f492a0ef8c1c1afb781efc6c5f37de105790a7debcf87", Grant:"", Revoke:"10"}, btcjson.IssuanceTxs{TxID:"6afffd7d14060b6e79c504c5f17596616eb99356b71d301b5b37df1df065b9a0", Grant:"100", Revoke:""}}}
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
Result: 
&WhcGetGrantsResult{PropertyID:0x3, Name:"test_token1", Issuer:"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5", CreateTxID:"1c3f95acbd6eb38e2a7c26b12dc9138b4523c355a20944874bdc3c82f4c5e4e1", TotalTokens:"90", Issuance:[]btcjson.IssuanceTxs{btcjson.IssuanceTxs{TxID:"a1f2d0126a04296aad6f492a0ef8c1c1afb781efc6c5f37de105790a7debcf87", Grant:"", Revoke:"10"}, btcjson.IssuanceTxs{TxID:"6afffd7d14060b6e79c504c5f17596616eb99356b71d301b5b37df1df065b9a0", Grant:"100", Revoke:""}}}
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
&WhcGetInfoResultResult{VersionInt:10000000, Version:"0.1.0", BitcoinVersion:"0.17.2", Block:1265112, BlockTime:1540802803, Txs:0, TotalTrades:0, TotalTxs:2694, Alerts:[]btcjson.AlertInfo{}}
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
Result: 
&WhcGetPayloadResult{Payload:"000000360100020000000074657374206d616e6167656420746f6b656e2032007465737400746573745f746f6b656e31007777772e746573746d616e61676564746f6b656e2e636f6d006d79206461746100", PayloadSize:82}
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
Result: 
&WhcGetPropertyResult{PropertyID:0x7b, Name:"212", Category:"", Subcategory:"", Data:"1", Url:"1", Precision:1, Issuer:"bchtest:qq6qag6mv2fzuq73qanm6k60wppy23djnv7ddk3lpk", CreateTxID:"5b757ee98840b97ef14445fe25ace4dc5f2aba9eb2d447e7bc886db20bfa65b7", FixedIssuance:true, ManagedIssuance:false, FreezingEnabled:(*bool)(nil), TotalTokens:"1.0"}
```

------

##### `WhcListTransactions`

获取指定token的信息

**Arguments**

1. address `address` optinal
2. count `int64` optional
3. skip `int64` optional
4. startblock `int64` optional
5. endblock `int64` optional

##### 返回值

##### `pointer`

##### 示例

```Go
height := 1437719290
rpc.NewRPCInstance().WhcListTransactions(nil, nil, nil, &height, nil)

// result:
Result: 
&WhcGetPropertyResult{PropertyID:0x7b, Name:"212", Category:"", Subcategory:"", Data:"1", Url:"1", Precision:1, Issuer:"bchtest:qq6qag6mv2fzuq73qanm6k60wppy23djnv7ddk3lpk", CreateTxID:"5b757ee98840b97ef14445fe25ace4dc5f2aba9eb2d447e7bc886db20bfa65b7", FixedIssuance:true, ManagedIssuance:false, FreezingEnabled:(*bool)(nil), TotalTokens:"1.0"}
```

------

#####  