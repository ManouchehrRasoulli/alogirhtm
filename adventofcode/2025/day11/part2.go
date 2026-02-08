package day11

func uniquePathsContainsItems(input string) int {
	g := parseGraph(input, start)
	return g.start.uniquePathsPart2(false, false)
}
