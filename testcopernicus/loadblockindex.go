package main

import (
	"github.com/copernet/copernicus/model/utxo"
	"github.com/copernet/copernicus/model/blockindex"
	"github.com/copernet/copernicus/persist/blkdb"
	"github.com/copernet/copernicus/persist"
	"github.com/copernet/copernicus/model/chain"
	"sort"
	"github.com/copernet/copernicus/log"
	"github.com/copernet/copernicus/util"
	"fmt"
	"github.com/copernet/copernicus/conf"
	"github.com/copernet/copernicus/persist/db"
	"github.com/copernet/copernicus/logic/lblockindex"
	"github.com/copernet/copernicus/logic/lchain"
	"github.com/copernet/copernicus/crypto"
	"github.com/copernet/copernicus/model/mempool"
	"github.com/copernet/copernicus/persist/disk"
	"os"
	"path/filepath"
	"encoding/json"
)

func appInitMain(args []string) {
	//init config
	conf.Cfg = conf.InitConfig(args)
	fmt.Println("Current data dir:\033[0;32m", conf.DataDir, "\033[0m")

	//init log
	logDir := filepath.Join(conf.DataDir, log.DefaultLogDirname)
	if !conf.ExistDataDir(logDir) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			panic("logdir create failed: " + err.Error())
		}
	}

	logConf := struct {
		FileName string `json:"filename"`
		Level    int    `json:"level"`
		Daily    bool   `json:"daily"`
	}{
		FileName: logDir + "/" + conf.Cfg.Log.FileName + ".log",
		Level:    log.GetLevel(conf.Cfg.Log.Level),
		Daily:    false,
	}

	configuration, err := json.Marshal(logConf)
	if err != nil {
		panic(err)
	}
	log.Init(string(configuration))

	// Init UTXO DB
	utxoConfig := utxo.UtxoConfig{Do: &db.DBOption{FilePath: conf.Cfg.DataDir + "/chainstate", CacheSize: (1 << 20) * 8}}
	utxo.InitUtxoLruTip(&utxoConfig)

	chain.InitGlobalChain()

	// Init blocktree DB
	blkdbCfg := blkdb.BlockTreeDBConfig{Do: &db.DBOption{FilePath: conf.Cfg.DataDir + "/blocks/index", CacheSize: (1 << 20) * 8}}
	blkdb.InitBlockTreeDB(&blkdbCfg)

	persist.InitPersistGlobal()

	// Load blockindex DB
	lblockindex.LoadBlockIndexDB()
	lchain.InitGenesisChain()

	mempool.InitMempool()
	crypto.InitSecp256()
}

func main() {
	args := []string{"--datadir=" + conf.DataDir, "--testnet"}
	appInitMain(args)
	gChain := chain.GetInstance()
	globalBlockIndexMap := make(map[util.Hash]*blockindex.BlockIndex)
	branch := make([]*blockindex.BlockIndex, 0, 20)

	// Load blockindex from DB
	ok := blkdb.GetInstance().LoadBlockIndexGuts(globalBlockIndexMap, gChain.GetParams())
	fmt.Println("load block index guts: ", ok)

	sortedByHeight := make([]*blockindex.BlockIndex, 0, len(globalBlockIndexMap))
	for _, index := range globalBlockIndexMap {
		sortedByHeight = append(sortedByHeight, index)

	}
	//sort by decrease
	sort.SliceStable(sortedByHeight, func(i, j int) bool {
		return sortedByHeight[i].Height < sortedByHeight[j].Height
	})
	for _, index := range sortedByHeight {
		//fmt.Printf("the block index value is:%v\n", index)
		blk, ok := disk.ReadBlockFromDisk(index, gChain.GetParams())
		if ok {
			fmt.Printf("block height is:%v===block size value is:%v\n", index.Height, blk.EncodeSize())
		}
	}
	fmt.Printf("LoadBlockIndexDB, BlockIndexMap len:%d, Branch len:%d, Orphan len:%d\n",
		len(globalBlockIndexMap), len(branch), gChain.ChainOrphanLen())
}
