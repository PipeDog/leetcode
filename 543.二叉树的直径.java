import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=543 lang=java
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (56.33%)
 * Likes:    954
 * Dislikes: 0
 * Total Accepted:    193.4K
 * Total Submissions: 342.7K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 * 
 * 
 * 
 * 示例 :
 * 给定二叉树
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \     
 * ⁠     4   5    
 * 
 * 
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 * 
 * 
 * 
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {

    // 每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和

    private int mMaxDiameter = 0;

    public int diameterOfBinaryTree(TreeNode root) {
        diameterOfBinaryTreeCore(root);
        return mMaxDiameter;
    }

    private void diameterOfBinaryTreeCore(TreeNode root) {
        if (root == null) {
            return;
        }

        int leftDepth = maxDepth(root.left);
        int rightDepth = maxDepth(root.right);
        int diameter = leftDepth + rightDepth;

        if (diameter > mMaxDiameter) {
            mMaxDiameter = diameter;
        }

        diameterOfBinaryTreeCore(root.left);
        diameterOfBinaryTreeCore(root.right);
    }

    private int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int leftDepth = maxDepth(root.left);
        int rightDepth = maxDepth(root.right);
        return 1 + Math.max(leftDepth, rightDepth);
    }

}
// @lc code=end

