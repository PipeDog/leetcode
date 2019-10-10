/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// 解决思路：https://mp.weixin.qq.com/s/liL1ewy9QQ4EwuxWNtl-NA
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	left, right := 0, len(height) - 1
	leftMax, rightMax := height[left], height[right]

	area := 0

	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if leftMax <= rightMax {
			area += leftMax - height[left]
			left++
		} else {
			area += rightMax - height[right]
			right--
		}
	}

	return area
}

