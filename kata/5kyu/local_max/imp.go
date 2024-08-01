package kata

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

type direction string

const (
	up       direction = "u"
	down     direction = "d"
	right    direction = "r"
	dontCare direction = "-"
)

func PickPeaks(array []int) PosPeaks {
	fmt.Println(array)

	p := PosPeaks{
		Pos:   make([]int, 0),
		Peaks: make([]int, 0),
	}

	if len(array) == 0 {
		return p
	}

	dArray := make([]direction, len(array))
	dArray[0] = dontCare

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			dArray[i] = up
			continue
		}
		if array[i] == array[i-1] {
			dArray[i] = right
			continue
		}
		dArray[i] = down
	}

	fmt.Println(dArray)

	var (
		lastUp      = 0
		lastUpPos   = 0
		lastUpAdded = false
	)
	for i := 0; i < len(array); i++ {
		if dArray[i] == up {
			lastUp = array[i]
			lastUpPos = i
			lastUpAdded = false
			continue
		}

		if dArray[i] == down && !lastUpAdded && dArray[lastUpPos] != dontCare {
			p.Pos = append(p.Pos, lastUpPos)
			p.Peaks = append(p.Peaks, lastUp)
			lastUpAdded = true
		}
	}

	return p
}
