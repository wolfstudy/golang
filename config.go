package main

import (
	"runtime"
	"fmt"
)

type config struct {
	ShowVersion  bool   `short:"V" long:"version" description:"Display version information and exit"`
	ListCommands bool   `short:"l" long:"listcommands" description:"List all of the supported commands and exit"`
	ConfigFile   string `short:"C" long:"configfile" description:"Path to configuration file"`
	RPCUser      string `short:"u" long:"rpcuser" description:"RPC username"`
	RPCPassword  string `short:"P" long:"rpcpass" default-mask:"-" description:"RPC password"`
	RPCServer    string `short:"s" long:"rpcserver" description:"RPC server to connect to"`
	RPCCert      string `short:"c" long:"rpccert" description:"RPC server certificate chain for validation"`
}

func main() {
	i := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(i)
}
