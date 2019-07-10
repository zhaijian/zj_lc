package main

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//Example:
//
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5
//
//Note:
//
//Only constant extra memory is allowed.
//You may not alter the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//分段处理，每一段进行反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for {
		count := 0
		pre := cur
		for cur != nil && count < k {
			cur = cur.Next
			count++
		}
		if cur == nil {
			break
		}

		//起始为pre,终点为cur
		curTemp := pre.Next
		tail := cur.Next
		for pre.Next != cur {
			temp := pre.Next
			pre.Next = temp.Next
			cur.Next = temp
			temp.Next = tail
			tail = temp

		}
		cur = curTemp
	}
	return dummy.Next
}
