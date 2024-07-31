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
}
