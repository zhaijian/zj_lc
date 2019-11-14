package main

// https://leetcode-cn.com/problems/spiral-matrix/

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
//
// Example 1:
//
// Input:
// [
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:
//
// Input:
// [
// [1, 2, 3, 4],
// [5, 6, 7, 8],
// [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var arr []int
	r, r1 := 0, len(matrix)-1
	c, c1 := 0, len(matrix[0])-1

	for r <= r1 && c <= c1 {
		for i := c; i <= c1; i++ {
			arr = append(arr, matrix[r][i])
		}
		for i := r + 1; i <= r1; i++ {
			arr = append(arr, matrix[i][c1])
		}
		if r < r1 && c < c1 {
			for i := c1 - 1; i > c; i-- {
				arr = append(arr, matrix[r1][i])
			}
			for i := r1; i > r; i-- {
				arr = append(arr, matrix[i][c])
			}
		}

		r++
		r1--
		c++
		c1--
	}
	return arr
}
