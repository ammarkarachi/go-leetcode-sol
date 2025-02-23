// https://leetcode.com/problems/merge-sorted-array/description/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days
package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	copyNums1 := append([]int(nil), nums1...)
	p1 := 0
	p2 := 0
	for p := 0; p < m+n; p++ {
		val := 0
		if p2 >= n || (p1 < m && copyNums1[p1] < nums2[p2]) {
			val = copyNums1[p1]
			p1++
		} else {
			val = nums2[p2]
			p2++
		}
		nums1[p] = val
	}
}
