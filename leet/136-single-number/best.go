package singlenumber

func singleNumberBest(nums []int) int {
	r := map[int]struct{}{}

	for i := 0; i < len(nums); i++ {
		if _, ok := r[nums[i]]; !ok {
			r[nums[i]] = struct{}{}
			continue
		}
		delete(r, nums[i])
	}

	for k := range r {
		return k
	}
	return 0
}
