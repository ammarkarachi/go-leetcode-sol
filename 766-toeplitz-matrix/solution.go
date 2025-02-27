package solution

func isToeplitzMatrix(matrix [][]int) bool {
	rows := len(matrix)
	columns := len(matrix[0])

	for i := range columns {
		col := i
		row := 0
		first := matrix[row][col]
		for col < columns && row < rows {
			if first != matrix[row][col] {
				return false
			}
			col++
			row++
		}
	}

	for i := 1; i < rows; i++ {
		col := 0
		row := i
		first := matrix[row][col]

		for col < columns && row < rows {
			if first != matrix[row][col] {
				return false
			}
			col++
			row++
		}
	}

	return true
}
