package solution

func isNumber(s string) bool {

	nextStateMap := []map[string]int{
		{"digit": 1, "sign": 2, "dot": 3},     // a digit or sign or dot is allowed
		{"digit": 1, "dot": 4, "exponent": 5}, // after a digit a digit dot or exponent is allowed
		{"digit": 1, "dot": 3},                // after a sign digit or dot is allowed
		{"digit": 4},                          // after a dot only a digit is allowed but _4
		{"digit": 4, "exponent": 5},           // digit or exponent
		{"sign": 6, "digit": 7},               // after exponent and digit a sign and digit is allowed
		{"digit": 7},                          // after sign after exponent  , only digits allowed
		{"digit": 7},                          // after sign after exponent after digit only digits allowed
	}

	currentState := 0

	for _, c := range s {
		currentCharType := whatIsChar(c)

		if currentCharType == "not-valid" {
			return false
		}
		state, exists := nextStateMap[currentState][currentCharType]
		if !exists {
			return false
		} else {
			currentState = state
		}
	}

	return currentState == 4 || currentState == 7 || currentState == 1
}

func whatIsChar(c rune) string {

	if c == '+' || c == '-' {
		return "sign"
	}

	if c == '.' {
		return "dot"
	}

	if c == 'e' || c == 'E' {
		return "exponent"
	}

	if '0' <= c && c <= '9' {
		return "digit"
	}

	return "not-valid"

}
