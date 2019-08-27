/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mergeCore := func(l1 *ListNode, l2 *ListNode) *ListNode {
		mergeHead := &ListNode{}
		currentNode := mergeHead

		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				currentNode.Next = l1
				l1 = l1.Next
			} else {
				currentNode.Next = l2
				l2 = l2.Next
			}
			currentNode = currentNode.Next
		}

		if l1 != nil {
			currentNode.Next = l1
		}
		if l2 != nil {
			currentNode.Next = l2
		}
		return mergeHead.Next
	}

	var mergeList *ListNode
	var merge func(*ListNode, []*ListNode, int)

	merge = func(list *ListNode, lists []*ListNode, index int) {
		if index >= len(lists) {
			mergeList = list
			return
		}

		newList := mergeCore(list, lists[index])
		merge(newList, lists, index + 1)
	}

	merge(lists[0], lists, 1)
	return mergeList
}

