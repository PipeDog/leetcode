/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	ret := make([][]int, 0)
	
	curLevelCount := 1
	nextLevelCount := 0

	queue = append(queue, root)
	var curLevelVals []int
	curLevelVals = make([]int, 0)

	for len(queue) > 0 {
		pop := queue[0]
		queue = append(queue[1:])
		curLevelCount--

		curLevelVals = append(curLevelVals, pop.Val)
		
		if pop.Left != nil {
			queue = append(queue, pop.Left)
			nextLevelCount++
		}
		if pop.Right != nil {
			queue = append(queue, pop.Right)
			nextLevelCount++
		}

		if curLevelCount == 0 {
			curLevelCount = nextLevelCount
			nextLevelCount = 0

			ret = append(ret, curLevelVals[:])
			curLevelVals = make([]int, 0)
		}
	}

	return ret
}

