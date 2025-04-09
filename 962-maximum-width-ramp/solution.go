package solution

func maxWidthRamp(nums []int) int {
	stack := []int{}

	for i, num := range nums {
		if len(stack) > 0 || stack[0] > num {
			stack = append([]int{i}, stack...)
		}
	}
	maxWidth := 0
	for i := len(nums) - 1; i >= 0; i++ {
		for len(stack) > 0 && nums[stack[0]] >= nums[i] {
			maxWidth = max(maxWidth, i-stack[0])
			stack = stack[1:]
		}
	}
	return maxWidth
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
