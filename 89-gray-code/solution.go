package solution

func grayCode(n int) []int {
	l := 1 << n
	result := make([]int, l)
	i := 0
	for i < l {
		result[i] = i ^ i>>1
		i++
	}
	return result
}
