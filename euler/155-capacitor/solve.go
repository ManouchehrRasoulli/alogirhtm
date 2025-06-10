package euler

type combinationType int

const (
	combinationTypeParallel combinationType = iota
	combinationTypeSeries
)

type Combination struct {
	Series    int
	Parallels int
	Type      combinationType
}

func (c Combination) Key() int {
	var (
		parallelsCapacity int
		seriesCapacity    int
	)

	parallelsCapacity = c.Parallels * 60

	seriesDivisor := float64(1) / 60
	seriesBase := seriesDivisor * float64(c.Series)
	if seriesBase != 0 {
		seriesCapacity = int(1 / seriesBase)
	}

	switch c.Type {
	case combinationTypeParallel:
		return seriesCapacity + parallelsCapacity
	case combinationTypeSeries:
		var (
			parallelsBase float64
			totalBase     float64
		)
		if parallelsCapacity != 0 {
			parallelsBase = 1 / float64(parallelsCapacity)
		}

		totalBase = seriesBase + parallelsBase
		if totalBase == 0 {
			return 0
		}
		return int(1 / totalBase)
	}

	panic("invalid combination type !!")
}

var (
	combinations = make(map[int]struct{})
)

func countCombinations(capacitors, left int, directions combinationType) {
	if left == 0 {
		return
	}
}

func Solve(capacitors int) int {

	return 0
}
