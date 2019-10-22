/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// https://leetcode-cn.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	/*
	var isValidString func(string) bool

	isValidString = func(str string) bool {
		if len(str) % 2 != 0 {
			return false
		}

		stack := ""
		strLen := len(str)

		for i := 0; i < strLen; i++ {
			curChar := str[i: i+1]

			// push stack
			if curChar == "(" {
				stack += curChar
				continue
			}

			// pop stack
			if curChar == ")" {
				if len(stack) == 0 {
					return false
				}

				if stack[len(stack)-1:] == "(" {
					stack = stack[: len(stack)-1]
				} else {
					return false
				}
				continue
			}

			return false
		}

		return (len(stack) == 0)
	}

	ret := make([]string, 0)
	filter := make(map[string]bool)

	// 当前的str
	// 总长度
	// 当前位置
	var combine func(string, int, int)

	combine = func(str string, length, loc int) {
		if loc >= length {
			if isValidString(str) {
				if _, ok := filter[str]; !ok {
					filter[str] = true
					ret = append(ret, str)
				}
			}
			return
		}

		for i := 0; i < 2; i++ {
			if i == 0 {
				str += "("
			} else {
				str += ")"
			}
			combine(str, length, loc + 1)
			str = str[: len(str)-1]
		}
	}

	combine("(", n * 2, 1)
	return ret
	*/

	ret := make([]string, 0)
	var generate func(str string, n, left, right int)

	generate = func(str string, n, left, right int) {
		if right == n {
			ret = append(ret, str)
			return
		}

		if left < n {
			generate(str + "(", n, left + 1, right)
		}
		if right < left {
			generate(str + ")", n, left, right + 1)
		}
	}

	generate("", n, 0, 0)
	return ret
}

