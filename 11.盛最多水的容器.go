/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	getArea := func(vals []int, left, right int) int {
		minHeight := vals[left]
		if minHeight > vals[right] {
			minHeight = vals[right]
		}

		return (right - left) * minHeight
	}

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		area := getArea(height, left, right)
		if maxArea < area {
			maxArea = area
		}

		// 将高度较小的一端向另一端移动
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

