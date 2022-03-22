import java.util.ArrayList;
import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=77 lang=java
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (76.97%)
 * Likes:    910
 * Dislikes: 0
 * Total Accepted:    306.4K
 * Total Submissions: 398.1K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 * 
 * 你可以按 任何顺序 返回答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 
 * 
 */

// @lc code=start
class Solution {

    // 注意避免重复

    private List<List<Integer>> result = new ArrayList<>();
    private LinkedList<Integer> backtrack = new LinkedList<>();
    private boolean[] used;

    public List<List<Integer>> combine(int n, int k) {
        used = new boolean[n];
        int[] nums = new int[n];
        for (int i = 1; i <= n; i++) {
            nums[i - 1] = i;
        }
        backtrackCore(nums, 0, k);
        return result;
    }

    private void backtrackCore(int[] nums, int start, int k) {
        if (backtrack.size() == k) {
            result.add(new LinkedList(backtrack));
            return;
        }

        for (int i = start; i < nums.length; i++) {
            if (used[i]) {
                continue;
            }

            used[i] = true;
            backtrack.add(nums[i]);
            backtrackCore(nums, i, k);
            used[i] = false;
            backtrack.removeLast();
        }
    }

}
// @lc code=end

