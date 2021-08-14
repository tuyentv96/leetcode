func accountsMerge(accounts [][]string) [][]string {
	accountEmails := make(map[string][]int)
	visitedAccounts := make([]bool, len(accounts))

	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			accountEmails[accounts[i][j]] = append(accountEmails[accounts[i][j]], i)
		}
	}

	var result [][]string
	added := make(map[string]bool)

	for i := 0; i < len(accounts); i++ {
		if visitedAccounts[i] {
			continue
		}

		var emails []string
		name := accounts[i][0]

		dfs(accounts, accountEmails, visitedAccounts, i, &emails, added)

		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})

		account := append([]string{name}, emails...)
		result = append(result, account)
	}

	return result
}

func dfs(accounts [][]string, accountEmails map[string][]int, visited []bool, index int, emails *[]string, added map[string]bool) {
	if visited[index] {
		return
	}

	visited[index] = true

	for i := 1; i < len(accounts[index]); i++ {
		email := accounts[index][i]

		if !added[email] {
			*emails = append(*emails, email)
			added[email] = true
		}

		for _, v := range accountEmails[email] {
			dfs(accounts, accountEmails, visited, v, emails, added)
		}
	}

	return
}
