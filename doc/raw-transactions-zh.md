##### `WhcCreateRawTxChange`

向未签名交易输出集合的指定位置追加一个交易输出

**Arguments**

1. rawtx `string`
2. prev `array`
3. address `string`
4. fee `string`
5. position `pointer`

##### 返回值

##### `string`

##### 示例

```json
inputs := []btcjson.PrevTx{
    {
        TxID:         "6779a710fcd5f6fb0883ea3306360c3ad8c0a3c5de902768ec57ef3104e65eb1",
        Vout:         4,
        ScriptPubKey: "76a9147b25205fd98d462880a3e5b0541235831ae959e588ac",
        Value:        0.00068257,
    },
}
rpc.NewRPCInstance().WhcCreateRawTxChange("0100000001b15ee60431ef57ec682790dec5a3c0d83a0c360633ea8308fbf6d5fc10a779670400000000ffffffff025c0d00000000000047512102f3e471222bb57a7d416c82bf81c627bfcd2bdc47f36e763ae69935bba4601ece21021580b888ff56feb27f17f08802ebed26258c23697d6a462d43fc13b565fda2dd52aeaa0a0000000000001976a914946cb2e08075bcbaf157e47bcb67eb2b2339d24288ac00000000", inputs, "bchtest:qrv2054hm0rswdnyaskxsyxaq6j58x479ynsgl79xp", "0.0001", nil)

// result:
0100000001b15ee60431ef57ec682790dec5a3c0d83a0c360633ea8308fbf6d5fc10a779670400000000ffffffff038bcb0000000000001976a914d8a7d2b7dbc7073664ec2c6810dd06a5439abe2988ac5c0d00000000000047512102f3e471222bb57a7d416c82bf81c627bfcd2bdc47f36e763ae69935bba4601ece21021580b888ff56feb27f17f08802ebed26258c23697d6a462d43fc13b565fda2dd52aeaa0a0000000000001976a914946cb2e08075bcbaf157e47bcb67eb2b2339d24288ac00000000
```

------

##### `WhcCreateRawTxInput`

向未签名的交易追加一个交易输入

**Arguments**

1. rawtx `string`
2. tx hash `string`
3. index `int`

##### 返回值

##### `string`

##### 示例

```json
rpc.NewRPCInstance().WhcCreateRawTxInput("01000000000000000000", "b006729017df05eda586df9ad3f8ccfee5be340aadf88155b784d1fc0e8342ee", 0)

// result:
0100000001ee42830efcd184b75581f8ad0a34bee5feccf8d39adf86a5ed05df17907206b00000000000ffffffff0000000000
```

------

##### `WhcCreateRawTxOpReturn`

将wormhole协议的载荷数据作为新输出的脚本追加在未签名交易中

**Arguments**

1. rawtx `string`
2. payload `string`

##### 返回值

##### `string`

##### 示例

```json
rpc.NewRPCInstance().WhcCreateRawTxOpReturn("01000000000000000000", "00000000000000020000000006dac2c0")

// result:
0100000000010000000000000000166a140877686300000000000000020000000006dac2c000000000
```

------

##### `WhcCreateRawTxReference`

向未签名交易追加一个交易输出

**Arguments**

1. rawtx `string`
2. address `string`
3. amount `pointer` optional

##### 返回值

##### `string`

##### 示例

```json
rpc.NewRPCInstance().WhcCreateRawTxReference("0100000001a7a9402ecd77f3c9f745793c9ec805bfa2e14b89877581c734c774864247e6f50400000000ffffffff03aa0a0000000000001976a9146d18edfe073d53f84dd491dae1379f8fb0dfe5d488ac5c0d0000000000004751210252ce4bdd3ce38b4ebbc5a6e1343608230da508ff12d23d85b58c964204c4cef3210294cc195fc096f87d0f813a337ae7e5f961b1c8a18f1f8604a909b3a5121f065b52aeaa0a0000000000001976a914946cb2e08075bcbaf157e47bcb67eb2b2339d24288ac00000000", "bchtest:qrpdz05x8n4fxs80w07gsdxp5qe208kg6us3r2l27l", nil)

// result:
0100000001a7a9402ecd77f3c9f745793c9ec805bfa2e14b89877581c734c774864247e6f50400000000ffffffff04aa0a0000000000001976a9146d18edfe073d53f84dd491dae1379f8fb0dfe5d488ac5c0d0000000000004751210252ce4bdd3ce38b4ebbc5a6e1343608230da508ff12d23d85b58c964204c4cef3210294cc195fc096f87d0f813a337ae7e5f961b1c8a18f1f8604a909b3a5121f065b52aeaa0a0000000000001976a914946cb2e08075bcbaf157e47bcb67eb2b2339d24288ac22020000000000001976a914c2d13e863cea9340ef73fc8834c1a032a79ec8d788ac00000000
```

------

##### `WhcDecodeTransaction`

解析wormhole的原始交易

**Arguments**

1. rawtx `string`
2. prev `pointer` optional
3. height `pointer` optional

##### 返回值

##### `string`

##### 示例

```json
rpc.NewRPCInstance().WhcDecodeTransaction("02000000019687c2769d84850b1d7331af1f37e03f19054c5850cb4c2dd148a4aafa6267f1000000006b4830450221008845964425c3600f6be736dae925e4ef289705897ea34110008aeb659ec5b4d602206cc1a781ed0e2c5f16970023fef9dc9a5eadbe2319fc7a035533b1c7b43bfb04412103511b46817ae12cf145e4c4430859a8a3b635ec8dd8c7d83352ed69841ca12582ffffffff0345899800000000001976a914385fdeb20e1f40d0a14df85eec8c2f18dc8cf28c88ac0000000000000000166a1408776863000000000000000100000002540be40022020000000000001976a914103b582493e8f718523b3b8ec75361085020db7c88ac00000000", nil, nil)

// result:
{
    "txid": "35f3210fd2963f901f7d49a4bff2a0f537ed9f27698c5372253a2e1d1d9f9a4c",
    "fee": "520",
    "sendingaddress": "bchtest:qqu9lh4jpc05p59pfhu9amyv9uvder8j3sa2up95vs",
    "referenceaddress": "bchtest:qqgrkkpyj050wxzj8vaca36nvyy9qgxm0s93z4fqac",
    "ismine": true,
    "version": 0,
    "type_int": 0,
    "type": "Simple Send",
    "propertyid": 1,
    "precision": "8",
    "amount": "100.00000000",
    "subsends": null
}
```

------

