func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}