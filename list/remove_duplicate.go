package list

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
//
// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
// 示例 1:
//
// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
// 示例 2:
//
// 输入: 1->1->1->2->3
// 输出: 2->3

// 思路：创建一个哑节点，并创建一个curr指向哑节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	curr := dummyNode

	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			tmp := curr.Next
			for tmp != nil && tmp.Next != nil && tmp.Val == tmp.Next.Val {
				// 3—3-4-4 这种情况会指向4，下次循环时4-4会被过滤
				tmp = tmp.Next
			}
			curr.Next = tmp.Next
		} else {
			curr = curr.Next
		}
	}
	return dummyNode.Next
}
