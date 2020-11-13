func getPermutation(n int, k int) string {
	nums := []int{}
	nums = append(nums, 1)

	count := make([]int, n)
	count[0] = 1

	for i := 1; i < n; i++ {
		nums = append(nums, i+1)
		count[i] = count[i-1] * nums[i]
	}

	numStrs := map[int]string {
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}

	searchIndex := func(nums []int, n int) int {
		left := 0
		right := len(nums) - 1

		for left < right {
			mid := left + (right - left) / 2
			midVal := nums[mid]

			if midVal == n {
				return mid
			} else if midVal < n {
				left = mid
			} else {
				right = mid
			}
		}
		return -1
	}

	ret := ""

	for len(nums) > 0 {
		fmt.Printf("+++")

		curIndex := len(nums) - 1
		curCount := count[curIndex]

		numIndex := (k - 1) / curCount + 1
		fmt.Printf("numIndex=%d,", numIndex)
		digit := nums[numIndex]
		fmt.Printf("digit=%d,", digit)
		// digit := (k - 1) / curCount
		// digit := (k - 1) / preCount + 1

		// digit := (k - 1) / (curCount / nums[curIndex]) + 1
		k = k % nums[curIndex]
		fmt.Printf("k=%d,", k)

		ret = ret + numStrs[digit]
		fmt.Printf("ret=%s,", ret)

		removeIndex := searchIndex(nums, digit)
		fmt.Printf("removeIndex=%d,", removeIndex)

		leftVals := nums[:removeIndex]
		rightVals := nums[removeIndex+1:]
		nums = append(leftVals, rightVals...)
	}

	return ret
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (46.61%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 33.5K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 * 
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * 给定 n 和 k，返回第 k 个排列。
 * 
 * 说明：
 * 
 * 
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 * 
 * 
 * 示例 1:
 * 
 * 输入: n = 3, k = 3
 * 输出: "213"
 * 
 * 
 * 示例 2:
 * 
 * 输入: n = 4, k = 9
 * 输出: "2314"
 * 
 * 
 */

// @lc code=start