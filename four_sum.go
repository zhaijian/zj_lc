package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/4sum/
//
//Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target?
//Find all unique quadruplets in the array which gives the sum of target.
//
//Note:
//
//The solution set must not contain duplicate quadruplets.
//
//Example:
//
//Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			return [][]int{nums}
		}
		return [][]int{}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int
	for i := 0; i <= len(nums)-4; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j <= len(nums)-3; j++ {
			rest := target - (nums[i] + nums[j])
			start, end := j+1, len(nums)-1

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			if nums[start]+nums[start+1] > rest {
				continue
			}

			if nums[end]+nums[end-1] < rest {
				continue
			}

			for start < end {
				if start > j+1 && nums[start] == nums[start-1] {
					start++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] {
					end--
					continue
				}

				if rest == nums[start]+nums[end] {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
				} else if nums[start]+nums[end] > rest {
					end--
				} else if nums[start]+nums[end] < rest {
					start++
				}
			}
		}
	}

	return res
}
