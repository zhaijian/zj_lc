package main

//Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
//If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
//
//The replacement must be in-place and use only constant extra memory.
//
//Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
//
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

//158476531
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}

	reverseNum(nums, i+1)
}

//7 6 5 4 3 2 1
func reverseNum(nums []int, start int) {
	end := len(nums) - 1
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
