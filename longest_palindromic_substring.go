package main

//Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
//
//Example 1:
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//Example 2:
//
//Input: "cbbd"
//Output: "bb"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	var start, end int
	for i := 0; i < len(s)-1; i++ {
		//第i个查询最长
		ls1, le1 := getSubMaxString(s, i, i)
		//第i到i+1查询最长
		ls2, le2 := getSubMaxString(s, i, i+1)

		if end-start < le1-ls1 {
			end = le1
			start = ls1
		}

		if end-start < le2-ls2 {
			end = le2
			start = ls2
		}
	}
	return s[start : end+1]
}

func getSubMaxString(s string, start, end int) (int, int) {
	//已start,end为中心向两边延伸
	for start >= 0 && end < len(s) {
		if s[start] == s[end] {
			start--
			end++
		} else {
			break
		}
	}
	//之所以加1是因为循环加多了
	return start + 1, end - 1
}
