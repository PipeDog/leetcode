/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */
func findSubstring(s string, words []string) []int {
	if len(s) == 0 {
		return -1
	}
	if len(words) == 0 {
		return 0
	}

	wordLen := len(words[0])
	wordsCount := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if _, ok := wordsCount[words[i]]; ok {
			count := wordsCount[words[i]]
			count++
			wordsCount[words[i]] = count
		} else {
			wordsCount[words[i]] = 1
		}
	}

	
	for i := 0; i < strLen; {
		word := words[i]

		if _, ok := wordsCount[word]; ok {
			i += wordLen

		} else {
			i++
		}
	}
}

