/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}

	val := x
	ret := 0

	for ; val > 0; val /= 10 {
		ret = ret * 10 + val % 10
	}

	return (ret == x)
}

