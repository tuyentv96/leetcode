func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	end := false
	queue = append(queue, root)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {
			end = true
		} else {
			if end {
				return false
			}

			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}

	return true
}