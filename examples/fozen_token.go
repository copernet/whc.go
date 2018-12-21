package main

import (
	"fmt"
	"github.com/copernet/whc.go/rpcclient"
)

func getClient() *rpcclient.Client {
	// rpc client instance
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:18332",
		User:         "rpc-username",
		Pass:         "rpc-password",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	c, err := rpcclient.New(connCfg, nil)
	if err != nil {
		panic(err)
	}

	return c
}

func main() {
	c := getClient()

	//res, err := c.WhcSendUnFreeze("bchreg:qzrck6dmz5lgs7v87dr5lp4g56aldg9knucnuqyl2g", 5, "1", "bchreg:qzuy3es55tygnmmeydh5uqc39sfkke6hlqf2dv26h0")
	//fmt.Println(res)
	//fmt.Println(err)
	//
	//balance, err := c.WhcGetFrozenBalance("qzuy3es55tygnmmeydh5uqc39sfkke6hlqf2dv26h0", 5)
	//fmt.Println(err)
	//fmt.Println(balance)
	//
	//balances ,err :=c.WhcGetFrozenBalanceForId(5)
	//fmt.Println(err)
	//fmt.Println(balances)
	//
	list,err :=c.WhcGetFrozenBalanceForAddress("qzuy3es55tygnmmeydh5uqc39sfkke6hlqf2dv26h0")
	fmt.Println(err)
	fmt.Println(list)


}
