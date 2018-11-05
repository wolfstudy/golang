package main

import (
	"github.com/bcext/cashutil"
	"github.com/bcext/gcash/chaincfg"
)

func main() {
	addrs := "bchtest:qq0ae7jqqvr87gex4yk3ukppvnm0w7ftqqpzv0lcqa"
	addr, err := cashutil.DecodeAddress(addrs, &chaincfg.TestNet3Params)
	if err != nil {

	}
	addr.EncodeAddress(true)
}
