package queue

import (
	"fmt"
	"testing"
)

func TestQueue_DeQueue(t *testing.T) {
	q := NewQueue()
	q.EnQueue("1")
	q.EnQueue("2")

	str := q.DeQueue()
	fmt.Println(str)//1
	str1 := q.DeQueue()
	fmt.Println(str1)//2

	n := q.DeQueue()
	fmt.Println(n)//nil
}
