package kata

import (
	"fmt"
	"testing"
)

func TestFitBiggestCircle(t *testing.T) {
	input := []float64{0, 0, 10, 0, 10, 10, 0, 10} // square
	result := FitBiggestCircle(input)
	fmt.Printf("Center: (%.2f, %.2f), Radius: %.2f\n", result[0], result[1], result[2])
}
