/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	/*
    if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	stack := []*TreeNode{}
	retainStack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		retainStack = append(retainStack, node)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	for i := 0; i < len(retainStack) ; i++ {
		if i == len(retainStack) - 1 {
			cur := retainStack[i]
			cur.Left = nil
			cur.Right = nil
		} else {
			cur := retainStack[i]
			next := retainStack[i+1]

			cur.Left = nil
			cur.Right = next
		}
	}
	*/

	var handler func(*TreeNode)
	var next *TreeNode

	handler = func(node *TreeNode) {
		if node == nil {
			return
		}

		handler(node.Right)
		handler(node.Left)

		node.Left = nil
		node.Right = next
		next = node
	}

	handler(root)
}

