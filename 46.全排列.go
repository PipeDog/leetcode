/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
func permute(nums []int) [][]int {
	ret := [][]int{}
	paths := []int{}
	visited := make([]bool, len(nums))

	deepCopy := func(items []int) []int {
		copyItems := make([]int, len(items))
		copy(copyItems, items)
		return copyItems
	}

	var permuteCore func(nums, paths []int, visited []bool, index int)

	permuteCore = func(nums, paths []int, visited []bool, index int) {
		if index >= len(nums) {
			ret = append(ret, deepCopy(paths))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true
				paths = append(paths, nums[i])
				permuteCore(nums, paths, visited, index + 1)
				paths = paths[: len(paths)-1]
				visited[i] = false
			}
		}
	}

	permuteCore(nums, paths, visited, 0)
	return ret
}

