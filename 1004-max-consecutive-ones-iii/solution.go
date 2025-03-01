package solution

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			k--
		}

		if k < 0 {
			k += 1 - nums[l]
			l++
		}

		r++
	}

	return r - l
}
