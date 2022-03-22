/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (36.29%)
 * Likes:    4909
 * Dislikes: 0
 * Total Accepted:    940K
 * Total Submissions: 2.6M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "cbbd"
 * 输出："bb"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 * 
 * 
 */

// @lc code=start
class Solution {

    public String longestPalindrome(String s) {
        String result = "";
        for (int i = 0; i < s.length(); i++) {
            String s1 = longestPalindrome(s, i, i);
            String s2 = longestPalindrome(s, i, i + 1);
            System.out.println("s1 = " + s1 + ", s2 = " + s2 + ",");
            result = result.length() >= s1.length() ? result : s1;
            result = result.length() >= s2.length() ? result : s2;            
        }
        return result;
    }

    public String longestPalindrome(String s, int leftCenter, int rightCenter) {
        while (
            leftCenter >= 0 && 
            rightCenter < s.length() &&
            s.charAt(leftCenter) == s.charAt(rightCenter)) {
            leftCenter--;
            rightCenter++;
        }

        // 恢复最后一次遍历的位置
        leftCenter++;
        // "hamburger".substring(4, 8) returns "urge"
        return s.substring(leftCenter, rightCenter);
    }

    /*
    // 其实应该用马拉车，但是不会写...
    // 这里是用了马拉车算法进行补码的方式解决了单、双数的问题

    private static final char PLACEHOLDER_CHAR = '_';

    public String longestPalindrome(String s) {
        if (s == null || s.length() <= 1) {
            return s;
        }

        // 插入补位码，如原始字符串为 "ABAC"，
        // 插入补位码后为 "_A_B_A_C_"
        int len = s.length() * 2 + 1;
        char[] charArray = new char[len];
        char[] sChars = s.toCharArray();

        for (int i = 0; i < s.length(); i++) {
            charArray[i * 2] = PLACEHOLDER_CHAR;
            charArray[i * 2 + 1] = sChars[i];            
        }

        charArray[len - 1] = PLACEHOLDER_CHAR;

        // 求最长回文串
        String maxLenString = "";
        for (int i = 1; i < len - 1; i += 1) {
            String currentString = longestPalindromeAtCenter(charArray, i, s);
            if (maxLenString.length() < currentString.length()) {
                maxLenString = currentString;
            }
        }

        return maxLenString;
    }

    private String longestPalindromeAtCenter(char[] charArray, int center, String s) {
        int left = center - 1, right = center + 1;

        while (left >= 0 && 
            right < charArray.length && 
            charArray[left] == charArray[right]) {
            left--;
            right++;
        }

        // 恢复最后一次符合条件的位置
        left += 1;
        right -= 1;

        if (charArray[left] == PLACEHOLDER_CHAR) {
            left = left / 2;
            right = right / 2;
        } else {
            left = left / 2 + 1;
            right = right / 2 + 1;
        }

        return s.substring(left, right);
    }
    */

}
// @lc code=end

