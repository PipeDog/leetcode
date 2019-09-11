/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	nums := map[string]int{
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

	strs := map[int]string {
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}

	// 求个位数与多位数的乘积，如123 * 9
	// num 完整的数
	// unit 个位数
	// suffix 后缀需要拼接的0的个数
	calProduct := func(numStr string, digitsStr string, suffix int) string {
		fmt.Printf("suffix=%d,", suffix)

		suffixStr := ""
		for i := 0; i < suffix; i++ {
			suffixStr += "0"
		}

		ret := ""
		unit := nums[digitsStr]
		carry := 0

		for i := len(numStr) - 1; i >= 0; i-- {
			curPro := nums[numStr[i: i+1]] * unit
			curPro += carry

			digits := curPro % 10 
			ret = strs[digits] + ret

			carry = curPro / 10
		}

		if carry > 0 {
			ret = strs[carry] + ret
		}
		return ret + suffixStr
	}

	// 求两个字符串的和
	calSum := func(str1, str2 string) string {
		if len(str1) < len(str2) {
			diff := len(str2) - len(str1)
			for i := 0; i < diff; i++ {
				str1 = "0" + str1
			}
		} else if len(str1) > len(str2) {
			diff := len(str1) - len(str2)
			for i := 0; i < diff; i++ {
				str2 = "0" + str2
			}
		}

		str := ""
		carry := 0

		for i := len(str1) - 1; i >= 0; i-- {
			val1 := nums[str1[i: i+1]]
			val2 := nums[str2[i: i+1]]

			sum := val1 + val2
			sum += carry

			str = strs[sum % 10] + str
			carry = sum / 10
		}

		if carry > 0 {
			str = strs[carry] + str
		}
		return str
	}
	

	products := []string{}
	num2Len := len(num2)

	for i := num2Len - 1; i >= 0; i-- {
		product := calProduct(num1, num2[i: i+1], num2Len - 1 - i)
		products = append(products, product)
	}

	sumStr := products[0]

	for i := 1; i < len(products); i++ {
		sumStr = calSum(sumStr, products[i])
	}

	return sumStr
}

