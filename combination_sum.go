package main

//https://leetcode-cn.com/problems/combination-sum/

//Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
//The same repeated number may be chosen from candidates unlimited number of times.
//
//Note:
//
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//Example 1:
//
//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//[7],
//[2,2,3]
//]
//Example 2:
//
//Input: candidates = [2,3,5], target = 8,
//A solution set is:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]

func combinationSum(candidates []int, target int) [][]int {
	var combine func(arr, cur []int, rest int)
	var res [][]int
	combine = func(arr, cur []int, rest int) {
		if len(arr) == 0 {
			return
		}

		for i, v := range arr {
			if v == rest {
				t := make([]int, len(cur))
				copy(t, cur)
				t = append(t, v)
				res = append(res, t)
				continue
			}

			if v > rest {
				continue
			}

			cur = append(cur, v)
			combine(arr[i:], cur, rest-v)
			cur = cur[:len(cur)-1]
		}
	}

	combine(candidates, nil, target)
	return res
}
