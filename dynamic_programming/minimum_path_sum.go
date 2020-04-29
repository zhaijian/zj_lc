package dynamic_programming

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例:
//
// 输入:
// [
//   [1,3,1],
// [1,5,1],
// [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

// 思路 动态规划，i=0或者j=0先初始化，然后循环算出最小即可
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			tmp := grid[i-1][j]
			if grid[i][j-1] < tmp {
				tmp = grid[i][j-1]
			}
			grid[i][j] += tmp
		}
	}
	return grid[m-1][n-1]
}
