package list

// https://leetcode-cn.com/problems/add-two-numbers-ii/
//
// 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 进阶:
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
// 示例:
//
// 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出: 7 -> 8 -> 0 -> 7

// 使用两个栈分别压入数据相加

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := newStack()
	for cur := l1; cur != nil; cur = cur.Next {
		stack1.push(cur)
	}
	stack2 := newStack()
	for cur := l2; cur != nil; cur = cur.Next {
		stack2.push(cur)
	}

	var newNode *ListNode
	var inc int
	for !stack1.empty() && !stack2.empty() {
		c1 := stack1.pop()
		c2 := stack2.pop()
		cur := &ListNode{Val: (c1.Val + c2.Val + inc) % 10}
		cur.Next = newNode
		newNode = cur
		inc = (c1.Val + c2.Val + inc) / 10
	}

	for !stack1.empty() {
		c1 := stack1.pop()
		cur := &ListNode{Val: (c1.Val + inc) % 10}
		cur.Next = newNode
		newNode = cur
		inc = (c1.Val + inc) / 10
	}

	for !stack2.empty() {
		c2 := stack2.pop()
		cur := &ListNode{Val: (c2.Val + inc) % 10}
		cur.Next = newNode
		newNode = cur
		inc = (c2.Val + inc) / 10
	}

	if inc > 0 {
		cur := &ListNode{Val: inc}
		cur.Next = newNode
		newNode = cur
	}
	return newNode
}

type Stack struct {
	nodes []*ListNode
}

func newStack() *Stack {
	return &Stack{nodes: []*ListNode{}}
}

func (p *Stack) push(n *ListNode) {
	p.nodes = append(p.nodes, n)
}

func (p *Stack) empty() bool {
	return len(p.nodes) == 0
}

func (p *Stack) pop() *ListNode {
	if len(p.nodes) == 0 {
		return nil
	}
	tmp := p.nodes[len(p.nodes)-1]
	p.nodes = p.nodes[0 : len(p.nodes)-1]
	return tmp
}
