package day9

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	dataRune := ReadInput(data)

	t.Log(string(dataRune))

	partition := GeneratePartitionPart1(dataRune)
	t.Log(partition)

	CompactPartitionPart1(partition)

	t.Log(partition)

	t.Log(ChecksumPart1(partition)) // 1928
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	dataRune := ReadInput(data)

	t.Log(string(dataRune))

	partition := GeneratePartitionPart1(dataRune)
	CompactPartitionPart1(partition)
	t.Log(partition)
	t.Log(ChecksumPart1(partition)) // 6330095022244
}
