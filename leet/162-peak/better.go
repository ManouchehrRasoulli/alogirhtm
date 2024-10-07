package peak

func findPeakElement_(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (right + left) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
