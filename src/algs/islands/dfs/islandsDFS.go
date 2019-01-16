/* Given a 2D find the number of islands
This Approach uses Depth First Search
*/

package main

func numIslands(grid [][]byte) int {
	count := 0
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return count
	}
	row, col := len(grid), len(grid[0])
	// iterate row first then col
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if string(grid[i][j]) == "1" {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	row, col := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= row || j >= col || string(grid[i][j]) != "1" {
		return
	}
	grid[i][j] = byte('#')
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func main() {
}
