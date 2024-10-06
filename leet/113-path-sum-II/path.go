package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func hasPath(root *TreeNode, targetSum, currentSum int) (bool, [][]int) {
	if root == nil {
		return false, nil
	}

	v := currentSum + root.Val
	if isLeaf(root) && v == targetSum {
		return true, [][]int{{root.Val}}
	}

	paths := make([][]int, 0)
	leftOk, leftPaths := hasPath(root.Left, targetSum, v)
	if leftOk {
		for _, p := range leftPaths {
			pp := make([]int, 0)
			pp = append(pp, root.Val)
			pp = append(pp, p...)
			paths = append(paths, pp)
		}
	}
	rightOk, rithPaths := hasPath(root.Right, targetSum, v)
	if rightOk {
		for _, p := range rithPaths {
			pp := make([]int, 0)
			pp = append(pp, root.Val)
			pp = append(pp, p...)
			paths = append(paths, pp)
		}
	}

	if len(paths) == 0 {
		return false, nil
	}

	return true, paths
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ok, paths := hasPath(root, targetSum, 0)
	if ok {
		return paths
	}

	return [][]int{}
}
