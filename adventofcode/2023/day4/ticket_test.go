package day4

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	tickets := ReadTickets(data)

	total := 0
	for _, ticket := range tickets {
		total += ticket.PointsPart1()
	}

	t.Log(total) // 13
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	tickets := ReadTickets(data)

	total := 0
	for _, ticket := range tickets {
		total += ticket.PointsPart1()
	}

	t.Log(total) // 19135
}
