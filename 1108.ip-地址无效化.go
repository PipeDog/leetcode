/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

 // https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	loc := 0

	for loc < len(address) {
		c := address[loc: loc+1]
		if c == "." {
			left := address[: loc]
			right := address[loc+1:]

			address = left + "[.]" + right
			loc += 3
		} else {
			loc++
		}
	}

	return address
}

