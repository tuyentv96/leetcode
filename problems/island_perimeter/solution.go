func islandPerimeter(grid [][]int) int {
	size := len(grid)
	perimiter := 0
	for i := range grid {
		for j := range grid[i] {
			sizeJ := len(grid[i])
			if grid[i][j] != 1 {
				continue
			}

			if i == 0 || grid[i-1][j] != 1 {
				perimiter++
			}

			if i == size-1 || grid[i+1][j] != 1 {
				perimiter++
			}

			if j == 0 || grid[i][j-1] != 1 {
				perimiter++
			}

			if j == sizeJ-1 || grid[i][j+1] != 1 {
				perimiter++
			}
		}
	}

	return perimiter
}
