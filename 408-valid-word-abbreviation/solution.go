package solution

func validWordAbbreviation(word string, abbr string) bool {
	i := 0
	j := 0
	for j < len(abbr) {
		if i >= len(word) {
			return false
		}
		if word[i] == abbr[j] {
			i++
			j++
		} else if '0' <= abbr[j] && abbr[j] <= '9' {
			val := 0
			if abbr[j] == '0' {
				return false
			}
			for j < len(abbr) && '0' <= abbr[j] && abbr[j] <= '9' {
				val = val*10 + int(abbr[j]-'0')
				j++
			}
			i += val

		} else {
			return false
		}

	}

	if !(i == len(word) && j == len(abbr)) {
		return false
	}

	return true
}
