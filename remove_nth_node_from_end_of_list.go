package main

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

//Given a linked list, remove the n-th node from the end of list and return its head.
//
//Example:
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//
//Given n will always be valid.
//
//Follow up:
//
//Could you do this in one pass?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	var preNode *ListNode
	for curLen, cur := 1, head; cur != nil; cur, curLen = cur.Next, curLen+1 {
		if curLen == n {
			preNode = dummy
		} else if curLen > n {
			preNode = preNode.Next
		}
	}

	if preNode.Next == head {
		return head.Next
	}
	preNode.Next = preNode.Next.Next

	return head
}
