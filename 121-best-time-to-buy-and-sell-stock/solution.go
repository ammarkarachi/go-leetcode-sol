package solution

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(price-minPrice, maxProfit)
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
