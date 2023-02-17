package main

import (
	"fmt"
)

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func main() {
	fmt.Print(replaceSpace("We are happy."))
}
func replaceSpace(s string) string {
	str := []rune(s)
	var ans []rune
	for _, v := range str {
		if v != ' ' {
			ans = append(ans, v)
		} else {
			ans = append(ans, '%')
			ans = append(ans, '2')
			ans = append(ans, '0')
		}
	}
	return string(ans)
}
