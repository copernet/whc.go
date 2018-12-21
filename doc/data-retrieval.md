##### `WhcGetAllBalancesForAddress`

Returns a list of all token balances for a given address.

**Arguments**

1. address `string` required

##### Result

`array`

##### Examples

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

Returns a list of token balances for a given currency or property identifier.

**Arguments**

1. property id `uint64` required: the property identifier

##### Result

`array`

##### Examples

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

Returns the token balance for a given address and property.

**Arguments**

1. address `uint64` required
2. property id `uint64 ` required: the property identifier

##### Result

`pointer`

##### Examples

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

Returns a hash of the balances for the property.

**Arguments**

1. property id `uint64` required: the property to hash balances for

##### Result

`pointer`

##### Examples

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

Returns information about a crowdsale.

**Arguments**

1. property id `uint64` required: the identifier of the crowdsale
2. verbose `pointer` optional: list crowdsale participants (default: false)

##### Result

`pointer`

##### Examples

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

Returns the consensus hash for all balances for the current block.

##### Result

`pointer`

##### Examples

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

Returns information about granted and revoked units of managed tokens.

**Arguments**

1. property id `uint64` required: the identifier of the managed tokens to lookup

##### Result

`pointer`

##### Examples

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

Returns various state information of the client and protocol.

##### Result

`pointer`

##### Examples

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

Get the payload for an Omni transaction.

**Arguments**

1. tx hash `string` required: he hash of the transaction to retrieve payload

##### Result

`pointer`

##### Examples

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

Returns details for about the tokens or smart property to lookup.

**Arguments**

1. property id `uint64` required: the identifier of the tokens or property

##### Result

`pointer`

##### Examples

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

List wallet transactions, optionally filtered by an address and block boundaries.

**Arguments**

1. address `address` optinal
2. count `int64` optional
3. skip `int64` optional
4. startblock `int64` optional
5. endblock `int64` optional

##### Result

`array`

##### Examples

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

Get detailed information about an Omni transaction.

**Arguments**

1. tx hash `string`

##### Result

`pointer`

##### Examples

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

Returns a list of unconfirmed Omni transactions, pending in the memory pool.

An optional filter can be provided to only include transactions which involve the given address.

Note: the validity of pending transactions is uncertain, and the state of the memory pool may change at any moment. It is recommended to check transactions after confirmation, and pending transactions should be considered as invalid.

**Arguments**

1. address `pointer` optional: address filter (default: "" for no filter)

##### Result

`array`

##### Examples

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

Lists all tokens or smart properties.

##### Result

`array`

##### Examples

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

##### `WhcGetFrozenBalance`

Returns the frozen token balance for a given address and property.

**Arguments**

1. address `string` the address
2. propertyid `int64` the property identifier


##### Result

`pointer`

##### Examples

```json
rpc.NewRPCInstance().WhcGetFrozenBalance("qzuy3es55tygnmmeydh5uqc39sfkke6hlqf2dv26h0", 459)

// result:
{
  "frozen": true,
  "balance": "100.00000000"
}
```
------

##### `WhcGetFrozenBalanceForId`

Returns the frozen token balance for a given address and property.

**Arguments**

1. propertyid `int64` the property identifier


##### Result

`array`

##### Examples

```json
rpc.NewRPCInstance().WhcGetFrozenBalanceForId(459)

// result:
[
  {
    "address": "bchtest:qqu9lh4jpc05p59pfhu9amyv9uvder8j3sa2up95vs",
    "balance": "100.00000000"
  }
]
```
------

##### `WhcGetFrozenBalanceForAddress`

Returns the frozen token balance for a given address and property.

**Arguments**

1. address `string` the address


##### Result

`array`

##### Examples

```json
rpc.NewRPCInstance().WhcGetFrozenBalanceForAddress("bchtest:qqu9lh4jpc05p59pfhu9amyv9uvder8j3sa2up95vs")

// result:
[
  {
    "propertyid": 459,
    "balance": "100.00000000"
  }
]

```
##### 

##### `whc_getERC721AddressTokens`

Returns details for about the tokens or smart property to lookup.

**Arguments**

- address `string` required:the address of the query.
- propertyid`string`required:the identifier of the ERC721 property

**Result**

- `tokenid`the identifier of the token
- `attribute`the name of the tokens
- `tokenurl`the url of the tokens
- `creationtxid`the hex-encoded creation transaction hash

**Examples**

```go
//request 
rpc.NewRPCInstance().WhcGetERC721AddressTokens("qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","6")

//result
[{"tokenid": "1","attribute": "0000000000000000000000000000000000000000000000000000000000edabef", "tokenurl": "http://lll","creationtxid": "b4437159f1443793f2fc3a4a98951a945c2e6d6fcf4839e1c2d39626821b911b"}
]
```

##### `whc_getERC721PropertyDestroyTokens`

Returns details for about the destroy tokens to lookup.

**Arguments**

- propertyid `string`required:the identifier of the ERC721 property

**Result**

- `tokenid`the identifier of the token 
- `attribute`the name of the tokens
- `tokenurl` the url of the tokens
- `creationtxid`the hex-encoded creation transaction hash

**Examples**

```go
//request
rpc.NewRPCInstance().WhcGetERC721PropertyDestoryTokens("9")
//result
[{ "tokenid": "1",
    "attribute": "0000000000000000000000000000000000000000000000000000000000edabef",
    "tokenurl": "http://lll",
    "creationtxid": "d6c220d2b4ca898f92de9975a32585b285a0c71a401a56b7e6c6cd50b8cfe2c3"
  }
]
```

##### `whc_getERC721PropertyNews`

Returns details for about the tokens or smart property to lookup

**Arguments**

- propertyid`string`required： the identifier of the ERC721 property

**Result**

- `propertyid`the identifier of the property
- `owner`the owner address of the ERC721 Property
- `creationtxid`the hex-encoded creation transaction hash
- `creationblock`the hex-encoded creation block hash
- `name`the name of the property
- `symbol`property symbol
- `data`remark
- `propertyurl`the url of the property
- `totalTokenNumber`the amount of tokens that will be issued
- `haveIssuedNumber`the amount of tokens that have issued
- `currentValidIssuedNumer`the amount of tokens that still available

**Examples**

```go
//request
rpc.NewRPCInstance().WhcGetERC721PropertyNews("7")
//result
{
  "propertyid": "7",
  "owner": "bchtest:qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt",
  "creationtxid": "e930f27993881bd74d89572ca8c2273538ea62dec5980190f7c191cb56f02db1",
  "creationblock": "00000000000000e204e8ac7c0bdf2914067e75f9f7b109a87bda9762afc339c1",
  "name": "property test",
  "symbol": "a",
  "data": "ttt",
  "propertyurl": "http://test",
  "totalTokenNumber": 88,
  "haveIssuedNumber": 1,
  "currentValidIssuedNumer": 1
}
```

##### `whc_getERC721TokenNews`

Returns details for about the tokens or smart property to lookup.

**Arguments**

- propertyid `string`required:the identifier of the ERC721 property
- tokenid`string` required:the identifier of the ERC721 token

**Result**

- `propertyid`the identifier of the property
- `tokenid`the identifier of the token 
- `owner`the owner address of token 
- `creationtxid`the hex-encoded creation transaction hash
- `creationblock`the hex-encoded creation block hash
- `attribute`the name of the tokens
- `tokenurl`the url of the tokens

**Examples**

```go
//request
rpc.NewRPCInstance().WhcGetERC721TokenNews("7","1")

//result
{
  "propertyid": "6",
  "tokenid": "1",
  "owner": "bchtest:qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt",
  "creationtxid": "7e88e4fdd9938f37fb94c7fa587287d57f800af6cf5f9cb8f15dca5edbad1f73",
  "creationblock": "000000000001481bc37330297e0053f0051125574364aa40e72a965da13d1e3e",
  "attribute": "000000000000000000000000000000000000000000000000000000000000000e",
  "tokenurl": "http://lll"
}
```

##### `whc_listERC721PropertyTokens`

List all tokens information for the specified ERC721Property.

**Arguments**

- propertyid`string` required:the identifier of the ERC721 property

**Result**

- `tokenid`the identifier of the token 
- `owner`the owner address of the token 

**Examples**

```go
//request
rpc.NewRPCInstance().WhcListERC721PropertyTokens("9")

//result
[{
    "tokenid": "1",
    "owner": "bchtest:qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt"
  },
  {
    "tokenid": "2",
    "owner": "bchtest:qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt"
  }
]

```

##### `whc_ownerOfERC721Token`

Query whether the Token's owner is the specified address.

**Arguments**

- propertyid `string`required:the identifier of the ERC721 property
- tokenid`string`required:the identifier of the ERC721 token
- address`string` query address for the specified ERC721 Token 

**Result**

- `ownwhether the query address owner the specified ERC721 Token`

**Examples**

```go
//request
rpc.NewRPCInstance().WhcOwnerOfERC721Token("7","1","qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt")

//result
{
  "own": true
}
```

##### 
##### `whc_destroyERC721Token`

Destroy ERC721 Token.

**Arguments**

- senderAddress`string`required:The bitcoincash address of the token owner
- propertyId`string`required:The propertyid within the token that will be destory
- tokenId`string`required:The tokenid that will be destory

**Result**

- `hash`the hex-encoded transaction hash

**Examples**

```go
//request
rpc.NewRPCInstance().WhcDestoryERC721Token("qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","9","1")

//result
2c82940e9f9d637fc456a4c651fef7b9dabeb3187571b562088a43aacbb9d80e
```

##### `whc_issuanceERC721Token`

Issue ERC721 Token.

**Arguments**

- issueAddress`string`required:The BitcoinCash address will issue a token in special property
- receiveaddress`string`required:The address of receiver will be received new created token
- propertyID`string`required:The ID of the special property that will be issued token
- tokenID`string`required:The tokenID that will be issued
- tokenAttributes`string`The Attributes of the new created token
- tokenURL`string`The URL of the new created token

**Result**

- `hash`the hex-encoded transaction hash

**Examples**

```go
//request
rpc.NewRPCInstance().WhcIssuanceERC721Token("qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","9","1","00000000000000000000000000EDABEF","http://lll")

//result
b4437159f1443793f2fc3a4a98951a945c2e6d6fcf4839e1c2d39626821b911b
```

##### `whc_issuanceERC721property`

Issue ERC721 property

**Arguments**

- issueAddress`string`required:The BitcoinCash address will issue ERC721 property
- name`string`required:the name of created property 
- symbol`string`required:the symbol of created property 
- data`string`required:the Data of created property 
- url`string`required:the URL of created property 
- totalNumber`string`required: the number of token that created property will issued in the future

**Result**

- `hash`the hex-encoded transaction hash

**Examples**

```go
//request
rpc.NewRPCInstance().WhcIssuanceERC721Property("qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","property test4","a","ttt","http://test","8")

//result
e5efb94c13d9fbbdbe3decfeb88300d61cf259447d225b641f00251df0bc6db6
```

##### `whc_transferERC721Token`

Transfer ERC721 Token.

**Arguments**

- ownerAddress`string`The bitcoincash address of the token owner
- receiveaddress`string`The redeem bitcoin address of the token receiver
- propertyId`string`The propertyid within the token that will be transfer
- tokenId`string`The tokenid that will be transfer

**Result**

- `hash`the hex-encoded transaction hash

**Examples**

```go
//request
rpc.NewRPCInstance().WhcTransferERC721Token("qrqfsslh4l27kwgxfq959gq9tpqn6j739ydn527mlt","qpd485y7slsn9mnudjm6n477m4yxdx9u5yf0ff3z55","8","1")

//result
025c1382cc049e6db911969313f42ec34b338c4ed21e0d83c25e1dc066db27d5
```

