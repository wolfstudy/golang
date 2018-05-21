package errProcess

import (
	"fmt"
)

const errDescFmt string = "module: [%s], inner err desc: [%s]"

type ProjectError struct {
	Module string
	Code   int
	Desc   string
}

func (e ProjectError) Error() string {
	return fmt.Sprintf("module: %s, global errcode: %v,  errdesc: %s", e.Module, e.Code, e.Desc)
}

func getFinalCode(errCode fmt.Stringer) (int, string) {
	code := 0
	name := ""

	switch t := errCode.(type) {
	case RpcErr:
		code = int(t) + 1000
		name = "rpc"
	case MemPoolErr:
		name = "mempool"
		code = int(t) + 2000
	default:
	}

	return code, name
}

func IsErrorCode(err error, errCode fmt.Stringer) bool {
	e, ok := err.(ProjectError)
	icode, _ := getFinalCode(errCode)
	return ok && icode == e.Code
}

func New(errCode fmt.Stringer) error {
	code, name := getFinalCode(errCode)

	return ProjectError{
		Module: name,
		Code:   code,
		Desc:   errCode.String(),
	}
}