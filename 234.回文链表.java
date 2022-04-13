/*
 * @lc app=leetcode.cn id=234 lang=java
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (51.25%)
 * Likes:    1340
 * Dislikes: 0
 * Total Accepted:    408.6K
 * Total Submissions: 797.2K
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,2,1]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1,2]
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 * 
 * 
 * 
 * 
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {

    // 如 1->2->3->2->1
    // 快慢指针，找到中间位置（2->1），然后翻转后边的链表
    // 遍历比较头节点和中间位置开始的节点
    // 如果都相等则是回文链表

    public boolean isPalindrome(ListNode head) {
        if (head == null || head.next == null) {
            return true;
        }

        // 找到中间位置
        ListNode slow = head, fast = head;
        while (fast != null) {
            if (fast.next != null && fast.next.next != null) {
                fast = fast.next;
                fast = fast.next;
            } else {
                break;
            }
            slow = slow.next;
        }

        // 慢指针向前前进一步
        slow = slow.next;

        // 反转链表后一半
        ListNode reverseHead = reverse(slow);

        // 逐个节点比较
        while (reverseHead != null) {
            if (head.val != reverseHead.val) {
                return false;
            }
            head = head.next;
            reverseHead = reverseHead.next;
        }

        return true;
    }

    // 翻转单链表
    private ListNode reverse(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode prev = null;
        ListNode cur = head, next = head;

        while (cur != null) {
            next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }

        return prev;
    }

    private void logList(String tag, ListNode head) {
        System.out.print(String.format("[%s] ", tag));
        ListNode node = head;
        while (node != null) {
            System.out.print(String.format("%d -> ", node.val));
            node = node.next;
        }
        System.out.print("\n");
    }

}
// @lc code=end
