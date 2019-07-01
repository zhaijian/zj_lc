package main

//https://leetcode-cn.com/problems/generate-parentheses/

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a solution set is:
//
//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	var l []string

	var gen func(prev string, start, end int)

	gen = func(prev string, start, end int) {
		if n*2 == len(prev) {
			l = append(l, prev)
			return
		}

		if start < n {
			gen(prev+"(", start+1, end)
		}

		if end < start {
			gen(prev+")", start, end+1)
		}
	}

	gen("", 0, 0)

	return l
}
