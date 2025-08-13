package day10

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	t.Log(field, start)
}
