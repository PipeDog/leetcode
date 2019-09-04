/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	needCarry := false
	for i := len(digits) - 1; i >= 0; i-- {
		num := 1
		if needCarry && i == len(digits) - 1 {
			num++
		}

		curNum := digits[i] + num
		if curNum < 10 {
			digits[i] = curNum
			needCarry = false
			break
		} else {
			digits[i] = curNum % 10
			needCarry = true
		}
	}

	if needCarry {
		ret := append([]int{1}, digits[:]...)
		return ret
	}
	return digits
}

