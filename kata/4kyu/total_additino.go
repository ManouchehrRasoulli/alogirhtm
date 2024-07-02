package _kyu

import (
	"fmt"
)

var (
	cache           map[int]int
	allPermutations [][]int
)

func join(nums []int) string {
	s := ""
	for _, n := range nums {
		s = fmt.Sprintf("%s%d", s, n)
	}
	return s
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func areNumsZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func isInList(i int, list []int) bool {
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return false
}

func reverseNumber(n int) int {
	x := make([]int, 0)
	for n != 0 {
		x = append(x, n%10)
		n /= 10
	}

	for i := 1; i <= len(x); i++ {
		n = (n * 10) + x[i-1]
	}

	return n
}

func makeList(limit int, number []int) []int {
	list := make([]int, 0)

	for i, v := range number {
		number[i] = reverseNumber(v)
	}

isZero:
	for !areNumsZero(number) {
		for i, v := range number {
			if v != 0 && !isInList(v%10, list) {
				list = append(list, v%10)
			}
			number[i] = v / 10
			if len(list) == limit {
				break isZero
			}
		}
	}

	return list
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	if v, ok := cache[n]; ok {
		return v
	}
	t := 1
	for ; n != 0; n-- {
		t *= n
	}
	cache[n] = t
	return t
}

func totalPermutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func cp(nums []int) []int {
	c := make([]int, len(nums), len(nums))
	for i, n := range nums {
		c[i] = n
	}
	return c
}

func cpn(nums []int, n int) []int {
	if n > len(nums) {
		return nums
	}
	c := make([]int, n, n)
	for i := 0; i < n; i++ {
		c[i] = nums[i]
	}
	return c
}

func permutation(numbers []int, j int) {
	if j == len(numbers) {
		allPermutations = append(allPermutations, cp(numbers))
		return
	}
	for i := j; i < len(numbers); i++ {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		permutation(numbers, j+1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func permutationSum(n int) int {
	x := make(map[string]struct{})
	total := 0
	for _, p := range allPermutations {
		cpnn := cpn(p, n)
		j := join(cpnn)
		if _, ok := x[j]; !ok {
			x[j] = struct{}{}
			total += sum(cpnn)
		}
	}
	return total
}

func Gta(limit int, numbers []int) int {
	cache = make(map[int]int)
	allPermutations = make([][]int, 0)

	fmt.Println(limit, numbers)
	list := makeList(limit, numbers)
	if limit != len(list) {
		fmt.Println(limit, numbers, list)
		return 0
	}
	fmt.Println(list)

	permutation(list, 0)
	fmt.Println("all-permutations", len(allPermutations))

	total := 0
	for i := 1; i <= len(list); i++ {
		total += permutationSum(i)
	}

	fmt.Println("total", total)

	cache = nil
	allPermutations = nil
	return total
}
