/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[left] >= target {
		return left
	}
	return left + 1
}

