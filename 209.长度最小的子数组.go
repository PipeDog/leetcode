/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (39.42%)
 * Likes:    262
 * Dislikes: 0
 * Total Accepted:    42.3K
 * Total Submissions: 99.8K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s
 * 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。
 * 
 * 示例: 
 * 
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 * 
 * 
 * 进阶:
 * 
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 * 
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	len := len(nums)
	sum := 0

	for i := 0; i < len; i++ {
		if nums[i] >= s {
			return 1
		}

		sum += nums[i]
	}

	if sum < s {
		return 0
	}

	left, right := 0, len - 1
	retLen := len

	for ; left < right; {
		tmpLeft := sum - nums[left]
		tmpRight := sum - nums[right]

		if tmpLeft <= s && tmpRight <= s {
			fmt.Printf("1_l=%d,r=%d,len=%d", left, right, retLen)
			// sum -= nums[left]
			// sum -= nums[right]
			// left++; right--
			retLen--; retLen--
			return retLen
		} else if tmpLeft <= s {
			sum -= nums[right]
			right--; retLen--
		} else {
			sum -= nums[left]
			left++; retLen--
		}
		//  else if tmpRight <= s {
		// 	sum -= nums[left]
		// 	left++; retLen--
		// } else {
		// 	// sum -= nums[left]
		// 	// // sum -= nums[right]
		// 	// left++//; right--
		// 	// retLen--//; retLen--
		// }
	}

	fmt.Printf("2_l=%d,r=%d,len=%d", left, right, retLen)
	return retLen
}
// @lc code=end

