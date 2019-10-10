/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
    if len(nums) <= 1 {
		return 0
	}

	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	// 贪心算法、最长无重复子字符串的方法相结合
	stepCount := 0
	endLoc := 0
	maxLoc := 0

	for loc := 0; loc < len(nums) - 1; loc++ {
		stepLen := nums[loc]
		maxLoc = max(maxLoc, loc + stepLen)

		if loc == endLoc {
			endLoc = maxLoc
			stepCount++
		}
	}

	return stepCount
}
// @lc code=end

