/*
 * @lc app=leetcode.cn id=962 lang=golang
 *
 * [962] 最大宽度坡
 *
 * https://leetcode-cn.com/problems/maximum-width-ramp/description/
 *
 * algorithms
 * Medium (32.24%)
 * Likes:    59
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 13.6K
 * Testcase Example:  '[6,0,8,2,1,5]'
 *
 * 给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
 * 
 * 找出 A 中的坡的最大宽度，如果不存在，返回 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[6,0,8,2,1,5]
 * 输出：4
 * 解释：
 * 最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
 * 
 * 
 * 示例 2：
 * 
 * 输入：[9,8,1,0,1,9,4,0,4,1]
 * 输出：7
 * 解释：
 * 最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= A.length <= 50000
 * 0 <= A[i] <= 50000
 * 
 * 
 * 
 * 
 */

// @lc code=start
func maxWidthRamp(A []int) int {
	stack := []int{}
	count := len(A)
    	
	for i := 0; i < count; i++ {
		if len(stack) == 0 || A[stack[len(stack) - 1]] > A[i] {
			stack = append(stack, i)
		}
	}

	ret, i := 0, count - 1
	for i > ret {
		for len(stack) > 0 && A[stack[len(stack) - 1]] <= A[i] {
			if i - stack[len(stack) - 1] > ret {
				ret = i - stack[len(stack) - 1]
			}
			stack = stack[: len(stack) - 1]
		}
		i--
	}
    
    return ret
}
// @lc code=end

