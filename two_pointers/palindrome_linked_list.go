package two_pointers

// https://leetcode-cn.com/problems/palindrome-linked-list/
//
// 请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
// 输出: false
// 示例 2:
//
// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 方法1 用额外的空间，把值拷贝到数组中，链表的正序和数组的倒序做比对
// 方法2 第一步：用快慢指针找到前半部分的最后一个结点（需要学习快慢指针这种方法）
//      第二步：反转链表的后半部（需要学习用o(n)的方法反转的做法）
//      第三步：循环判断链表
//      第四步：恢复链表

func isPalindromeII(head *ListNode) bool {
	if head == nil {
		return true
	}
	getHalfEnd := func() *ListNode {
		slow, fast := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	reverse := func(node *ListNode) *ListNode {
		var prev *ListNode
		cur := node
		for cur != nil {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}
		return prev
	}

	halfEndNode := getHalfEnd()
	startSecondHalf := reverse(halfEndNode.Next)

	before, end := head, startSecondHalf
	result := true
	for end != nil {
		if before.Val != end.Val {
			result = false
			break

		}
		before = before.Next
		end = end.Next
	}

	halfEndNode.Next = reverse(startSecondHalf)
	return result
}
