package day7

import (
	"testing"
)

func TestManifoldDimensionsSample1(t *testing.T) {
	t.Log(ManifoldDimensions(sample1)) // 40
}

func TestManifoldDimensionsPuzzle(t *testing.T) {
	t.Log(ManifoldDimensions(puzzle)) // 95408386769474
}
