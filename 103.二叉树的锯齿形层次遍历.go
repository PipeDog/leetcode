/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ret := [][]int{}
	queue := []*TreeNode{}

	ret = append(ret, []int{})
	queue = append(queue, root)

	curLevelCount := 1
	nextLevelCount := 0
	leftToRight := true

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		curLevelCount--

		nums := ret[len(ret)-1]
		nums = append(nums, head.Val)
		ret[len(ret)-1] = nums

		if head.Left != nil {
			queue = append(queue, head.Left)
			nextLevelCount++
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
			nextLevelCount++
		}

		if curLevelCount == 0 {
			curLevelCount = nextLevelCount
			nextLevelCount = 0

			if leftToRight {
				// Do nothing
			} else {
				// Reverse
				for i := 0; i < len(nums) / 2; i++ {
					nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
				}
				ret[len(ret)-1] = nums
			}

			leftToRight = !leftToRight

			if len(queue) > 0 {
				ret = append(ret, []int{})
			}
		}
	}

	return ret
}

