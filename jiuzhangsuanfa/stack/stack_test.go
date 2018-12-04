package stack

import (
	"fmt"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := NewStack()

	s.Push("1")
	s.Push("2")

	e := s.Pop()
	fmt.Println(e)
}
