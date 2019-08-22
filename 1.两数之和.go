/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	numsLen := len(nums)
	
	if numsLen == 0 {
		return ret
	}

	indexs := make(map[int]int)

	for i := 0; i < numsLen; i++ {
		indexs[nums[i]] = i
	}

	for i := 0; i < numsLen; i ++ {
		num1 := nums[i]
		num2 := target - num1

		if _, ok := indexs[num2]; ok {
			index1 := i
			index2 := indexs[num2]

			if index1 == index2 {
				continue
			}

			ret[0] = index1
			ret[1] = index2
			return ret
		}
	}

	return ret
}

