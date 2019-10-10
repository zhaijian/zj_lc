package main

import (
	"strings"
)

//https://leetcode-cn.com/problems/n-queens/
//
//The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
//
//
//
//Given an integer n, return all distinct solutions to the n-queens puzzle.
//
//Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
//
//Example:
//
//Input: 4
//Output: [
//[".Q..",  // Solution 1
//"...Q",
//"Q...",
//"..Q."],
//
//["..Q.",  // Solution 2
//"Q...",
//"...Q",
//".Q.."]
//]
//Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

func solveNQueens(n int) [][]string {
	rowMap := map[int]int{}
	colMap := map[int]int{}
	leftMap := map[int]int{}
	rightMap := map[int]int{}

	initArr := func() [][]string {
		raw := make([][]string, n)
		for i := 0; i < n; i++ {
			raw[i] = make([]string, 0)
			for j := 0; j < n; j++ {
				raw[i] = append(raw[i], ".")
			}
		}
		return raw
	}
	raw := initArr()

	exist := func(row, col int) bool {
		if rowMap[row] != 0 {
			return true
		}

		if colMap[col] != 0 {
			return true
		}

		if leftMap[row+col] != 0 {
			return true
		}

		return rightMap[row-col] != 0
	}

	add := func(arr [][]string, row, col int) {
		rowMap[row] = 1
		colMap[col] = 1
		leftMap[row+col] = 1
		rightMap[row-col] = 1
		arr[row][col] = "Q"
	}

	remove := func(arr [][]string, row, col int) {
		rowMap[row] = 0
		colMap[col] = 0
		leftMap[row+col] = 0
		rightMap[row-col] = 0
		arr[row][col] = "."
	}

	var res [][]string
	var f func(arr [][]string, i int)

	f = func(arr [][]string, row int) {
		if row == n {
			var items []string
			for i := 0; i < n; i++ {
				str := strings.Join(arr[i], "")
				items = append(items, str)
			}
			res = append(res, items)
			return
		}

		for i := 0; i < n; i++ {
			if exist(row, i) {
				continue
			}
			add(arr, row, i)
			f(arr, row+1)
			remove(arr, row, i)
		}
	}

	f(raw, 0)
	return res
}
