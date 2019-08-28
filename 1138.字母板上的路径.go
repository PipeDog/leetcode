/*
 * @lc app=leetcode.cn id=1138 lang=golang
 *
 * [1138] 字母板上的路径
 */

// https://leetcode-cn.com/problems/alphabet-board-path/solution/100java-100-by-cznczai/
func alphabetBoardPath(target string) string {
	if len(target) == 0 {
		return ""
	}

	ret := make([]byte, 0)
	bytes := []byte(target)

	length := len(bytes)

	vLast := 0
	hLast := 0

	for loc := 0; loc < length; loc++ {
		curNum := int(bytes[loc] - 'a')

		vCur := curNum / 5
		hCur := curNum % 5

		// 左上下右
		if hCur < hLast {
			for i := 0; i < hLast - hCur; i++ {
				ret = append(ret, 'L')
			}
		}
		if vCur < vLast {
			for i := 0; i < vLast - vCur; i++ {
				ret = append(ret, 'U')
			}
		}
		if vCur > vLast {
			for i := 0; i < vCur - vLast; i++ {
				ret = append(ret, 'D')
			}
		}
		if hCur > hLast {
			for i := 0; i < hCur - hLast; i++ {
				ret = append(ret, 'R')
			}
		}

		ret = append(ret, '!')

		vLast = vCur
		hLast = hCur
	}
	
	return string(ret)
}

