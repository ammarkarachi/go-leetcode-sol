package solution

func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	columns := len(mat[0])
	i := 0
	total := rows * columns
	goingDown := false
	increasingRow := true
	ans := make([]int, total)
	row := 0
	column := 0

	for i < total {

		currentRow := row
		currentColumn := column
		tempArray := []int{}
		for currentRow >= 0 && currentColumn < columns {
			if goingDown {
				tempArray = append([]int{mat[currentRow][currentColumn]}, tempArray...)
			} else {
				tempArray = append(tempArray, mat[currentRow][currentColumn])
			}
			currentRow--
			currentColumn++
		}
		j := 0

		for j < len(tempArray) {
			ans[j+i] = tempArray[j]
			j++
		}
		i += j

		if row == rows-1 {
			increasingRow = false
		}

		if increasingRow {
			row++
		} else {
			column++
		}
		goingDown = !goingDown
	}

	return ans

}
