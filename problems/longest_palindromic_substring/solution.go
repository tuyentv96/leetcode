func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	i := 0
	maxlen := 0
	begin := 0
	for i = 0; i < len(s); i++ {
		b, length := findPalindrome(s, i, i)
		if length > maxlen {
			begin = b
			maxlen = length
		}

		if i+1 < len(s) && s[i] == s[i+1] {
			b, length := findPalindrome(s, i, i+1)
			if length > maxlen {
				begin = b
				maxlen = length
			}
		}
	}

	return s[begin : begin+maxlen]
}
func findPalindrome(subRune string, i, j int) (int, int) {
	for i >= 0 && j < len(subRune) && subRune[i] == subRune[j] {
		i--
		j++
	}

	b := i + 1
	return b, j - b
}