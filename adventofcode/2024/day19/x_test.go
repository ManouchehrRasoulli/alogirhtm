package day19

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTest1Part1Naive(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	patterns, towels := ParseInput(data)

	matchCount := 0
	for _, towel := range towels {
		if NaiveMatch(towel, patterns) {
			matchCount++
			t.Log(towel)
		}
	}

	t.Log(matchCount) // 6
}

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	patterns, towels := ParseInput(data)

	matchCount := 0
	for _, towel := range towels {
		if Match(towel, patterns) {
			matchCount++
			t.Log(towel)
		}
	}

	t.Log(matchCount) // 6
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	patterns, towels := ParseInput(data)

	matchCount := 0
	for i, towel := range towels {
		if Match(towel, patterns) {
			matchCount++
			t.Log(i, towel)
		}
	}

	t.Log(matchCount) // 327
}

func TestInputTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	patterns, towels := ParseInput(data)

	matchCount := 0
	for _, towel := range towels {
		matchCount += MatchCount(towel, patterns)
	}

	t.Log(matchCount) // 16
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	patterns, towels := ParseInput(data)

	matchCount := 0
	for _, towel := range towels {
		matchCount += MatchCount(towel, patterns)
	}

	t.Log(matchCount) // 772696486795255
}
