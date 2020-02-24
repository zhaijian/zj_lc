package binary_search

// https://leetcode-cn.com/problems/find-the-duplicate-number/
//
// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
//
// 示例 1:
//
// 输入: [1,3,4,2,2]
// 输出: 2
// 示例 2:
//
// 输入: [3,1,3,4,2]
// 输出: 3
// 说明：
//
// 不能更改原数组（假设数组是只读的）。
// 只能使用额外的 O(1) 的空间。
// 时间复杂度小于 O(n2) 。
// 数组中只有一个重复的数字，但它可能不止重复出现一次。

// 思路1： 额外空间一个mapset，空间复杂度不满足
// 思路2： 排序，相邻相等 违反 不能更改原数组
// 思路3： 官方推荐，快慢指针成环，太不好理解
// 思路4： 二分法，先找一个中间数，循环比较，如果total>mid,则肯定在[1，mid]，否则在(mid, len(nums)-1]
//        最后start和end碰头时则就是那个重复的数字，时间复杂度o(nlogn)

func findDuplicate(nums []int) int {
	start, end := 1, len(nums)-1

	for start < end {
		mid := (start + end) / 2
		total := 0
		for _, num := range nums {
			if num <= mid {
				total++
			}
		}

		if total > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
