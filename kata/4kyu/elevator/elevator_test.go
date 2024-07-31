package kata

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s := NewSlice[Person]()

	p := Person{
		From: 1,
		To:   10,
	}
	s.Push(p)
	p.To = 9
	s.Push(p)
	p.From = 3
	s.Push(p)

	t.Log(s)

	t.Log(s.At(2))
	t.Log(s.At(4))

	t.Log(s.Pop(1))
	t.Log(s)

	t.Log(s.Pop(1))
	t.Log(s)

	t.Log(s.Pop(0))
	t.Log(s.Pop(0))
}

func TestBuilding(t *testing.T) {
	b := NewBuilding()
	p := Person{
		From: 1,
		To:   10,
	}
	b.Add(p)
	p.To = 9
	b.Add(p)
	p.From = 3
	b.Add(p)

	t.Log(b)
	t.Log(b.PersonsInDirection(1, 3, Up))
	t.Log(b.AnyPersonLeft())

	t.Log(b.PersonsInDirection(3, 0, Up))
	t.Log(b.AnyPersonLeft())
}

func TestElevator(t *testing.T) {
	b := NewBuilding()
	p := Person{
		From: 1,
		To:   4,
	}
	b.Add(p)
	p.To = 2
	b.Add(p)
	p.From = 2
	p.To = 4
	b.Add(p)

	e := NewElevator(1, Up, 5, b)
	t.Log(e.NextStop())
	t.Log(e.NextStop())
	t.Log(e.NextStop())
	t.Log(e.NextStop())
	t.Log(e.NextStop())
	t.Log(e.NextStop())
	t.Log(e.NextStop())
}

func TestOrder1(t *testing.T) {
	startingFloor := 1
	puzzle := []Person{
		{From: 3, To: 2}, // Al
		{From: 5, To: 2}, // Betty
		{From: 2, To: 1}, // Charles
		{From: 2, To: 5}, // Dan
		{From: 4, To: 3}, // Ed
	}
	solution := []int{2, 5, 4, 3, 2, 1}

	s := Order(startingFloor, puzzle)

	t.Log(s)
	t.Log(solution)
}
