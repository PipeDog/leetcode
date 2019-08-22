/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// https://leetcode-cn.com/problems/longest-common-prefix/solution/
func longestCommonPrefix(strs []string) string {
    if strs == nil || len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i: i+1]

		// j从1开始，以strs[0]为比较对象，所以不需要对strs[0]再次遍历
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || ch != strs[j][i: i+1] {
				return strs[0][: i]
			}
		}
	}

	return strs[0]
}

