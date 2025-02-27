package solution

import (
	"math/rand/v2"
)

var (
	frequency map[int]int
	uniq      []int
)

func topKFrequent(nums []int, k int) []int {

	frequency = make(map[int]int)

	for _, num := range nums {
		if freq, ok := frequency[num]; ok {
			frequency[num] = freq + 1
		} else {
			frequency[num] = 1
		}
	}

	uniq = make([]int, len(frequency))
	i := 0
	for key := range frequency {
		uniq[i] = key
		i++
	}

	quickSelect(0, len(uniq)-1, len(uniq)-k)

	return uniq[len(uniq)-k:]
}

func partition(left, right, pivot int) int {

	freq, _ := frequency[uniq[pivot]]

	// send pivot to the right of the array
	uniq[pivot], uniq[right] = uniq[right], uniq[pivot]

	// get the index that is left most and store it
	storeIndex := left
	for i := left; i <= right; i++ {
		freqi, _ := frequency[uniq[i]]
		// if the frequency is less move them to left
		if freqi < freq {
			uniq[i], uniq[storeIndex] = uniq[storeIndex], uniq[i]
			storeIndex++
		}
	}
	uniq[storeIndex], uniq[right] = uniq[right], uniq[storeIndex]
	return storeIndex
}

func quickSelect(left, right, kSmall int) {

	if left == right {
		return
	}

	pivot := left + rand.Int()%(right-left+1)

	pivot = partition(left, right, pivot)

	if pivot == kSmall {
		return
	} else if pivot < kSmall {
		quickSelect(pivot+1, right, kSmall)
	} else {
		quickSelect(left, pivot-1, kSmall)
	}

}
