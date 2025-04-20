package day6

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	input, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	time, distance := ReadRaceAndTiming(input)
	t.Log(time, distance)

	t.Log(Part1(time, distance)) // 288
}

func TestInputPuzzlePart1(t *testing.T) {
	input, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	time, distance := ReadRaceAndTiming(input)
	t.Log(time, distance)

	t.Log(Part1(time, distance)) // 2269432
}

func TestInputTestPart2(t *testing.T) {
	input, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	time, distance := ReadRaceAndTimingJoined(input)
	t.Log(time, distance)

	t.Log(Part2(time, distance)) // 71503
}

func TestInputPuzzlePart2(t *testing.T) {
	input, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	time, distance := ReadRaceAndTimingJoined(input)
	t.Log(time, distance)

	t.Log(Part2(time, distance)) // 35865985
}
