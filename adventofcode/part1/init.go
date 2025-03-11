package part1

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	sources []int64
	syncs   []int64
)

func readInput() {
	sources = make([]int64, 0)
	syncs = make([]int64, 0)

	lines := strings.Split(input, "\n")

	for _, l := range lines {
		items := strings.Split(l, "   ")
		source, _ := strconv.ParseInt(items[0], 10, 64)
		sync, _ := strconv.ParseInt(items[1], 10, 64)
		sources = append(sources, source)
		syncs = append(syncs, sync)
	}
}
