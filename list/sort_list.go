package list

// https://leetcode-cn.com/problems/sort-list/
//
// 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
// 示例 1:
//
// 输入: 4->2->1->3
// 输出: 1->2->3->4
// 示例 2:
//
// 输入: -1->5->3->4->0
// 输出: -1->0->3->4->5

// 思路：快慢指针找到中心点
// 左链表和右链表递归分别排序，并合并

func sortList(head *ListNode) *ListNode {
	split := func(x *ListNode) *ListNode {
		var prev *ListNode
		slow, fast := x, x
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}

		if prev != nil {
			prev.Next = nil
		}

		return slow
	}

	merge := func(left, right *ListNode) *ListNode {
		dummyNode := &ListNode{}
		cur := dummyNode
		x, y := left, right
		for x != nil && y != nil {
			if x.Val <= y.Val {
				cur.Next = x
				x = x.Next
			} else {
				cur.Next = y
				y = y.Next
			}
			cur = cur.Next
		}

		if x != nil {
			cur.Next = x
		}

		if y != nil {
			cur.Next = y
		}
		return dummyNode.Next
	}

	var innerSort func(x *ListNode) *ListNode
	innerSort = func(x *ListNode) *ListNode {
		if x == nil || x.Next == nil {
			return x
		}

		mid := split(x)
		left := innerSort(x)
		right := innerSort(mid)

		return merge(left, right)
	}

	return innerSort(head)
}
