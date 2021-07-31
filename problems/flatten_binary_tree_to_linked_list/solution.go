func flatten(root *TreeNode) {
	var prev *TreeNode
	dfs(root, &prev)
}

func dfs(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Right, prev)
	dfs(root.Left, prev)
	root.Right = *prev
	root.Left = nil
	*prev = root
}
