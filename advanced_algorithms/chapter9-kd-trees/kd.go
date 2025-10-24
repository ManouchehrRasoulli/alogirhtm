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

func remove(root, target *node, level int) *node {
	if root == nil {
		return nil
	}

	if len(root.point) != len(target.point) {
		panic("dimension mismatch in remove")
	}

	if equal(root, target) {
		if root.repeat > 1 {
			root.repeat--
			return root
		}

		if root.right != nil {
			k := len(root.point)
			minNode := findMin(root.right, root.level%k, root.level+1, k)
			root.point = append([]float32(nil), minNode.point...)
			root.tag = minNode.tag
			root.repeat = minNode.repeat
			root.right = remove(root.right, minNode, root.level+1)
			return root
		}

		if root.left != nil {
			k := len(root.point)
			minNode := findMin(root.left, root.level%k, root.level+1, k)
			root.point = append([]float32(nil), minNode.point...)
			root.tag = minNode.tag
			root.repeat = minNode.repeat
			root.left = remove(root.left, minNode, root.level+1)
			root.right = root.left
			root.left = nil
			return root
		}

		return nil
	}

	switch compare(root, target) {
	case 1:
		root.left = remove(root.left, target, level+1)
	case -1:
		root.right = remove(root.right, target, level+1)
	default:
		root.left = remove(root.left, target, level+1)
	}
	return root
}

func findMin(n *node, axis, depth, k int) *node {
	if n == nil {
		return nil
	}

	cd := depth % k

	if cd == axis {
		leftMin := findMin(n.left, axis, depth+1, k)
		return minNode(n, leftMin, axis)
	}

	leftMin := findMin(n.left, axis, depth+1, k)
	rightMin := findMin(n.right, axis, depth+1, k)

	m := minNode(n, leftMin, axis)
	return minNode(m, rightMin, axis)
}

func minNode(a, b *node, axis int) *node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.point[axis] <= b.point[axis] {
		return a
	}
	return b
}
