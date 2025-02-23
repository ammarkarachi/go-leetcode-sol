// https://leetcode.com/problems/maximum-swap/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days
package solution

import (
	"strconv"
)

func maximumSwap(num int) int {
	numString := strconv.Itoa(num)
	n := len(numString)
	maxDigitIndex, swap1, swap2 := -1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if maxDigitIndex == -1 || numString[i] > numString[maxDigitIndex] {
			maxDigitIndex = i
		} else if numString[i] < numString[maxDigitIndex] {
			swap1 = i
			swap2 = maxDigitIndex
		}
	}
	if swap1 != -1 && swap2 != -2 {
		newNumber, _ := strconv.Atoi(swapLetters(numString, swap1, swap2))
		return newNumber
	}

	return num
}

func swapLetters(s string, i int, j int) string {
	letters := []rune(s)
	temp := letters[i]
	letters[i] = letters[j]
	letters[j] = temp
	return string(letters)
}
