package day6

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestParseTestInputPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	room, guard := ParseRoom(data)
	t.Log(room)
	t.Log(guard)

	for guard.Move(room) {
	}

	t.Log(room)
	t.Log(guard)

	t.Log(len(guard.DistinctLocations())) // 41
}

func TestParsePuzzleInputPart1(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	room, guard := ParseRoom(data)
	t.Log(room)
	t.Log(guard)

	count := 0
	for guard.Move(room) {
		count++
	}

	t.Log(room)
	t.Log(guard)

	t.Log(len(guard.DistinctLocations()), count) // 5080, 5641
}
