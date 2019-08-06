package main

import "sort"

//https://leetcode-cn.com/problems/combination-sum-ii/

//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note:
//
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//Example 1:
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8,
//A solution set is:
//[
//[1, 7],
//[1, 2, 5],
//[2, 6],
//[1, 1, 6]
//]
//Example 2:
//
//Input: candidates = [2,5,2,1,2], target = 5,
//A solution set is:
//[
//  [1,2,2],
//  [5]
//]

//[]int{10, 1, 2, 7, 6, 1, 5}, 8
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var res [][]int
	var combine func(arr, cur []int, rest int)

	combine = func(arr, cur []int, rest int) {
		if len(arr) == 0 {
			return
		}

		for i := 0; i < len(arr); i++ {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}

			if arr[i] == rest {
				temp := make([]int, len(cur))
				copy(temp, cur)
				temp = append(temp, arr[i])
				res = append(res, temp)
				continue
			}

			if arr[i] > rest {
				continue
			}

			cur = append(cur, arr[i])
			combine(arr[i+1:], cur, rest-arr[i])
			cur = cur[:len(cur)-1]
		}
	}

	combine(candidates, nil, target)
	return res
}
