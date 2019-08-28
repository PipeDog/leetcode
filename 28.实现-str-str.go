/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

/*
	KMP
	BM
	Sunday
 */

// https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	// Sunday
	hLen := len(haystack)
	nLen := len(needle)
	offset := make(map[string]int)

	for i := 0; i < nLen; i++ {
		offset[needle[i: i+1]] = nLen - i
	}

	hIndex := 0
	nIndex := 0

	for hIndex <= hLen - nLen {
		nIndex = 0
		
		for haystack[hIndex+nIndex: hIndex+nIndex+1] == needle[nIndex: nIndex+1] {
			nIndex++
			if nIndex == nLen {
				return hIndex
			}
		}

		if hIndex + nLen >= hLen {
			return -1
		}

		curByte := haystack[hIndex+nLen: hIndex+nLen+1]
		if _, ok := offset[curByte]; ok {
			hIndex += offset[curByte]
		} else {
			hIndex += nLen + 1
		}
	}

	return -1
}

