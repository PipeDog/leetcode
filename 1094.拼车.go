/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */
func carPooling(trips [][]int, capacity int) bool {
	nums := [1001]int{0}

	for _, trip := range trips {
		nums[trip[1]] += trip[0]
		nums[trip[2]] -= trip[0]
	}

	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += nums[i]
		if sum > capacity {
			return false
		}
	}
	return true
}

