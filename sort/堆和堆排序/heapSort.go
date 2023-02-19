package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5, 3, 2, 1, 4}))
	fmt.Println(sortArray([]int{5, 3, 221, 2, 42, 1, 4}))
}
func sortArray(nums []int) []int {
	stack := NewStack()
	return stack.heapSort(nums)
}

type maxStack struct {
	arr  []int
	size int
}

func NewStack() *maxStack {
	return &maxStack{arr: make([]int, 0)}
}
func (s *maxStack) heapSort(arr []int) []int {
	//快速建堆
	s.arr = arr
	s.size = len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		s.Heapify(i, s.size)
	}
	//pop 排序
	for i := 0; i < len(arr); i++ {
		s.Pop()
	}
	return s.arr
}
func (s *maxStack) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}
func (s *maxStack) heapInsert(v int) {
	for s.arr[v] > s.arr[(v-1)/2] {
		s.swap(v, (v-1)/2)
		v = (v - 1) / 2
	}
}
func (s *maxStack) Push(v int) {
	if s.size == len(s.arr) {
		s.arr = append(s.arr, v)
	} else {
		s.arr[s.size] = v
	}
	s.heapInsert(s.size)
	s.size++

}
func (s *maxStack) Heapify(i, size int) {
	left := 2*i + 1

	for left < size {
		max := left
		if left+1 < size && s.arr[left+1] > s.arr[left] {
			max = left + 1
		}
		if s.arr[max] <= s.arr[i] {
			break
		}
		s.swap(i, max)
		i = max
		left = 2*i + 1
	}
}
func (s *maxStack) Pop() int {
	if s.size == 0 {
		return -1
	}
	ans := s.arr[0]
	s.size--
	s.swap(0, s.size)
	s.Heapify(0, s.size)
	return ans
}
