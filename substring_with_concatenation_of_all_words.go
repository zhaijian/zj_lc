package main

//https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
//
//You are given a string, s, and a list of words, words, that are all of the same length.
//Find all starting indices of substring(s) in s that is a concatenation of each word in
//words exactly once and without any intervening characters.
//
//Example 1:
//
//Input:
//s = "barfoothefoobarman",
//words = ["foo","bar"]
//Output: [0,9]
//Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
//The output order does not matter, returning [9,0] is fine too.
//Example 2:
//
//Input:
//s = "wordgoodgoodgoodbestword",
//words = ["word","good","best","word"]
//Output: []

//注意是"that are all of the same length" 他们的长度都相等
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	totalLen := len(words) * len(words[0])
	strMap := map[string]int{}
	for _, str := range words {
		strMap[str] += 1
	}
	if totalLen > len(s) {
		return []int{}
	}

	res := []int{}
	batch := len(words[0])
	for i := 0; i < len(s)-totalLen+1; i++ {
		batchMap := map[string]int{}
		for j := i; j < i+totalLen-batch+1; j += batch {
			if rawCount, ok := strMap[s[j:j+batch]]; !ok {
				break
			} else {
				batchMap[s[j:j+batch]] += 1
				if batchMap[s[j:j+batch]] > rawCount {
					break
				}
			}
		}

		hit := true
		for raw, count := range strMap {
			if curCount, ok := batchMap[raw]; !ok {
				hit = false
				break
			} else if curCount != count {
				hit = false
			}
		}
		if hit {
			res = append(res, i)
		}
	}

	return res
}
