func lengthOfLongestSubstring(s string) int {
	rs := 0
	start, end := 0, 0
	var visited [256]bool
	for end < len(s) {
		for visited[s[end]] && start <= end {
			visited[s[start]] = false
			start++
		}
		visited[s[end]] = true
		rs = max(rs, end-start+1)
		end++
	}

	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
