package jz09_两个栈实现队列

type CQueue struct {
	s1, s2 []int
}

func Constructor() CQueue {
	return CQueue{s1: []int{}, s2: []int{}}
}

// impl
func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

// impl
func (this *CQueue) DeleteHead() (val int) {
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	if len(this.s2) == 0 {
		return -1
	}
	ans := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return ans
}
