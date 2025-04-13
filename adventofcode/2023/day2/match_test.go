package day2

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")
	games := make([]*Game, 0)
	for _, line := range lines {
		g := ReadGame(line)
		games = append(games, g)
	}

	count := 0
	for _, g := range games {
		count += g.Match(12, 14, 13)
	}

	t.Log(count) // 8
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")
	games := make([]*Game, 0)
	for _, line := range lines {
		g := ReadGame(line)
		games = append(games, g)
	}

	count := 0
	for _, g := range games {
		count += g.Match(12, 14, 13)
	}

	t.Log(count) // 2285
}
