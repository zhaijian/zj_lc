package main

import "fmt"

//https://leetcode-cn.com/problems/count-and-say/
//
//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//
//Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//Example 1:
//
//Input: 1
//Output: "1"
//Example 2:
//
//Input: 4
//Output: "1211"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	convert := func(str string) string {
		var res string
		var i, j int
		for j < len(str) {
			times := 0
			for j < len(str) && str[j] == str[i] {
				times++
				j++
			}
			res += fmt.Sprintf("%v%v", times, str[i]-'0')
			i = j
		}
		return res
	}

	return convert(countAndSay(n - 1))
}
