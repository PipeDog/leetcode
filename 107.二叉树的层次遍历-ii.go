/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{}
	ret := [][]int{}

	queue = append(queue, root)
	levelNums := []int{}

	curLevelCount := 1
	nextLevelCount := 0

	for len(queue) > 0 {
		node := queue[0]
		levelNums = append(levelNums, node.Val)
		queue = queue[1:]
		curLevelCount-- 

		if node.Left != nil {
			queue = append(queue, node.Left)
			nextLevelCount++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextLevelCount++
		}

		if curLevelCount == 0 {
			curLevelCount = nextLevelCount
			nextLevelCount = 0

			nums := make([]int, len(levelNums))
			copy(nums, levelNums)

			ret = append(ret, nums)

			levelNums = []int{}
		}
	}

	for i := 0; i < len(ret) / 2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}

	return ret
}

