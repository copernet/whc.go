sdk.go
---

sdk.go provides the full wormhole and bitcoin cash relative operation and query services via RPC protocol. 

#### Usages:

Please refer to usage examples: [examples/whcinfo.go](https://github.com/copernet/whc.go/examples/whcinfo.go)

#### Support API:

|   #   |  Method    |   Description   |
| ---- | ---- | ---- |
| 1 |whc_setautocommit| Determine whether to automatically commit transactions. |
| 2 |whc_getactivecrowd| Returns details for about the active crowd with special address. |
| 3 |whc_getallbalancesforaddress| Returns a list of all token balances for a given address. |
| 4 |whc_getallbalancesforid| Returns a list of token balances for a given currency or property identifier. |
| 5 |whc_getbalance| Returns the token balance for a given address and property. |
| 6 |whc_getbalanceshash| Returns a hash of the balances for the property. |
| 7 |whc_getcrowdsale| Returns information about a crowdsale. |
| 8 |whc_getcurrentconsensushash| Returns the consensus hash for all balances for the current block. |
| 9 |whc_sendsto| Create and broadcast a send-to-owners transaction. |
| 10 |whc_getgrants| Returns information about granted and revoked units of managed tokens. |
| 11 |whc_getinfo| Returns various state information of the client and protocol. |
| 12 |whc_getpayload| Get the payload for an Wormhole transaction. |
| 13 |whc_getproperty| Returns details for about the tokens or smart property to lookup. |
| 14 |whc_getseedblocks| Returns a list of blocks containing Omni transactions for use in seed block filtering. |
| 15 |whc_getsto| Get information and recipients of a send-to-owners transaction. |
| 16 |whc_gettransaction| Get detailed information about an Omni transaction. |
| 17 |whc_listblocktransactions| Lists all Wormhole transactions in a block. |
| 18 |whc_listpendingtransactions| Returns a list of unconfirmed Wormhole transactions, pending in the memory pool. |
| 19 |whc_listproperties| Lists all tokens or smart properties. |
| 20 |whc_listtransactions| List wallet transactions, optionally filtered by an address and block boundaries. |
| 21 |whc_createpayload_burnbch| Creates the payload to burn bch to get whc. |
| 22 |whc_createpayload_changeissuer| Creats the payload to change the issuer on record of the given tokens. |
| 23 |whc_createpayload_closecrowdsale| Creates the payload to manually close a crowdsale. |
| 24 |whc_createpayload_grant| Creates the payload to issue or grant new units of managed tokens. |
| 25 |whc_createpayload_issuancecrowdsale| Creates the payload for a new tokens issuance with crowdsale. |
| 26 |whc_createpayload_issuancefixed| Creates the payload for a new tokens issuance with fixed supply. |
| 27 |whc_createpayload_issuancemanaged| Creates the payload for a new tokens issuance with manageable supply. |
| 28 |whc_createpayload_particrowdsale| Create the payload for a participate crowsale transaction. |
| 29 |whc_createpayload_revoke| Creates the payload to revoke units of managed tokens. |
| 30 |whc_createpayload_sendall| Create the payload for a send all transaction. |
| 31 |whc_createpayload_simplesend| Create the payload for a simple send transaction. |
| 32 |whc_createpayload_sto| Creates the payload for a send-to-owners transaction. |
| 33 |whc_createrawtx_change| Adds a change output to the transaction. |
| 34 |whc_createrawtx_input| Adds a transaction input to the transaction. |
| 35 |whc_createrawtx_opreturn| Adds a payload with class C (op-return) encoding to the transaction. |
| 36 |whc_createrawtx_reference| Adds a reference output to the transaction. |
| 37 |whc_decodetransaction| Decodes an Omni transaction. |
| 38 |whc_burnbchgetwhc| Burn BCH to get WHC |
| 39 |whc_particrowsale| Create and broadcast a participate crowsale transaction. |
| 40 |whc_send| Create and broadcast a simple send transaction. |
| 41 |whc_sendall| Transfers all available tokens in the given ecosystem to the recipient. |
| 42 |whc_sendchangeissuer| Change the issuer on record of the given tokens. |
| 43 |whc_sendclosecrowdsale| Manually close a crowdsale. |
| 44 |whc_sendgrant| Issue or grant new units of managed tokens. |
| 45 |whc_sendissuancecrowdsale| Create new tokens as crowdsale. |
| 46 |whc_sendissuancefixed| Create new tokens with fixed supply. |
| 47 |whc_sendissuancemanaged| Create new tokens with manageable supply. |
| 48 |whc_sendrawtx| Broadcasts a raw Wormhole Layer transaction. |
| 49 |whc_sendrevoke| Revoke units of managed tokens. |



