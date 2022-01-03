func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([]int, cols+1)
	max := 0
	prev := 0

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min3(dp[j-1], prev, dp[j]) + 1

				if dp[j] > max {
					max = dp[j]
				}
			} else {
				dp[j] = 0
			}

			prev = temp
		}
	}

	return max * max
}

func min3(a, b, c int) int {
	result := a
	if result > b {
		result = b
	}

	if result > c {
		result = c
	}

	return result
}
