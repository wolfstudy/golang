package skiplist

import (
	"math"
	"math/rand"
	"time"
	
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type node struct {
	Lesser
	nexts []*node
}

type skiplist struct {
	header node
	size   int
}

// New create a skiplist instance
func New(estimatedNodeCount int) MapContainer {
	if estimatedNodeCount < 2 {
		estimatedNodeCount = 2
	}
	return &skiplist{header: node{nexts: make([]*node,
		int64(math.Ceil(math.Log2(float64(estimatedNodeCount)))))}}
}

func isEqual(ls1 Lesser, ls2 Lesser) bool {
	return !ls1.Less(ls2) && !ls2.Less(ls1)
}

// ReplaceOrInsert insert a element or replace the element if it already exists
func (sl *skiplist) ReplaceOrInsert(elem Lesser) Lesser {
	insertToHeight := calcInsertionHeight(len(sl.header.nexts))
	newNode := &node{elem, make([]*node, insertToHeight)}
	prevs, removed := sl.removeIfAny(elem)
	if removed == nil {
		sl.size++
	}

	for i := 0; i < len(newNode.nexts); i++ {
		newNode.nexts[i] = prevs[i].nexts[i]
		prevs[i].nexts[i] = newNode
	}
	return newNode
}

func (sl *skiplist) getMaxHeight() int {
	return len(sl.header.nexts)
}

func (sl *skiplist) searchPrev(target Lesser) []*node {
	prevs := make([]*node, sl.getMaxHeight()+1)
	prevs[len(prevs)-1] = &sl.header
	for h := sl.getMaxHeight() - 1; h >= 0; h-- {
		prevs[h] = prevs[h+1]
		for pv := prevs[h].nexts[h]; pv != nil && pv.Lesser.Less(target); pv = pv.nexts[h] {
			prevs[h] = pv
		}
	}
	return prevs[:len(prevs)-1]
}

// Search search an element in skiplist
func (sl *skiplist) Search(target Lesser) (Lesser, bool) {
	prevs := sl.searchPrev(target)
	if prevs[0].nexts[0] != nil && isEqual(prevs[0].nexts[0].Lesser, target) {
		return prevs[0].nexts[0].Lesser, true
	}
	return nil, false
}
func remove(prevs []*node, target Lesser) (deleted *node) {
	for i := 0; i < len(prevs) && prevs[i].nexts[i] != nil && isEqual(prevs[i].nexts[i].Lesser, target); i++ {
		deleted = prevs[i].nexts[i]
		prevs[i].nexts[i] = prevs[i].nexts[i].nexts[i]
	}
	return
}
func (sl *skiplist) removeIfAny(target Lesser) (prevs []*node, deleted *node) {
	prevs = sl.searchPrev(target)
	deleted = remove(prevs, target)
	return
}

func calcInsertionHeight(maxHeight int) int {
	res := 1
	for i := 0; i < maxHeight-1; i++ {
		upDown := rand.Intn(2)
		if upDown == 0 {
			break
		}
		res++
	}
	return res
}

// Delete delete the target from skiplist
func (sl *skiplist) Delete(target Lesser) (deleted Lesser, found bool) {
	_, deletedNode := sl.removeIfAny(target)
	if deletedNode != nil {
		sl.size--
		return deletedNode.Lesser, true
	}
	return nil, false
}

// Min get the element with minimum value
func (sl *skiplist) Min() (Lesser, bool) {
	if sl.header.nexts[0] == nil {
		return nil, false
	}
	return sl.header.nexts[0].Lesser, true
}

// Ascend loop over all element and run @iterHandler
func (sl *skiplist) Ascend(iterHandler func(i Lesser) bool) {
	for next := sl.header.nexts[0]; next != nil &&
		iterHandler(next.Lesser); next = next.nexts[0] {
	}
}

// Len return element count in skiplist
func (sl *skiplist) Len() int {
	return sl.size
}

func (sl *skiplist) max() ([]*node, bool) {
	prevs := make([]*node, sl.getMaxHeight()+1)
	prevs[len(prevs)-1] = &sl.header
	for h := sl.getMaxHeight() - 1; h >= 0; h-- {
		prevs[h] = prevs[h+1]
		for pv := prevs[h].nexts[h]; pv != nil &&
			pv.nexts[h] != nil; pv = pv.nexts[h] {

			prevs[h] = pv
		}
	}
	if sl.size != 0 {
		return prevs[:len(prevs)-1], true
	}
	return nil, false
}

// Max return maximum element in skiplist
func (sl *skiplist) Max() (less Lesser, found bool) {
	prevs, found := sl.max()
	if found {
		return prevs[0].nexts[0].Lesser, true
	}
	return nil, false
}

// DeleteMax delete the element with maximum value.
func (sl *skiplist) DeleteMax() (less Lesser, found bool) {
	prevs, found := sl.max()
	if found {
		sl.size--
		maxLesser := prevs[0].nexts[0].Lesser
		return remove(prevs, maxLesser).Lesser, true
	}
	return nil, false
}
