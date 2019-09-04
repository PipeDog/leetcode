/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	var check func(node *TreeNode) (int, bool)

	check = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		left, ok := check(node.Left)
		if !ok {
			return 0, false
		}

		right, ok := check(node.Right)
		if !ok {
			return 0, false
		}

		if diff := abs(left - right); diff > 1 {
			return 0, false
		}

		return max(left, right) + 1, true
	}

	_, ok := check(root)
	return ok
}

