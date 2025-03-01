package solution

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	ans := []rune{}
	referenceString := strs[0]
	j := 0
	for j < len(referenceString) {
		i := 1
		for i < len(strs) {
			if j >= len(strs) || strs[i][j] != referenceString[j] {
				return string(ans)
			}
			i++
		}
		ans = append(ans, rune(referenceString[j]))
		j++
	}
	return string(ans)
}
