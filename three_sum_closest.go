package main

import "sort"

//https://leetcode-cn.com/problems/3sum-closest/

//Given an array nums of n integers and an integer target,
//find three integers in nums such that the sum is closest to target.
//Return the sum of the three integers.
//You may assume that each input would have exactly one solution.
//
//Example:
//
//Given array nums = [-1, 2, 1, -4], and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

//先排序，然后游标分别指向前面和后面，取和target相减的绝对值
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return nums[0] + nums[1]
	}

	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	abs := func(i int) int {
		if i < 0 {
			return -i
		}

		return i
	}

	var start, end int
	var total int
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		rest := target - nums[i]
		start, end = i+1, len(nums)-1
		for start < end {
			total = nums[start] + nums[end] + nums[i]
			if abs(total-target) < abs(ans-target) {
				ans = total
			}
			if nums[start]+nums[end] > rest {
				end--
			} else {
				start++
			}
		}
	}
	return ans
}
