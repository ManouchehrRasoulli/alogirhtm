package kata

import "math"

func RectangleRotation(a, b int) int {
	var x = int(float64(a) / math.Sqrt(2))
	var y = int(float64(b) / math.Sqrt(2))
	var r = (x+1)*(y+1) + x*y
	return r + r%2 - 1
}
