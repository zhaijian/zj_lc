package main

//https://leetcode-cn.com/problems/sudoku-solver/
//
//Write a program to solve a Sudoku puzzle by filling the empty cells.
//
//A sudoku solution must satisfy all of the following rules:
//
//Each of the digits 1-9 must occur exactly once in each row.
//Each of the digits 1-9 must occur exactly once in each column.
//Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
//Empty cells are indicated by the character '.'.
//
//
//A sudoku puzzle...
//
//
//...and its solution numbers marked in red.
//
//Note:
//
//The given board contain only digits 1-9 and the character '.'.
//You may assume that the given Sudoku puzzle will have a single unique solution.
//The given board size is always 9x9.

func solveSudoku(board [][]byte) {
	var row []map[byte]int
	var col []map[byte]int
	var box []map[byte]int

	getBoxId := func(i, j int) byte {
		return byte((i/3)*3 + j/3)
	}

	for i := 0; i < 9; i++ {
		row = append(row, map[byte]int{})
		col = append(col, map[byte]int{})
		box = append(box, map[byte]int{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			row[i][board[i][j]] = 1
			col[j][board[i][j]] = 1
			box[getBoxId(i, j)][board[i][j]] = 1
		}
	}

	inc := func(r, c int, data byte) {
		board[r][c] = data
		row[r][data]++
		col[c][data]++
		box[getBoxId(r, c)][data]++
	}

	dec := func(r, c int, data byte) {
		board[r][c] = '.'
		row[r][data]--
		col[c][data]--
		box[getBoxId(r, c)][data]--
	}

	var backTrace func(r, c int) bool
	backTrace = func(r, c int) bool {
		rNext, cNext := r, c+1
		if c == 8 {
			rNext, cNext = r+1, 0
		}

		if board[r][c] != '.' {
			if r == 8 && c == 8 {
				return true
			}
			return backTrace(rNext, cNext)
		}

		for i := 1; i <= 9; i++ {
			data := byte(i) + 48
			if row[r][data]+col[c][data]+box[getBoxId(r, c)][data] > 0 {
				continue
			}
			inc(r, c, data)

			if r == 8 && c == 8 {
				return true
			}

			if f := backTrace(rNext, cNext); !f {
				dec(r, c, data)
				continue
			}
			return true
		}

		return false
	}

	backTrace(0, 0)
}
