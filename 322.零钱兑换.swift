/*
 * @lc app=leetcode.cn id=322 lang=swift
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (45.09%)
 * Likes:    1770
 * Dislikes: 0
 * Total Accepted:    381.5K
 * Total Submissions: 844.3K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 * 
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 * 
 * 你可以认为每种硬币的数量是无限的。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3 
 * 解释：11 = 5 + 5 + 1
 * 
 * 示例 2：
 * 
 * 
 * 输入：coins = [2], amount = 3
 * 输出：-1
 * 
 * 示例 3：
 * 
 * 
 * 输入：coins = [1], amount = 0
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 * 
 * 
 */

// @lc code=start
class Solution {

    func coinChange(_ coins: [Int], _ amount: Int) -> Int {
        return coinChangeCore(coins, amount)
    }


    // 自顶向下，暴力遍历，未优化（超时）

    /*
    func coinChangeCore(_ coins: [Int], _ amount: Int) -> Int {
        if amount == 0 {
            return 0
        }
        if amount < 0 {
            return -1
        }

        var retVal = Int.max

        for coin in coins {
            var subRet = coinChangeCore(coins, amount - coin)
            if subRet == -1 {
                continue
            }

            retVal = retVal < (subRet + 1) ? retVal : (subRet + 1)
        }

        return retVal != Int.max ? retVal : -1
    }
    */


    // 自顶向下，进行备忘录缓存剪枝

    /*
    private var dpTable = [Int : Int]()

    func coinChangeCore(_ coins: [Int], _ amount: Int) -> Int {
        if amount == 0 {
            return 0
        }
        if amount < 0 {
            return -1
        }

        var retVal = Int.max

        for coin in coins {
            var subRet = 0
            if dpTable.keys.contains(amount - coin) {
                subRet = dpTable[amount - coin]!
            } else {
                subRet = coinChangeCore(coins, amount - coin)
                dpTable[amount - coin] = subRet
            }

            if subRet == -1 {
                continue
            }

            retVal = retVal < (subRet + 1) ? retVal : (subRet + 1)
        }

        return retVal == Int.max ? -1 : retVal
    }
    */


    // 迭代解法，时间复杂度为 O(n)

    func coinChangeCore(_ coins: [Int], _ amount: Int) -> Int {
        if amount == 0 {
            return 0
        }
        if amount < 0 {
            return -1
        }

        var dpTable = [Int : Int]()
        dpTable[0] = 0

        for optionAmount in 0...amount {
            for coin in coins {
                if optionAmount - coin < 0 {
                    continue
                }

                if !dpTable.keys.contains(optionAmount) {
                    dpTable[optionAmount] = amount + 1
                }

                dpTable[optionAmount] = (
                    dpTable[optionAmount]! < (1 + dpTable[optionAmount - coin]!) ?
                    dpTable[optionAmount]! : (1 + dpTable[optionAmount - coin]!)
                )
            }
        }

        return dpTable[amount]! == amount + 1 ? -1 : dpTable[amount]!
    }

}
// @lc code=end

