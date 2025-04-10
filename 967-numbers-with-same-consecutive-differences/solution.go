package solution

func numsSameConsecDiff(n int, k int) []int {
	lower := intPow(10, n-1)
	upper := intPow(10, n)
	ans := []int{}
	for i := range 10 {

		if i == 0 {
			continue
		}
		queue := []int{i}

		for len(queue) > 0 {
			dequeue := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if dequeue >= lower && dequeue < upper {
				ans = append(ans, dequeue)
				continue
			}
			lsd := dequeue % 10

			if k == 0 {
				queue = append(queue, dequeue*10+lsd)
				continue
			}

			if lsd+k <= 9 {
				queue = append(queue, dequeue*10+(lsd+k))
			}

			if lsd-k >= 0 {
				queue = append(queue, dequeue*10+(lsd-k))
			}
		}
	}
	return ans
}
func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
