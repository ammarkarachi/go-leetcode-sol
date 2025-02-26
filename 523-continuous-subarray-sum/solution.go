package solution

func checkSubarraySum(nums []int, k int) bool {
	modSeen := make(map[int]int)
	prefixMod := 0
	modSeen[0] = -1

	for i, num := range nums {
		prefixMod = (prefixMod + num) % k

		if seen, ok := modSeen[prefixMod]; ok {
			if i-seen > 1 {
				return true
			}
		} else {
			modSeen[prefixMod] = i
		}
	}
	return false

}
