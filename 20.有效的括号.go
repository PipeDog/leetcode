/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s) % 2 == 1 {
		return false
	}

    brackets := map[string]string {
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := ""
	
	for i := 0; i < len(s); i++ {
		ch := s[i: i+1]
		top := ""
		if i > 0 && len(stack) > 0 {
			top = stack[len(stack)-1:]
		}

		if i == 0 {
			// push stask
			stack += ch
		} else if len(s) - 1 == i {
			if brackets[top] == ch {
				return true
			} else {
				return false
			}
		} else {
			if ch == "(" || ch == "{" || ch == "[" {
				// push stask
				stack += ch
			} else {
				if brackets[top] == ch {
					// pop stack
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return (stack == "")
}

