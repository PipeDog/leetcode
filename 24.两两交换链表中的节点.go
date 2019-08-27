/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := head.Next
	
	curNode := head
	var lastNode *ListNode

	for ; curNode.Next != nil; {
		ptr1 := curNode
		ptr2 := ptr1.Next

		ptr1.Next = ptr1.Next.Next
		ptr2.Next = ptr1

		if lastNode != nil {
			lastNode.Next = ptr2
		}
		lastNode = curNode
		curNode = curNode.Next

		if curNode == nil {
			break
		}
	}
	return result
}

