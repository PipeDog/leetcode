/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	
	var pathSum func(node *TreeNode, curSum int)

	pathSum = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}
		curSum = curSum * 10 + node.Val

		if node.Left == nil && node.Right == nil {
			sum += curSum
			return
		}

		pathSum(node.Left, curSum)
		pathSum(node.Right, curSum)
	}

	pathSum(root, 0)
	return sum
}

