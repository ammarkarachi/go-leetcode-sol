package solution

type Solution struct {
	IndexMap map[int][]int
}

func Constructor(nums []int) Solution {
	indexMap := make(map[int][]int)
	for i, num := range nums {
		indices, ok := indexMap[num]
		if !ok {
			indexMap[num] = []int{i}
		} else {
			indexMap[num] = append(indices, i)
		}
	}
	return Solution{
		IndexMap: indexMap,
	}
}

func (this *Solution) Pick(target int) int {
	indices, ok := this.IndexMap[target]

	if ok {
		index := indices[0]
		indices = append(indices[1:], index)
		this.IndexMap[target] = indices
		return index
	}

	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
