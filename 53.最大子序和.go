/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if maxSum < curSum {
			maxSum = curSum
		}
	}

	return maxSum
}

