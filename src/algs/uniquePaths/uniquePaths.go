// Unique Paths -> determaines the number of unique paths from (m,n) at the top left on a grid to
// the bottom right, end
// Note -> m rows n columns  -> m by n
// Problem TO DO
package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	waysToMake := make([][]int, n)
	for i := 0; i < n; i++ {
		waysToMake[i] = make([]int, m)
	}
	if m == 1 || n == 1 {
		return 1
	} else {
		// rows
		for i := 0; i < n; i++ {
			waysToMake[i][0] = 1
		}
		// cols
		for j := 0; j < m; j++ {
			waysToMake[0][j] = 1
		}
		// 1 accounts for the offset
		for i := 1; i < n; i++ {
			for j := 1; j < m; j++ {
				waysToMake[i][j] = waysToMake[i-1][j] + waysToMake[i][j-1]
			}
		}
	}
	return waysToMake[n-1][m-1]
}

func main() {
	// Unique Paths I
	n := 3
	m := 7
	uniquePathNum := uniquePaths(m, n)
	fmt.Println(uniquePathNum)
	// Unique Paths II
	input := [][]int{{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0}}

}

// Just use dp -> solution for uniquePaths one and uniquePaths two
// 1. want to create m by n template to iterate
// 2. keep track of all paths we go through
// 3. run bfs or dfs
// note we do not need to find the min shortest path
// we want to develop a shortest path algorithm s.t
// we find all possible solutions
// use a dfs algorithm ? or bfs
// pros cons
// bfs find all neighboring distance, similar to shortest path but without weights
// dfs explores all nodes from top to its depth
// can only go up and down i.e left and right
// from m what are all the possible paths
// from m + 1 what are all the posible paths
// from n what are all the possible paths
// from n + 1 what are all the possible paths
// from m,n what are all the possible paths
// from m - 1, n - 1 what are all the possible paths ?
// how can we use our resources to generate a path
// want to run the path n time ?

// given a 2 * 2 grid
/*
(1,0)
  x  x
  x  x

  Possible paths ->
	   right  right down
  1 -> (0,0), (1,0), (1,1)
		down  right right
  2 -> (0,0), (0,1), (1,1)
  2 total paths
  row, col
  start position -> (0, n)
  Given that we start at (0,n) find all possible paths to (m,n)
  path from (m, n - 1)
  path from (m - 1, n) down
  start = 0 right condition
  (0,n) -> (1,n) -> (2,n) -> (start + 1) < m
  (0, n+1) ()
  // keep track of indexes
  from (m,n-1) have we reached the next node before ?
  from (m, n-1) or (m - 1, n) is there a path to top left


*/
