/*
 * @lc app=leetcode.cn id=92 lang=java
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.24%)
 * Likes:    1220
 * Dislikes: 0
 * Total Accepted:    274.5K
 * Total Submissions: 496.7K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点数目为 n
 * 1 
 * -500 
 * 1 
 * 
 * 
 * 
 * 
 * 进阶： 你可以使用一趟扫描完成反转吗？
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

    public ListNode reverseBetween(ListNode head, int left, int right) {
        // 因为 left 是从 1 开始，所以当 left 为 1 时这道题就变为了反转链表的前 n 个节点
        if (left == 1) {
            return reverseN(head, right - left + 1/* 结果仍为 right */);
        }

        // 因为是把下一个节点开始当成一个新的链表进行处理，所以这里需要重新拼接修改后的结果，组成新的链表结构
        head.next = reverseBetween(head.next, left - 1, right - 1);
        return head;
    }

    // input: 1 -> 2 -> 3 -> 4, n = 2
    // ouput: 2 -> 1 -> 3 -> 4
    // 则 3 为 successor 节点
    private ListNode successor = null;

    // 反转链表的前 n 个节点
    private ListNode reverseN(ListNode head, int n) {
        if (n == 1) {
            successor = head.next;
            return head;
        }

        ListNode last = reverseN(head.next, n - 1);
        head.next.next = head;
        head.next = successor;
        return last;
    }

}
// @lc code=end
