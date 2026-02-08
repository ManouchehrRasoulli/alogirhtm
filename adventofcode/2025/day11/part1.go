package day11

func uniquePaths(input string) int {
	g := parseGraph(input, you)
	return g.start.uniquePathsPart1()
}
