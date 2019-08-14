package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/multiply-strings/
//
//Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
//Example 1:
//
//Input: num1 = "2", num2 = "3"
//Output: "6"
//Example 2:
//
//Input: num1 = "123", num2 = "456"
//Output: "56088"
//Note:
//
//The length of both num1 and num2 is < 110.
//Both num1 and num2 contain only digits 0-9.
//Both num1 and num2 do not contain any leading zero, except the number 0 itself.
//You must not use any built-in BigInteger library or convert the inputs to integer directly.

func multiply(num1 string, num2 string) string {
	tmp := make([]int32, len(num1)+len(num2))
	multi := func(n string, i int32, start int) {
		cursor := len(tmp) - start - 1
		var carry int32
		for j := len(n) - 1; j >= 0; j-- {
			total := tmp[cursor] + (int32(n[j]-'0'))*(i-'0') + carry
			tmp[cursor] = total % 10
			carry = total / 10
			cursor--
		}
		tmp[cursor] += carry
	}

	var j int
	for i := len(num1) - 1; i >= 0; i-- {
		multi(num2, int32(num1[i]), j)
		j++
	}

	var res string
	var b bool
	for _, t := range tmp {
		if !b {
			if t == 0 {
				continue
			}
		}
		b = true

		res += fmt.Sprintf("%v", t)
	}

	if res == "" {
		res = "0"
	}

	return res
}
