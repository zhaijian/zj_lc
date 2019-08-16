package main

//Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.
//
//'?' Matches any single character.
//'*' Matches any sequence of characters (including the empty sequence).
//The matching should cover the entire input string (not partial).
//
//Note:
//
//s could be empty and contains only lowercase letters a-z.
//p could be empty and contains only lowercase letters a-z, and characters like ? or *.
//Example 1:
//
//Input:
//s = "aa"
//p = "a"
//Output: false
//Explanation: "a" does not match the entire string "aa".
//Example 2:
//
//Input:
//s = "aa"
//p = "*"
//Output: true
//Explanation: '*' matches any sequence.
//Example 3:
//
//Input:
//s = "cb"
//p = "?a"
//Output: false
//Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
//Example 4:
//
//Input:
//s = "adceb"
//p = "*a*b"
//Output: true
//Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
//Example 5:
//
//Input:
//s = "acdcb"
//p = "a*c?b"
//Output: false

//尝试递归解法，但是超时了
func isMatchV2(s string, p string) bool {
	if len(s) > 0 && len(p) == 0 {
		return false
	}

	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if p[0] == '*' {
		var start int
		for start < len(p) {
			if p[start] != p[0] {
				break
			}
			start++
		}
		return len(s) > 0 && (isMatchV2(s[1:], p[start-1:]) || isMatchV2(s[1:], p[start:])) || isMatchV2(s, p[start:])
	}

	return len(s) > 0 && (s[0] == p[0] || p[0] == '?') && isMatchV2(s[1:], p[1:])
}

//动态规划
func isMatchV3(s string, p string) bool {
	tmp := make([][]bool, len(s)+1)
	for i := range tmp {
		tmp[i] = make([]bool, len(p)+1)
	}

	//tmp[i][j]表示s的前i个字符，和p的前j个字符，能否匹配
	tmp[0][0] = true
	//初始化
	for j := 1; j <= len(p); j++ {
		tmp[0][j] = tmp[0][j-1] && p[j-1] == '*'
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				tmp[i][j] = tmp[i][j-1] || tmp[i-1][j]
			} else {
				tmp[i][j] = tmp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}

	return tmp[len(s)][len(p)]
}
