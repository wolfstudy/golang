package errProcess

import (
	"fmt"
)

type MemPoolErr int

const (
	A MemPoolErr = iota
	B
	C
	D
	E
	F
	G
)

var merrToString = map[MemPoolErr]string{
	A: "fffffffffffff",
	B: "bxxx fdsafdsa",
	C: "fewafewafewa",
	D: "fdsafewafewa",
	E: "fdsafewafewa",
	F: "fdsafewafewa",
	G: "fdsafewafewa",
}

func (me MemPoolErr) String() string {
	if s, ok := merrToString[me]; ok {
		return s
	}
	return fmt.Sprintf("Unknown code (%d)", me)
}
