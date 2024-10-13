package sum

import (
	"math"
	"slices"
)

func diff(target, total int) int {
	dif := target - total
	if dif < 0 {
		return -dif
	}
	return dif
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort[[]int](nums)

	dif := math.MaxInt64
	v := 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			df := diff(target, total)
			if df <= dif {
				dif = df
				v = total
			}

			if dif == 0 {
				return v
			}

			if total < target {
				left++
				continue
			}

			right--
		}
	}

	return v
}
