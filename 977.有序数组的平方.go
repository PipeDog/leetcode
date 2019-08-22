/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(A []int) []int {
	B := make([]int, len(A))

	for i, val := range A {
		A[i] = val * val
	}

	left := 0
	right := len(B) - 1

	// 
	/*
		注意：这里必须从后向前赋值
		@eg:
			ori nums: [-4,-1,0,3,10]
			after nums: [16,1,0,9,100]

			如果从前向后赋值，则情形如下：
			A[left] = 16, A[right] = 100, B[0] = 16
			A[left] = 1,  B[right] = 100, B[1] = 1
			A[left] = 0,  B[right] = 100, B[2] = 0
			A[left] = 9,  B[right] = 100, B[3] = 9
			A[left] = 100,B[right] = 100, B[4] = 100

		所以需要从后向前赋值
	 */
	for i := right; left <= right; i-- {
		if A[left] <= A[right] {
			B[i] = A[right]
			right--
		} else {
			B[i] = A[left]
			left++
		}
	}

	return B
}

