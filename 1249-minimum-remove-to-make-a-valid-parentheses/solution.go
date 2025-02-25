package solution

func minRemoveToMakeValid(s string) string {
	indexSet := make(map[int]bool)
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				indexSet[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for _, s := range stack {
		indexSet[s] = true
	}

	runes := make([]rune, len(s)-len(indexSet))
	count := 0
	for i := 0; i < len(s); i++ {
		if _, ok := indexSet[i]; ok {
			continue
		}
		runes[count] = rune(s[i])
		count++
	}
	return string(runes)
}
