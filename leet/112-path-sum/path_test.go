package pathsum

import "testing"

func Test_hasPathSum(t *testing.T) {
	t.Log(hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 3))

	t.Log(hasPathSum(&TreeNode{
		Val:  -2,
		Left: nil,
		Right: &TreeNode{
			Val: -3,
		},
	}, -5))
}
