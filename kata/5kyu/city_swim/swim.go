package kata

func RainVolume(towers []int) int {
	if len(towers) < 2 {
		return 0
	}

	var (
		left     = 0
		right    = len(towers) - 1
		maxLeft  = 0
		maxRight = 0
		water    = 0
	)

	for left < right {
		if towers[left] < towers[right] {
			if towers[left] >= maxLeft {
				maxLeft = towers[left]
			} else {
				water += maxLeft - towers[left]
			}
			left++
		} else {
			if towers[right] >= maxRight {
				maxRight = towers[right]
			} else {
				water += maxRight - towers[right]
			}
			right--
		}
	}

	return water
}
