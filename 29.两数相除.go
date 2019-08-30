/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0 // Error
	}
	if dividend == 0 {
		return 0
	}

	handleVal := func(x int) int {
		if x > 2147483647 {
			return 2147483647
		}
		if x < -2147483648 {
			return -2147483648
		}
		return x
	}

	subCount := 0
	if dividend > 0 && divisor > 0 {
		for dividend >= divisor {
			dividend -= divisor
			subCount++
		}
	} else if dividend > 0 && divisor < 0 {
		for dividend + divisor > 0 {
			dividend += divisor
			subCount++
		}
		dividend += divisor
		subCount++
	} else if dividend < 0 && divisor < 0 {
		for dividend - divisor < 0 {
			dividend -= divisor
			subCount++
		}
		dividend -= divisor
		subCount++
	} else if dividend < 0 && divisor > 0 {
		for dividend + divisor < 0 {
			dividend += divisor
			subCount++
		}
		dividend += divisor
		dividend = -dividend
		subCount++
	}

	return handleVal(subCount)
}

