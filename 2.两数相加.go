/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	curNode := ret
	eachSum := 0

	for l1 != nil || l2 != nil || eachSum > 0 {
		if l1 != nil {
			eachSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			eachSum += l2.Val
			l2 = l2.Next
		}

		curNode.Next = &ListNode{Val: eachSum % 10}
		curNode = curNode.Next
		eachSum /= 10
	}

	return ret.Next
}


