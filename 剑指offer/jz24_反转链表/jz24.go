package jz24_反转链表

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/submissions/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
