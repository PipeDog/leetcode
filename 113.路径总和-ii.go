/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var paths [][]int

	var matchPath func(*TreeNode, int, []int)

	matchPath = func(node *TreeNode, curSum int, path []int) {
		if node == nil {
			return
		}
		if path == nil {
			path = []int{}
		}

		curSum += node.Val
		path = append(path, node.Val)

		if curSum == sum {
			if node.Left == nil && node.Right == nil {
				nums := make([]int, len(path))
				copy(nums, path)

				paths = append(paths, nums)
				return
			}
		}

		if node.Left != nil {
			matchPath(node.Left, curSum, path)
		}
		if node.Right != nil {
			matchPath(node.Right, curSum, path)
		}
	}

	matchPath(root, 0, nil)
	return paths
}

