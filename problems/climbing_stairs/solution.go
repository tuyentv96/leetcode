func climbStairs(n int) int {
	dp := make([]int, n)
	if n <= 1 {
		return 1
	}
	
	if n == 2 {
		return 2
	}
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}