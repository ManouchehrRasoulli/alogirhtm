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
	floors   map[int]*Slice[Person]
}

func NewBuilding() Building {
	b := Building{
		MaxFloor: math.MinInt,
		MinFloor: math.MaxInt,
		floors:   make(map[int]*Slice[Person]),
	}
	return b
}

func (b *Building) Add(p Person) {
	b.MinFloor = minItem(b.MinFloor, p.From, p.To)
	b.MaxFloor = maxItem(b.MaxFloor, p.From, p.To)
	if _, ok := b.floors[p.From]; !ok {
		s := NewSlice[Person]()
		b.floors[p.From] = &s
	}
	b.floors[p.From].Push(p)
}

func (b *Building) AnyPersonLeft() bool {
	for _, f := range b.floors {
		if f.Len() > 0 {
			return true
		}
	}

	return false
}

func (b *Building) PersonsInDirection(floor int, count int, direction bool) []Person {
	persons := b.floors[floor]
	if persons == nil {
		return nil
	}

	p := make([]Person, 0)
	for i := 0; i < persons.Len() && len(p) < count; {
		person, err := persons.At(i)
		if err != nil {
			break
		}

		if (direction == Up && person.From < person.To) ||
			(direction == Down && person.From > person.To) { // person moving up
			p = append(p, person)
			_, _ = persons.Pop(i)
			continue
		}

		i++
	}

	return p
}

const (
	Up   = true
	Down = false
)

type Elevator struct {
	currentFloor int
	capacity     int
	onboard      int
	to           Slice[Person]
	building     Building
	direction    bool
}

func NewElevator(currentFloor int, direction bool, capacity int, building Building) Elevator {
	e := Elevator{
		currentFloor: currentFloor,
		capacity:     capacity,
		onboard:      0,
		to:           NewSlice[Person](),
		building:     building,
		direction:    direction,
	}

	return e
}

func (e *Elevator) HowManyCanBoard() int {
	return e.capacity - e.onboard
}

func (e *Elevator) ToNextFloor() {
	if e.direction == Up {
		e.currentFloor++
	} else {
		e.currentFloor--
	}
	if e.currentFloor > e.building.MaxFloor {
		e.direction = Down
		e.currentFloor = e.building.MaxFloor
	}
	if e.currentFloor < e.building.MinFloor {
		e.direction = Up
		e.currentFloor = e.building.MinFloor
	}
}

func (e *Elevator) AnyOneCouldBoard() bool {
	persons := e.building.PersonsInDirection(e.currentFloor, e.HowManyCanBoard(), e.direction)
	for _, p := range persons {
		e.to.Push(p)
		e.onboard++
		fmt.Printf("enter << %+v, floor %d, capacity %d\n", p, e.currentFloor, e.HowManyCanBoard())
	}
	return len(persons) != 0
}

func (e *Elevator) AnyOneCouldExit() bool {
	x := e.onboard
	anyOneExited := false
	index := 0

	for i := 0; i < x; i++ {
		p, err := e.to.At(index)
		if err != nil {
			break
		}

		if p.To == e.currentFloor {
			fmt.Printf("exit >> %+v - floor %d\n", p, e.currentFloor)
			_, _ = e.to.Pop(index)
			e.onboard--
			anyOneExited = true
			continue
		}

		index++
	}

	return anyOneExited
}

func (e *Elevator) NextStop() (int, error) {
	for e.building.AnyPersonLeft() || e.onboard != 0 {
		if e.AnyOneCouldBoard() || e.AnyOneCouldExit() {
			return e.currentFloor, nil
		}

		e.ToNextFloor()
	}

	return -1, fmt.Errorf("noone left")
}

func Order(level int, queue []Person) []int {
	fmt.Printf("levle %d, queue %+v\n", level, queue)

	b := NewBuilding()
	for _, p := range queue {
		b.Add(p)
	}
	e := NewElevator(level, Up, 5, b)

	stops := make([]int, 0)
	for {
		lv, err := e.NextStop()
		if err != nil {
			break
		}

		stops = append(stops, lv)
	}

	reducedStops := make([]int, 0)
	for i := 1; i < len(stops); i++ {
		if stops[i] != stops[i-1] {
			reducedStops = append(reducedStops, stops[i-1], stops[i])
			i += 2
		}
	}

	return reducedStops
}

type Person struct {
	From int
	To   int
}
