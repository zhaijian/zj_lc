package main

// https://leetcode-cn.com/problems/rotate-list/

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Example 1:
//
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL
// Example 2:
//
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	if head.Next == nil {
		return head
	}

	getTotal := func() int {
		var total int
		for p := head; p != nil; p = p.Next {
			total++
		}
		return total
	}

	total := getTotal()
	if total <= k {
		k = k % total
		if k == 0 {
			return head
		}
	}

	p1, p2 := head, head

	var len int
	for p2.Next != nil {
		if len < k {
			len++
		} else {
			p1 = p1.Next
		}

		p2 = p2.Next
	}

	p3 := p1.Next
	p2.Next = head
	p1.Next = nil

	return p3
}
