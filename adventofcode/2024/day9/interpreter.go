package day9

import (
	"log"
	"strconv"
)

func GeneratePartitionPart1(data []rune) []int {
	var (
		partition = make([]int, 0)
	)

	for i, v := range data {
		vn, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}

		if i%2 != 0 { // free space
			for range vn {
				partition = append(partition, -1)
			}
			continue
		}

		// data
		for range vn {
			partition = append(partition, i/2)
		}
	}

	return partition
}

func CompactPartitionPart1(data []int) {
	var (
		i = 0
		j = len(data) - 1
	)

	for {
		if i > j {
			break
		}

		if data[i] != -1 {
			i++
			continue
		}

		if data[j] == -1 {
			j--
			continue
		}

		data[i], data[j] = data[j], data[i]
	}
}

func CompactPartitionPart2(data []int) {
	var (
		blockStart = len(data)
		blockEnd   = len(data)

		emptyStart = 0
		emptyEnd   = 0
	)

	for {
		blockStart, blockEnd = NextDataBlock(blockStart-1, data)
		if blockStart == -1 || blockEnd == -1 {
			break // no file left
		}

		fileSize := blockEnd + 1 - blockStart

		emptyStart, emptyEnd = NextEmptyBlockWithSize(fileSize, data)
		if emptyStart == -1 || emptyEnd == -1 {
			continue
		}

		if emptyStart > blockEnd {
			// don't copy
			continue
		}

		i := 0
		for j := emptyStart; j <= emptyEnd && blockStart+i <= blockEnd; j++ {
			data[j], data[blockStart+i] = data[blockStart+i], data[j]
			i++
		}
	}
}

func NextEmptyBlockWithSize(size int, data []int) (int, int) {
	var (
		start = -1
		end   = -1
	)

	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			start = i
		} else {
			continue
		}
		for j := i; j < len(data); j++ {
			if data[j] == -1 {
				end = j
			} else {
				break
			}
		}

		if end+1-start >= size {
			return start, end
		}
	}

	return -1, -1
}

func NextDataBlock(from int, data []int) (int, int) {
	var (
		start = 0
		end   = 0
	)

	if from >= len(data) {
		from = len(data) - 1
	}

	for i := from; i >= 0; i-- {
		if data[i] == -1 {
			continue
		}

		end = i
		for j := i; j >= 0; j-- {
			if data[j] == data[i] {
				continue
			}

			start = j + 1
			break
		}

		return start, end
	}

	return -1, -1
}

func ChecksumPart1(partition []int) int {
	checksum := 0

	for i, id := range partition {
		if id != -1 {
			checksum += id * i
		}
	}

	return checksum
}
