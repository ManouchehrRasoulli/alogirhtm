package day16

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestMazeTest0Part1(t *testing.T) {
	data, err := helper.ReadAll("test0.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart1()) // 2008
}

func TestMazeTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart1()) // 7036
}

func TestMazeTest2Part1(t *testing.T) {
	data, err := helper.ReadAll("test2.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart1()) // 11048
}

func TestMazePuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart1()) // 135536
}

func TestMazeTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart2()) // 7036
	t.Log(maze.PathCount())  // 45

	t.Log(maze)
}

func TestMazeTest2Part2(t *testing.T) {
	data, err := helper.ReadAll("test2.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart2()) // 11048
	t.Log(maze.PathCount())  // 64

	t.Log(maze)
}

func TestMazePuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.SolvePart2()) // 135536
	t.Log(maze.PathCount())  // 583

	t.Log(maze)
}
