/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */
func convert(s string, numRows int) string {
	if len(s) == 0 {
		return s
	}
	if numRows == 1 {
		return s
	}

	strs := make([]string, numRows)
	curRow := 0
	goDown := true

	for i := 0; i < len(s); i++ {
		curChar := s[i: i+1]
		strs[curRow] = strs[curRow] + curChar

		if goDown {
			curRow++
		} else {
			curRow--
		}

		if curRow == 0 || curRow == numRows - 1 {
			goDown = !goDown
		}
	}

	ret := ""

	for i := 0; i < numRows; i++ {
		ret = ret + strs[i]
	}

	return ret
}

