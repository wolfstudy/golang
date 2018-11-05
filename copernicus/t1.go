package main

import (
	"encoding/hex"
	"fmt"

	"github.com/copernet/copernicus/crypto"
)

func main() {
	crypto.InitSecp256()
	bs, err := hex.DecodeString("03ffd03de44a6e11b9917f3a29f9443283d9871c9d743ef30d5eddcd37094b64d1")
	pubkey, err := crypto.ParsePubKey(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("pubkey=%+v\n", pubkey)
}
