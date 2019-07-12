package main

//https://leetcode-cn.com/problems/implement-strstr/

//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

//错了N次，思路，暴力循环，第一次循环haystack， 内层循环needle
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	} else if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			return -1
		}

		if haystack[i] != needle[0] {
			continue
		}

		hit := true
		for start, j := i, 0; j < len(needle); j++ {
			if haystack[start] != needle[j] {
				hit = false
				break
			}
			start++
		}

		if hit {
			return i
		}
	}
	return -1
}
