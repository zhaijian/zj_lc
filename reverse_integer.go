package main

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed
//integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0
//when the reversed integer overflows.

//判断大于7或者小于-8,是最后一位是7或者-8
func reverse(x int) int {
	var r int
	for x != 0 {
		pop := x % 10
		if r > 2147483647/10 || (r == 2147483647/10 && pop > 7) {
			return 0
		}

		if r < -2147483648/10 || (r == -2147483648/10 && pop < -8) {
			return 0
		}

		r = r*10 + pop
		x /= 10
	}
	return r
}
