/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 *
 * https://leetcode-cn.com/problems/valid-triangle-number/description/
 *
 * algorithms
 * Medium (45.94%)
 * Likes:    50
 * Dislikes: 0
 * Total Accepted:    3.7K
 * Total Submissions: 7.9K
 * Testcase Example:  '[2,2,3,4]'
 *
 * 给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [2,2,3,4]
 * 输出: 3
 * 解释:
 * 有效的组合是: 
 * 2,3,4 (使用第一个 2)
 * 2,3,4 (使用第二个 2)
 * 2,2,3
 * 
 * 
 * 注意:
 * 
 * 
 * 数组长度不超过1000。
 * 数组里整数的范围为 [0, 1000]。
 * 
 * 
 */

// @lc code=start
func triangleNumber(nums []int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}

	// Sort and reverse
	var sort func(nums []int, left, right int)

	sort = func(nums []int, left, right int) {
		if left >= right {
			return
		}

		i, j := left, right
		key := nums[i]

		for i < j {
			for i < j && nums[j] >= key {
				j--
			}
			nums[i] = nums[j]
			for i < j && nums[i] <= key {
				i++
			}
			nums[j] = nums[i]
		}

		nums[i] = key
		sort(nums, left, i - 1)
		sort(nums, i + 1, right)
	}

	reverse := func(nums []int) {
		for i := 0; i < len(nums) / 2; i++ {
			nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
		}
	}

	sort(nums, 0, len(nums) - 1)
	reverse(nums)

	if nums[0] <= 0 {
		return 0
	}

	// Add up
	count := 0

	for i := 0; i < len(nums) - 2; i++ {
		left, right := i + 1, len(nums) - 1

		for left < right {
			// 
			/*
				@eg: [4, 4, 3, 2, 2, 2]
					  ^  ^           ^
					  |  |           |
					  i left       right
				
				current loop count = right - left
			*/
			if nums[left] + nums[right] > nums[i] {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}
// @lc code=end

