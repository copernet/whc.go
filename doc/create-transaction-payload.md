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

##### Arguments

- propertyid `int64` ： token的property ID

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

##### Arguments

- propertyid `int64` ： token的property ID

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

##### Arguments

- eco `int64`：the ecosystem to create the tokens in, must be 1
- precision `int64`： the precision of the tokens to create:[0, 8]
- preId `int64`： an identifier of a predecessor token (0 for new crowdsales)
- desiredID`int64`：the identifier of a token eligible to participate in the crowdsale
- deadline   `int64`：the deadline of the crowdsale as Unix timestamp
- earlyBonus `int64`：an early bird bonus for participants in percent per week
- issuerPercentage `int64`：(number, required) the value must be 0
- category `string `： a category for the new tokens (can be "")
- subcategory `string`：a subcategory for the new tokens  (can be "")
- name`string` the name of the new tokens to create
- url `string`： an URL for further information about the new tokens (can be "")
- data `string`： a description for the new tokens (can be "")
- tokensPerunit `string`：the amount of tokens granted per unit invested in the crowdsale
- totalNumber `string`： (string, required) the number of tokens to create

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

##### Arguments

- eco `int64`：the ecosystem to create the tokens in, must be 1
- precision `int64`： the precision of the tokens to create:[0, 8]
- preId `int64`： an identifier of a predecessor token (0 for new crowdsales)
- category `string `： a category for the new tokens (can be "")
- subcategory `string`：a subcategory for the new tokens  (can be "")
- name`string` the name of the new tokens to create
- url `string`： an URL for further information about the new tokens (can be "")
- data `string`： a description for the new tokens (can be "")
- amount `string`： (string, required) the number of tokens to create

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

##### Arguments

- eco `int64`：the ecosystem to create the tokens in, must be 1
- precision `int64`： the precision of the tokens to create:[0, 8]
- preId `int64`： an identifier of a predecessor token (0 for new crowdsales)
- category `string `： a category for the new tokens (can be "")
- subcategory `string`：a subcategory for the new tokens  (can be "")
- name`string` the name of the new tokens to create
- url `string`： an URL for further information about the new tokens (can be "")
- data `string`： a description for the new tokens (can be "")

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

##### Arguments

- amount `string`：the amount of WHC to particrowsale

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

##### Arguments

- id `int64`：the identifier of the tokens to revoke
- amount `string`：the amount of tokens to revok
- note `*string` **optional**：a text note attached to this transaction (none by default)

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

##### Arguments

- eco `int64`： the ecosystem to create the tokens in, must be 1

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

##### Arguments

- id `int64`：the identifier of the tokens to send
- amount `string`：the amount to send

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

##### Arguments

- fromId `int64`：the identifier of the tokens to distribute
- amount `string`：the amount to distribute
- toID `*int64`：the identifier of the property holders to distribute to

##### Result

`string`：生成的wormhole 协议payload数据

##### Examples

```
client.WhcCreatePayloadSto(35, "23", nil)

// result:
000000000000002300000000000000e6
```
