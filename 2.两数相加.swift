/*
 * @lc app=leetcode.cn id=2 lang=swift
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init() { self.val = 0; self.next = nil; }
 *     public init(_ val: Int) { self.val = val; self.next = nil; }
 *     public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
 * }
 */
class Solution {
    func addTwoNumbers(_ l1: ListNode?, _ l2: ListNode?) -> ListNode? {
        if l1 == nil {
            return l2
        }
        if l2 == nil {
            return l1
        }

        var ptr1 = l1, ptr2 = l2
        // 用于保存两个节点值相加的进位值
        var carry = 0
        let prePtr = ListNode(0)
        var currentPtr = prePtr

        while ptr1 != nil || ptr2 != nil {
            let val1 = ptr1 != nil ? ptr1!.val : 0
            let val2 = ptr2 != nil ? ptr2!.val : 0
            let sum = val1 + val2 + carry
            let currentVal = sum % 10
            carry = sum / 10            

            currentPtr.next = ListNode(currentVal)
            if let nextPtr = currentPtr.next {
                currentPtr = nextPtr
            }

            ptr1 = ptr1?.next
            ptr2 = ptr2?.next
        }

        if carry > 0 {
            currentPtr.next = ListNode(carry)
        }

        return prePtr.next
    }
}
// @lc code=end

