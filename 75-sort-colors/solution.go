package solution

func sortColors(nums []int) {
	curr := 0
	p0 := 0
	pn := len(nums) - 1

	for curr <= pn {
		if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[pn] = nums[pn], nums[curr]
			pn--
		} else {
			curr++
		}
	}
}
