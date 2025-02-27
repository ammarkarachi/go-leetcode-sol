package solution

func calculate(s string) int {
	answer := 0
	currentNumber := 0
	lastNumber := 0
	operand := '+'
	for i, b := range s {
		if isDigit(b) {
			currentNumber = currentNumber*10 + toDigit(b)
		}
		if (!isDigit(b) && !isSpace(b)) || i == len(s)-1 {
			if operand == '+' || operand == '-' {
				answer += lastNumber
				if operand == '-' {
					lastNumber = currentNumber * -1
				} else {
					lastNumber = currentNumber
				}
			}
			if operand == '*' {
				lastNumber = lastNumber * currentNumber
			}
			if operand == '/' {
				lastNumber = lastNumber / currentNumber
			}
			currentNumber = 0
			operand = b
		}
	}
	answer += lastNumber
	return answer
}

func isSpace(s rune) bool {
	return s == ' '
}

func toDigit(s rune) int {
	return int(s - '0')
}

func isDigit(s rune) bool {
	return (s >= '0' && s <= '9')
}
