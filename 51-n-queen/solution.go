package solution

func solveNQueens(n int) [][]string {

	queue := [][]int{
		{0},
		{1},
		{2},
		{3},
	}
	locations := [][]int{}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		rowToPlace := len(top)
		inValidLocations := make(map[int]bool)
		for i, location := range top {
			inValidLocations[location] = true
			inValidLocations[location-(rowToPlace-i)] = true
			inValidLocations[location+(rowToPlace-i)] = true
		}
		placed := false
		for i := range n {
			_, exists := inValidLocations[i]
			if !exists {
				placed = true
				top = append(top, i)
				break
			}
		}

		if len(top) == n {
			locations = append(locations, top)
			continue
		}

		if placed {
			queue = append(queue, top)
		}

	}
	ans := make([][]string, n)
	for i, location := range locations {
		ans[i] = make([]string, n)
		for j, pos := range location {
			runes := make([]rune, n)
			runes[pos] = 'Q'
			ans[i][j] = string(runes)
		}
	}
	return ans
}
