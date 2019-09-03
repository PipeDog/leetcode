/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}

	ret := []int{}
	stack := []*TreeNode{root}
	var last *TreeNode

	for len(stack) > 0 {
		for stack[len(stack)-1].Left != nil {
			stack = append(stack, stack[len(stack)-1].Left)
		}

		for len(stack) > 0 {
			if stack[len(stack)-1].Right == nil ||
				stack[len(stack)-1].Right == last {
				last = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ret = append(ret, last.Val)
			} else if stack[len(stack)-1].Right != nil {
				stack = append(stack, stack[len(stack)-1].Right)
				break
			}
		}
	}

	return ret
}

