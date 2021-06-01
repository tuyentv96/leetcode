func partition(s string) [][]string {
	result := [][]string{}
	dp := make(map[int]map[int]bool)
	current := make([]string, 0, len(s))
	dfs(s, dp, 0, current, &result)
	return result
}

func dfs(s string, dp map[int]map[int]bool, i int, cur []string, result *[][]string) {
	if i == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for j := i; j < len(s); j++ {
		if dp[i][j] || par(s[i:j+1]) {
			if dp[i] == nil {
				dp[i] = map[int]bool{}
			}
			dp[i][j] = true
			dfs(s, dp, j+1, append(cur, s[i:j+1]), result)
		}
	}
}

func par(s string) bool {
	if len(s) <= 1 {
		return true
	}
	a, b := 0, len(s)-1
	for a < b {
		if s[a] != s[b] {
			return false
		}
		a++
		b--
	}
	return true
}
