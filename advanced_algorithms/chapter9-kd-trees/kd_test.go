package kd

import (
	"math/rand"
	"testing"
)

func Test_Compare(t *testing.T) {
	a := &node{point: []float32{2, 3}, level: 0}
	b := &node{point: []float32{5, 4}, level: 0}

	if compare(a, b) != -1 {
		t.Errorf("expected a < b along axis 0")
	}

	a.level = 1
	if compare(a, b) != -1 {
		t.Errorf("expected a < b along axis 1")
	}
}

func Test_Equal(t *testing.T) {
	a := &node{point: []float32{2, 3}}
	b := &node{point: []float32{2, 3}}
	c := &node{point: []float32{3, 2}}

	if !equal(a, b) {
		t.Errorf("expected a == b")
	}
	if equal(a, c) {
		t.Errorf("expected a != c")
	}
}

func Test_DistanceSquared(t *testing.T) {
	a := &node{point: []float32{0, 0}}
	b := &node{point: []float32{3, 4}}
	d := distanceSquared(a, b)
	if d != 25 {
		t.Errorf("expected 25, got %v", d)
	}
}

func Test_CompareDistance(t *testing.T) {
	ref := &node{point: []float32{0, 0}}
	a := &node{point: []float32{1, 1}}
	b := &node{point: []float32{2, 2}}

	if compareDistance(a, b, ref) != -1 {
		t.Errorf("expected a closer than b")
	}
}

func Test_InsertAndSearch(t *testing.T) {
	points := [][]float32{
		{3, 6}, {6, 12}, {9, 1}, {3, 6}, {17, 15}, {13, 15}, {2, 7}, {10, 19},
	}

	randomTag := func() string {
		return string(rune('a' + rand.Intn(26)))
	}

	var root *node
	for _, p := range points {
		root = insert(root, newNode(randomTag(), p), 0)
	}

	PrintKDTree(root)

	target := &node{point: []float32{10, 19}}
	found := search(root, target)
	if found == nil || !equal(found, target) {
		t.Fatalf("expected to find node %v", target.point)
	}

	missing := &node{point: []float32{99, 99}}
	if search(root, missing) != nil {
		t.Fatalf("expected not to find missing node")
	}
}

func TestBuildKDTreeDiagram(t *testing.T) {
	/*
		          A (0,5)
		         /      \
		   C(-1,6)      B(1,-1)
		     /              \
		  D(-1,1)           E(2,-5)
		   /   \
		G(-1.5,-2)  F(-0.5,0)
	*/

	var root *node

	// Insert in the logical order (A first as root)
	root = insert(root, newNode("A", []float32{0, 5}), 0)
	root = insert(root, newNode("B", []float32{1, -1}), 0)
	root = insert(root, newNode("C", []float32{-1, 6}), 0)
	root = insert(root, newNode("D", []float32{-1, 1}), 0)
	root = insert(root, newNode("E", []float32{2, -5}), 0)
	root = insert(root, newNode("F", []float32{-0.5, 0}), 0)
	root = insert(root, newNode("G", []float32{-1.5, -2}), 0)

	t.Log("✅ KD-tree successfully built")
	PrintKDTree(root)

	root = insert(root, newNode("H", []float32{2.5, -3}), 0)
	t.Log("✅ KD-tree add new Node H")

	PrintKDTree(root)
}
