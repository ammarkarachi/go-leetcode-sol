package solution

//https://leetcode.com/problems/find-peak-element/description/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days

func findPeakElement(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, l int, r int) int {
	if l == r {
		return l
	}
	mid := (l + r) / 2
	if nums[mid] > nums[mid+1] {
		return search(nums, l, mid)
	} else {
		return search(nums, mid+1, r)
	}
}
