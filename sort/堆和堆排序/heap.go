package main

import "fmt"

func main() {
	max := NewHeap()
	max.Push(1)
	max.Push(3)
	max.Push(7)
	max.Push(9)
	max.Push(2)
	max.Push(2)
	max.Pop()
	max.Pop()
	max.Push(3)
	max.Push(8)
	for !max.IsEmpty() {
		fmt.Print(max.Pop())
	}
}

type MaxHeap struct {
	heap     []int
	heapSize int
}

func NewHeap() *MaxHeap {
	return &MaxHeap{heap: make([]int, 0)}
}
func (h *MaxHeap) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *MaxHeap) Push(v int) {
	if h.heapSize == len(h.heap) {
		h.heap = append(h.heap, v)
	} else {
		h.heap[h.heapSize] = v
	}
	h.heapInsert(h.heapSize)
	h.heapSize++
}
func (h *MaxHeap) heapInsert(i int) {
	//不断往上判断
	//i == 0， （i-1）/2 = 0
	for h.heap[i] > h.heap[(i-1)/2] {
		h.heap[i], h.heap[(i-1)/2] = h.heap[(i-1)/2], h.heap[i]
		i = (i - 1) / 2
	}
}

func (h *MaxHeap) Pop() int {
	ans := h.heap[0]
	//堆的大小-1，此时h.heap[heapsize] 就是最后一个元素
	h.heapSize--
	if h.heapSize > 0 {
		h.heap[0], h.heap[h.heapSize] = h.heap[h.heapSize], h.heap[0]
		h.heapify(0, h.heapSize)
	}
	return ans
}
func (h *MaxHeap) heapify(i, heapSize int) {
	left := 2*i + 1
	for left < heapSize {
		large := left
		//左右最大的
		if left+1 < heapSize && h.heap[left+1] > h.heap[left] {
			large = left + 1
		}
		//父、孩子的比较，如果不如父亲 直接返回 不需要下沉了
		if h.heap[large] <= h.heap[i] {
			break
		}
		h.heap[i], h.heap[large] = h.heap[large], h.heap[i]
		i = large
		left = 2*i + 1
	}
}
