package day6

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
	"testing"
)

func TestParseTestInputPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	room, guard := ParseRoomWithSimpleGuard(data)
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

	room, guard := ParseRoomWithSimpleGuard(data)
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

func TestParseTestInputPart2(t *testing.T) {
	data, err := helper.ReadAll("input_test_part2.txt")
	if err != nil {
		t.Fatal(err)
	}

	room, guard := ParseRoomWithCircuitGuard(data)
	t.Log(room)
	t.Log(guard)

	log.Println(guard.CountCircuits(room)) // 6

	t.Log(room)
	t.Log(guard)
}

func TestParsePuzzleInputPart2(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	room, guard := ParseRoomWithCircuitGuard(data)
	t.Log(room)
	t.Log(guard)

	log.Println(guard.CountCircuits(room))

	t.Log(room)
	t.Log(guard)
}
