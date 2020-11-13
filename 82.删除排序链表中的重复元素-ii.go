/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (43.13%)
 * Likes:    162
 * Dislikes: 0
 * Total Accepted:    22.9K
 * Total Submissions: 51.7K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rear := head
	prev := head//.Next

	var needle *ListNode
	var ret *ListNode

	for prev != nil {
		if prev.Val == rear.Val && prev.Val != prev.Next.Val {
			if needle == nil {
				needle = prev.Next
				ret = needle
			} else {
				needle.Next = prev.Next
				needle = needle.Next
			}
			rear = prev.Next
			prev = prev.Next
			fmt.Printf("{n=%d,p=%d,r=%d},", needle.Val, prev.Val, rear.Val)
		} else {
			prev = prev.Next
			fmt.Printf("[n=%d,p=%d,r=%d],", needle.Val, prev.Val, rear.Val)
		}

	}

	return ret
}
// @lc code=end

