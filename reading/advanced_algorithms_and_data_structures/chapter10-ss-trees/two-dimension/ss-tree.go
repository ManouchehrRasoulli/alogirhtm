package main

import (
	"fmt"
	"math"
)

const (
	maxPointsPerNode = 3 // capacity before splitting
)

// -------------------------------
// Node definition
// -------------------------------
type SSTreeNode struct {
	isLeaf   bool
	points   [][2]float64
	children []*SSTreeNode
	center   [2]float64
	radius   float64
}

// -------------------------------
// Utility functions
// -------------------------------

// distance between 2D points
func dist(a, b [2]float64) float64 {
	dx, dy := a[0]-b[0], a[1]-b[1]
	return math.Sqrt(dx*dx + dy*dy)
}

// recompute bounding sphere for a node
func (n *SSTreeNode) recomputeBoundingSphere() {
	if n.isLeaf {
		if len(n.points) == 0 {
			n.center = [2]float64{0, 0}
			n.radius = 0
			return
		}
		// compute mean as center
		var cx, cy float64
		for _, p := range n.points {
			cx += p[0]
			cy += p[1]
		}
		n.center = [2]float64{cx / float64(len(n.points)), cy / float64(len(n.points))}

		// radius = max distance to center
		var r float64
		for _, p := range n.points {
			d := dist(n.center, p)
			if d > r {
				r = d
			}
		}
		n.radius = r
	} else {
		if len(n.children) == 0 {
			n.center = [2]float64{0, 0}
			n.radius = 0
			return
		}
		// mean of child centers
		var cx, cy float64
		for _, c := range n.children {
			cx += c.center[0]
			cy += c.center[1]
		}
		n.center = [2]float64{cx / float64(len(n.children)), cy / float64(len(n.children))}

		// radius = max(distance from center to child center + child radius)
		var r float64
		for _, c := range n.children {
			d := dist(n.center, c.center) + c.radius
			if d > r {
				r = d
			}
		}
		n.radius = r
	}
}

// -------------------------------
// Insertion
// -------------------------------

func (n *SSTreeNode) insert(p [2]float64) *SSTreeNode {
	if n.isLeaf {
		n.points = append(n.points, p)
		n.recomputeBoundingSphere()

		// Split if over capacity
		if len(n.points) > maxPointsPerNode {
			return n.splitLeaf()
		}
		return n
	}

	// Internal node: choose child needing least radius expansion
	var best *SSTreeNode
	var minIncrease float64 = math.MaxFloat64
	for _, c := range n.children {
		d := dist(c.center, p)
		inc := 0.0
		if d > c.radius {
			inc = d - c.radius
		}
		if inc < minIncrease {
			minIncrease = inc
			best = c
		}
	}

	// Insert into best child
	newChild := best.insert(p)
	if newChild != best {
		// best was split → replace with two
		n.replaceChild(best, newChild)
	}

	n.recomputeBoundingSphere()

	// Split if over capacity
	if len(n.children) > maxPointsPerNode {
		return n.splitInternal()
	}
	return n
}

// replace a child with its two split children
func (n *SSTreeNode) replaceChild(old *SSTreeNode, newNode *SSTreeNode) {
	// find old child
	idx := -1
	for i, c := range n.children {
		if c == old {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	// replace old with new's children
	left, right := newNode.children[0], newNode.children[1]
	// remove old
	n.children = append(n.children[:idx], n.children[idx+1:]...)
	// append both
	n.children = append(n.children, left, right)
}

// -------------------------------
// Splitting
// -------------------------------

// split leaf into two new leaves
func (n *SSTreeNode) splitLeaf() *SSTreeNode {
	// pick two farthest points as seeds
	maxD := -1.0
	var a, b int
	for i := 0; i < len(n.points); i++ {
		for j := i + 1; j < len(n.points); j++ {
			d := dist(n.points[i], n.points[j])
			if d > maxD {
				maxD = d
				a, b = i, j
			}
		}
	}

	left := &SSTreeNode{isLeaf: true}
	right := &SSTreeNode{isLeaf: true}
	left.points = append(left.points, n.points[a])
	right.points = append(right.points, n.points[b])

	// assign remaining points to nearest seed
	for i, p := range n.points {
		if i == a || i == b {
			continue
		}
		if dist(p, n.points[a]) < dist(p, n.points[b]) {
			left.points = append(left.points, p)
		} else {
			right.points = append(right.points, p)
		}
	}

	left.recomputeBoundingSphere()
	right.recomputeBoundingSphere()

	parent := &SSTreeNode{
		isLeaf:   false,
		children: []*SSTreeNode{left, right},
	}
	parent.recomputeBoundingSphere()
	return parent
}

// split internal node (same logic but on children)
func (n *SSTreeNode) splitInternal() *SSTreeNode {
	maxD := -1.0
	var a, b int
	for i := 0; i < len(n.children); i++ {
		for j := i + 1; j < len(n.children); j++ {
			d := dist(n.children[i].center, n.children[j].center)
			if d > maxD {
				maxD = d
				a, b = i, j
			}
		}
	}

	left := &SSTreeNode{isLeaf: false}
	right := &SSTreeNode{isLeaf: false}
	left.children = append(left.children, n.children[a])
	right.children = append(right.children, n.children[b])

	for i, c := range n.children {
		if i == a || i == b {
			continue
		}
		if dist(c.center, n.children[a].center) < dist(c.center, n.children[b].center) {
			left.children = append(left.children, c)
		} else {
			right.children = append(right.children, c)
		}
	}

	left.recomputeBoundingSphere()
	right.recomputeBoundingSphere()

	parent := &SSTreeNode{
		isLeaf:   false,
		children: []*SSTreeNode{left, right},
	}
	parent.recomputeBoundingSphere()
	return parent
}

// -------------------------------
// Search & Delete
// -------------------------------

func (n *SSTreeNode) search(p [2]float64, epsilon float64) bool {
	if n.isLeaf {
		for _, q := range n.points {
			if dist(q, p) < epsilon {
				return true
			}
		}
		return false
	}

	for _, c := range n.children {
		if dist(c.center, p) <= c.radius+epsilon {
			if c.search(p, epsilon) {
				return true
			}
		}
	}
	return false
}

func (n *SSTreeNode) delete(p [2]float64, epsilon float64) bool {
	if n.isLeaf {
		for i, q := range n.points {
			if dist(q, p) < epsilon {
				n.points = append(n.points[:i], n.points[i+1:]...)
				n.recomputeBoundingSphere()
				return true
			}
		}
		return false
	}

	for _, c := range n.children {
		if dist(c.center, p) <= c.radius+epsilon {
			if c.delete(p, epsilon) {
				n.recomputeBoundingSphere()
				return true
			}
		}
	}
	return false
}

// -------------------------------
// Demo
// -------------------------------

/*
Search (7,8): true
Search (5,5): false
After delete (7,8): false
*/
func main() {
	root := &SSTreeNode{isLeaf: true}

	points := [][2]float64{
		{1, 2}, {3, 4}, {2, 5}, {7, 8},
		{8, 9}, {9, 10}, {10, 11},
	}

	for _, p := range points {
		root = root.insert(p)
	}

	fmt.Println("Search (7,8):", root.search([2]float64{7, 8}, 1e-6))
	fmt.Println("Search (5,5):", root.search([2]float64{5, 5}, 1e-6))

	root.delete([2]float64{7, 8}, 1e-6)
	fmt.Println("After delete (7,8):", root.search([2]float64{7, 8}, 1e-6))
}
