package binratytreepreorderinorder

import "testing"

func Test_Sliding(t *testing.T) {
	i := []int{0, 1, 2, 3, 4, 5, 6}
	t.Log(i[:1])
	t.Log(i[1:])

	t.Log(i[:6])
	t.Log(i[6:])
}

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Log(StringTree(tree))

	tree = buildTree([]int{-1}, []int{-1})
	t.Log(StringTree(tree))

	tree = buildTree([]int{1, 2}, []int{1, 2})
	t.Log(StringTree(tree))

	tree = buildTree([]int{1, 2, 3}, []int{3, 2, 1})
	t.Log(StringTree(tree))
}
