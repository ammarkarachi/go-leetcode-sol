package solution

func findBuildings(heights []int) []int {
	length := len(heights)
	i := length - 1
	maxHeight := -1
	indices := make([]int, length)
	count := length - 1
	for i >= 0 {
		if maxHeight < heights[i] {
			indices[count] = i
			maxHeight = heights[i]
			count--
		}
		i--
	}

	return indices[(count + 1):]
}
