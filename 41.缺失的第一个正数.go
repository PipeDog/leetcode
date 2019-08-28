/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// https://leetcode-cn.com/problems/first-missing-positive/
func firstMissingPositive(nums []int) int {
	numsLen := len(nums)

	for i := 0; i < numsLen; {
		if nums[i] >= 1 && nums[i] <= numsLen && nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < numsLen; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	return numsLen + 1
}

