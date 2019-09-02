/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		for stack[len(stack)-1].Left != nil {
			left := stack[len(stack)-1].Left
			stack = append(stack, left)
		}

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			ret = append(ret, top.Val)
			stack = stack[:len(stack)-1]

			if top.Right != nil {
				stack = append(stack, top.Right)
				break
			}
		}
	}

	return ret
}

