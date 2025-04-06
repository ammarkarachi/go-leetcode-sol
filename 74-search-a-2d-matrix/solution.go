package solution

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	l := 0
	h := (m * n) - 1
	for l <= h {
		mid := (l + h) / 2
		i, j := getCoordinate(mid, n)
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func getCoordinate(num, cols int) (int, int) {
	row := num / cols
	col := num % cols
	return row, col
}
