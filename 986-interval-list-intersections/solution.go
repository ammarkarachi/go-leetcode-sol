package solution

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	lenA := len(firstList)
	lenB := len(secondList)
	i := 0
	j := 0
	ans := [][]int{}
	for i < lenA && j < lenB {

		l := max(firstList[i][0], secondList[j][0])
		h := min(firstList[i][1], secondList[j][1])
		if l <= h {
			ans = append(ans, []int{l, h})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}

	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
