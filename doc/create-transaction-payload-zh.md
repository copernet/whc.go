##### `WhcCreatePayloadBurnBCH`

燃烧BCH，获取WHC

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```Go
rpc.NewRPCInstance().WhcCreatePayloadBurnBCH()

// result:
00000044
```

---

##### `WhcCreatePayloadChangeIssuer`

修改token的发行者

##### 参数

- propertyid `int64` ： token的property ID

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadChangeIssuer(34)

// result:
0000004600000022
```

---

##### `WhcCreatePayloadIssuanceCrowdSale`

发行可众筹的token

##### 参数

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

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadIssuanceCrowdSale(1, 8, 0, 1, 1540470809, 23, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12.34", "12734782")

// result:
0000003301000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d000000000100000000498d5880000000005bd1b819170000048638bfbd7e00
```

---

##### `WhcCreatePayloadPartiCrowdSale`

参与众筹

##### 参数

- amount `string`：the amount of WHC to particrowsale

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadPartiCrowdSale("90.23")

// result:
00000001000000010000000219d00dc0
```

---

##### `WhcCreatePayloadCloseCrowdSale`

关闭众筹

##### 参数

- propertyid `int64` ： token的property ID

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadCloseCrowdSale(46)

// result:
000000350000002e
```

---

##### `WhcCreatePayloadIssuanceFixed`

发行固定数量的token

##### 参数

- eco `int64`：the ecosystem to create the tokens in, must be 1
- precision `int64`： the precision of the tokens to create:[0, 8]
- preId `int64`： an identifier of a predecessor token (0 for new crowdsales)
- category `string `： a category for the new tokens (can be "")
- subcategory `string`：a subcategory for the new tokens  (can be "")
- name`string` the name of the new tokens to create
- url `string`： an URL for further information about the new tokens (can be "")
- data `string`： a description for the new tokens (can be "")
- amount `string`： (string, required) the number of tokens to create

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadIssuanceFixed(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12734782")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadIssuanceManaged`

发行可管理的token

##### 参数

- eco `int64`：the ecosystem to create the tokens in, must be 1
- precision `int64`： the precision of the tokens to create:[0, 8]
- preId `int64`： an identifier of a predecessor token (0 for new crowdsales)
- category `string `： a category for the new tokens (can be "")
- subcategory `string`：a subcategory for the new tokens  (can be "")
- name`string` the name of the new tokens to create
- url `string`： an URL for further information about the new tokens (can be "")
- data `string`： a description for the new tokens (can be "")

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadIssuanceManaged(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadGrant`

增发指定数量的可管理token

##### 参数

- id `int64`：the identifier of the tokens to revoke
- amount `string`：the amount of tokens to revok
- note `*string` **optional**：a text note attached to this transaction (none by default)

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadGrant(4, "90.23", nil)

// result:
000000370000000400000000000008fc00
```

---

##### `WhcCreatePayloadRevoke`

销毁指定数量的可管理token

##### 参数

- id `int64`：the identifier of the tokens to revoke
- amount `string`：the amount of tokens to revok
- note `*string` **optional**：a text note attached to this transaction (none by default)

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadRevoke(4, "90.23", nil)

// result:
0000003800000004000000000000233f00
```

---

##### `WhcCreatePayloadSendAll`

发送指定地址的所有token至另一个地址

##### 参数

- eco `int64`： the ecosystem to create the tokens in, must be 1

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadSendAll(1)

// result:
0000000401
```

---

##### `WhcCreatePayloadSimpleSend`

token转账

##### 参数

- id `int64`：the identifier of the tokens to send
- amount `string`：the amount to send

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadSimpleSend(35, "23")

// result:
000000000000002300000000000000e6
```

---

##### `WhcCreatePayloadSto`

空投

##### 参数

- fromId `int64`：the identifier of the tokens to distribute
- amount `string`：the amount to distribute
- toID `*int64` **optional**：the identifier of the property holders to distribute to

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client.WhcCreatePayloadSto(35, "23", nil)

// result:
000000000000002300000000000000e6
```
---

##### `WhcCreatePayloadFreeze`

冻结可管理资产

##### 参数

- toaddress `int64` ： token的拥有者地址
- propertyid `int64` ： token的property ID(资产类型必须是管理资产且启动冻结选项)
- amount `string ` ： 冻结数量（备注：该参数当前未启用）

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client. WhcCreatePayloadFreeze("bchreg:qzrck6dmz5lgs7v87dr5lp4g56aldg9knucnuqyl2g", 5, "1")

// result:
0000004600000022
```

---
##### `WhcCreatePayloadUnFreeze `

解冻可管理资产

##### 参数

- toaddress `int64` ： token的拥有者地址
- propertyid `int64` ： token的property ID(资产类型必须是管理资产且启动冻结选项)
- amount `string ` ： 冻结数量（备注：该参数当前未启用）

##### 返回值

`string`：生成的wormhole 协议payload数据

##### 示例

```
client. WhcCreatePayloadUnFreeze("bchreg:qzrck6dmz5lgs7v87dr5lp4g56aldg9knucnuqyl2g", 5, "1")

// result:
0000004600000022
```

---
