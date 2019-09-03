/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	wordLen := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			if wordLen == 0 {
				continue
			} else {
				break
			}
		} else {
			wordLen++
		}
	}

	return wordLen
}

