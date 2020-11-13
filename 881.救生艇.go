/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 *
 * https://leetcode-cn.com/problems/boats-to-save-people/description/
 *
 * algorithms
 * Medium (37.49%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    8.4K
 * Total Submissions: 19.6K
 * Testcase Example:  '[1,2]\n3'
 *
 * 第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
 * 
 * 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
 * 
 * 返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：people = [1,2], limit = 3
 * 输出：1
 * 解释：1 艘船载 (1, 2)
 * 
 * 
 * 示例 2：
 * 
 * 输入：people = [3,2,2,1], limit = 3
 * 输出：3
 * 解释：3 艘船分别载 (1, 2), (2) 和 (3)
 * 
 * 
 * 示例 3：
 * 
 * 输入：people = [3,5,3,4], limit = 5
 * 输出：4
 * 解释：4 艘船分别载 (3), (3), (4), (5)
 * 
 * 提示：
 * 
 * 
 * 1 <= people.length <= 50000
 * 1 <= people[i] <= limit <= 30000
 * 
 * 
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {
	if len(people) == 0 {
		return 0
	}
	if len(people) == 1 {
		return 1
	}

	var sort func([]int, int, int)
	sort = func(vals []int, left int, right int) {
		if left >= right {
			return
		}
		i := left
		j := right
		key := vals[i]

		for i < j {
			for i < j && vals[j] >= key {
				j--
			}
			vals[i] = vals[j]
			for i < j && vals[i] <= key {
				i++
			}
			vals[j] = vals[i]
		}

		vals[i] = key
		sort(vals, left, i - 1)
		sort(vals, i + 1, right)
	}

	sort(people, 0, len(people) - 1)

	left, right := 0, len(people) - 1
	numberOfBoats := 0

	for left <= right {
		sum := people[left] + people[right]

		if sum <= limit {
			numberOfBoats++; left++; right--
		} else {
			numberOfBoats++; right--
		}
	}

	return numberOfBoats
}
// @lc code=end

