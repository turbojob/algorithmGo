package main

import (
	"log"
)

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
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, l, r int) {
	if r-l+1 < 2 {
		return
	}
	info := partition(nums, l, r, nums[l])
	quickSort(nums, l, info[0]-1)
	quickSort(nums, info[1]+1, r)
}
func partition(nums []int, l, r, pivot int) []int {
	info := make([]int, 2)
	//l...r 进行partition操作
	//返回等于区的左右侧
	//less 小于区的右边界  more 大于区的左边界
	less, more, cur := l-1, r+1, l
	for {
		if cur >= more {
			break
		}
		if nums[cur] < pivot {
			less++
			swap(nums, less, cur)
			cur++
		} else if nums[cur] > pivot {
			more--
			swap(nums, cur, more)
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
