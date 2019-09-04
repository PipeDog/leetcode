/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	var handler func(low, high int, nums []int) *TreeNode

	handler = func(low, high int, nums []int) *TreeNode {
		if low > high {
			return nil
		}

		mid := (low + high) / 2

		node := &TreeNode{}
		node.Val = nums[mid]
		node.Left = handler(low, mid - 1, nums)
		node.Right = handler(mid + 1, high, nums)

		return node
	}

	return handler(0, len(nums) - 1, nums)
}

