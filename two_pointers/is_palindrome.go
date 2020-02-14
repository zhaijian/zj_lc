package two_pointers

// https://leetcode-cn.com/problems/valid-palindrome/
//
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:
//
// 输入: "race a car"
// 输出: false

// 思路:先循环把字符把合法的字符放到一个byte数组里
// 然后前后判断即可。
// 把一个字符变小写的方法：item - 'A' + 'a'
// 把一个字符变大写的方法：'A' + (item - 'a')

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	valid := func(item uint8) bool {
		if item >= 'a' && item <= 'z' {
			return true
		}

		if item >= 'A' && item <= 'Z' {
			return true
		}

		if item >= '0' && item <= '9' {
			return true
		}
		return false
	}

	lower := func(item byte) byte {
		if item >= 'A' && item <= 'Z' {
			return item + 'a' - 'A'
		}
		return item
	}

	c := make([]byte, len(s))
	var index int
	for i := 0; i < len(s); i++ {
		if valid(s[i]) {
			c[index] = lower(s[i])
			index++
		}
	}

	i, j := 0, index-1
	for i < j {
		if c[i] != c[j] {
			return false
		}
		i++
		j--
	}
	return true
}
