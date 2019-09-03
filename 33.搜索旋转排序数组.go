/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums) - 1

	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] > nums[right] {
			left++; right--
		} else {
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

