package day10

import (
	"container/heap"
	"fmt"
	"strings"
)

type State struct {
	val []int
}

type Item struct {
	state State
	cost  int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func key(v []int) string {
	return fmt.Sprint(v)
}

func minPresses(buttons []button, target []int) int {
	start := make([]int, len(target))
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{state: State{start}, cost: 0})

	dist := make(map[string]int)
	dist[key(start)] = 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		cur := item.state.val
		cost := item.cost

		if key(cur) == key(target) {
			return cost
		}

		for _, b := range buttons {
			next := make([]int, len(cur))
			copy(next, cur)

			valid := true
			for _, idx := range b {
				next[idx]++
				if next[idx] > target[idx] {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			k := key(next)
			if d, ok := dist[k]; !ok || cost+1 < d {
				dist[k] = cost + 1
				heap.Push(pq, &Item{
					state: State{next},
					cost:  cost + 1,
				})
			}
		}
	}
	return -1
}

func Part2(input string) int {
	var (
		lines = strings.Split(input, "\n")
		total = 0
	)

	for _, line := range lines {
		m, err := ParseMachine(line)
		if err != nil {
			panic(err)
		}

		total += minPresses(m.buttons, m.joltage)
	}
	return total
}
