// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2018 The copernet developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/copernet/whc.go/rpcclient"
)

func main() {
	// Connect to local wormhole RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:18332",
		User:         "rpc-username",
		Pass:         "rpc-password",
		HTTPPostMode: true, // wormhole only supports HTTP POST mode
		DisableTLS:   true, // wormhole does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

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
	txHash, err := client.WhcSendIssuanceCrowdSale(addressFrom, ecosystem, precision, previousID, category,
		subCatetory, name, url, data, desiredID, tokensPerUnit, deadline, earlyBonus, issuerPercentage, amount)

	log.Println(txHash)
}
