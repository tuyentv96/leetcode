func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := uint8(0)
	builder := strings.Builder{}
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += num1[i] - '0'
		}

		if j >= 0 {
			sum += num2[j] - '0'
		}

		carry = sum / 10
		builder.WriteRune(rune((sum % 10) + '0'))

		i--
		j--
	}

	if carry > 0 {
		builder.WriteRune(rune(carry + '0'))
	}

	return reverse(builder.String())
}

func reverse(str string) string {
	sb := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteRune(rune(str[i]))
	}

	return sb.String()
}
