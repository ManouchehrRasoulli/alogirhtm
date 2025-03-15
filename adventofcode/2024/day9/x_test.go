package day9

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
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

func TestNextEmptyBlockWithSize(t *testing.T) {
	data := []int{0, 0, -1, -1, 12, 12, 12, -1, -1, -1}
	start, end := NextEmptyBlockWithSize(2, data)
	t.Log(start, end) // 2, 3
	start, end = NextEmptyBlockWithSize(3, data)
	t.Log(start, end) // 7, 9
	start, end = NextEmptyBlockWithSize(4, data)
	t.Log(start, end) // -1, -1
}

func TestNextDataBlock(t *testing.T) {
	data := []int{0, 0, -1, -1, 12, 12, 12, -1, -1, -1}
	start, end := NextDataBlock(len(data), data)
	t.Log(start, end) // 4, 6
	start, end = NextDataBlock(start-1, data)
	t.Log(start, end) // 0, 1
	start, end = NextDataBlock(start-1, data)
	t.Log(start, end) // -1, -1
}

func TestInputTestPart2(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	dataRune := ReadInput(data)
	partition := GeneratePartitionPart1(dataRune)
	CompactPartitionPart2(partition)
	log.Println(ChecksumPart1(partition)) // 2858
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	dataRune := ReadInput(data)

	partition := GeneratePartitionPart1(dataRune)
	CompactPartitionPart2(partition)
	t.Log(ChecksumPart1(partition)) // 6359491814941
}
