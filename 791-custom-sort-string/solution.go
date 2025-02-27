package solution

func customSortString(order string, s string) string {
	histogram := buildHistogram(s)
	i := 0
	runeAnswer := make([]rune, len(s))

	for _, c := range order {
		n := histogram[int(c-'a')]
		j := 0
		for j < n {
			runeAnswer[i+j] = c
			j++
		}
		histogram[int(c-'a')] = 0
		i += j
	}

	for a, count := range histogram {
		j := 0
		for j < count {
			runeAnswer[i+j] = rune(a + int('a'))
			j++
		}
		i += j
	}

	return string(runeAnswer)
}

func buildHistogram(s string) [26]int {
	hist := [26]int{}
	for _, c := range s {
		hist[int(c-'a')]++
	}
	return hist
}
