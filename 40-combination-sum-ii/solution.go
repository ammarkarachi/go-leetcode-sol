package solution

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	temp := []int{}
	slices.Sort(candidates)
	var backtrack func(result *[][]int, tempResult []int, candidates *[]int, left int, index int)
	backtrack = func(result *[][]int, tempResult []int, candidates *[]int, left int, index int) {
		if left == 0 {
			r := *result
			r = append(r, tempResult)
			*result = r
		} else if left > 0 {
			c := *candidates
			for i := index; i < len(c) && left >= c[i]; i++ {
				if i > index && c[i] == c[i-1] {
					continue
				}
				tempResult = append(tempResult, c[i])

				backtrack(result, tempResult, candidates,
					left-c[i], i+1)

				tempResult = tempResult[:len(tempResult)-1]
			}
		}

	}
	backtrack(&ans, temp, &candidates, target, 0)
	return ans

}
