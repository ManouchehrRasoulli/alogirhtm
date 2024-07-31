package kata

import (
	"fmt"
	"math"
)

func minItem(item ...int) int {
	m := item[0]
	for _, x := range item {
		if x < m {
			m = x
		}
	}

	return m
}

func maxItem(item ...int) int {
	m := item[0]
	for _, x := range item {
		if x > m {
			m = x
		}
	}

	return m
}

type Slice[Item any] struct {
	items []Item
}

func NewSlice[Item any]() Slice[Item] {
	s := Slice[Item]{
		items: make([]Item, 0),
	}
	return s
}

func (s *Slice[Item]) Len() int {
	return len(s.items)
}

func (s *Slice[Item]) Push(item Item) {
	s.items = append(s.items, item)
}

func (s *Slice[Item]) At(index int) (Item, error) {
	var item Item
	if index < 0 || index >= s.Len() || s.Len() == 0 {
		return item, fmt.Errorf("index out of range")
	}
	item = s.items[index]
	return item, nil
}

func (s *Slice[Item]) Pop(index int) (Item, error) {
	item, err := s.At(index)
	if err != nil {
		return item, err
	}

	s.items = append(s.items[:index], s.items[index+1:]...)
	return item, nil
}

type Building struct {
	MaxFloor int
	MinFloor int
	Floor    map[int]*Slice[Person]
}

func NewBuilding() Building {
	b := Building{
		MaxFloor: math.MinInt,
		MinFloor: math.MaxInt,
		Floor:    make(map[int]*Slice[Person]),
	}
	return b
}

func (b *Building) Add(p Person) {
	b.MinFloor = minItem(b.MinFloor, p.From, p.To)
	b.MaxFloor = maxItem(b.MaxFloor, p.From, p.To)
	if _, ok := b.Floor[p.From]; !ok {
		b.Floor[p.From] = &Slice[Person]{
			items: make([]Person, 0),
		}
	}
	b.Floor[p.From].Push(p)
}

func (b *Building) PersonsMovingUp(floor int, count int) []Person {
	panic("not implemented")
}

const (
	Up   = true
	Down = false
)

type Elevator struct {
	capacity  int
	members   Slice[Person]
	building  *Building
	direction bool
}

func Order(level int, queue []Person) []int {
	return []int{}
}

type Person struct {
	From int
	To   int
}
