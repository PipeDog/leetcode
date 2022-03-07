/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
    for ;; {
		num := (rand7() - 1) * 7 + rand7()
		if num <= 40 {
			return num % 10 + 1
		}

		num = (num - 40 - 1) * 7 + rand7()
		if num <= 60 {
			return num % 10 + 1
		}

		num = (num - 60 - 1) * 7 + rand7()
		if num <= 20 {
			return num % 10 + 1
		}
	}
}
// @lc code=end

