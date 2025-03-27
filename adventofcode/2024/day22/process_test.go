package day22

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strconv"
	"strings"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	numbers := make([]int, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	t.Log(part1(numbers)) // 37327623
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	numbers := make([]int, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	t.Log(part1(numbers)) // 19927218456
}
