package _kyu

import (
	"fmt"
	"testing"
)

func Test_join(t *testing.T) {
	s := []int{1, 2, 3, 4}
	t.Log(join(s))
	t.Log(join(s[:2]))
}

func Test_reverse(t *testing.T) {
	t.Log(reverseNumber(127391))
}

func Test_makeList(t *testing.T) {
	t.Log(makeList(2, []int{12, 11, 12318}))
	t.Log(makeList(10, []int{99841, 57794, 586}))
	t.Log(makeList(9, []int{123489, 5, 67}))
}

func Test_factorial(t *testing.T) {
	t.Log(factorial(7))
	t.Log(factorial(8))
	t.Log(factorial(7))
}

func Test_permutation(t *testing.T) {
	permutation([]int{1, 4, 3}, 0)
	fmt.Println(allPermutations)
	fmt.Println(len(allPermutations))
}

func Test_totalPermutation(t *testing.T) {
	t.Log(totalPermutation(7, 1))
	t.Log(totalPermutation(7, 2))
	t.Log(totalPermutation(7, 3))
	t.Log(totalPermutation(7, 4))
	t.Log(totalPermutation(7, 5))
	t.Log(totalPermutation(7, 6))
	t.Log(totalPermutation(7, 7))
}

func Test_Gta(t *testing.T) {
	t.Log(Gta(7, []int{123489, 5, 67}))
}

func Test_Gta2(t *testing.T) {
	t.Log(Gta(8, []int{12348, 47, 3639}))
}
