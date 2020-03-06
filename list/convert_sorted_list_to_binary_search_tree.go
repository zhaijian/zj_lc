package list

// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
//
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
// 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：快慢指针找到中间点，分别对左右节点做递归，直到为空为止

func sortedListToBST(head *ListNode) *TreeNode {
	getMid := func(start *ListNode) *ListNode {
		slow := start
		var prev *ListNode
		fast := start
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

	var buildTree func(start *ListNode) *TreeNode
	buildTree = func(start *ListNode) *TreeNode {
		if start == nil {
			return nil
		}
		midNode := getMid(start)
		node := &TreeNode{Val: midNode.Val}

		if start == midNode {
			return node
		}

		node.Left = buildTree(start)
		node.Right = buildTree(midNode.Next)
		return node
	}

	return buildTree(head)
}
