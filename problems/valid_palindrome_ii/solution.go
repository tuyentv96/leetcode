func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			if valid(s, start+1, end) {
				return true
			}

			if valid(s, start, end-1) {
				return true
			}

			return false
		}

		start++
		end--
	}

	return true
}

func valid(str string, start, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}

		start++
		end--
	}

	return true
}
