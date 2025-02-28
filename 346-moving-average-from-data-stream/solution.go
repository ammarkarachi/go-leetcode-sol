package solution

type MovingAverage struct {
	Size      int
	HeadIndex int
	List      []int
	Capacity  int
	Total     int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size:      size,
		HeadIndex: 0,
		List:      make([]int, size),
		Capacity:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.addValue(val)
	return this.getAverage()
}
func (this *MovingAverage) addValue(val int) {
	currentTotal := this.Total + val - this.List[this.HeadIndex]
	this.List[this.HeadIndex] = val
	this.incrementElementCount()
	this.HeadIndex = (this.HeadIndex + 1) % this.Size
	this.Total = currentTotal
}

func (this *MovingAverage) incrementElementCount() {
	if this.Capacity < this.Size {
		this.Capacity++
	}
}
func (this *MovingAverage) elementCount() int {
	if this.Capacity < this.Size {
		return this.Capacity
	}
	return this.Size
}

func (this *MovingAverage) getAverage() float64 {
	return float64(this.Total) / float64(this.elementCount())
}
