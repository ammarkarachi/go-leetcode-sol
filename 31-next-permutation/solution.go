package solution

func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}
func reverse(nums []int, start int) {
	l, r := start, len(nums)-1
	for l < r {
		swap(nums, l, r)
		l++
		r--
	}

}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
