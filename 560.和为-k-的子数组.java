import java.util.HashMap;
import java.util.Map;

/*
 * @lc app=leetcode.cn id=560 lang=java
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.98%)
 * Likes:    1417
 * Dislikes: 0
 * Total Accepted:    210.5K
 * Total Submissions: 468K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 * 
 * 
 */

// @lc code=start
class Solution {

    // 前缀和，O(n^2) 优化为 O(n)

    public int subarraySum(int[] nums, int k) {
        int len = nums.length;
        int[] preSums = new int[len + 1];
        Map<Integer, Integer> countTable = new HashMap<>();
        // 注意：base case
        countTable.put(0, 1);

        int result = 0;

        for (int i = 1; i <= len; i++) {
            preSums[i] = preSums[i - 1] + nums[i - 1];

            int need = preSums[i] - k;
            if (countTable.containsKey(need)) {
                System.out.print(String.format("preSums[%d]=%d, need=%d, count=%d,\n", i, preSums[i], need,
                        countTable.get(need)));
                result += countTable.get(need);
            }

            int currentCount = countTable.getOrDefault(preSums[i], 0);
            currentCount++;
            countTable.put(preSums[i], currentCount);
        }

        return result;
    }

}
// @lc code=end
