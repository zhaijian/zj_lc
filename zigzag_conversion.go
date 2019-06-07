package main

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you
//may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//And then read line by line: "PAHNAPLSIIGYIR"
//
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string s, int numRows);
//Example 1:
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//Example 2:
//
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//
//P     I    N
//A   L S  I G
//Y A   H R
//P     I

func convert(s string, numRows int) string {
	genRows := numRows
	if len(s) < genRows {
		genRows = len(s)
	}

	if genRows <= 1 {
		return s
	}

	rows := make([]string, genRows)
	down := true
	cur := 0

	for i := 0; i < len(s); {
		//向下走，如果已经到头就向上走
		if down {
			if cur < genRows {
				rows[cur] += string(s[i])
				cur++
				i++
			} else {
				//之所以-2是因为最后一个已经走过了，需要走到前一个
				cur = genRows - 2
				down = false
			}
		} else {
			if cur >= 0 {
				rows[cur] += string(s[i])
				cur--
				i++
			} else {
				//之所以为1是因为第一个已经走过了，需要走到第二个
				cur = 1
				down = true
			}
		}
	}

	var str string
	for _, item := range rows {
		str += item
	}

	return str
}
