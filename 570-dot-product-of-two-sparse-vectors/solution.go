package solution

type SparseVector struct {
	Map map[int]int
}

func Constructor(nums []int) SparseVector {
	sparseMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			sparseMap[i] = nums[i]
		}
	}
	vec := SparseVector{
		Map: sparseMap,
	}
	return vec

}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	sum := 0
	for k, v := range this.Map {
		if val, ok := vec.Map[k]; ok {
			sum += val * v
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
