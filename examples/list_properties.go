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

	// Fetch all wormhole properties.
	list, err := client.WhcListProperties()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(list)
}
