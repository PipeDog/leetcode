/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	res := head
	front := head
	back := head

	alreadyScan := 0
	loc := 0

	for front.Next != nil {
		if alreadyScan >= n {
			back = back.Next
			loc++
		}
		alreadyScan++
		front = front.Next
	}

	if loc == 0 && alreadyScan - loc != n {
		res = res.Next
	} else {
		back.Next = back.Next.Next
	}

	return res
}

