package kata

import "fmt"

const (
	flat = 0
	up   = 1
	down = -1
)

func RainVolume(towers []int) int {
	fmt.Println(towers)

	if len(towers) < 2 {
		return 0
	}

	var (
		direction         = flat
		peaks             = make([]int, 0)
		maxValueIndexFunc = func(i, j int) int {
			if towers[i] > towers[j] {
				return i
			}
			return j
		}
		difToDirection = func(i, j int) int {
			if towers[i] > towers[j] {
				return down
			}
			if towers[i] < towers[j] {
				return up
			}
			return flat
		}
	)

	direction = difToDirection(0, 1)
	if direction == down {
		peaks = append(peaks, 0)
	}

	for i := 2; i < len(towers); i++ {
		currentDirection := difToDirection(i-1, i)
		if currentDirection == direction {
			continue
		}

		if direction == up && currentDirection == down {
			peaks = append(peaks, maxValueIndexFunc(i-1, i))
			direction = currentDirection
		}

		if direction == down && currentDirection == up {
			direction = currentDirection
		}
	}

	if direction == up {
		peaks = append(peaks, len(towers)-1)
	}

	fmt.Println(peaks)

	return 0
}
