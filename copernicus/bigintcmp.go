package main

import (
	"fmt"
	"github.com/copernet/copernicus/model/pow"
	"github.com/copernet/copernicus/util"
	"sort"
)

type sortTxs []*util.Hash

func (s sortTxs) Len() int {
	return len(s)
}
func (s sortTxs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortTxs) Less(i, j int) bool {
	h1 := pow.HashToBig(s[i])
	h2 := pow.HashToBig(s[j])
	return h1.Cmp(h2) < 0
}

func main() {
	tmpTxs := make([]*util.Hash, 0)
	hash1 := util.HashFromString("00000000000001bcd6b635a1249dfbe76c0d001592a7219a36cd9bbd002c7238")
	//h1 := pow.HashToBig(hash1)
	hash2 := util.HashFromString("00000000000001bcd6b635a1249dfbe76c0d001592a7219a36cd9bbd002c7239")
	//h2 := pow.HashToBig(hash2)
	hash3 := util.HashFromString("00000000000001bcd6b635a1249dfbe76c0d001592a7219a36cd9bbd002c7237")
	//h3 := pow.HashToBig(hash3)
	hash4 := util.HashFromString("00000000000001bcd6b635a1249dfbe76c0d001592a7219a36cd9bbd002c7236")
	//h4 := pow.HashToBig(hash4)
	tmpTxs = append(tmpTxs, hash1)
	tmpTxs = append(tmpTxs, hash2)
	tmpTxs = append(tmpTxs, hash4)
	tmpTxs = append(tmpTxs, hash3)
	//not sort
	fmt.Printf("%v", tmpTxs)

	tmpt := sortTxs(tmpTxs)

	fmt.Println()
	//sort
	sort.Sort(tmpt)
	fmt.Printf("%v", tmpTxs)
}
