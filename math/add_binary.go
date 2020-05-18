package math

// https://leetcode-cn.com/problems/add-binary/

// 给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
// 输入: a = "11", b = "1"
// 输出: "100"
// 示例 2:
//
// 输入: a = "1010", b = "1011"
// 输出: "10101"
//
//
// 提示：
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。

// 注意byte不能参与加减，10进制才支持
func addBinary(a string, b string) string {
	length := len(a) + 1
	if len(a) < len(b) {
		length = len(b) + 1
	}
	arr := make([]byte, length)
	i, j := len(a)-1, len(b)-1
	k := length - 1
	tmp := '0'
	var total int
	for i >= 0 || j >= 0 {
		total = int(tmp - '0')
		if i >= 0 {
			total += int(a[i] - '0')
		}
		if j >= 0 {
			total += int(b[j] - '0')
		}
		if total >= 2 {
			total -= 2
			tmp = '1'
		} else {
			tmp = '0'
		}
		arr[k] = byte(total + '0')
		i--
		j--
		k--
	}
	if tmp == '1' {
		arr[k] = '1'
		return string(arr[k:])
	}
	return string(arr[k+1:])
}
