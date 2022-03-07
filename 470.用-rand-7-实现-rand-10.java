/*
 * @lc app=leetcode.cn id=470 lang=java
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
/**
 * The rand7() API is already defined in the parent class SolBase.
 * public int rand7();
 * @return a random integer in the range 1 to 7
 */
class Solution extends SolBase {
    public int rand10() {
        // return (rand7() + rand7() + rand7() + rand7() + rand7() + rand7() + rand7() + rand7() + rand7() + rand7()) % 10 + 1;
        while (true) {
            int num = (rand7() - 1) * 7 + rand7();
            if (num <= 40) {
                return num % 10 + 1;
            }
            num = (num - 40 - 1) * 7 + rand7();
            if (num <= 60) {
                return num % 10 + 1;
            }
            num = (num - 60 - 1) * 7 + rand7();
            if (num <= 20) {
                return num % 10 + 1;
            }
        }
    }
}
// @lc code=end

