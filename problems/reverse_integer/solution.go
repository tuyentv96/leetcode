func reverse(x int) int {
	const MaxInt = 2147483647
	const MinInt = -2147483648

	var res int
	for {
		var pop = x % 10
		x = x / 10

		if res > MaxInt/10 || (res == MaxInt/10 && pop > 7) {
			return 0
		}

		if res < MinInt/10 || (res == MinInt/10 && pop < -8) {
			return 0
		}

		res = (res * 10) + pop
		if x == 0 {
			break
		}
	}

	return res
}