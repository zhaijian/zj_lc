package main

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
//
//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
//You are given a target value to search. If found in the array return its index, otherwise return -1.
//
//You may assume no duplicate exists in the array.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//Example 1:
//
//Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//Example 2:
//
//Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] && target >= nums[start] && target < nums[mid] {
			//左半边是有序的start-mid二分查找
			return binarySearch(nums, start, end, target)
		} else if nums[start] <= nums[mid] && (target > nums[mid] || target < nums[start]) {
			//5 1 3   3
			start = mid + 1
		} else if nums[mid] <= nums[end] && target > nums[mid] && target <= nums[end] {
			//右半边是有序的mid-end二分查找
			return binarySearch(nums, start, end, target)
		} else {
			end = mid - 1
		}
	}

	return -1
}

func binarySearch(nums []int, start, end, target int) int {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
