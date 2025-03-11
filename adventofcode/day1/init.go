package day1

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	sources []int64
	syncs   []int64
)

func init() {
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

	slices.Sort(sources)
	slices.Sort(syncs)
}
