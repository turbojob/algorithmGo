package jz06_从尾到头打印链表

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var ans, tmp []int
	hair := ListNode{Val: 0}
	hair.Next = head
	cur := hair
	for {
		if cur.Next == nil {
			break
		}
		tmp = append(tmp, cur.Next.Val)
		cur = *cur.Next
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		ans = append(ans, tmp[i])
	}
	return ans
}
