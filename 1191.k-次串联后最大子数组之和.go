/*
 * @lc app=leetcode.cn id=1191 lang=golang
 *
 * [1191] K 次串联后最大子数组之和
 */

// @lc code=start
func kConcatenationMaxSum(arr []int, k int) int {
	if len(arr) == 0 || k == 0 {
		return 0
	}

	fetchMaxSum := func(num []int, count int) int {
		curSum, maxSum := 0, 0

		for i := 0; i < count; i++ {
			for j := 0; j < len(num); j++ {
				if curSum < 0 {
					curSum = arr[j]
				} else {
					curSum += arr[j]
				}
		
				if curSum > maxSum {
					maxSum = curSum
				}
			}
		}
	
		return maxSum
	}

	fetchAllSum := func(num []int) int {
		sum := 0
		for i := 0; i < len(num); i++ {
			sum += num[i]
		}
		return sum
	}

	mod, maxSum := 1000000007, 0

	if k == 1 {
		maxSum = fetchMaxSum(arr, 1)
	} else if k == 2 {
		maxSum = fetchMaxSum(arr, 2)
	} else if k > 2 {
		oneAllSum := fetchAllSum(arr)

		if oneAllSum > 0 {
			maxSum = (k - 2) * oneAllSum + fetchMaxSum(arr, 2)
		} else {
			maxSum = fetchMaxSum(arr, 2)
		}
	} else {
		// invalid arguments `k`
		return 0
	}

	return maxSum % mod
}
// @lc code=end

