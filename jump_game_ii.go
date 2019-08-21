package main

//https://leetcode-cn.com/problems/jump-game-ii/
//
//Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
//Each element in the array represents your maximum jump length at that position.
//
//Your goal is to reach the last index in the minimum number of jumps.
//
//Example:
//
//Input: [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2.
//Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Note:
//
//You can assume that you can always reach the last index.
//方法一，动态规划
func jump(nums []int) int {
	stepArr := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		max := len(nums) - i - 1
		if max > nums[i] {
			max = nums[i]
		}
		steps := stepArr[i+1] + 1
		for j := 1; j <= max; j++ {
			if stepArr[i+j]+1 < steps {
				steps = stepArr[i+j] + 1
			}
		}
		stepArr[i] = steps
	}
	return stepArr[0]
}

//贪心算法
func jumpV1(nums []int) int {
	getMax := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	//2,3,1,1,4
	var maxLength, end, steps int
	for i := 0; i < len(nums)-1; i++ {
		maxLength = getMax(maxLength, nums[i]+i)
		if i == end {
			end = maxLength
			steps++
		}
	}
	return steps
}
