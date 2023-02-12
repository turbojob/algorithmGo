package main

import "log"

func main() {
	arr := []int{10, 3, 0, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	sortArray(arr)
	log.Println(arr)
}

func sortArray(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l, mid, r int) {
	//arr[l...mid]  arr[mid+1...r]
	tmp := make([]int, r-l+1)
	index := 0
	p1, p2 := l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			tmp[index] = arr[p1]
			index++
			p1++
		} else {
			tmp[index] = arr[p2]
			index++
			p2++
		}
	}
	for ; p1 <= mid; p1++ {
		tmp[index] = arr[p1]
		index++
	}
	for ; p2 <= r; p2++ {
		tmp[index] = arr[p2]
		index++
	}
	for i := 0; i < len(tmp); i++ {
		arr[i+l] = tmp[i]
	}
}
