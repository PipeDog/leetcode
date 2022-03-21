/*
 * @lc app=leetcode.cn id=142 lang=java
 *
 * [142] 环形链表 II
 *
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (55.76%)
 * Likes:    1467
 * Dislikes: 0
 * Total Accepted:    396.7K
 * Total Submissions: 710.9K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 * 
 * 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos
 * 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos
 * 不作为参数进行传递，仅仅是为了标识链表的实际情况。
 * 
 * 不允许修改 链表。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：返回索引为 1 的链表节点
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：head = [1,2], pos = 0
 * 输出：返回索引为 0 的链表节点
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：head = [1], pos = -1
 * 输出：返回 null
 * 解释：链表中没有环。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目范围在范围 [0, 10^4] 内
 * -10^5 <= Node.val <= 10^5
 * pos 的值为 -1 或者链表中的一个有效索引
 * 
 * 
 * 
 * 
 * 进阶：你是否可以使用 O(1) 空间解决此题？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {

    // https://labuladong.github.io/algo/1/8/

    // fast -> 2k 步
    // slow -> k 步
    // 如果存在环，则两指针相遇一定在环上，即 k (fast - slow) 一定是环长度的整数倍
    // 可以先判断是否会相遇，如果相遇，让 fast 指针从头开始以和 slow 指针相同的速度向前
    // 再次相遇的位置就是环开始的位置

    public ListNode detectCycle(ListNode head) {
        if (head == null) {
            return null;
        }
        if (head.next == null) {
            return null;
        }

        ListNode fast = head, slow = head;

        // 确定是否有环
        while (fast != null && slow != null) {
            fast = fast.next;
            // 无环
            if (fast == null) {
                return null;
            }

            if (fast.next != null) {
                fast = fast.next;
            } else {
                // 无环
                return null;
            }

            slow = slow.next;

            if (fast == slow) {
                break;
            }
        }

        // 这里可以确定有环
        fast = head;

        while (fast != slow) {
            fast = fast.next;
            slow = slow.next;
        }
        return fast;
    }

}
// @lc code=end

