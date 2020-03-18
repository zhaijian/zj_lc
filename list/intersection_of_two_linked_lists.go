package list

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// //
// // 编写一个程序，找到两个单链表相交的起始节点。
// //
// // 注意：
// //
// // 如果两个链表没有交点，返回 null.
// // 在返回结果后，两个链表仍须保持原有的结构。
// // 可假定整个链表结构中没有循环。
// // 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

// 思路a和b走过的总路程相等就一定是相交点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB
	for curA != nil || curB != nil {
		if curA == nil {
			curA = headB
		}

		if curB == nil {
			curB = headA
		}

		if curA == curB {
			return curA
		}

		curA = curA.Next
		curB = curB.Next
	}

	return nil
}
