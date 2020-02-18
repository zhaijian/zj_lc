package two_pointers

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
//
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
// 说明：不允许修改给定的链表。
//
// 示例 1：
//
// 输入：head = [3,2,0,-4], pos = 1
// 输出：tail connects to node index 1
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
// 输入：head = [1,2], pos = 0
// 输出：tail connects to node index 0
// 解释：链表中有一个环，其尾部连接到第一个节点。
//
// 示例 3：
//
// 输入：head = [1], pos = -1
// 输出：no cycle
// 解释：链表中没有环。
//
//
// 进阶：
// 你是否可以不用额外空间解决此题？

// 思路1 和linked_list_cycle一样，创建一个map,循环即可
// 思路2 第一步，slow和fast指针，得到相遇的结点，如果不存在，则直接返回
// 第二步，head和相遇的结点分别向前迭代相遇时就是 成环处
// 比如起点到成环点距离为a, 成环点到相遇点距离为b, 成环点到相遇点另一部分为c
// 则 2* (a + b) = a + b + c + b，那么 a = c，由此第二步成立。

func detectCycle(head *ListNode) *ListNode {
	getNode := func() *ListNode {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return slow
			}
		}
		return nil
	}

	node := getNode()
	if node == nil {
		return nil
	}

	for head != node {
		head = head.Next
		node = node.Next
	}
	return head
}
