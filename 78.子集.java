import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=78 lang=java
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (80.38%)
 * Likes:    1539
 * Dislikes: 0
 * Total Accepted:    400.6K
 * Total Submissions: 498.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 * 
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0]
 * 输出：[[],[0]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10 
 * nums 中的所有元素 互不相同
 * 
 * 
 */

// @lc code=start
class Solution {

    private LinkedList<Integer> backtrack = new LinkedList<>();
    private List<List<Integer>> result = new ArrayList<>();

    public List<List<Integer>> subsets(int[] nums) {
        backtrackCore(nums, 0);
        return result;
    }

    private void backtrackCore(int[] nums, int start) {
        result.add(new LinkedList(backtrack));

        for (int i = start; i < nums.length; i++) {
            // 做选择
            backtrack.add(nums[i]);
            backtrackCore(nums, i + 1);
            // 撤销选择
            backtrack.removeLast();
        }
    }

}
// @lc code=end

