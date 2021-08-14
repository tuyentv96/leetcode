/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	dfs(root, &result, 0)
	return result
}

func dfs(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*result) <= level {
		*result = append(*result, []int{})
	}

	if level%2 != 0 {
		// insert at begin
		(*result)[level] = append([]int{root.Val}, (*result)[level]...)
	} else {
		// append to end
		(*result)[level] = append((*result)[level], root.Val)
	}

	dfs((*TreeNode)(root.Left), result, level+1)
	dfs((*TreeNode)(root.Right), result, level+1)
}
