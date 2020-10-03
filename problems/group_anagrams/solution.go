func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	count := [26]int{}
	for _, str := range strs {
		count = [26]int{}
		for i := range str {
			count[str[i]-'a'] += 1
		}

		keyBuilder := strings.Builder{}
		for i := 0; i < 26; i++ {
			keyBuilder.WriteString("#")
			keyBuilder.WriteString(strconv.Itoa(count[i]))
		}

		key := keyBuilder.String()
		m[key] = append(m[key], str)
	}

	var res [][]string
	for _, group := range m {
		res = append(res, group)
	}

	return res
}