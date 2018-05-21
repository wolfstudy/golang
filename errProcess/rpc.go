package errProcess

type RpcErr int

const (
	Rpc1 RpcErr = iota
	Rpc2
	Rpc3
	Rpc4
	Rpc5
	Rpc6
	Rpc7
)

var rpcDesc = [...]string{
	Rpc1: "rpc11111111",
	Rpc2: "bxxx fdsafdsa",
	Rpc3: "fewafewafewa",
	Rpc4: "fdsafewafewa",
	Rpc5: "fdsafewafewa",
	Rpc6: "fdsafewafewa",
	Rpc7: "fdsafewafewa",
}

func (re RpcErr) String() string {
	return rpcDesc[re]
}
