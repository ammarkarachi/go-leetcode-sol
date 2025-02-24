package solution

import "math/rand"

func findKthLargest(nums []int, k int) int {
	pivot := nums[rand.Int()%len(nums)]
	left := make([]int, 0)
	right := make([]int, 0)
	mid := make([]int, 0)

	for _, num := range nums {
		if num < pivot {
			right = append(right, num)
		} else if num > pivot {
			left = append(left, num)
		} else {
			mid = append(mid, num)
		}
	}

	if k <= len(left) {
		return findKthLargest(left, k)
	}

	if len(left)+len(mid) < k {
		return findKthLargest(right, k-len(left)-len(mid))
	}
	return pivot
}

func minMaxArray(nums []int) (int, int) {
	max := nums[0]
	min := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}

	return min, max
}
