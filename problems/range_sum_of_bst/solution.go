func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	dfs(root, low, high, &sum)
	return sum
}

func dfs(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= low && root.Val <= high {
		*sum += root.Val
	}

	if root.Val > low {
		dfs(root.Left, low, high, sum)
	}

	if root.Val < high {
		dfs(root.Right, low, high, sum)
	}
}