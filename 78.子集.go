/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
func subsets(nums []int) [][]int {
	sets := [][]int{}
	sets = append(sets, []int{})

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]

		setsLen := len(sets)
		for j := 0; j < setsLen; j++ {
			set := []int{num}
			set = append(set, sets[j]...)
			sets = append(sets, set)
		}
	}

	return sets
}

