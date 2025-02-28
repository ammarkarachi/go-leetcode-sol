package solution

import (
	"strings"
)

func groupStrings(strings []string) [][]string {
	patternMap := make(map[string][]string)
	for _, s := range strings {
		pattern := shiftPattern(&s)
		p, exists := patternMap[pattern]
		if !exists {
			patternMap[pattern] = []string{s}
		} else {
			patternMap[pattern] = append(p, s)
		}
	}
	ans := make([][]string, len(patternMap))
	i := 0
	for _, v := range patternMap {
		ans[i] = v
		i++
	}
	return ans
}

func shiftPattern(s1 *string) string {
	s := *s1
	n := len(s)

	if n == 0 {
		return ""
	}

	shift := int(s[0] - 'a')
	var sb strings.Builder
	sb.Grow(n)

	for _, c := range s {
		// add 26 to ensure we don't get a negative value
		// take modulo to avoid overflow
		// this basically gives you the index from a <-> z
		// then add a to it
		ch := byte((int(c-'a')-shift+26)%26 + int('a'))
		sb.WriteByte(ch)
	}
	return sb.String()
}
