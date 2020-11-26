/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
func longestMountain(arr []int) int {
	numLen := len(arr)
	if numLen < 3 {
		return 0
	}

	// up or down
	unknown, up, down := 0, 1, 2
	curState := unknown
	mountNum := 0

	// current mount length, max mount length
	curLen, maxLen := 0, 0
	
	for index := 1; index < numLen; index++ {
		preNum := arr[index - 1]
		curNum := arr[index]
		oldState := curState

		if preNum < curNum { // up
			if oldState == down {
				if curLen > maxLen {
					maxLen = curLen
				}	
				curLen = 2
			} else {
				if index == 1 {
					curLen = 1
				}
				curLen += 1
			}

			curState = up
		} else if preNum > curNum { // go down
			if oldState == up {
				mountNum += 1		
			}
			curLen += 1
			curState = down
		} else {
			if oldState == down && curLen > maxLen {
				maxLen = curLen
			}
			curState = unknown
			curLen = 1
		}
	}

	if curState == down && curLen > maxLen {
		maxLen = curLen
	}
	if mountNum == 0 {
		return 0
	}
	return maxLen
}
// @lc code=end

