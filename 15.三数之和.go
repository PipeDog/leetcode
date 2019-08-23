/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
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
			nums[i] = nums[j]
			for i < j && vals[i] <= key {
				i++
			}
			nums[j] = nums[i]
		}

		nums[i] = key
		sort(vals, left, i - 1)
		sort(vals, i + 1, right)
	}

	sort(nums, 0, len(nums) - 1)

	ret := make([][]int, 0)

	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				effect := []int{nums[i], nums[left], nums[right]}
				ret = append(ret, effect)

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++; right--
			}
		}
	}

	return ret
}



