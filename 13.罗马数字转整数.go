/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
func romanToInt(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		// I、X、C
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	if _, ok := roman[s]; ok {
		return roman[s]
	}

	ret := 0

	for i := 0; i < len(s); i++ {
		if i == len(s) - 1 {
			ret += roman[s[i: i + 1]]
			break
		}

		switch s[i] {
		case 'I':
			fallthrough
		case 'X':
			fallthrough
		case 'C':
			// 如果是'I'、'X'、'C'三种字符，则先去查看当前位与下一位，如果有匹配值，则向前进2位，否则，正常进一位
			if val, ok := roman[s[i: i + 2]]; ok {
				ret += val
				i++
			} else {
				ret += roman[s[i: i + 1]]
			}
		default:
			ret += roman[s[i: i + 1]]
		}
	}

	return ret
}

