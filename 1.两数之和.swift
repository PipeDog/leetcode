/*
 * @lc app=leetcode.cn id=1 lang=swift
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        if nums == nil {
            return [Int]()
        }

        var recordMap : [Int : Int] = [Int : Int]()
        for (index, value) in nums.enumerated() {
            let missValue = target - value
            if recordMap.keys.contains(missValue) {
                return [index, recordMap[missValue]!]
            }
            recordMap[value] = index
        }

        return [Int]()
    }
}
// @lc code=end

