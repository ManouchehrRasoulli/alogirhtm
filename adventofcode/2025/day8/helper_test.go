package day8

import (
	"testing"
)

func TestDistance(t *testing.T) {
	a := Location{X: 162, Y: 817, Z: 812}
	b := Location{X: 431, Y: 825, Z: 988}
	d := distance(a, b)
	if d != 321.5603 {
		t.Fatal("invalid distance", d)
	}
}

func TestDistanceCollector(t *testing.T) {
	locations := ParseLocations(sample1)
	dc := newDistanceCollection()
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			dc.insert(locations[i], locations[j])
		}
	}

	t.Log(dc.next())
	t.Log(dc.next())
	t.Log(dc.next())
}
