import java.util.PriorityQueue;

import javax.management.ListenerNotFoundException;

/*
 * @lc app=leetcode.cn id=23 lang=java
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (56.57%)
 * Likes:    1833
 * Dislikes: 0
 * Total Accepted:    420.2K
 * Total Submissions: 741.7K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 * 
 * 
 * 示例 2：
 * 
 * 输入：lists = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 输入：lists = [[]]
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {

    // 两种解法
    // 1、多次合并两个有序链表
    // 2、优先队列（二叉堆）

    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        }
        if (lists.length == 1) {
            return lists[0];
        }

        ListNode resultList = new ListNode(0);
        ListNode lPtr = resultList;
    
        PriorityQueue<ListNode> queue = new PriorityQueue<>(
            lists.length, 
            (a, b) -> (a.val - b.val)
        );

        for (ListNode list : lists) {
            if (list != null) {
                queue.offer(list);
            }
        }

        while (!queue.isEmpty()) {
            ListNode list = queue.poll();
            if (list.next != null) {
                queue.offer(list.next);
            }

            lPtr.next = list;
            lPtr = lPtr.next;
        }

        return resultList.next;
    }

    /*
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        }
        if (lists.length == 1) {
            return lists[0];
        }

        ListNode list1 = lists[0];

        for (int i = 1; i < lists.length; i++) {
            ListNode list2 = lists[i];
            list1 = merge2Lists(list1, list2);
        }

        return list1;
    }

    private ListNode merge2Lists(ListNode list1, ListNode list2) {
        if (list1 == null) {
            return list2;
        }
        if (list2 == null) {
            return list1;
        }

        ListNode resultList = new ListNode(0);
        ListNode l1 = list1, l2 = list2, l3 = resultList;

        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                l3.next = l1;
                l1 = l1.next;
            } else {
                l3.next = l2;
                l2 = l2.next;
            }
            l3 = l3.next;
        }

        if (l1 != null) {
            l3.next = l1;
        }  
        if (l2 != null) {
            l3.next = l2;
        }

        return resultList.next;
    }
    */

}
// @lc code=end

