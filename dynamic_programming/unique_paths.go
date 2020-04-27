package dynamic_programming

// https://leetcode-cn.com/problems/unique-paths/

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 问总共有多少条不同的路径？
//
// 示例 1:
//
// 输入: m = 3, n = 2
// 输出: 3
// 解释:
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向右 -> 向下
// 2. 向右 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向右
// 示例 2:
//
// 输入: m = 7, n = 3
// 输出: 28

// 思路 边上肯定都为1中，其余的都是arr[x][y] = arr[x-1]+arr[y-1]
func uniquePaths(m int, n int) int {
	arr := initArray(m, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
	}
	return arr[m-1][n-1]
}

func initArray(m, n int) [][]int {
	arr := [][]int{}
	for i := 0; i < m; i++ {
		arrX := make([]int, n)
		arr = append(arr, arrX)
	}

	for i := 0; i < m; i++ {
		arr[i][0] = 1
	}

	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}
	return arr
}
