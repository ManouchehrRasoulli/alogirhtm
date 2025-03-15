package day11

import (
	"log"
	"strconv"
)

type Stones struct {
	numbers []int

	// for part2 if we continue with first solution memory limits will be exceeded
	numbersCount map[int]int
}

func (s *Stones) MultipleBlinksPart1(n int) int {
	log.Println("Start blinking part 1 for :: ", n)

	for i := 0; i < n; i++ {
		log.Println("Blink :: ", i, len(s.numbers))
		s.BlinkPart1()
	}

	return len(s.numbers)
}

func (s *Stones) BlinkPart1() {
	newNumbers := make([]int, 0)

	for _, n := range s.numbers {
		newNumbers = append(newNumbers, ModifyNumberPart(n)...)
	}

	s.numbers = newNumbers
}

func (s *Stones) MultipleBlinksPart2(n int) int {
	log.Println("Start blinking part 2 for :: ", n)

	// use memorization inorder to increase
	// computation resource and also decrease amount of memory usage
	for _, n := range s.numbers {
		s.numbersCount[n] += 1
	}

	for i := 0; i < n; i++ {
		log.Println("Blink :: ", i, len(s.numbers))
		s.BlinkPart2()
	}

	count := 0
	for _, n := range s.numbersCount {
		count += n
	}

	return count
}

func (s *Stones) BlinkPart2() {
	clone := CloneMap(s.numbersCount)
	for stone, count := range clone {
		s.numbersCount[stone] -= count
		numbers := ModifyNumberPart(stone)
		for _, n := range numbers {
			s.numbersCount[n] += count
		}
	}
}

func ModifyNumberPart(i int) []int {
	if i == 0 {
		return []int{1}
	}

	is := strconv.Itoa(i)
	if len(is)%2 == 1 {
		return []int{i * 2024}
	}

	isP1, _ := strconv.Atoi(is[:len(is)/2])
	isP2, _ := strconv.Atoi(is[len(is)/2:])

	return []int{isP1, isP2}
}

func CloneMap(m map[int]int) map[int]int {
	clone := make(map[int]int)
	for k, v := range m {
		clone[k] = v
	}
	return clone
}
