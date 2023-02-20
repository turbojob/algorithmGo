package jz10_斐波那契数列

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof
func fib2(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib2(n-1) + fib2(n-2)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	ans := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmp := ans[len(ans)-1] + ans[len(ans)-2]
		tmp = tmp % 1000000007
		ans = append(ans, tmp)
	}
	return ans[len(ans)-1]
}
