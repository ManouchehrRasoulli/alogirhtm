package kd

type node struct {
	tag    string
	point  []float32
	level  int
	repeat int
	left   *node
	right  *node
}

func newNode(tag string, point []float32) *node {
	return &node{
		repeat: 1,
		tag:    tag,
		point:  point,
		left:   nil,
		right:  nil,
	}
}

func compare(a, b *node) int {
	if len(a.point) != len(b.point) {
		panic("could not compare the nodes")
	}

	var axis int
	if a.level != 0 {
		axis = a.level % len(a.point)
	}

	if a.point[axis] < b.point[axis] {
		return -1
	} else if a.point[axis] > b.point[axis] {
		return 1
	}

	return 0
}

func equal(a, b *node) bool {
	if len(a.point) != len(b.point) {
		panic("could not compare the nodes")
	}

	for i := 0; i < len(a.point); i++ {
		if a.point[i] != b.point[i] {
			return false
		}
	}

	return true
}

func distanceSquared(a, b *node) float32 {
	sum := float32(0)

	for i := range a.point {
		d := a.point[i] - b.point[i]
		sum += d * d
	}

	return sum
}

func compareDistance(a, b, ref *node) int {
	da := distanceSquared(a, ref)
	db := distanceSquared(b, ref)

	if da < db {
		return -1
	} else if da > db {
		return 1
	}

	return 0
}

func search(root, target *node) *node {
	if root == nil {
		return nil
	}

	if equal(root, target) {
		return root
	}

	switch compare(root, target) {
	case 1:
		return search(root.left, target)
	case -1:
		return search(root.right, target)
	default:
		return search(root.left, target)
	}
}

func insert(root, newNode *node, level int) *node {
	newNode.level = level
	if root == nil {
		return newNode
	}

	if equal(root, newNode) {
		root.repeat++
		return root
	}

	switch compare(root, newNode) {
	case 1:
		root.left = insert(root.left, newNode, level+1)
	case -1:
		root.right = insert(root.right, newNode, level+1)
	default:
		root.left = insert(root.left, newNode, level+1)
	}

	return root
}
