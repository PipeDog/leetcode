/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	letters := map[string]string {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	ret := make([]string, 0)

	if len(digits) == 0 {
		return ret
	}

	var printLetters func(string, string, int)

	printLetters = func(combine string, digits string, loc int) {
		if loc == len(digits) {
			ret = append(ret, combine)
			return
		}

		letterStr := letters[digits[loc: loc+1]]

		for i := 0; i < len(letterStr); i++ {
			combine += letterStr[i: i+1]
			printLetters(combine, digits, loc+1)
			combine = combine[: len(combine)-1]
		}
	}

	// first letter
	letterStr := letters[digits[0: 1]]

	for i := 0; i < len(letterStr); i++ {
		combine := letterStr[i: i+1]
		printLetters(combine, digits, 1)
	}

	return ret
}

