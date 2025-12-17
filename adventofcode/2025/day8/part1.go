package day8

func ConnectedCircuitsCapacityCount(input string, connect, topProduct int) int {
	locations := ParseLocations(input)
	dc := newDistanceCollection()
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			dc.insert(locations[i], locations[j])
		}
	}

	cl := newCluster()
	for i := 0; i < connect; i++ {
		nx := dc.next()
		cl.insert(nx.a, nx.b)
	}

	return cl.topProduct(topProduct)
}
