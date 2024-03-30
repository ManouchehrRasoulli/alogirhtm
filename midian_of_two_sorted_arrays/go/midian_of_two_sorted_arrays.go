package midian

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1Size := len(nums1)
	nums2Size := len(nums2)

	if nums1Size > nums2Size {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := nums1Size + nums2Size

	if nums1Size == 0 {
		if total%2 == 1 {
			return float64(nums2[total/2])
		} else {
			return float64(nums2[total/2]+nums2[total/2-1]) / 2
		}
	}

	nums := make([]int, total)

	num1Counter := 0
	num2Counter := 0
	index := 0

	for i := 0; i < total; i++ {
		if nums1[num1Counter] < nums2[num2Counter] {
			nums[i] = nums1[num1Counter]
			num1Counter += 1
		} else if nums2[num2Counter] < nums1[num1Counter] {
			nums[i] = nums2[num2Counter]
			num2Counter += 1
		} else {
			nums[i] = nums1[num1Counter]
			num1Counter += 1
		}

		if num1Counter == nums1Size || num2Counter == nums2Size {
			index = i + 1
			break
		}
	}

	if num2Counter == nums2Size {
		for i := index; i < total; i++ {
			nums[i] = nums1[num1Counter]
			num1Counter += 1
		}
	} else {
		for i := index; i < total; i++ {
			nums[i] = nums2[num2Counter]
			num2Counter += 1
		}
	}

	if total%2 == 1 {
		return float64(nums[total/2])
	} else {
		return float64(nums[total/2]+nums[total/2-1]) / 2
	}
}
