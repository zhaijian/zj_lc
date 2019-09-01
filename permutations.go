package main

//https://leetcode-cn.com/problems/permutations/
//
//Given a collection of distinct integers, return all possible permutations.
//
//Example:
//
//Input: [1,2,3]
//Output:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	copySlice := func(ban int, arr []int) []int {
		if len(arr) == 1 {
			return []int{}
		}

		if ban == 0 {
			return arr[1:]
		}

		if ban == len(arr)-1 {
			return arr[:len(arr)-1]
		}
		var c []int
		c = append(c, arr[0:ban]...)
		c = append(c, arr[ban+1:]...)
		return c
	}

	var res [][]int

	var f func(dest, curr []int)

	f = func(dest, curr []int) {
		if len(curr) == 0 {
			a := make([]int, len(dest))
			copy(a, dest)
			res = append(res, a)
		}
		for i := 0; i < len(curr); i++ {
			dest = append(dest, curr[i])
			f(dest, copySlice(i, curr))
			dest = dest[0 : len(dest)-1]
		}
	}
	f(nil, nums)
	return res
}
