import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=752 lang=java
 *
 * [752] 打开转盘锁
 *
 * https://leetcode-cn.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (53.04%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    83.7K
 * Total Submissions: 157.8K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 * 
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 * 
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 * 
 * 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：无法旋转到目标数字且不被锁定。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= deadends.length <= 500
 * deadends[i].length == 4
 * target.length == 4
 * target 不在 deadends 之中
 * target 和 deadends[i] 仅由若干位数字组成
 * 
 * 
 */

// @lc code=start
class Solution {

    // 初始值 0000，画一个多叉树（8叉树），其实就是 BFS 遍历

    private Set<String> mDeadendsSet = new HashSet<>();
    private Set<String> mVisistedSet = new HashSet<>();
    private List<String> mResultList = new ArrayList<>();

    public int openLock(String[] deadends, String target) {
        if (target == null) {
            return -1;
        }
        if (target.length() < 4) {
            return -1;
        }
        if (target.equals("0000")) {
            return 0;
        }

        for (int i = 0; i < deadends.length; i++) {
            mDeadendsSet.add(deadends[i]);
        }

        String startPath = "0000";
        return openLockCore(startPath, target);
    }

    private int openLockCore(String path, String target) {
        Queue<String> queue = new LinkedList<>();
        queue.offer(path);
        int step = 0;

        while (!queue.isEmpty()) {
            List<String> list = new ArrayList<>();
            while (!queue.isEmpty()) {
                String temp = queue.poll();

                if (mDeadendsSet.contains(temp)) {
                    continue;
                }
                if (temp.equals(target)) {
                    return step;
                }

                list.add(temp);
            }

            for (String curStep : list) {
                for (int i = 0; i < 4; i++) {
                    String nextStep = plus(curStep, i);
                    if (!mVisistedSet.contains(nextStep)) {
                        mVisistedSet.add(nextStep);
                        queue.offer(nextStep);
                    }
                    
                    nextStep = minus(curStep, i);
                    if (!mVisistedSet.contains(nextStep)) {
                        mVisistedSet.add(nextStep);
                        queue.offer(nextStep);
                    }
                }
            }

            step++;
        }

        return -1;
    }

    private String plus(String path, int index) {
        char[] chars = path.toCharArray();
        if (chars[index] == '9') {
            chars[index] = '0';
        } else {
            chars[index] += 1;
        }
        return new String(chars);
    }

    private String minus(String path, int index) {
        char[] chars = path.toCharArray();
        if (chars[index] == '0') {
            chars[index] = '9';
        } else {
            chars[index] -= 1;
        }
        return new String(chars);
    }

}
// @lc code=end

