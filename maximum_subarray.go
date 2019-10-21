package main

//https://leetcode-cn.com/problems/maximum-subarray/

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//Follow up:
//
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if ans < sum {
			ans = sum
		}
	}
	return ans

}
