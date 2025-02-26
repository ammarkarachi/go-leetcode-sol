package solution

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	currentIndex := 0
	nextIndex := 1

	for nextIndex < len(intervals) {

		if intervals[currentIndex][1] >= intervals[nextIndex][0] {
			intervals[currentIndex][1] = max(intervals[currentIndex][1], intervals[nextIndex][1])
		} else {
			currentIndex++
			intervals[currentIndex] = intervals[nextIndex]
		}
		nextIndex++
	}
	return intervals[:currentIndex+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
