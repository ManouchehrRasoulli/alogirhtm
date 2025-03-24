package day18

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	p := ReadPuzzle(data, 12, 6, 6)
	t.Log(p)

	t.Log(p.BfsWithPath())

	t.Log(p)
	t.Log(p.PathCount()) // 22
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	p := ReadPuzzle(data, 1024, 70, 70)
	t.Log(p)

	t.Log(p.BfsWithPath())

	t.Log(p)
	t.Log(p.PathCount()) // 416
}

func TestInputTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	filled := 12
	p := ReadPuzzle(data, filled, 6, 6)
	t.Log(p)

	locations := ReadLocations(data)
	for i := filled; i < len(locations); i++ {
		x, y := locations[i].Get()
		p.WriteValue(x, y, usedSpace)
		if !p.Bfs() {
			x, y = locations[i].Get()
			t.Log("no path !!!", y, ",", x) // 6,1
			break
		}
	}
	t.Log(p)
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	filled := 1024
	p := ReadPuzzle(data, filled, 70, 70)
	t.Log(p)

	locations := ReadLocations(data)
	for i := filled; i < len(locations); i++ {
		x, y := locations[i].Get()
		p.WriteValue(x, y, usedSpace)
		if !p.Bfs() {
			x, y = locations[i].Get()
			t.Log("no path !!!", y, ",", x) // 50,23
			break
		}
	}
	t.Log(p)
}
