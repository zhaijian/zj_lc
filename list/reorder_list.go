package list

// https://leetcode-cn.com/problems/reorder-list/
//
// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1:
//
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 示例 2:
//
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

// 第一步：找到中心点（快慢指针）
// 第二步：将后半部分反转
// 第三步：双指针，将后半部分循环放到前面
func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow == nil {
		return
	}

	var prev *ListNode
	curr := slow.Next
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	ss := head
	for prev != nil {
		prevTmp := prev.Next
		tmp := ss.Next
		ss.Next = prev
		prev.Next = tmp

		prev = prevTmp
		ss = tmp
	}
	slow.Next = nil
}
