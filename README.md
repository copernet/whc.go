whc.go
---

whc.go provides the full wormhole and bitcoin cash relative SDK written in GO.

#### Requirements:

- Install `wormhole ` ([Earth-0.1.0](https://github.com/copernet/wormhole/releases) ) according to the installation instructions:[https://github.com/copernet/wormhole/tree/master/doc](https://github.com/copernet/wormhole/tree/master/doc). 

- [Go](http://golang.org/) 1.8 or newer required.

- Install:

  - Recommended method:

    ```
    go get github.com/copernet/whc.go
    ```

  - Optional method: 

    1. clone repository

    ```
    git clone https://github.com/copernet/whc.go.git $GOPATH/src/github.com/copernet/whc.go
    ```

    2. dependencies installation:

    ```
    go get github.com/btcsuite/go-socks
    go get github.com/btcsuite/websocket
    go get github.com/btcsuite/btclog
    go get github.com/bcext/gcash
    go get github.com/bcext/cashutil
    ```


#### Usage Example:

1. Install and run a wormhole client and configurate correctly, `rpcuser` , `rpcpassword`  and `rpcallowip` fields are  necessary. The following is a simple example for wormhole configuration(default in ~/.bitcoin/bitcoin.conf file):

   ```
   rpcuser=D313FF53C1
   rpcpassword=E51DD672DED3FA7D3FB00E946C10C681EF164
   # server ip of your project using this sdk
   rpcallowip=54.78.09.134
   startclean=1
   ```

2. Connect to wormhole server.

   ```
   func main() {
   	// Connect to local wormhole RPC server using HTTP POST mode.
   	connCfg := &rpcclient.ConnConfig{
   		Host:         "127.0.0.1:18332",
   		User:         "D313FF53C1",
   		Pass:         "E51DD672DED3FA7D3FB00E946C10C681EF164",
   		HTTPPostMode: true, 		// wormhole only supports HTTP POST mode
   		DisableTLS:   true, 		// wormhole does not provide TLS by default
   	}
   	// Notice the notification parameter is nil since notifications are
   	// not supported in HTTP POST mode.
   	client, err := rpcclient.New(connCfg, nil)
   	if err != nil {
   		log.Fatal(err)
   	}
   	defer client.Shutdown()
   }
   ```

   >  Please refer to usage examples: [examples/whcinfo.go](https://github.com/copernet/whc.go/blob/master/examples/whcinfo.go). More APIs are supported as the followings. Open a [issue](https://github.com/copernet/whc.go/issues/new) if necessary.

3. Usage

   - Fetch the current block height for wormhole client

     ```
     // Get the current wormhole relative information.
     info, err := client.WhcGetInfo()
     if err != nil {
     	log.Fatal(err)
     }
     
     log.Printf("Block count: %d", info.Block)
     ```

   - Fetch properties list in the current wormhole system

     ```
     // Fetch all wormhole properties.
     list, err := client.WhcListProperties()
     if err != nil {
     	log.Fatal(err)
     }
     
     log.Println(list)
     ```

   - Create a crowdsale transaction

     ```
     // necessary field
     addressFrom := "bchtest:qqg2fwfzd4xeywf8h2zajqy77357gk0v7yvsvhd4xu"
     ecosystem := int64(1) // must be 1
     precision := int64(8)
     previousID := int64(0) // must be 0
     category := "Blockchain research"
     subCatetory := "Bitcoin cash"
     name := "wormhole"
     url := "https://www.wormhole.cash"
     data := "working for the future"
     desiredID := int64(1) // must be 1
     tokensPerUnit := "100"
     deadline := int64(1582772366)
     earlyBonus := int64(2)
     issuerPercentage := int64(0) // must be 0
     amount := "10000000.987"
     
     // Create a crowdsalle transaction
     txHash, err := client.WhcSendIssuanceCrowdSale(addressFrom, ecosystem, precision, previousID, category,subCatetory, name, url, data, desiredID, tokensPerUnit, deadline, earlyBonus, issuerPercentage, amount)
     
     log.Println(txHash)
     ```

   > Notice, the wormhole client only supports `HTTP POST` mode. Completed expamle is [here](https://github.com/copernet/whc.go/blob/master/examples).

#### Supported API:

The following apis are fully supported by this sdk. `whc.go` intentionally been designed so it can be used as a standalone package which is convenient for import and usage.For more detailed wormhole rpc api, please refer [https://github.com/copernet/spec/blob/master/wormhole-RPC.md](https://github.com/copernet/spec/blob/master/wormhole-RPC.md)

|   #   |  Method    |   Description   |
| ---- | ---- | ---- |
| 1 |WhcSetAutoCommit| Determine whether to automatically commit transactions. |
| 2 |WhcGetActiveCrowd| Returns details for about the active crowd with special address. |
| 3 |WhcGetAllBalancesForAddress| Returns a list of all token balances for a given address. |
| 4 |WhcGetAllBalancesForID| Returns a list of token balances for a given currency or property identifier. |
| 5 |WhcGetBalances| Returns the token balance for a given address and property. |
| 6 |WhcGetBalanceHash| Returns a hash of the balances for the property. |
| 7 |WhcGetCrowdSale| Returns information about a crowdsale. |
| 8 |WhcGetCurrentConsensusHash| Returns the consensus hash for all balances for the current block. |
| 9 |WhcSendSto| Create and broadcast a send-to-owners transaction. |
| 10 |WhcGetGrants| Returns information about granted and revoked units of managed tokens. |
| 11 |WhcGetInfo| Returns various state information of the client and protocol. |
| 12 |WhcGetPayload| Get the payload for an Wormhole transaction. |
| 13 |WhcGetProperty| Returns details for about the tokens or smart property to lookup. |
| 14 |WhcGetSeedBlocks| Returns a list of blocks containing wormhole transactions for use in seed block filtering. |
| 15 |WhcGetSto| Get information and recipients of a send-to-owners transaction. |
| 16 |WhcGetTransaction| Get detailed information about an wormhole transaction. |
| 17 |WhcListBlockTransactions| Lists all Wormhole transactions in a block. |
| 18 |WhcListPendingTransactions| Returns a list of unconfirmed Wormhole transactions, pending in the memory pool. |
| 19 |WhcListProperties| Lists all tokens or smart properties. |
| 20 |WhcListTransactions| List wallet transactions, optionally filtered by an address and block boundaries. |
| 21 |WhcCreatePayloadBurnBCH| Creates the payload to burn bch to get whc. |
| 22 |WhcCreatePayloadChangeIssuer| Creats the payload to change the issuer on record of the given tokens. |
| 23 |WhcCreatePayloadCloseCrowdSale| Creates the payload to manually close a crowdsale. |
| 24 |WhcCreatePayloadGrant| Creates the payload to issue or grant new units of managed tokens. |
| 25 |WhcCreatePayloadIssuanceCrowdSale| Creates the payload for a new tokens issuance with crowdsale. |
| 26 |WhcCreatePayloadIssuanceFixed| Creates the payload for a new tokens issuance with fixed supply. |
| 27 |WhcCreatePayloadIssuanceManaged| Creates the payload for a new tokens issuance with manageable supply. |
| 28 |WhcCreatePayloadPartiCrowdSale| Create the payload for a participate crowsale transaction. |
| 29 |WhcCreatePayloadRevoke| Creates the payload to revoke units of managed tokens. |
| 30 |WhcCreatePayloadSendAll| Create the payload for a send all transaction. |
| 31 |WhcCreatePayloadSimpleSend| Create the payload for a simple send transaction. |
| 32 |WhcCreatePayloadSto| Creates the payload for a send-to-owners transaction. |
| 33 |WhcCreateRawTxChange| Adds a change output to the transaction. |
| 34 |WhcCreateRawTxInput| Adds a transaction input to the transaction. |
| 35 |WhcCreateRawTxOpReturn| Adds a payload with class C (op-return) encoding to the transaction. |
| 36 |WhcCreateRawTxReference| Adds a reference output to the transaction. |
| 37 |WhcDecodeTransaction| Decodes an wormhole transaction. |
| 38 |WhcBurnBCHGetWhc| Burn BCH to get WHC |
| 39 |WhcPartiCrowdSale| Create and broadcast a participate crowsale transaction. |
| 40 |WhcSend| Create and broadcast a simple send transaction. |
| 41 |WhcSendAll| Transfers all available tokens in the given ecosystem to the recipient. |
| 42 |WhcSendChangeIssuer| Change the issuer on record of the given tokens. |
| 43 |WhcSendCloseCrowdSale| Manually close a crowdsale. |
| 44 |WhcSendGrant| Issue or grant new units of managed tokens. |
| 45 |WhcSendIssuanceCrowdSale| Create new tokens as crowdsale. |
| 46 |WhcSendIssuanceFixed| Create new tokens with fixed supply. |
| 47 |WhcSendIssuanceManaged| Create new tokens with manageable supply. |
| 48 |WhcSendRawTx| Broadcasts a raw Wormhole Layer transaction. |
| 49 |WhcSendRevoke| Revoke units of managed tokens. |

All API supported sync and async request mode, and all async methods suffixed with `Async`. At some situations, using async request method is more efficient. The following example using async request: 

```
// Notice the notification parameter is nil since notifications are
// not supported in HTTP POST mode.
client, err := rpcclient.New(connCfg, nil)
if err != nil {
	log.Fatal(err)
}
defer client.Shutdown()

// use async mode
r := client.WhcGetCrowdSaleAsync(34, nil)
// exec a time-consuming task
// ...
// receive result
result, err := r.Receive()

// sync mode
r := client.WhcGetCrowdSale(34, nil)
```
