/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (56.18%)
 * Likes:    324
 * Dislikes: 0
 * Total Accepted:    38.7K
 * Total Submissions: 62.5K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
 * 请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
 * 
 * 
 * 
 * 示例：
 * 
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 * 
 * 返回 13。
 * 
 * 
 * 
 * 
 * 提示：
 * 你可以假设 k 的值永远是有效的，1 ≤ k ≤ n^2 。
 * 
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	h, v := 0, 0
	count := 1
	length := len(matrix)

	for ; h < length-1 && v < length-1 && count < k; {
		rightVal := matrix[v][h + 1]
		downVal := matrix[v + 1][h]
		fmt.Printf("||val=%d,r=%d,d=%d", matrix[v][h], rightVal, downVal)

		// if rightVal < downVal {
		// 	h++
		// } else if rightVal > downVal {
		// 	v++
		// } else {
		// 	// TODO: ...
		// }
		if rightVal < downVal {
			h++
		} else {
			v++
		}
		count++
	}

	fmt.Printf(" +++ val=%d,v=%d,h=%d,c=%d", matrix[v][h], v, h, count)

	if count == k {
		return matrix[v][h]
	}

	fmt.Printf(" &&& val=%d,v=%d,h=%d,c=%d", matrix[v][h], v, h, count)
	if h == length - 1 {
		for ; count < k; {
			v++; count++
		}
	}

	fmt.Printf(" >>> val=%d,v=%d,h=%d, c=%d", matrix[v][h], v, h, count)

	if v == length - 1 {
		for ; count < k; {
			h++; count++
		}
	}

	fmt.Printf(" $$$ val=%d,v=%d,h=%d,c=%d", matrix[v][h], v, h, count)

	return matrix[v][h]
}
// @lc code=end

