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

func ChecksumPart1(partition []int) int {
	checksum := 0

	for i, id := range partition {
		if id != -1 {
			checksum += id * i
		}
	}

	return checksum
}
