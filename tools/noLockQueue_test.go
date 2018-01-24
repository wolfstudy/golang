package main

import (
	"testing"
	"fmt"
	"sync/atomic"
	"time"
	"runtime"
	"sync"
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

func testQueuePutGet(t *testing.T, grp, cnt int) (
	put time.Duration, get time.Duration) {
	var wg sync.WaitGroup
	var id int32
	wg.Add(grp)
	q := NewQueue(1024 * 1024)
	start := time.Now()
	for i := 0; i < grp; i++ {
		go func(g int) {
			defer wg.Done()
			for j := 0; j < cnt; j++ {
				val := fmt.Sprintf("Node.%d.%d.%d", g, j, atomic.AddInt32(&id, 1))
				ok, _ := q.Put(&val)
				for !ok {
					time.Sleep(time.Microsecond)
					ok, _ = q.Put(&val)
				}
			}
		}(i)
	}
	wg.Wait()
	end := time.Now()
	put = end.Sub(start)

	wg.Add(grp)
	start = time.Now()
	for i := 0; i < grp; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < cnt; {
				_, ok, _ := q.Get()
				if !ok {
					runtime.Gosched()
				} else {
					j++
				}
			}
		}()
	}
	wg.Wait()
	end = time.Now()
	get = end.Sub(start)
	if q := q.Quantity(); q != 0 {
		t.Errorf("Grp:%v, Quantity Error: [%v] <>[%v]", grp, q, 0)
	}
	return put, get
}
