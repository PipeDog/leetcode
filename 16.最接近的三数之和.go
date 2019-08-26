/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// https://leetcode-cn.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	diffAbs := func(x, y int) int {
		ret := max(x, y) - min(x, y)
		if ret >= 0 {
			return ret
		}
		return -ret
	} 

	var sort func([]int, int, int)

	sort = func(vals []int, left, right int) {
		if left >= right {
			return
		}

		i := left
		j := right
		key := vals[i]

		for i < j {
			for i < j && vals[j] >= key {
				j--
			}
			vals[i] = vals[j]
			for i < j && vals[i] <= key {
				i++
			}
			vals[j] = vals[i]
		}

		vals[i] = key

		sort(vals, left, i - 1)
		sort(vals, i + 1, right)
	}

	sort(nums, 0, len(nums) - 1)

	sum := 0

	for i := 0; i < len(nums) - 2; i++ {
		left := i + 1
		right := len(nums) - 1

		if i == 0 {
			sum = nums[i] + nums[left] + nums[right]
		}

		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			if curSum == target {
				return curSum
			}

			need := target - nums[i]

			if diffAbs(curSum, target) < diffAbs(sum, target) {
				sum = curSum
			}

			if nums[left] + nums[right] < need {
				left++
			} else {
				right--
			}
		}
	}

	return sum
}

