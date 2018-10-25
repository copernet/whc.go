##### `WhcCreatePayloadBurnBCH`

燃烧BCH，获取WHC

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```Go
rpc.NewRPCInstance().WhcCreatePayloadBurnBCH()

// result:
00000044
```

---

##### `WhcCreatePayloadChangeIssuer`

修改token的发行者

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadChangeIssuer(34)

// result:
0000004600000022
```

---

##### `WhcCreatePayloadCloseCrowdSale`

关闭众筹

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadCloseCrowdSale(46)

// result:
000000350000002e
```

---

##### `WhcCreatePayloadIssuanceCrowdSale`

发行可众筹的token

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadIssuanceCrowdSale(1, 8, 0, 1, 1540470809, 23, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12.34", "12734782")

// result:
0000003301000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d000000000100000000498d5880000000005bd1b819170000048638bfbd7e00
```

---

##### `WhcCreatePayloadIssuanceFixed`

发行固定数量的token

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadIssuanceFixed(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12734782")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadIssuanceManaged`

发行可管理的token

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadIssuanceManaged(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadPartiCrowdSale`

参与众筹

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadPartiCrowdSale("90.23")

// result:
00000001000000010000000219d00dc0
```

---

##### `WhcCreatePayloadRevoke`

销毁指定数量的可管理token

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadRevoke(4, "90.23", nil)

// result:
0000003800000004000000000000233f00
```

---

##### `WhcCreatePayloadSendAll`

发送指定地址的所有token至另一个地址

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadSendAll(1)

// result:
0000000401
```

---

##### `WhcCreatePayloadSimpleSend`

token转账

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadSimpleSend(35, "23")

// result:
000000000000002300000000000000e6
```

---

##### `WhcCreatePayloadSto`

空投

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadSto(35, "23", nil)

// result:
000000000000002300000000000000e6
```
