package main

import (
	"testing"
)

var (
	value int = 1
)

func TestQueue(t *testing.T) {
	q := NewQueue(8)
	ok, quantity := q.Put(&value)
	if !ok {
		t.Error("TestStack Get.Fail")
		return
	} else {
		t.Logf("TestStack Put value:%d[%v], quantity:%v\n", &value, value, quantity)
	}

	val, ok, quantity := q.Get()
	if !ok {
		t.Error("TestStack Get.Fail")
		return
	} else {
		t.Logf("TestStack Get value:%d[%v], quantity:%v\n", val, *(val.(*int)), quantity)
	}
	if q := q.Quantity(); q != 0 {
		t.Errorf("Quantity Error: [%v] <>[%v]", q, 0)
	}
}
