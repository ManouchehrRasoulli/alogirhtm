package day1

import (
	"testing"
)

func Test_MinimumDifference(t *testing.T) {
	t.Log("minimum difference ...")

	readInput()

	t.Log(minDiff(sources, syncs))
}

func Test_SimilarityScore(t *testing.T) {
	t.Log("similarity score ...")

	readInput()

	t.Log(similarityScore(sources, syncs))
}
