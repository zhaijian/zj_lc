package main

//https://leetcode-cn.com/problems/merge-k-sorted-lists/

//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
//Example:
//
//Input:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//Output: 1->1->2->3->4->4->5->6

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	for i := 1; i < len(lists); i *= 2 {
		for j := 0; j < len(lists)-i; j += i * 2 {
			lists[j] = mergeTwoLists(lists[j], lists[j+i])
		}
	}

	return lists[0]
}
