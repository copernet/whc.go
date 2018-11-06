##### `WhcCreatePayloadBurnBCH`

Creates the payload to burn bch to get whc.

##### Result

`string`：the hex-encoded payload

##### Examples

```Go
rpc.NewRPCInstance().WhcCreatePayloadBurnBCH()

// result:
00000044
```

---

##### `WhcCreatePayloadChangeIssuer`

Creats the payload to change the issuer on record of the given tokens.

##### Arguments

- propertyid `int64` ：the identifier of the tokens

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadChangeIssuer(34)

// result:
0000004600000022
```

---

##### `WhcCreatePayloadIssuanceCrowdSale`

Creates the payload for a new tokens issuance with crowdsale.

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

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadIssuanceCrowdSale(1, 8, 0, 1, 1540470809, 23, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12.34", "12734782")

// result:
0000003301000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d000000000100000000498d5880000000005bd1b819170000048638bfbd7e00
```

---

##### `WhcCreatePayloadPartiCrowdSale`

Create the payload for a participate crowsale transaction.

##### Arguments

- amount `string`：the amount of WHC to particrowsale

##### Result

`string`：the hex-encoded payload

##### 示例

```
client.WhcCreatePayloadPartiCrowdSale("90.23")

// result:
00000001000000010000000219d00dc0
```

---

##### `WhcCreatePayloadCloseCrowdSale`

Creates the payload to manually close a crowdsale.

##### Arguments

- propertyid `int64` ：the identifier of the crowdsale to close

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadCloseCrowdSale(46)

// result:
000000350000002e
```

---

##### `WhcCreatePayloadIssuanceFixed`

Creates the payload for a new tokens issuance with fixed supply.

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

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadIssuanceFixed(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem", "12734782")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadIssuanceManaged`

Creates the payload for a new tokens issuance with manageable supply.

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

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadIssuanceManaged(1, 8, 0, "womhole expamle", "awesome", "crowdSaleToken", "https://wormhole.cash", "welcome to wormhole ecosystem")

// result:
0000003201000800000000776f6d686f6c6520657870616d6c6500617765736f6d650063726f776453616c65546f6b656e0068747470733a2f2f776f726d686f6c652e636173680077656c636f6d6520746f20776f726d686f6c652065636f73797374656d0000048638bfbd7e00
```

---

##### `WhcCreatePayloadGrant`

Creates the payload to issue or grant new units of managed tokens.

##### Arguments

- id `int64`：the identifier of the tokens to revoke
- amount `string`：the amount of tokens to revok
- note `*string` **optional**：a text note attached to this transaction (none by default)

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadGrant(4, "90.23", nil)

// result:
000000370000000400000000000008fc00
```

---

##### `WhcCreatePayloadRevoke`

Creates the payload to revoke units of managed tokens.

##### Arguments

- id `int64`：the identifier of the tokens to revoke
- amount `string`：the amount of tokens to revok
- note `*string` **optional**：a text note attached to this transaction (none by default)

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadRevoke(4, "90.23", nil)

// result:
0000003800000004000000000000233f00
```

---

##### `WhcCreatePayloadSendAll`

Create the payload for a send all transaction.

##### Arguments

- eco `int64`： the ecosystem to create the tokens in, must be 1

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadSendAll(1)

// result:
0000000401
```

---

##### `WhcCreatePayloadSimpleSend`

Create the payload for a simple send transaction.

##### Arguments

- id `int64`：the identifier of the tokens to send
- amount `string`：the amount to send

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadSimpleSend(35, "23")

// result:
000000000000002300000000000000e6
```

---

##### `WhcCreatePayloadSto`

Creates the payload for a send-to-owners transaction.

##### Arguments

- fromId `int64`：the identifier of the tokens to distribute
- amount `string`：the amount to distribute
- toID `*int64` **optional**：the identifier of the property holders to distribute to

##### Result

`string`：the hex-encoded payload

##### Examples

```
client.WhcCreatePayloadSto(35, "23", nil)

// result:
000000000000002300000000000000e6
```
---

##### `WhcCreatePayloadFreeze`

Creates the payload to freeze an address for a centrally managed token.

##### Arguments

- toaddress `string` ： the address to freeze tokens for
- propertyid `int64` ： the property to freeze tokens for (must be managed type and have freezing option enabled)
- amount `string ` ： the amount of tokens to freeze (note: this is unused - once frozen an address cannot send any transactions)

##### Result

`string`：the hex-encoded payload

##### Examples

```
client. WhcCreatePayloadFreeze("bchreg:qzrck6dmz5lgs7v87dr5lp4g56aldg9knucnuqyl2g", 5, "1")

// result:
0000004600000022
```

---

##### `WhcCreatePayloadUnFreeze`

Creates the payload to unfreeze an address for a centrally managed token.

##### Arguments

- toaddress `string` ： the address to freeze tokens for
- propertyid `int64` ： the property to freeze tokens for (must be managed type and have freezing option enabled)
- amount `string ` ： the amount of tokens to freeze (note: this is unused - once frozen an address cannot send any transactions)

##### Result

`string`：the hex-encoded payload

##### Examples

```
client. WhcCreatePayloadUnFreeze("bchreg:qzrck6dmz5lgs7v87dr5lp4g56aldg9knucnuqyl2g", 5, "1")

// result:
0000004600000022
```
