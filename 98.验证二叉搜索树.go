/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var validBST func(*bool, *TreeNode, *int, *int)

	validBST = func(res *bool, node *TreeNode, maxVal, minVal *int) {
		if node == nil || *res == false {
			return
		}

		if maxVal != nil && node.Val >= *maxVal {
			*res = false
			return
		}
		if minVal != nil && node.Val <= *minVal {
			*res = false
			return
		}

		v := node.Val
		validBST(res, node.Left, &v, minVal)
		validBST(res, node.Right, maxVal, &v)
	}

	res, v := true, root.Val
	validBST(&res, root.Left, &v, nil)
	validBST(&res, root.Right, nil, &v)
	return res
}

