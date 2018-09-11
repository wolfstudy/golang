package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	bs, err := hex.DecodeString("03ffd03de44a6e11b9917f3a29f9443283d9871c9d743ef30d5eddcd37094b64d1")
	key, err := btcec.ParsePubKey(bs, btcec.S256())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key=%+v\n", key)
}
