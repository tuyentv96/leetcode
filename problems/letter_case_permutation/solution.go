func letterCasePermutation(s string) []string {
	res := make([]string, 0)
	dfs(s, "", 0, &res)
	return res
}

func dfs(s string, cur string, index int, res *[]string) {
	if len(cur) == len(s) {
		*res = append(*res, cur)
		return
	}

	dfs(s, cur+string(s[index]), index+1, res)
	if s[index] >= 'A' && s[index] <= 'Z' {
		dfs(s, cur+string(s[index]+32), index+1, res)
	}

	if s[index] >= 'a' && s[index] <= 'z' {
		dfs(s, cur+string(s[index]-32), index+1, res)
	}
}
