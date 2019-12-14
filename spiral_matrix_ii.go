package main

// https://leetcode-cn.com/problems/spiral-matrix-ii/
//
// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
//
// Example:
//
// Input: 3
// Output:
// [
// [ 1, 2, 3 ],
// [ 8, 9, 4 ],
// [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}

	ii, jj := 0, n-1
	index := 1

	for ii < jj {
		for i := ii; i <= jj; i++ {
			arr[ii][i] = index
			index++
		}

		for i := ii + 1; i <= jj; i++ {
			arr[i][jj] = index
			index++
		}

		for i := jj - 1; i >= ii; i-- {
			arr[jj][i] = index
			index++
		}

		for i := jj - 1; i > ii; i-- {
			arr[i][ii] = index
			index++
		}

		ii++
		jj--
	}

	if ii == jj {
		arr[ii][jj] = n * n
	}

	return arr
}
