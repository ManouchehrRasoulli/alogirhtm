package helper

import (
	"container/heap"
	"testing"
)

type Item struct {
	v int
}

func (i *Item) Value() int {
	return i.v
}

var _ Comparable = &Item{}

func TestPriorityQueue(t *testing.T) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Item{v: 2})
	heap.Push(pq, &Item{v: 3})
	heap.Push(pq, &Item{v: 4})
	heap.Push(pq, &Item{v: 1})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		t.Log(item)
	}
}
