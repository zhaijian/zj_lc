package main

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

//Example:
//
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//Note:
//
//Although the above answer is in lexicographical order, your answer could be in any order you want.
//234
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	m := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

	res := make([]string, 0)

	var combine func(ds, pre string)

	combine = func(ds, pre string) {
		if len(ds) == 0 {
			res = append(res, pre)
			return
		}

		for _, l := range m[ds[0:1]] {
			combine(ds[1:], pre+string(l))
		}
	}

	combine(digits, "")
	return res
}
