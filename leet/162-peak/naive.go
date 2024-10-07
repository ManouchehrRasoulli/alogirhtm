package peak

func isPeak(nums []int, i int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	if i == 0 {
		return nums[i] > nums[i+1]
	}

	if i == len(nums)-1 {
		return nums[i] > nums[i-1]
	}

	return nums[i] > nums[i-1] && nums[i] > nums[i+1]
}

func findPeakElement(nums []int) int {
	maxPeak := nums[0]
	pos := 0

	for i := 1; i < len(nums); i++ {
		if isPeak(nums, i) {
			if nums[i] > maxPeak {
				maxPeak = nums[i]
				pos = i
			}
		}
	}

	return pos
}
