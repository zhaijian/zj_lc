package main

//https://leetcode-cn.com/problems/n-queens-ii/

func totalNQueens(n int) int {
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

	var count int
	var f func(arr [][]string, i int)

	f = func(arr [][]string, row int) {
		if row == n {
			count++
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
	return count

}
