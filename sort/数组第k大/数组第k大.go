package main

import "fmt"

//https://leetcode.cn/problems/kth-largest-element-in-an-array/
func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	//第k大
	//1大  --》n-1位置的
	//2大	-》n-2
	space := n - k
	return process(nums, 0, len(nums)-1, space)
}
func process(arr []int, l, r, space int) int {
	info := partition(arr, l, r, arr[l])
	if info[0] <= space && space <= info[1] {
		return arr[space]
	} else if info[0] > space {
		return process(arr, l, info[0]-1, space)
	} else {
		return process(arr, info[1]+1, r, space)
	}
}
func partition(arr []int, l, r, pivot int) []int {
	info := make([]int, 2)
	less, more, cur := l-1, r+1, l
	for {
		if cur >= more {
			break
		}
		if arr[cur] < pivot {
			less++
			swap(arr, cur, less)
			cur++
		} else if arr[cur] > pivot {
			more--
			swap(arr, cur, more)
		} else {
			cur++
		}
	}
	info[0], info[1] = less+1, more-1
	return info
}
func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}
