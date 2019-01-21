// Unique Paths -> determaines the number of unique paths from (m,n) at the top left on a grid to
// the bottom right, end
// Note -> m rows n columns  -> m by n
package main

func setUpGrid(m int, n int) int {
	return 0
}

func uniquePaths() int {

	return 0
}
func main() {

}

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
