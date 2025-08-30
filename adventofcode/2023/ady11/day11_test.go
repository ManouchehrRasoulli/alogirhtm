package ady11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestPart1Input1(t *testing.T) {
	input, err := helper.ReadAll("./input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("original universe")
	universe := ReadUniverse(input)
	PrintUniverse(universe)

	t.Log("gravity affected --> expanded universe")
	universe = ExpandUniverse(universe)
	PrintUniverse(universe)
}
