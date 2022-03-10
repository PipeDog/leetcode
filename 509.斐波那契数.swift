/*
 * @lc app=leetcode.cn id=509 lang=swift
 *
 * [509] 斐波那契数
 *
 * https://leetcode-cn.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.98%)
 * Likes:    413
 * Dislikes: 0
 * Total Accepted:    335.6K
 * Total Submissions: 502.1K
 * Testcase Example:  '2'
 *
 * 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
 * 
 * 
 * F(0) = 0，F(1) = 1
 * F(n) = F(n - 1) + F(n - 2)，其中 n > 1
 * 
 * 
 * 给定 n ，请计算 F(n) 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 2
 * 输出：1
 * 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3
 * 输出：2
 * 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 4
 * 输出：3
 * 解释：F(4) = F(3) + F(2) = 2 + 1 = 3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= n <= 30
 * 
 * 
 */

// @lc code=start
class Solution {
    
    // 自顶向下，时间复杂度最高的解法

    /*
    func fib(_ n: Int) -> Int {
        if n < 0 {
            return -1
        }
        if n == 0 {
            return 0
        }
        if n == 1 || n == 2 {
            return 1
        }

        return fib(n - 1) + fib(n - 2)        
    }
    */


    // 自顶向下，剪枝优化后时间复杂度为 O(n) 的解法

    /*
    private var dpDict : [Int : Int] = [Int : Int]()

    func fib(_ n: Int) -> Int {
        return fibCore(n)
    }

    func fibCore(_ n: Int) -> Int {
        if n < 0 {
            return -1
        }
        if n == 0 {
            return 0
        }
        if n == 1 || n == 2 {
            return 1
        }

        if dpDict.keys.contains(n) {
            return dpDict[n]!
        }

        dpDict[n] = fibCore(n - 1) + fibCore(n - 2)
        return dpDict[n]!
    }
    */


    // 自底向上，空间复杂度为 O(n) 的解法

    /*
    func fib(_ n: Int) -> Int {
        if n <= 0 {
            return n
        }
        if n == 1 || n == 2 {
            return 1
        }

        var dpDict : [Int : Int] = [Int : Int]()
        dpDict[1] = 1
        dpDict[2] = 1

        for i in 3...n {
            dpDict[i] = dpDict[i - 1]! + dpDict[i - 2]!
        }

        return dpDict[n]!
    }
    */


    // 自底向上，空间复杂度为 O(1) 的解法

    func fib(_ n: Int) -> Int {
        if n <= 0 {
            return n
        }
        if n == 1 || n == 2 {
            return 1
        }

        var val1 = 1, val2 = 1
        var ret = 0

        for i in 3...n {
            ret = val1 + val2
            val1 = val2
            val2 = ret
        }

        return ret
    }

}
// @lc code=end

