/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {

	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'
		dfs(grid, x, y-1)
		dfs(grid, x-1, y)
		dfs(grid, x, y+1)
		dfs(grid, x+1, y)
	}

	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	
	return num
}
// @lc code=end

