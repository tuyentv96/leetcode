func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
	return dfs(root, &prev)
}

func dfs(root *TreeNode, prev **TreeNode) bool {
	if root == nil {
		return true
	}

    if !dfs(root.Left,prev){
        return false
    }
    
    if (*prev!=nil && root.Val<=(*prev).Val){
        return false
    }
    *prev=root
    
    return dfs(root.Right,prev)
}
