package main

//Given a string, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

//双指针法，map存放
//如果 s[j]s[j] 在 [i, j)[i,j) 范围内有与 j' 重复的字符，我们不需要逐渐增加 i。 我们可以直接跳过 [i，j']范围内的所有元素，并将 i 变为 j' + 1
func lengthOfLongestSubstring(s string) int {
	var i, max int
	m := make(map[uint8]int)
	for j := 0; j < len(s); j++ {
		sign, ok := m[s[j]]
		if ok {
			if sign > i {
				i = sign
			}
		}
		if max < j-i+1 {
			max = j - i + 1
		}
		m[s[j]] = j + 1
	}
	return max
}
