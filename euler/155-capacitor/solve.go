package euler

type Combination struct {
	Series    int
	Parallels int
}

func (c Combination) Key() int {
	if c.Series == 0 {
		return c.Parallels * 60
	}

	if c.Parallels == 0 {
		seriesDivisor := float64(1) / 60
		seriesBase := seriesDivisor * float64(c.Series)
		if seriesBase == 0 {
			return 0
		}

		return int(1 / seriesBase)
	}

	parallelsBase := 1 / float64(c.Parallels*60)
	seriesDivisor := float64(1) / 60
	seriesBase := seriesDivisor * float64(c.Series)

	return int(1 / (parallelsBase + seriesBase))
}

const (
	parallel = iota
	series
)

var (
	combinations = make(map[int]struct{})
)

func countCombinations(capacitors, left int, directions int) {
	if left == 0 {
		return
	}

	switch directions {
	case parallel:
	case series:
	}
}

func Solve(capacitors int) int {

	return 0
}
