func findClosestElements(arr []int, k int, x int) []int {
	i := search(arr, x)
	l, r := i-1, i
	for k > 0 {
		if l < 0 || (r < len(arr) && abs(arr[l]-x) > abs(arr[r]-x)) {
			r++
		} else {
			l--
		}

		k--
	}

	return arr[l+1 : r]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func search(arr []int, k int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] == k {
			return m
		}

		if arr[m] > k {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}
