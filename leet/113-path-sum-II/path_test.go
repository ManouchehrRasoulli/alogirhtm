package pathsumii

import "testing"

func Test_pathSum(t *testing.T) {
	t.Log(pathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 3))

	t.Log(pathSum(&TreeNode{
		Val:  -2,
		Left: nil,
		Right: &TreeNode{
			Val: -3,
		},
	}, -5))
}
