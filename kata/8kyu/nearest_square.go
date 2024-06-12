package _kyu

import (
	"math"
)

func NearestSq(n int) int {
	// Code goes here
	square := math.Pow(float64(n), 0.5)
	squareInt := int(square)

	if int(square*10) == squareInt*10 { // perfect square
		return squareInt
	}
	oneBelow := math.Pow(float64(squareInt), 2)
	oneAbove := math.Pow(float64(squareInt)+1, 2)

	difBelow := n - int(oneBelow)
	difAbove := int(oneAbove) - n

	if difBelow < difAbove {
		return int(oneBelow)
	}

	return int(oneAbove)
}
