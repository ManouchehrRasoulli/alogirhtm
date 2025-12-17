package day8

import "log"

func ConnectedCircuitsAndReturnLastOne(input string) int {
	locations := ParseLocations(input)
	dc := newDistanceCollection()
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			dc.insert(locations[i], locations[j])
		}
	}

	cl := newCluster()
	for _, loc := range locations {
		cl.insert(loc)
	}

	log.Println("start to merge :: ", cl.count())

	var last *distanceNode
	for cl.count() > 1 {
		nx := dc.next()
		cl.merge(nx.a, nx.b)
		last = nx
	}

	log.Println("merge done :: ", cl.count())

	if last == nil {
		return 0
	}

	return int(last.a.X * last.b.X)
}
