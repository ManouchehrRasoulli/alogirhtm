package part1

import "slices"

func diff(a, b int64) int64 {
	if b > a {
		a, b = b, a
	}

	return a - b
}

func minDiff(sources, syncs []int64) int64 {
	slices.Sort(sources)
	slices.Sort(syncs)

	totalDiff := int64(0)

	for i := 0; i < len(sources); i++ {
		totalDiff += diff(sources[i], syncs[i])
	}

	return totalDiff
}

func similarityScore(sources, syncs []int64) int64 {
	syncMap := make(map[int64]int64)
	for _, sync := range syncs {
		syncMap[sync]++
	}

	similarity := int64(0)
	for i := 0; i < len(sources); i++ {
		v := sources[i]
		similarity += v * syncMap[v]
	}

	return similarity
}
