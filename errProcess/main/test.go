package main

import (
	"fmt"
	"golang/errProcess"
)

func f1() error {
	return errProcess.New(errProcess.D)
}

func f2(a int) (int, error) {
	if a%2 == 0 {
		return 0, errProcess.New(errProcess.Rpc5)
	}
	return 15, nil
}

func main() {
	if err := f1(); err != nil {
		fmt.Println("got error", err)
	}
	v, err := f2(122)
	if err != nil && errProcess.IsErrorCode(err, errProcess.Rpc5) {
		fmt.Println("got error", err)
		return
	}
	fmt.Printf("got v = %v\n", v)
}