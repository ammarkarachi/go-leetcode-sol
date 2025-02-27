package solution

func calculate(s string) int {
	answer := 0
	stack := []int{}
	currentNumber := 0
	operand := '+'
	for i, b := range s {
		if isDigit(b) {
			currentNumber = currentNumber*10 + toDigit(b)
		} else if !isDigit(b) && !isSpace(b) && i == len(s)-1 {
			if operand == '+' {
				push(&stack, currentNumber)
			}
			if operand == '-' {
				push(&stack, -currentNumber)
			}
			if operand == '*' {
				val := pop(&stack) * currentNumber
				push(&stack, val)
			}
			if operand == '/' {
				val := pop(&stack) / currentNumber
				push(&stack, val)
			}
			currentNumber = 0
			operand = b
		}
	}
	for len(stack) > 0 {
		answer += pop(&stack)
	}
	return answer
}
func push(stack *[]int, val int) {
	(*stack) = append([]int{val}, *stack...)
}
func pop(stack *[]int) int {
	val := (*stack)[len(*stack)-1]
	(*stack) = (*stack)[:len(*stack)-1]
	return val
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
