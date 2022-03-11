/*
 * @lc app=leetcode.cn id=46 lang=swift
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (78.49%)
 * Likes:    1832
 * Dislikes: 0
 * Total Accepted:    538K
 * Total Submissions: 685.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1]
 * 输出：[[1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 * 
 * 
 */

// @lc code=start
class Solution {

    private var result = [[Int]]()

    func permute(_ nums: [Int]) -> [[Int]] {
        // 已确定路径
        var paths = [Int]()

        // 标记已使用数字
        var used = [Int : Bool]()
        for num in nums {
            used[num] = false
        }
        
        // 开始回溯
        permuteCore(nums, &paths, &used)
        return result
    }

    // nums  : 可选列表
    // paths : 已确定路径
    // used  : 使用标记
    func permuteCore(_ nums: [Int], _ paths: inout [Int], _ used: inout [Int : Bool]) {
        if nums.count == paths.count {
            // Copy paths to new Array
            var copyPaths = paths
            result.append(copyPaths)
            return
        }

        for num in nums {
            // 过滤已选择数字
            let isUsed = used[num]!
            if isUsed {
                continue
            }

            // 做选择
            used[num] = true
            paths.append(num)

            // 回溯递归
            permuteCore(nums, &paths, &used)

            // 撤销选择
            paths.removeLast()
            used[num] = false
        }
    }

}
// @lc code=end

