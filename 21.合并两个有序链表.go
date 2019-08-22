/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	ret := &ListNode{}
	var curNode *ListNode

	for {
		if l1 == nil && l2 == nil {
			break
		}

		var tmpNode *ListNode

		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				tmpNode = l1
				l1 = l1.Next
			} else {
				tmpNode = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			tmpNode = l1
			l1 = l1.Next
		} else {
			tmpNode = l2
			l2 = l2.Next
		}

		if curNode == nil {
			curNode = tmpNode
			ret.Next = curNode
		} else {
			curNode.Next = tmpNode
			curNode = curNode.Next
		}
	}

	return ret.Next
}

