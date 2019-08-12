package main

//https://leetcode-cn.com/problems/trapping-rain-water/
//
//Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
//
//The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
//
//Example:
//
//Input: [0,1,0,2,1,0,1,3,2,1,2,1]
//Output: 6

func trap(height []int) int {
	i, j := 0, len(height)-1
	var leftMax, rightMax int
	var res int
	for i < j {
		if height[i] < height[j] {
			if height[i] <= leftMax {
				res += leftMax - height[i]
			} else {
				leftMax = height[i]
			}
			i++
		} else {
			if height[j] <= rightMax {
				res += rightMax - height[j]
			} else {
				rightMax = height[j]
			}
			j--
		}
	}

	return res
}
