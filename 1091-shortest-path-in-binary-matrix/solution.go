package solution

func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	if grid[0][0] != 0 || grid[rows-1][columns-1] != 0 {
		return -1
	}
	queue := make([][3]int, 0)
	visited := make(map[[2]int]bool)

	queue = append(queue, [3]int{0, 0, 1})
	visited[[2]int{0, 0}] = true

	for len(queue) > 0 {
		currentLocation := queue[0]
		queue = queue[1:]
		n := neighbors(rows, columns, currentLocation[0], currentLocation[1])
		if currentLocation[0] == rows-1 && currentLocation[1] == columns-1 {
			return currentLocation[2]
		}
		for _, neighbor := range n {
			_, visit := visited[neighbor]
			if grid[neighbor[0]][neighbor[1]] != 1 && !visit {
				queue = append(queue, [3]int{neighbor[0], neighbor[1], currentLocation[2] + 1})
				visited[neighbor] = true
			}
		}
	}

	return -1
}

func neighbors(rows, columns, i, j int) [][2]int {
	locations := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	n := [][2]int{}
	for _, loc := range locations {
		if i+loc[0] < rows && j+loc[1] < columns && i+loc[0] >= 0 && j+loc[1] >= 0 {
			n = append(n, [2]int{i + loc[0], j + loc[1]})
		}
	}
	return n
}
