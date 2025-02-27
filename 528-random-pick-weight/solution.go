package solution

import "math/rand"

type Solution struct {
	Total  int
	Weight []int
}

func Constructor(w []int) Solution {
	total := 0
	cumWeight := make([]int, len(w))
	for i, weight := range w {
		total += weight
		cumWeight[i] = total
	}
	return Solution{
		Total:  total,
		Weight: cumWeight,
	}
}

func (this *Solution) PickIndex() int {
	l, r := 0, len(this.Weight)
	target := rand.Int() % this.Total
	for l < r {
		mid := l + (r-l)/2
		if this.Weight[mid] <= target {
			l = mid - 1
		} else {
			r = mid
		}
	}
	return l
}
