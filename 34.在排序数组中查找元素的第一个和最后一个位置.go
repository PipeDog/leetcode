/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
func searchRange(nums []int, target int) []int {

	searchMinIndex := func(nums []int, target int) int {
		minIndex := -1
		// 注意right不是len - 1
		left, right := 0, len(nums)
	
		for left < right {
			mid := (left + right) / 2

			if nums[mid] >= target {
				right = mid

				if nums[mid] == target {
					minIndex = mid
				}
			} else {
				left = mid + 1
			}
		}

		return minIndex
	}

	searchMaxIndex := func(nums []int, target int) int {
		maxIndex := -1
		left, right := 0, len(nums)

		for left < right {
			mid := (left + right) / 2

			if nums[mid] <= target {
				left = mid + 1

				if nums[mid] == target {
					maxIndex = mid
				}
			} else {
				right = mid
			}
		}

		return maxIndex
	}

	minIndex := searchMinIndex(nums, target)
	maxIndex := searchMaxIndex(nums, target)
	return []int{minIndex, maxIndex}
}


