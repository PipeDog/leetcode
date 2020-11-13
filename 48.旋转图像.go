/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	minIndex := 0
	maxIndex := len(matrix) - 1

	for ; minIndex < maxIndex; {
		for offset := 0; offset < maxIndex - minIndex; offset++ {
			temp := matrix[maxIndex - offset][minIndex]
			matrix[maxIndex - offset][minIndex] = matrix[maxIndex][maxIndex - offset]
			matrix[maxIndex][maxIndex - offset] = matrix[minIndex + offset][maxIndex]
			matrix[minIndex + offset][maxIndex] = matrix[minIndex][minIndex + offset]
			matrix[minIndex][minIndex + offset] = temp; 
		}

		minIndex++
		maxIndex--
	}
}
// @lc code=end

