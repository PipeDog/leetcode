import java.util.Arrays;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=47 lang=java
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (64.36%)
 * Likes:    995
 * Dislikes: 0
 * Total Accepted:    288.7K
 * Total Submissions: 447.8K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 * 
 * 
 */

// @lc code=start
class Solution {

    private List<List<Integer>> resultList = new LinkedList<>();
    private LinkedList<Integer> trackList = new LinkedList<>();

    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);

        boolean[] usedList = new boolean[nums.length];
        permuteUniqueCore(nums, usedList);
        return resultList;
    }

    private void permuteUniqueCore(int[] nums, boolean[] usedList) {
        if (nums.length == trackList.size()) {
            resultList.add(new LinkedList<>(trackList));
            return;
        }

        for (int i = 0; i < nums.length; i++) {
            if (usedList[i] == true) {
                continue;
            }
            // usedList[i - 1] == false
            // @eg: [1, 2, 2', 2'']，首先对数组进行排序，保证相同值的元素处于相邻位置
            // 如果此时遍历到 2'，可以选择 2' 的情况应该是元素 2 已经被选择，才允许选择 2'
            // 否则应该跳过本次遍历
            if (i > 0 && nums[i - 1] == nums[i] && usedList[i - 1] == false) {
                continue;
            }

            usedList[i] = true;
            trackList.add(nums[i]);
            
            permuteUniqueCore(nums, usedList);

            trackList.removeLast();
            usedList[i] = false;
        }
    }

}
// @lc code=end

