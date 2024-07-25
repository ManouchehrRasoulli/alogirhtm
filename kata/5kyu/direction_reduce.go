package _kyu

import (
	"fmt"
)

var (
	empty = fmt.Errorf("EMPTY")
)

const (
	North string = "NORTH"
	South        = "SOUTH"
	East         = "EAST"
	West         = "WEST"
)

type item struct {
	direction string
	index     int
}

type stack struct {
	items []item
}

func (s *stack) Push(v item) {
	s.items = append(s.items, v)
}

func (s *stack) Pop() (item, error) {
	l := len(s.items)
	if l == 0 {
		return item{}, empty
	}
	item := s.items[l-1]
	s.items = s.items[:l-1]

	return item, nil
}

func DirReduc(arr []string) []string {

	s := stack{
		items: make([]item, 0),
	}

	for i := 0; i < len(arr); i++ {
		it, err := s.Pop()
		if err != nil {
			s.Push(item{
				direction: arr[i],
				index:     i,
			})
			continue
		}
		switch arr[i] {
		case North:
			if it.direction == South { // reduce-able
				fmt.Println("reduce-able north with south index: ", i, it.index)
				continue
			}
		case South:
			if it.direction == North { // reduce-able
				fmt.Println("reduce-able south with north index: ", i, it.index)
				continue
			}
		case West:
			if it.direction == East { // reduce-able
				fmt.Println("reduce-able west with east index: ", i, it.index)
				continue
			}
		case East:
			if it.direction == West { // reduce-able
				fmt.Println("reduce-able west with east index: ", i, it.index)
				continue
			}
		}
		s.Push(it)
		s.Push(item{
			direction: arr[i],
			index:     i,
		})
	}

	reduce := make([]string, 0)
	for i := 0; i < len(s.items); i++ {
		reduce = append(reduce, s.items[i].direction)
	}

	return reduce
}
