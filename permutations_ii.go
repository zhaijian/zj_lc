package main

//https://leetcode-cn.com/problems/permutations-ii/
//
//Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
//Example:
//
//Input: [1,1,2]
//Output:
//[
//[1,1,2],
//[1,2,1],
//[2,1,1]
//]

func permuteUnique(nums []int) [][]int {
	getRest := func(arr []int, i int) []int {
		if len(arr) == 0 {
			return arr
		}
		if i == 0 {
			return arr[1:]
		}
		if i == len(arr)-1 {
			return arr[:len(arr)-1]
		}

		var s []int
		s = append(s, arr[0:i]...)
		s = append(s, arr[i+1:]...)
		return s
	}

	dup := func(cur []int, i int) bool {
		for j := i - 1; j >= 0; j-- {
			if cur[j] == cur[i] {
				return true
			}
		}
		return false
	}

	var permute func(dest, cur []int)

	var res [][]int
	permute = func(dest, cur []int) {
		if len(cur) == 0 {
			a := make([]int, len(dest))
			copy(a, dest)
			res = append(res, a)
			return
		}
		for i := 0; i < len(cur); i++ {
			if dup(cur, i) {
				continue
			}
			dest = append(dest, cur[i])
			permute(dest, getRest(cur, i))
			dest = dest[0 : len(dest)-1]
		}
	}

	permute(nil, nums)
	return res
}
