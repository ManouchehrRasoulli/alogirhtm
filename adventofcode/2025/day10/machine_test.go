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
