package main

type NumArray struct {
	Left  *NumArray
	Right *NumArray
	Start int
	End   int
	Info  int
}

func (this *NumArray) init(nums []int) {
	if this.Start >= this.End {
		this.Info = nums[this.Start]
		return
	}
	mid := (this.Start + this.End) / 2
	this.Left = &NumArray{
		Start: this.Start,
		End:   mid,
	}
	this.Left.init(nums)
	this.Right = &NumArray{
		Start: mid + 1,
		End:   this.End,
	}
	this.Right.init(nums)
	this.Info = this.Left.Info + this.Right.Info
}

func (this *NumArray) updateSingle(idx int, val int) {
	if idx < this.Start || idx > this.End {
		return
	}
	if idx == this.Start && idx == this.End {
		this.Info = val
		return
	}
	if this.Left != nil {
		this.Left.updateSingle(idx, val)
		this.Right.updateSingle(idx, val)
		this.Info = this.Left.Info + this.Right.Info
	}
}

func (this *NumArray) queryRange(left int, right int) int {
	if left <= this.Start && right >= this.End {
		return this.Info
	}
	if this.Left != nil {
		return this.Left.queryRange(left, right) + this.Right.queryRange(left, right)
	}
	return 0
}

func Constructor1(nums []int) NumArray {
	na := &NumArray{
		Start: 0,
		End:   len(nums) - 1,
	}
	na.init(nums)
	return *na
}

func (this *NumArray) Update(index int, val int) {
	this.updateSingle(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.queryRange(left, right)
}

func main() {
	obj := Constructor1([]int{1, 3, 5})
	print(obj.SumRange(0, 2))
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
