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

	p.Bfs()

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

	p.Bfs()

	t.Log(p)
	t.Log(p.PathCount()) // 416
}
