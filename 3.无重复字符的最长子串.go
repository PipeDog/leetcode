/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	locs := make(map[string]int)
	curLen := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		ch := s[i:i+1]
		preLoc, ok := locs[ch]

		if !ok || i - preLoc > curLen {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = i - preLoc
		}

		locs[ch] = i
	}
	if curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

