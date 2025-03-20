package helper

import (
	"container/heap"
	"log"
)

type Comparable interface {
	Value() int
}

type PriorityQueue []Comparable

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].Value() < (*p)[j].Value()
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x any) {
	cp, ok := x.(Comparable)
	if !ok {
		log.Fatal("not comparable type")
	}

	*p = append(*p, cp)
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[0]
	*p = old[1:n]
	return item
}

var _ heap.Interface = &PriorityQueue{}
