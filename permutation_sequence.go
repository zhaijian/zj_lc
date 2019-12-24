package main

import "fmt"

// https://leetcode-cn.com/problems/permutation-sequence/
//
// The set [1,2,3,...,n] contains a total of n! unique permutations.
//
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
// Given n and k, return the kth permutation sequence.
//
// Note:
//
// Given n will be between 1 and 9 inclusive.
// Given k will be between 1 and n! inclusive.
// Example 1:
//
// Input: n = 3, k = 3
// Output: "213"
// Example 2:
//
// Input: n = 4, k = 9
// Output: "2314"

func getPermutation(n int, k int) string {
	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	rmIndex := func(curr []int, i int) []int {
		if len(curr) == 1 {
			return []int{}
		}

		if i == 0 {
			return curr[1:]
		}

		if i == len(curr)-1 {
			return curr[:len(curr)-1]
		}
		var c []int
		c = append(c, curr[0:i]...)
		c = append(c, curr[i+1:]...)

		return c
	}

	var offset int
	var f func(result string, curr []int) string
	f = func(result string, curr []int) string {
		if len(curr) == 0 {
			offset++
			fmt.Println(result, offset)
			if offset == k {
				return result
			}
		}

		for i := 0; i < len(curr); i++ {
			copied := fmt.Sprintf(`%s%d`, result, curr[i])
			if str := f(copied, rmIndex(curr, i)); str != "" {
				return str
			}

		}

		return ""
	}
	return f("", arr)
}
