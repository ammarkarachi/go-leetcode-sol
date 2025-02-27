package solution

func subarraySum(nums []int, k int) int {
	sumOccurence := make(map[int]int)
	sum, count := 0, 0
	sumOccurence[0] = 1
	for _, num := range nums {
		sum += num

		if idx, ok := sumOccurence[sum-k]; ok {
			count += idx
		}
		index, ok := sumOccurence[sum]
		if ok {
			sumOccurence[sum] = index + 1
		} else {
			sumOccurence[sum] = 1
		}
	}

	return count
}
