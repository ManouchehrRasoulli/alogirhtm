package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func hasPath(root *TreeNode, targetSum, currentSum int) bool {
	if root == nil {
		return false
	}

	v := currentSum + root.Val
	if isLeaf(root) && v == targetSum {
		return true
	}

	return hasPath(root.Left, targetSum, v) || hasPath(root.Right, targetSum, v)

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPath(root, targetSum, 0)
}
