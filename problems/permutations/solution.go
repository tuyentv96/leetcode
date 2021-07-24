func permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	var cur []int
	var result [][]int
	dfs(nums, visited, cur, &result)
	return result
}

func dfs(nums []int, visited []bool, cur []int, result *[][]int) {
	if len(cur) == len(nums) {
		c := make([]int, len(cur))
		copy(c, cur)
		*result = append(*result, c)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		cur = append(cur, nums[i])
		dfs(nums, visited, cur, result)
		cur = cur[:len(cur)-1]
		visited[i] = false
	}
}
