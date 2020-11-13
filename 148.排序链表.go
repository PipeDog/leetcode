/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (61.68%)
 * Likes:    315
 * Dislikes: 0
 * Total Accepted:    28.8K
 * Total Submissions: 46.1K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 * 
 * 示例 1:
 * 
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 * 
 * 
 * 示例 2:
 * 
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head
	
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	
	mid := slow.Next

	left := sortList(head)
	right := sortList(mid)

	tmp := &ListNode{0, nil}

	ret := tmp
	
	for left != nil && right != nil {
		if left.Val < right.Val {
			tmp.Next = left
			left = left.Next
		} else {
			tmp.Next = right
			right = right.Next
		}

		tmp = tmp.Next
	}

	if left != nil {
		tmp.Next = left
	}
	if right != nil {
		tmp.Next = right
	}

	return ret.Next
}
// @lc code=end

