/*
 * @lc app=leetcode.cn id=1109 lang=java
 *
 * [1109] 航班预订统计
 *
 * https://leetcode-cn.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (60.31%)
 * Likes:    334
 * Dislikes: 0
 * Total Accepted:    75.8K
 * Total Submissions: 125.7K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 * 
 * 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
 * firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
 * 
 * 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 * 解释：
 * 航班编号        1   2   3   4   5
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       20  20
 * 预订记录 3 ：       25  25  25  25
 * 总座位数：      10  55  45  25  25
 * 因此，answer = [10,55,45,25,25]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：bookings = [[1,2,10],[2,2,15]], n = 2
 * 输出：[10,25]
 * 解释：
 * 航班编号        1   2
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       15
 * 总座位数：      10  25
 * 因此，answer = [10,25]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 2 * 10^4
 * 1 <= bookings.length <= 2 * 10^4
 * bookings[i].length == 3
 * 1 <= firsti <= lasti <= n
 * 1 <= seatsi <= 10^4
 * 
 * 
 */

// @lc code=start
class Solution {

    // 差分数组

    private int[] diff;

    public int[] corpFlightBookings(int[][] bookings, int n) {

        createDiffArray(n);

        for (int i = 0; i < bookings.length; i++) {
            int[] booking = bookings[i];
            int start = booking[0] - 1;
            int end = booking[1] - 1;
            int val = booking[2];
            System.out.print(String.format("start=%d, end=%d, val=%d\n", start, end, val));
            increment(start, end, val);
        }

        return getResult();
    }

    /**
     * 创建查分 diff 数组
     */
    private void createDiffArray(int n) {
        diff = new int[n];
    }

    /**
     * 增减操作
     * 
     * @param start 开始位置
     * @param end   结束为止
     * @param val   修改的值，正负数皆可
     */
    private void increment(int start, int end, int val) {
        diff[start] += val;

        if (end + 1 < diff.length) {
            diff[end + 1] -= val;
        }

        // logList("diff", diff);
        // logList("nums", getResult());
        // System.out.print("\n");
    }

    /**
     * 获取 N 次计算后的原始数组结果（根据 diff 数组还原）
     */
    private int[] getResult() {
        int[] nums = new int[diff.length];
        nums[0] = diff[0];

        for (int i = 1; i < nums.length; i++) {
            nums[i] = diff[i] + nums[i - 1];
        }

        return nums;
    }

    private void logList(String tag, int[] nums) {
        System.out.print(tag + " [");

        for (int i = 0; i < nums.length; i++) {
            System.out.print(String.format("%d, ", nums[i]));
        }

        System.out.print("]\n");
    }

}
// @lc code=end
