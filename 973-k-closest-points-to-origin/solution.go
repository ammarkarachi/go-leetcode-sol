package solution

import "container/heap"

type PointHeap [][]int

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] < h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {

	h := PointHeap(points)
	heap.Init(&h)
	i := 0
	length := min(k, len(points))
	ans := make([][]int, length)
	for i < length {
		point := heap.Pop(&h).([]int)
		ans[i] = point
		i++
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
