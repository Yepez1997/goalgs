package main

func uniquePathsII(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	// dp table 1. first create the rows 2. create the columns in the row
	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, col)
	}

	if grid[0][0] == 1 {
		return 0
	} else {
		grid[0][0] = 1 // set to a valid path
	}

	// check first col
	for i := 0; i < row; i++ {
		if grid[i][0] == 1 || grid[i-1][0] == 0 {
			result[i][0] = 0
		} else {
			result[i][0] = 1
		}
	}

	// check first row
	for j := 0; j < col; j++ {
		if grid[0][j] == 1 || grid[0][j-1] == 0 {
			result[0][j-1] = 0
		} else {
			result[0][j] = 1
		}
	}

	// start at 1,1 checking for all solutions
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if grid[i][j] != 1 {
				result[i][j] = result[i-1][j] + result[i][j-1]
			} else {
				result[i][j] = 0
			}
		}
	}
	return result[row-1][col-1]
}
