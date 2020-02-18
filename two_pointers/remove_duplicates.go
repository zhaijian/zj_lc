package two_pointers

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/submissions/

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
// 示例 1:
//
// 给定 nums = [1,1,1,2,2,3],
//
// 函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
//
// 你不需要考虑数组中超出新长度后面的元素。
// 示例 2:
//
// 给定 nums = [0,0,1,1,1,1,2,3,3],
//
// 函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。
//
// 你不需要考虑数组中超出新长度后面的元素。
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
// int len = removeDuplicates(nums);
//
// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
// for (int i = 0; i < len; i++) {
//     print(nums[i]);
// }

// 思路： 快指针i, 慢指针j，count，
// 如果没有连续重复超过三个或者不相同
// 则i=j,且 nums[j]覆盖nums[i]
// 否则nums[j]覆盖最新的不重复的元素（nums[i]）
// 然后j会和i保持一致

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count, j := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}