package kata

import (
	"fmt"
	"sort"
)

var (
	prod            map[int]struct{}
	alreadyComputed map[string]struct{}
)

func sm(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func pr(nums []int) int {
	p := 1
	for _, v := range nums {
		p *= v
	}
	return p
}

func cp(nums []int) []int {
	c := make([]int, len(nums), len(nums))
	for i, n := range nums {
		c[i] = n
	}
	return c
}

func parts(nums []int) {
	ncm := cp(nums)
	sort.Ints(ncm)
	v := fmt.Sprintf("%v", ncm)
	if _, ok := alreadyComputed[v]; ok {
		return
	}

	alreadyComputed[v] = struct{}{}
	prod[pr(nums)] = struct{}{}
	nm := cp(nums)

	for i := 0; i < len(nm); i++ {
		if nm[i]-1 < 1 {
			continue
		}

		nm[i]--

		for j := len(nm) - 1; j > i; j-- {
			if nm[j]+1 <= nm[i] {
				nm[j]++
				parts(nm)
				nm[j]--
			}
		}

		nm = append(nm, 1)
		parts(nm)

		nm[i]++
		nm = nm[:len(nm)-1]
	}
}

func Part(n int) string {
	// your code
	fmt.Printf("input.txt number: %d\n", n)

	prod = make(map[int]struct{}, 0)
	alreadyComputed = make(map[string]struct{})
	parts([]int{n})

	s := make([]int, 0)
	for k := range prod {
		s = append(s, k)
	}

	sort.Ints(s)
	fmt.Println(s)

	rng := s[len(s)-1] - s[0]
	avg := float64(sm(s)) / float64(len(s))
	var mid float64

	if len(s)%2 == 0 {
		mid = float64(s[len(s)/2-1]+s[len(s)/2]) / 2
	} else {
		mid = float64(s[len(s)/2])
	}

	prod = nil
	alreadyComputed = nil

	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rng, avg, mid)
}
