func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node,_:=dfs(root)
	return node
}

func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	l, ld := dfs(root.Left)
	r, rd := dfs(root.Right)
	if ld > rd {
		return l, ld + 1
	}

	if rd > ld {
		return r, rd + 1
	}

	return root, ld + 1
}