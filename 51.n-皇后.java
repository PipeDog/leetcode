import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=51 lang=java
 *
 * [51] N 皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (73.80%)
 * Likes:    1211
 * Dislikes: 0
 * Total Accepted:    189K
 * Total Submissions: 256.1K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 * 
 * 
 * 
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：[["Q"]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 9
 * 
 * 
 * 
 * 
 */

// @lc code=start
class Solution {

    // 满足条件的 N 个棋盘的数据
    private List<List<String>> result = new ArrayList<>();
    // 一个棋盘的数据
    private List<String> backtrack = new ArrayList<>();

    public List<List<String>> solveNQueens(int n) {
        if (n <= 0) {
            return result;
        }

        // 初始化原始数据，所有都为 "."
        for (int i = 0; i < n; i++) {
            StringBuilder builder = new StringBuilder();
            for (int j = 0; j < n; j++) {
                builder.append(".");
            }
            backtrack.add(builder.toString());
        }

        solveNQueensCore(n, 0);
        return result;
    }

    private void solveNQueensCore(int n, int row) {
        if (row == backtrack.size()) {
            List<String> copy = new ArrayList<>(backtrack);
            result.add(copy);
            System.out.println(result.size());
            return;
        }

        for (int col = 0; col < n; col++) {            
            if (!isValid(row, col, n)) {
                continue;
            }

            // 此处需要注意 Java 的字符串和字符之间的转换，char[].toString() 并不是将 char[] 转换成 String 类型
            // 需要使用 new String(char[]) 这一构造方法，因此这里的开销会相比其他语言大一些

            char[] chars = backtrack.get(row).toCharArray();
            chars[col] = 'Q';
            backtrack.set(row, new String(chars));

            solveNQueensCore(n, row + 1);
            chars[col] = '.';
            backtrack.set(row, new String(chars));
        }
    }

    private boolean isValid(int row, int col, int n) {
        // 检查同列数据
        for (int i = 0; i < row; i++) {
            String content = backtrack.get(i);
            if (content.charAt(col) == 'Q') {
                return false;
            }
        }

        // 检查右上角数据
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
            String content = backtrack.get(i);
            if (content.charAt(j) == 'Q') {
                return false;
            }
        }

        // 检查左上角数据
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
            String content = backtrack.get(i);
            if (content.charAt(j) == 'Q') {
                return false;
            }
        }

        return true;
    }

}
// @lc code=end

