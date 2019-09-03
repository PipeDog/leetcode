/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// Only left or right subtree
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}

	// Both left and right subtrees, get min val + 1 (current ndoe)
	if left < right {
		return left + 1
	}
	return right + 1
}

