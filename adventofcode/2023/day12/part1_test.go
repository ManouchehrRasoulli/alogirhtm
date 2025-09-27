package day12

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest0Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := readInput(data)
	for _, i := range items {
		t.Log(i)
	}
}
