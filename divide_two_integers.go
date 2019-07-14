package main

//https://leetcode-cn.com/problems/divide-two-integers/
//
//Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
//
//Return the quotient after dividing dividend by divisor.
//
//The integer division should truncate toward zero.
//
//Example 1:
//
//Input: dividend = 10, divisor = 3
//Output: 3
//Example 2:
//
//Input: dividend = 7, divisor = -3
//Output: -2
//Note:
//
//Both dividend and divisor will be 32-bit signed integers.
//The divisor will never be 0.
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
//	For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

func divide(dividend int, divisor int) int {
	min, max := -2<<30, 2<<30-1
	if dividend == 0 {
		return 0
	} else if dividend == max && divisor == -1 {
		return max
	} else if dividend == min && divisor == -1 {
		return max
	}

	sum := 0
	if dividend > 0 && divisor > 0 {
		for dividend >= divisor {
			sum++
			dividend -= divisor
		}
		return sum
	} else if dividend < 0 && divisor < 0 {
		for dividend <= divisor {
			sum++
			dividend -= divisor
		}
		return sum
	} else if dividend > 0 && divisor < 0 {
		//7 -3
		for dividend+divisor >= 0 {
			sum++
			dividend += divisor
		}
		return 0 - sum
	} else {
		// -18 7
		for dividend+divisor <= 0 {
			sum++
			dividend += divisor
		}
		return 0 - sum
	}
	return sum
}
