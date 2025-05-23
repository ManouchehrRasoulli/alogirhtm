package day1

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(Add(data)) // 142
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(Add(data)) // 55017
}

func TestInputTestPart2(t *testing.T) {
	data, err := helper.ReadAll("test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(AddWithStr(data)) // 281
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(AddWithStr(data)) // 53539
}
