package main

// https://leetcode-cn.com/problems/length-of-last-word/

// Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
//
// If the last word does not exist, return 0.
//
// Note: A word is defined as a character sequence consists of non-space characters only.
//
// Example:
//
// Input: "Hello World"
// Output: 5

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	var end int
	for end = len(s) - 1; end > 0; end-- {
		if s[end] != ' ' {
			break
		}
	}

	if end == 0 {
		if s[0] == ' ' {
			return 0
		} else {
			return 1
		}
	}

	var start int
	for start = end; start > 0; start-- {
		if s[start] == ' ' {
			break
		}
	}

	if s[start] == ' ' {
		return end - start
	} else {
		return end - start + 1
	}
}
