package day15

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestRoomTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	room := ParseRoom(data)

	t.Log(room)

	room.RunDirections()

	t.Log(room)
	t.Log(room.GPS()) // 2028
}

func TestRoomTest2Part1(t *testing.T) {
	data, err := helper.ReadAll("test2.part1")
	if err != nil {
		t.Fatal(err)
	}

	room := ParseRoom(data)

	t.Log(room)

	room.RunDirections()

	t.Log(room)
	t.Log(room.GPS()) // 10092
}

func TestRoomPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	room := ParseRoom(data)

	t.Log(room)

	room.RunDirections()

	t.Log(room)
	t.Log(room.GPS()) // 1465152
}
