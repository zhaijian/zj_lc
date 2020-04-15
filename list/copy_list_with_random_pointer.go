package list

// https://leetcode-cn.com/problems/copy-list-with-random-pointer/
//
// 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
//
// 要求返回这个链表的 深拷贝。
//
// 我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//
// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
//
//
// 示例 1：
//
//
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
// 示例 2：
//
//
//
// 输入：head = [[1,1],[2,1]]
// 输出：[[1,1],[2,1]]
// 示例 3：
//
//
//
// 输入：head = [[3,null],[3,0],[3,null]]
// 输出：[[3,null],[3,0],[3,null]]
// 示例 4：
//
// 输入：head = []
// 输出：[]
// 解释：给定的链表为空（空指针），因此返回 null。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 思路1： 将所有的指针放到map里，key 为当前node，value为复制的结点
// 通过m[cur].Next = m[cur.Next]  m[cur].Random = m[cur.Random]

// 思路2： 1 迭代复制下一个结点，2 迭代复制random结点， 3 分离结点
func copyRandomList(head *Node) *Node {
	copyNext(head)
	copyRandom(head)
	return split(head)
}

func copyNext(head *Node) {
	cur := head
	for cur != nil {
		tmp := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = tmp
		cur = cur.Next.Next
	}
}

func copyRandom(head *Node) {
	cur := head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
}

func split(head *Node) *Node {
	dummyNode := &Node{}
	cur := dummyNode
	cur1 := head
	for cur1 != nil {
		tmp := cur1.Next
		cur1.Next = cur1.Next.Next
		cur.Next = tmp
		cur1 = cur1.Next
		cur = cur.Next
	}
	return dummyNode.Next
}
