package solution

type SparseVector struct {
	IndexValue [][]int
}

func Constructor(nums []int) SparseVector {
	ivPairs := [][]int{}
	for i := 0; i < len(nums); i++ {
		ivPairs = append(ivPairs, []int{i, nums[i]})
	}
	vec := SparseVector{
		IndexValue: ivPairs,
	}
	return vec

}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	sum, i, j := 0, 0, 0
	for i < len(this.IndexValue) && j < len(vec.IndexValue) {
		if this.IndexValue[i][0] == this.IndexValue[j][0] {
			sum += this.IndexValue[i][1] * vec.IndexValue[j][1]
			i++
			j++
		} else if this.IndexValue[i][0] < vec.IndexValue[j][0] {
			i++
		} else {
			j++
		}

	}
	return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
