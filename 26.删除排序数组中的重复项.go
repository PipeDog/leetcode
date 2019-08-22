/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// @eg => [1,1,2,3,4]
	// retLen需要从1开始，如果从0开始的话，会把第一个0覆盖掉，会删除掉第一个值
	retLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[retLen] = nums[i]
			retLen++
		}
	}

	return retLen
}

