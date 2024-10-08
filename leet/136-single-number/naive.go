package singlenumber

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func singleNumber(nums []int) int {
	x := abs(nums[0])
	for i := 1; i < len(nums); i++ {
		v := abs(nums[i])
		if v > x {
			x = v
		}
	}

	vp := make([]byte, x+1)
	vn := make([]byte, x+1)
	for _, i := range nums {
		if i < 0 {
			vn[abs(i)] += 1
			continue
		}
		vp[i] += 1
	}

	for _, i := range nums {
		if i < 0 {
			if vn[abs(i)] == 1 {
				return i
			}
			continue
		}
		if vp[i] == 1 {
			return i
		}
	}

	return -1
}
