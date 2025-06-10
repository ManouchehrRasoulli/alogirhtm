package euler

import (
	"testing"
)

func TestCombination(t *testing.T) {
	c := Combination{
		Series:    1,
		Parallels: 2,
		Type:      combinationTypeSeries,
	}
	t.Log(c.Key()) // 40

	c = Combination{
		Series:    2,
		Parallels: 1,
		Type:      combinationTypeParallel,
	}
	t.Log(c.Key()) // 90

	c = Combination{
		Series:    0,
		Parallels: 3,
		Type:      combinationTypeSeries,
	}
	t.Log(c.Key()) // 180

	c = Combination{
		Series:    2,
		Parallels: 0,
		Type:      combinationTypeSeries,
	}
	t.Log(c.Key()) // 30

	c = Combination{
		Series:    3,
		Parallels: 0,
		Type:      combinationTypeSeries,
	}
	t.Log(c.Key()) // 20
}
