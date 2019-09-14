package main

import (
	"sort"
	"strings"
)

//https://leetcode-cn.com/problems/group-anagrams/
//
//Given an array of strings, group anagrams together.
//
//Example:
//
//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//Note:
//
//All inputs will be in lowercase.
//The order of your output does not matter.

func groupAnagrams(strs []string) [][]string {
	f := func(s string) string {
		//还可以用统计计数的方法
		arr := strings.Split(s, "")
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		return strings.Join(arr, "")
	}

	m := make(map[string][]string)
	for _, str := range strs {
		rs := f(str)
		m[rs] = append(m[rs], str)
	}

	var arr [][]string
	for _, s := range m {
		arr = append(arr, s)
	}
	return arr
}
