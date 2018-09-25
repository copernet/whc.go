whc.go
---

whc.go provides the full wormhole and bitcoin cash relative operation and query services via RPC protocol. Now wormhole client with the latest version and Go environment are required before using this wormhole sdk.

#### Requirements:

- Now, you need a wormhole client with the latest version. Install `wormhole` according to the installation instructions here: [https://github.com/copernet/wormhole/tree/master/doc](https://github.com/copernet/wormhole/tree/master/doc). There are many different document files for different operation system. Example: ubuntu users need refer to [https://github.com/copernet/wormhole/blob/master/doc/build-unix.md](https://github.com/copernet/wormhole/blob/master/doc/build-unix.md)

- This sdk requires [Go](http://golang.org/) 1.8 or newer.

- Install:

  1. clone this repository. In order to clone and install dependencies at the same time, you should use `go get github.com/copernet/whc.go` and this way is recommended.

     ```
     git clone https://github.com/copernet/whc.go.git $GOPATH/src/github.com/copernet/whc.go
     ```

  2. install dependencies. If you install this repository via `go get github.com/copernet/whc.go `, the following instructions should be ignored:

     ```
     go get github.com/btcsuite/go-socks
     go get github.com/btcsuite/websocket
     go get github.com/btcsuite/btclog
     go get github.com/bcext/gcash
     go get github.com/bcext/cashutil
     ```

  3. Now, import this sdk package in your own project like this: `import github.com/copernet/whc.go/rpcclient`. And the detailed usage examples as the followings.

#### Examples:

1. Now, you should run a wormhole client and configurate correctly, `rpcuser` and `rpcpassword` fields are  necessary. If your project servers on a differect server machine from wormhole, the `rpcallowip`field is necessary. The following is a simple example for wormhole configuration:

   ```
   rpcuser=D313FF53C1
   rpcpassword=E51DD672DED3FA7D3FB00E946C10C681EF164
   # server ip of your project using this sdk
   rpcallowip=54.78.09.134
   startclean=1
   ```

2. In your project, a cofiguration with necessary fields is recommended. For simple usage, the following example uses hard coded configuration.

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
   
   	// Get the current wormhole relative information.
   	info, err := client.WhcGetInfo()
   	if err != nil {
   		log.Fatal(err)
   	}
   	log.Printf("Block count: %d", info.Block)
   }
   
   // possible output:
   1251782
   ```

   >  Please refer to usage examples: [examples/whcinfo.go](https://github.com/copernet/whc.go/blob/master/examples/whcinfo.go). More APIs are supported as the followings. Open a [issue](https://github.com/copernet/whc.go/issues/new) if necessary.

#### Supported API:

Now the following apis are fully supported by this sdk and will be synced with the latest version wormhole upgrade. `whc.go` intentionally been designed so it can be used as a standalone package for any projects needing the ability query and create wormhole relative transactions. 

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
| 14 |WhcGetSeedBlocks| Returns a list of blocks containing Omni transactions for use in seed block filtering. |
| 15 |WhcGetSto| Get information and recipients of a send-to-owners transaction. |
| 16 |WhcGetTransaction| Get detailed information about an Omni transaction. |
| 17 |WhcListBlockTransactions| Lists all Wormhole transactions in a block. |
| 18 |WhcListPendingTransactions| Returns a list of unconfirmed Wormhole transactions, pending in the memory pool. |
| 19 |WhcListProperties| Lists all tokens or smart properties. |
| 20 |WhcListTransactions| List wallet transactions, optionally filtered by an address and block boundaries. |
| 21 |WhcBurnBCHGetWhc| Creates the payload to burn bch to get whc. |
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
| 37 |WhcDecodeTransaction| Decodes an Omni transaction. |
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

All API supported sync and async access request, and all async methods suffixed with `Async`. At some situations, using async request method is more efficient. The following example using async request: 

```
// Notice the notification parameter is nil since notifications are
// not supported in HTTP POST mode.
client, err := rpcclient.New(connCfg, nil)
if err != nil {
	log.Fatal(err)
}
defer client.Shutdown()

r := client.WhcGetCrowdSaleAsync(34, nil)
// do some time-consuming code
time.Sleep(3 * time.Second)

result, err := r.Receive()
```

Now, you will find the program is time-saving because the program will exec a async job and the wormhole server works at the same time.