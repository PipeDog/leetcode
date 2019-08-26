/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// 马拉车算法 https://zhuanlan.zhihu.com/p/70532099

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 中心扩展法
	// 获取当前位置的回文子串长度
	curLocLen := func(str string, left, right int) int {
		for {
			if left >= 0 && right < len(str) && str[left] == str[right] {
				left--
				right++
			} else {
				break
			}
		}

		return right - left - 1
	}

	beginLoc := 0
	endLoc := 0

	for i := 0; i < len(s); i++ {
		// 以字符为中心的回文子串长度，如"abac"，lenOnNum表示的是以'a' 'b' 'c' 'd' 为中心的回文子串长度
		lenOnChar := curLocLen(s, i, i)
		// 以字符间隔为中心的回文子串长度，如"abac"，lenBetweenChars表示的是以'a' 'b'之间为中心的回文子串长度
		lenBetweenChars := curLocLen(s, i, i + 1)

		len := lenOnChar
		if lenBetweenChars > len {
			len = lenBetweenChars
		}

		if len > endLoc - beginLoc {
			beginLoc = i - (len - 1) / 2
			endLoc = i + len / 2
		}
	}

	return s[beginLoc: endLoc + 1]
}

