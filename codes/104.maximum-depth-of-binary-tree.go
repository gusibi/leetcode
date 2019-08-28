package codes

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (59.40%)
 * Total Accepted:    468K
 * Total Submissions: 786.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its maximum depth.
 *
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its depth = 3.
 * 思路一：递归，取左子树和右子树最大深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func getDepth(node *TreeNode, depth int) int {
	// depth 为node 节点的深度
	if node.Left == nil && node.Right == nil {
		return depth
	}
	var ldepth, rdepth int
	depth++
	if node.Left != nil {
		ldepth = getDepth(node.Left, depth)
	}
	if node.Right != nil {
		rdepth = getDepth(node.Right, depth)
	}
	if ldepth > rdepth {
		return ldepth
	} else {
		return rdepth
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getDepth(root, 1)
}

// func main() {
// 	/*
// 	   1
// 	  / \
// 	 2   3
// 	    /
// 	   5
// 	    \
// 	     6
// 	*/
// 	proot := &TreeNode{Val: 1}
// 	fmt.Println(maxDepth(proot))
// 	pleft := &TreeNode{Val: 2}
// 	fmt.Println(maxDepth(proot))
// 	pright := &TreeNode{Val: 3}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	fmt.Println(maxDepth(proot))
// 	qleft := &TreeNode{Val: 5}
// 	qright := &TreeNode{Val: 6}
// 	pright.Left = qleft
// 	fmt.Println(maxDepth(proot))
// 	qleft.Right = qright
// 	fmt.Println(maxDepth(proot))
// }
