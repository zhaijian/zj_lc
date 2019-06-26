package main

//https://leetcode-cn.com/problems/merge-two-sorted-lists/

//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	cur := dummy
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		} else {
			cur.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	}

	if cur2 != nil {
		cur.Next = cur2
	}

	return dummy.Next
}
