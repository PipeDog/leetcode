/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	num1 := 1
	num2 := 1

	for i := 2; i <= n; i++ {
		sum := num1 + num2
		num2 = num1
		num1 = sum
	}
	return num1
}

