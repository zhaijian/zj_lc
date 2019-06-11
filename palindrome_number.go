package main

//Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
//Example 1:
//
//Input: 121
//Output: true
//Example 2:
//
//Input: -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//
//Input: 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	var rest int
	for x > rest {
		//如果为x=121
		if x/10 == rest {
			return true
		}

		rest = rest*10 + x%10
		x /= 10
		//如果为1221
		if x == rest {
			return true
		}
	}

	return false
}
