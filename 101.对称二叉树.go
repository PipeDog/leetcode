/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	var isSymmetricCore func(*TreeNode, *TreeNode) bool

	isSymmetricCore = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else {
			if left.Val != right.Val {
				return false
			}

			return isSymmetricCore(left.Left, right.Right) && isSymmetricCore(left.Right, right.Left)
		}
	}

	return isSymmetricCore(root.Left, root.Right)
}

