/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {

	/*
    if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
	*/

	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	curLevelNodes := 1
	nextLevelNodes := 0
	depth := 0

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		curLevelNodes--

		if head.Left != nil {
			queue = append(queue, head.Left)
			nextLevelNodes++
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
			nextLevelNodes++
		}

		if curLevelNodes == 0 {
			curLevelNodes = nextLevelNodes
			nextLevelNodes = 0

			depth++
		}
	}

	return depth
}

