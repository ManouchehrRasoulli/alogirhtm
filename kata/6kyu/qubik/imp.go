package kata

import (
	"math"
)

func SolomonsQuest(moves [][3]int) [2]int {
	x, y := 0, 0
	layer := 0

	for _, move := range moves {
		layerChange := move[0]
		direction := move[1]
		relDist := move[2]

		layer += layerChange
		if layer < 0 {
			layer = 0
		}

		// Calculate standard distance
		standardDistance := float64(relDist) * math.Pow(2, float64(layer))

		switch direction {
		case 0: // North
			y += int(standardDistance)
		case 1: // East
			x += int(standardDistance)
		case 2: // South
			y -= int(standardDistance)
		case 3: // West
			x -= int(standardDistance)
		}
	}

	return [2]int{x, y}
}
