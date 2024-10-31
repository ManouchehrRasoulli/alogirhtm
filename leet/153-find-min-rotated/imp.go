package findminrotated

import "fmt"

func findMin(nums []int) int {
	var x func(nums []int, start, end int) int

	x = func(nums []int, start int, end int) int {
		if start+1 == end || start == end {
			return min(nums[start], nums[end])
		}

		mid := (end + start) / 2
		fmt.Println("----", start, end, nums)
		fmt.Println(mid)

		if nums[mid] < nums[end] {
			return x(nums, start, mid)
		}

		if nums[mid] > nums[start] {
			return x(nums, mid, end)
		}

		return -1
	}

	return x(nums, 0, len(nums)-1)
}
