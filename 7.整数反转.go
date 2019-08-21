/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	val := x
	ret := 0
	isPlus := true
	if x < 0 {
		isPlus = false
		val = -x
	}
	
	for ; val > 0; val /= 10 {
		ret = ret * 10 + val % 10
		if ret > 2147483647 {
			return 0
		}
	}

	if !isPlus {
		return -ret
	}
	return ret
}

