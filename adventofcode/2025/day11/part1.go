package day11

func uniquePaths(input string) int {
	g := parseGraphPart1(input)
	return g.start.uniquePathsPart1()
}
