/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	var islandArea func([][]int, int, int) int
	islandArea = func(grid [][]int, x, y int) int {
		if grid[x][y] == 0 {
			return 0
		}

		grid[x][y] = 0
		area := 1
		if y > 0 {
			area += islandArea(grid, x, y-1)
		}
		if x > 0 {
			area += islandArea(grid, x-1, y)
		}
		if y < len(grid[x]) - 1 {
			area += islandArea(grid, x, y+1)
		}
		if x < len(grid) - 1 {
			area += islandArea(grid, x+1, y)
		}
		return area
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			area := islandArea(grid, i, j)
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
// @lc code=end

