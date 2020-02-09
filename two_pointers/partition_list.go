package two_pointers

// https://leetcode-cn.com/problems/partition-list/
//
// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
// 你应当保留两个分区中每个节点的初始相对位置。
//
// 示例:
//
// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：设置两个指针，迭代node, currNode比x小的
// 放before，比x大的放after
// before：1->2->2 after: 4->3->5
// 设置before.Next=afterHead.Next即可连接
// 需要注意的是，防止形成环状链表，需要after.Next=nil

func partition(head *ListNode, x int) *ListNode {
	before := &ListNode{}
	beforeHead := before

	after := &ListNode{}
	afterHead := after

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	before.Next = afterHead.Next
	after.Next = nil
	return beforeHead.Next
}
