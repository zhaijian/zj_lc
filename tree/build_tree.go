package tree

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
//
// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：
//
//  3
// / \
// 9  20
// /  \
// 15   7
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先找先序遍历找root，然后根据root在中序遍历的位置确定root左子树的长度
// 这样就能确定root下面左节点和右节点的位置
// 左节点在cur+1, 右节点在cur+(idx-left)+1的位置，以此类推
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := map[int]int{}
	for idx, item := range inorder {
		m[item] = idx
	}
	return gen(preorder, m, 0, 0, len(inorder)-1)
}

// left， right 分别代表中序遍历的节点的起始和结束
func gen(preorder []int, inorderMap map[int]int, cur, left, right int) *TreeNode {
	if right < left {
		return nil
	}

	root := &TreeNode{
		Val: preorder[cur],
	}
	idx := inorderMap[preorder[cur]]
	root.Left = gen(preorder, inorderMap, cur+1, left, idx-1)
	root.Right = gen(preorder, inorderMap, cur+(idx-left)+1, idx+1, right)
	return root
}
