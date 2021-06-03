func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		for left < len(s) && !isValidChar(s[left]) {
			left++
		}

		for right > -1 && !isValidChar(s[right]) {
			right--
		}

		if left > len(s)-1 || right < 0 || left >= right {
			break
		}

		if lowerCase(s[left]) != lowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isValidChar(c uint8) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func lowerCase(c uint8) uint8 {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}
