func isAlphaNumeric(s uint8) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	}

	return false
}

func toLower(s uint8) uint8 {
	if s >= 'A' && s <= 'Z' {
		return s - 'A' + 'a'
	}

	return s
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	var left, right = 0, len(s) - 1
	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}