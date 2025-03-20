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

	t.Log(maze.Solve()) // 2008
}

func TestMazeTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.Solve()) // 7036
}

func TestMazeTest2Part1(t *testing.T) {
	data, err := helper.ReadAll("test2.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.Solve()) // 11048
}

func TestMazePuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	maze := MazeFromData(data)
	t.Log(maze)

	t.Log(maze.Solve()) // 135536
}
