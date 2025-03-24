package day19

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

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
