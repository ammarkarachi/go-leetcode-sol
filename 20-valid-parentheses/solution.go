package solution

func isValid(s string) bool {
	stack := []rune{}
	for _, b := range s {

		if isOpen(b) {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if b == getClosing(top) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}

	return len(stack) == 0
}

func isOpen(b rune) bool {
	return b == '(' || b == '{' || b == '['
}

func getClosing(b rune) rune {
	if b == '(' {
		return ')'
	}

	if b == '{' {
		return '}'
	}

	if b == '[' {
		return ']'
	}
	return '0'
}
