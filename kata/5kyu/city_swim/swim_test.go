package kata

import "testing"

func TestRainVolume(t *testing.T) {
	t.Log(RainVolume([]int{5, 2, 6}))                // 3
	t.Log(RainVolume([]int{1, 0, 5, 2, 6, 3, 10}))   // 7
	t.Log(RainVolume([]int{15, 0, 6, 10, 11, 2, 5})) // 20
	t.Log(RainVolume([]int{1, 5, 1}))                // 0
}
