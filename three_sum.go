package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/3sum/
//
//Given an array nums of n integers, are there elements a, b, c in nums
//such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note:
//
//The solution set must not contain duplicate triplets.
//
//Example:
//
//Given array nums = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

//思路，step1 做一些小于3的短路
//step2 排序
//step3 两次循环，求和，双指针
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{{nums[0], nums[1], nums[2]}}
		} else {
			return [][]int{}
		}
	}

	//step 1: sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	if nums[len(nums)-1] < 0 {
		return [][]int{}
	}

	if nums[0] > 0 {
		return [][]int{}
	}

	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		havePre := i >= 1
		if havePre && nums[i-1] == nums[i] {
			continue
		}

		twoSum := calcTwoSum(nums, i+1, len(nums)-1, nums[i])
		if len(twoSum) > 0 {
			res = append(res, twoSum...)
		}
	}
	return res
}

func calcTwoSum(nums []int, start, end, target int) [][]int {
	rest := 0 - target
	i, j := start, end
	var res [][]int
	for i < j {
		havePre := i > start
		if havePre && nums[i] == nums[i-1] {
			i++
			continue
		}
		haveEnd := j < end
		if haveEnd && nums[j] == nums[j+1] {
			j--
			continue
		}

		if nums[i]+nums[j] == rest {
			res = append(res, []int{target, nums[i], nums[j]})
			i++
			j--
		} else if nums[i]+nums[j] > rest {
			j--
		} else {
			i++
		}
	}

	return res
}
