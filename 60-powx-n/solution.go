package solution

import "math"

func myPow(x float64, n int) float64 {
	n1 := int(math.Abs(float64(n)))
	ans := 1.0

	if x == 1 || x == 0 {
		return x
	}

	pow := n1
	for pow != 0 {

		if pow%2 == 1 {
			ans = ans * x
			pow--
		}

		x = x * x
		pow = pow / 2
	}

	if int(n1) != n {
		return 1 / ans
	}

	return ans
}
