package solution

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i, _ := range grid {
		grid[i] = make([]int, n)
		for j := range n {
			grid[i][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]

}
