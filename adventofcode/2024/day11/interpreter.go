package day11

import "strconv"

type Stones struct {
	numbers []int
}

func (s *Stones) MultipleBlinksPart1(n int) int {
	for i := 0; i < n; i++ {
		s.BlinkPart1()
	}

	return len(s.numbers)
}

func (s *Stones) BlinkPart1() {
	newNumbers := make([]int, 0)

	for _, n := range s.numbers {
		newNumbers = append(newNumbers, ModifyNumberPart1(n)...)
	}

	s.numbers = newNumbers
}

func ModifyNumberPart1(i int) []int {
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
