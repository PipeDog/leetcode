/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
    trimSpace := func(s string) string {
		for i := 0; i < len(s); i++ {
			if s[i: i+1] != " " {
				return s[i:]
			}
		}
		return s
	}

	str = trimSpace(str)
	if len(str) == 0 {
		return 0
	}

	nums := map[string]int {
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	
	if _, ok := nums[str[0:1]]; !ok && str[0:1] != "+" && str[0:1] != "-" {
		return 0
	}

	intMax := 2147483647
	intMin := -2147483648

	ret := 0
	isMinus := (str[0:1] == "-")

	for i := 0; i < len(str); i++ {
		if _, ok := nums[str[i: i+1]]; ok {
			ret = ret * 10 + nums[str[i: i+1]]

			if ret > intMax {
				if isMinus {
					return intMin
				} else {
					return intMax
				}
			}
		} else {
			if i > 0 {
				break
			}
		}
	}

	if isMinus {
		ret = -ret
	}

	return ret
}

