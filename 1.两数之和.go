/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
    ret := make([]int, 2)

    if len(nums) == 0 {
		return ret
	}

	indexs := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		indexs[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
        
		num1 := nums[i]
		num2 := target - num1
                
		if _, ok := indexs[num2]; ok {
            
            if i == indexs[num2] {
                continue
            }

            ret[0] = i
            ret[1] = indexs[num2]
			return ret
		}
	}
    
	return ret
}

