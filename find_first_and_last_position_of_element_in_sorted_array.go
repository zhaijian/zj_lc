package main

//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
//
//Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//If the target is not found in the array, return [-1, -1].
//
//Example 1:
//
//Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4]
//Example 2:
//
//Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lStart := getRange(nums, target, true)
	if lStart == len(nums) || nums[lStart] != target {
		return []int{-1, -1}
	}
	//5,7,7,8,8,10 8
	lEnd := getRange(nums, target, false) - 1

	return []int{lStart, lEnd}
}

func getRange(nums []int, target int, left bool) int {
	start, end := 0, len(nums)
	for start < end {
		//1  2 2 3 4
		mid := (start + end) / 2
		if nums[mid] > target || (nums[mid] == target && left) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}
