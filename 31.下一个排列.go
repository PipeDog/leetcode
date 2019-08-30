/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
func nextPermutation(nums []int)  {
    if len(nums) <= 1 {
		return
	}
	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	// [1,2,3,4,5]
	// [1,2,3,5,4]

	left, right := 1, len(nums) - 1
	// []
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

