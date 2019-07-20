package main

//https://leetcode-cn.com/problems/longest-valid-parentheses/
//
//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//Example 1:
//
//Input: "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()"
//Example 2:
//
//Input: ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()"

func longestValidParentheses(s string) int {
	var items []int
	push := func(item int) {
		items = append(items, item)
	}

	pop := func() int {
		if len(items) > 0 {
			tmp := items[len(items)-1]
			items = items[0 : len(items)-1]
			return tmp
		}
		return 0
	}

	empty := func() bool {
		return len(items) == 0
	}
	//"()(((((()))))))))))))(())(())(())(())(())(((())))"
	max, start := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			push(i)
			continue
		}

		if empty() {
			start = i + 1
		} else {
			pop()
			if empty() {
				//)()()
				cur := i - start + 1
				if cur > max {
					max = cur
				}
			} else {
				//((()
				cur := i - items[len(items)-1]
				if cur > max {
					max = cur
				}
			}
		}

	}
	return max
}
