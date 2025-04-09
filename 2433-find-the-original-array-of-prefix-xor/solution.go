package solution

func findArray(pref []int) []int {
	ans := make([]int, len(pref))
	ans[0] = pref[0]
	i := 1
	for i < len(pref) {
		ans[i] = (pref[i] ^ pref[i-1])
		i++
	}

	return ans

}
