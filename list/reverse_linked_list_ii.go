package list

// https://leetcode-cn.com/problems/reverse-linked-list-ii/
//
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL

// 第一步：找到需要反转的起始和结束点
// 第二步：反转起始到结束点
// 第三步：将节点连接
// 需要注意的是：用哑节点防止空指针
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{Next: head}

	getReverseNode := func() (*ListNode, *ListNode) {
		var prev, tail *ListNode
		i := 0
		for cur := dummyNode; cur != nil; cur = cur.Next {
			if i == m-1 {
				prev = cur
			}

			if i == n {
				tail = cur
			}
			i++
		}

		return prev, tail
	}

	reverse := func(prev, tail *ListNode) (*ListNode, *ListNode) {
		cur := prev
		var pre *ListNode
		for cur != tail {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}

		tail.Next = pre
		return tail, prev
	}

	prev, tail := getReverseNode()
	tmp := tail.Next
	rPrev, rTail := reverse(prev.Next, tail)

	prev.Next = rPrev
	rTail.Next = tmp
	return dummyNode.Next
}
