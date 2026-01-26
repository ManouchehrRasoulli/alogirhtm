package day10

import "testing"

func TestMachine(t *testing.T) {
	m, err := ParseMachine("[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)

	m, err = ParseMachine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}

func TestMachine_Push(t *testing.T) {
	m, err := ParseMachine("[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("current")
	t.Log(m)

	t.Log("push 1")
	m.PushButton(1)
	t.Log(m)

	t.Log("push 0")
	m.PushButton(0)
	t.Log(m)

	t.Log("reset")
	m.Reset()
	t.Log(m)
}

func TestMachine_Push2(t *testing.T) {
	m, err := ParseMachine("[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("current")
	t.Log(m)

	t.Log("push 2")
	m.PushButton(2)
	t.Log(m)

	t.Log("push 3")
	m.PushButton(3)
	t.Log(m)

	t.Log("push 4")
	m.PushButton(4)
	t.Log(m)

	t.Log("reset")
	m.Reset()
	t.Log(m)
}
