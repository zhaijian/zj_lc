package main

//https://leetcode-cn.com/problems/longest-common-prefix/

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	firstItem := strs[0]
	var res string

	for i := 0; i < len(firstItem); i++ {
		for j := 1; j < len(strs); i++ {
			if len(strs[j]) < i+1 {
				return res
			}

			if strs[j][i] != firstItem[i] {
				return res
			}

			if j == len(strs)-1 {
				res += string(strs[j][i])
			}
		}
	}

	return res
}
