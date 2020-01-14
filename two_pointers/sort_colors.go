package two_pointers

// https://leetcode-cn.com/problems/sort-colors/

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，
// 使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 注意:
// 不能使用代码库中的排序函数来解决这道题。
//
// 示例:
//
// 输入: [2,0,2,1,1,0]
// 输出: [0,0,1,1,2,2]
// 进阶：
//
// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//

// 三个指针，cur先和左边比较，如果为0则交换，再和右边比较，如果为2则交换 1 则前进
func sortColors(nums []int) {
	var cur int
	left, right := 0, len(nums)-1
	// 是否有等于号 依据是所有的数据都需要比较
	for cur <= right {
		if nums[cur] == 0 {
			nums[cur] = nums[left]
			nums[left] = 0
			cur++
			left++
		} else if nums[cur] == 2 {
			// cur不增加是因为交换过来的数字还没有进行比对
			nums[cur] = nums[right]
			nums[right] = 2
			right--
		} else {
			cur++
		}
	}
}
