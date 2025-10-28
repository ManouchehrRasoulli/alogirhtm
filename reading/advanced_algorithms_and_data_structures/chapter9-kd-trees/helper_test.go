package kd

import "testing"

func Test_PrintKDTree(t *testing.T) {
	root := &node{tag: "a", point: []float32{5, 4}, level: 0}
	root.left = &node{tag: "b", point: []float32{2, 3}, level: 1}
	root.right = &node{tag: "c", point: []float32{9, 6}, level: 1}
	root.left.left = &node{tag: "d", point: []float32{1, 1}, level: 2}
	root.left.right = &node{tag: "e", point: []float32{3, 7}, level: 2}
	root.right.left = &node{tag: "f", point: []float32{8, 1}, level: 2}

	PrintKDTree(root)
}
