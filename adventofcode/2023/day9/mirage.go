package day9

import (
	"log"
	"strconv"
	"strings"
)

func ReadData(input string) (data [][]int) {
	var (
		lines = strings.Split(input, "\n")
	)

	for _, line := range lines {
		lineItemStr := strings.Split(line, " ")
		lineItemInt := make([]int, 0)
		for _, i := range lineItemStr {
			itemInt, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}

			lineItemInt = append(lineItemInt, itemInt)
		}

		data = append(data, lineItemInt)
	}

	return data
}

func findNextNaivePart1(data []int) []int {
	var (
		isAllZero = func(data []int) bool {
			for _, i := range data {
				if i != 0 {
					return false
				}
			}

			return true
		}

		generateNextData = func(data []int) []int {
			output := make([]int, 0)
			for i := 1; i < len(data); i++ {
				output = append(output, data[i]-data[i-1])
			}

			return output
		}
	)

	if isAllZero(data) { // last item is zero appended
		data = append(data, 0)
		return data
	}

	nextDataOutput := findNextNaivePart1(generateNextData(data))
	data = append(data, data[len(data)-1]+nextDataOutput[len(data)-1])
	return data
}

func Part1(items [][]int) int {
	var (
		sum = 0
	)

	for _, item := range items {
		data := findNextNaivePart1(item)
		sum += data[len(data)-1]
	}

	return sum
}

func findPreviousNaivePart2(data []int) []int {
	var (
		isAllZero = func(data []int) bool {
			for _, i := range data {
				if i != 0 {
					return false
				}
			}

			return true
		}

		generateNextData = func(data []int) []int {
			output := make([]int, 0)
			for i := 1; i < len(data); i++ {
				output = append(output, data[i]-data[i-1])
			}

			return output
		}
	)

	if isAllZero(data) { // last item is zero appended
		newData := make([]int, 0)
		newData = append(newData, 0)
		return append(newData, data...)
	}

	nextDataOutput := findPreviousNaivePart2(generateNextData(data))
	newData := make([]int, 0)
	newData = append(newData, data[0]-nextDataOutput[0])
	return newData
}

func Part2(items [][]int) int {
	var (
		sum = 0
	)

	for _, item := range items {
		data := findPreviousNaivePart2(item)
		sum += data[0]
	}

	return sum
}
