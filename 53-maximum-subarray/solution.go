package solution

func maxSubArray(nums []int) int {
	curr := nums[0]
	max := nums[0]
	for i := 0; i < len(nums); i++ {

		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}

		if curr > max {
			max = curr
		}

	}

	return max
}
