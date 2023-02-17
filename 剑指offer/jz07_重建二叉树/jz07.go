package main

import (
	"fmt"
)

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	root = buildTree([]int{-1}, []int{-1})
	fmt.Print(root.Val)
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	return process(preorder, inorder, 0, len(preorder)-1, 0, len(preorder)-1)
}
func process(pre, in []int, p1, p2, i1, i2 int) *TreeNode {
	if p1 > p2 || i1 > i2 {
		return nil
	}
	if p1 == p2 {
		return &TreeNode{Val: pre[p1]}
	}
	//前序遍历 root -l -r
	root := TreeNode{Val: pre[p1]}
	//find root in inArr
	space := i1
	for {
		if in[space] == pre[p1] {
			break
		}
		space++
	}
	//in  i1 ... space ... i2
	//len(left) = space - i1; len(right) = i2 - space
	//pre p1 p1+1 ... ?
	//? - p1 - 1 + 1 = space - i1
	//? = space - i1 + p1
	//pre p1 p1+1...? ?+1...p2
	left := process(pre, in, p1+1, space-i1+p1, i1, space-1)
	right := process(pre, in, space-i1+p1+1, p2, space+1, i2)
	root.Left = left
	root.Right = right
	return &root
}
