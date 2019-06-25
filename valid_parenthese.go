package main

//https://leetcode-cn.com/problems/valid-parentheses/

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.
//
//Example 1:
//
//Input: "()"
//Output: true
//Example 2:
//
//Input: "()[]{}"
//Output: true
//Example 3:
//
//Input: "(]"
//Output: false
//Example 4:
//
//Input: "([)]"
//Output: false
//Example 5:
//
//Input: "{[]}"
//Output: true

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s) == 0 {
		return false
	}

	if len(s)%2 != 0 {
		return false
	}

	goal := func(i, j int32) bool {
		if i == '(' && j == ')' {
			return true
		}

		if i == '[' && j == ']' {
			return true
		}

		if i == '{' && j == '}' {
			return true
		}

		return false
	}

	stack := make([]int32, 0)
	pop := func() int32 {
		item := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return item
	}
	push := func(item int32) {
		stack = append(stack, item)
	}

	for _, item := range s {
		push(item)

		if len(stack) <= 1 {
			continue
		}

		if goal(stack[len(stack)-2], stack[len(stack)-1]) {
			pop()
			pop()
		}
	}
	return len(stack) == 0
}
