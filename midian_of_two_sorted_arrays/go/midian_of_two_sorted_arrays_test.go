package midian

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{3}
	nums2 := []int{-2, -1}

	t.Log(findMedianSortedArrays(nums1, nums2))
}
