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

	minValue := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	getArea := func(vals []int, left, right int) int {
		high := minValue(vals[left], vals[right])
		return high * (right - left)
	}

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		curArea := getArea(height, left, right)
		if curArea > maxArea {
			maxArea = curArea
		}

		// 将高度较小的一端向另一端移动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

