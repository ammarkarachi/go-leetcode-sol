package solution

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	n := len(nums)
	ans := [][]int{}
	if n == 0 {
		return [][]int{
			{lower, upper},
		}
	}

	if lower < nums[0] {
		ans = append(ans, []int{lower, nums[0] - 1})
	}

	for i := range n - 1 {

		if nums[i+1]-nums[i] <= 1 {
			continue
		}

		ans = append(ans, []int{nums[i] + 1, nums[i+1] - 1})
	}

	if upper > nums[n-1] {
		ans = append(ans, []int{nums[n-1] + 1, upper})
	}
	return ans
}
