package main

//https://leetcode-cn.com/problems/swap-nodes-in-pairs/

//Given a linked list, swap every two adjacent nodes and return its head.
//
//You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//
//
//Example:
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	cur := head
	head = head.Next

	var prev *ListNode
	next := cur.Next
	tail := next.Next

	for cur != nil && next != nil {
		if prev != nil {
			prev.Next = next
		}
		next.Next = cur
		cur.Next = tail

		prev = cur
		cur = cur.Next
		if cur == nil {
			break
		}

		next = cur.Next
		if next != nil {
			tail = next.Next
		}
	}

	return head
}
