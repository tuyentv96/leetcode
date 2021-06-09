func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var str strings.Builder
	var carry int
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
		}

		str.WriteRune(rune((sum % 10) + '0'))
		carry = sum / 10

		i--
		j--
	}

	if carry > 0 {
		str.WriteRune(rune(carry + '0'))
	}

	return reverse(str.String())
}

func reverse(str string) string {
	strBuilder := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		strBuilder.WriteRune(rune(str[i]))
	}

	return strBuilder.String()
}
