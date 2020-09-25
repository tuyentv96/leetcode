func combinationSum3(k int, n int) [][]int {
	c, res := []int{}, [][]int{}
	findcombinationSum(0, 1, k, n, c, &res)
	return res
}

func findcombinationSum(level int, start int, k int, n int, c []int, res *[][]int) {
	if level == k || n < 0 {
		if n == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}

		return
	}

	for i := start; i <= 9; i++ {
		if i > n {
			break
		}

		c = append(c, i)
		findcombinationSum(level+1, i+1, k, n-i, c, res)
		c = c[:len(c)-1]
	}
}