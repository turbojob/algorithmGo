package jz03_数组重复的数字

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/?favorite=xb9nqhhg
func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] == false {
			m[v] = true
		} else {
			return v
		}
	}
	return -1
}
