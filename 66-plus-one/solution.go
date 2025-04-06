package solution

func plusOne(digits []int) []int {
	carry := 0
	i := len(digits) - 1
	ans := make([]int, len(digits))
	ans[i] = digits[i] + 1
	for i >= 0 {

		if ans[i]+carry > 10 {
			carry = 1
			ans[i] = ans[i] % 10
		}

		if carry == 1 {
			i--
		} else {
			break
		}
	}
	if carry == 1 {
		ans = append([]int{carry}, ans...)
	}
	return ans
}
