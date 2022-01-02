func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	prev := make([]int, m+1)
	for col := n - 1; col >= 0; col-- {
		dp := make([]int, m+1)
		for row := m - 1; row >= 0; row-- {
			if text2[col] == text1[row] {
				dp[row] = 1 + prev[row+1]
			} else {
				dp[row] = max(dp[row+1], prev[row])
			}
		}

		prev = dp
	}

	return prev[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
