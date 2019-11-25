package main

import "sort"

// https://leetcode-cn.com/problems/insert-interval/
//
// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
//
// You may assume that the intervals were initially sorted according to their start times.
//
// Example 1:
//
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
// Example 2:
//
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
// NOTE: input types have been changed on April 15, 2019.
// Please reset to default code definition to get new method signature.

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	canMerge := func(left, right []int) bool {
		return left[1] >= right[0]
	}

	merge := func(left, right []int) []int {
		if right[1] > left[1] {
			return []int{left[0], right[1]}
		}
		return left
	}

	if len(intervals) <= 1 {
		return intervals
	}

	var result [][]int
	m := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if canMerge(m, intervals[i]) {
			m = merge(m, intervals[i])
		} else {
			result = append(result, m)
			m = intervals[i]
		}

		if i == len(intervals)-1 {
			result = append(result, m)
			return result
		}
	}
	return result
}
