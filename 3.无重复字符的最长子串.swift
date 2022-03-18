/*
 * @lc app=leetcode.cn id=3 lang=swift
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
class Solution {
    func lengthOfLongestSubstring(_ s: String) -> Int {
        if s == nil {
            return 0
        }
        if s.isEmpty {
            return 0
        }

        let len = s.count
        var indexDict : [Character : Int] = [:]
        var maxLen = 0

        var left = 0, right = 0

        for index in s.indices {
            let currentChar = s[index]

            if let prevIndex = indexDict[currentChar] {
                left = left > prevIndex ? left : prevIndex
            }

            let currentLen = right - left + 1
            maxLen = maxLen > currentLen ? maxLen : currentLen
            // `right + 1` 表示从字符位置后一位才开始不重复
            indexDict[currentChar] = right + 1

            right += 1
        }

        return maxLen
    }
}
// @lc code=end

