/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	slow := 0
	valCount := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			valCount++
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return len(nums) - valCount
}

