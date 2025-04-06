package solution

func setZeroes(matrix [][]int) {
	cols := make(map[int]bool)
	rows := make(map[int]bool)

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				cols[j] = true
				rows[i] = true
			}
		}
	}

	for k, _ := range cols {
		for i := range len(matrix) {
			matrix[i][k] = 0
		}
	}

	for k, _ := range rows {
		for j := range len(matrix[0]) {
			matrix[k][j] = 0
		}
	}

}
