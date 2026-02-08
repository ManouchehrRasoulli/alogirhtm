package day11

func uniquePaths(input string) int {
	g := parseGraph(input)
	return g.start.uniquePaths()
}
