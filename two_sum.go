package main

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
	// 循环数组，查询map是否存在结果
	m := make(map[int]int)
	for i, num := range nums {
		rest := target - num
		// 如果存在，返回
		if sign, ok := m[rest]; ok {
			return []int{sign, i}
		}

		// 如果不存在，记录map, key为值，value为坐标
		m[num] = i
	}

	return []int{}
}
