package sum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort[[]int](nums)

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

				continue
			}

			if total < 0 {
				left++
				continue
			}

			right--
		}
	}

	return res
}
