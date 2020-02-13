package two_pointers

// https://leetcode-cn.com/problems/merge-sorted-array/
//
// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
// 说明:
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// 示例:
//
// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
//
// 输出: [1,2,2,3,5,6]

// 思路：两个指针i, j; 分别指向nums1, nums2的最后一个元素
// 从m + n - 1位置往回移动覆盖最大的元素
// 可能有遗漏的元素，覆盖nums前k个即可
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	index := m + n - 1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
		index--
	}

	for k := 0; k <= j; k++ {
		nums1[k] = nums2[k]
	}
}
