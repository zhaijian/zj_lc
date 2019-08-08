package main

//https://leetcode-cn.com/problems/first-missing-positive/
//
//Given an unsorted integer array, find the smallest missing positive integer.
//
//Example 1:
//
//Input: [1,2,0]
//Output: 3
//Example 2:
//
//Input: [3,4,-1,1]
//Output: 2
//Example 3:
//
//Input: [7,8,9,11,12]
//Output: 1
//Note:
//
//Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	//第一次循环，查询是否有1
	var contain bool
	for _, v := range nums {
		if v == 1 {
			contain = true
		}
	}

	if !contain {
		return 1
	}

	if len(nums) == 1 {
		return 2
	}

	//第二次循环，大于N和负数都替换为1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}

	abs := func(p int) int {
		if p < 0 {
			return 0 - p
		}

		return p
	}

	//第三次循环，数组游标为a的替换为负数
	for i := 0; i < len(nums); i++ {
		a := abs(nums[i])
		if a == len(nums) {
			nums[0] = abs(nums[0]) * -1
		} else {
			nums[a] = abs(nums[a]) * -1
		}
	}

	//第四次循环，返回值
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums) + 1
}
