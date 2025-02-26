package solution

func searchRange(nums []int, target int) []int {
	start := searchBound(nums, target, true)
	if start == -1 {
		return []int{-1, -1}
	}
	end := searchBound(nums, target, false)
	return []int{start, end}
}

func searchBound(nums []int, target int, first bool) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			if first {
				if l == mid || nums[mid-1] != target {
					return mid
				}
				r = mid - 1
			} else {
				if r == mid || nums[mid+1] != target {
					return mid
				}
				l = mid + 1
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
