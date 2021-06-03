func addBinary(a string, b string) string {
	sb := strings.Builder{}
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}

		if j >= 0 {
			sum += int(b[j] - '0')
		}

		sb.WriteRune(rune(sum%2 + '0'))
		carry = sum / 2

		i--
		j--
	}

	if carry != 0 {
		sb.WriteRune('1')
	}

	return reverse(sb.String())
}

func reverse(str string) string {
	sb := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteRune(rune(str[i]))
	}

	return sb.String()
}
