package main

//You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example:
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)   2 4 3    5  2  3    7  6  6 6 6 7
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针法
//注意：设置一个哑节点（dummy head node）
//如果没有哑结点，则必须编写额外的条件语句来初始化表头的值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	carry := 0
	var v1, v2 int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return head.Next
}
