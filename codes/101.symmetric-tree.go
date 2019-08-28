package codes

// import "fmt"

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (42.82%)
 * Total Accepted:    369.6K
 * Total Submissions: 861.9K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 *
 *
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 前序遍历(左中右) 1 2 3 4 2 4 3
 * 前序(右中左)遍历 1 2 3 4 2 4 3
 *
 * 中序遍历(左中右) 3 2 4 1 4 2 3
 * 中序(右中左)遍历 3 2 4 1 4 2 3
 *
 *
 * But the following [1,2,2,null,3,null,3]  is not:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 * 前序遍历(中左右) 1 2 3 2 3
 * 前序(中右左)遍历 1 2 3 2 3
 *
 * 中序遍历(左中右) 2 3 1 2 3
 * 中序(右中左)遍历 3 2 1 3 2
 *
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
 *
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

func inOrderTraversal1(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	} else if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	lsame := inOrderTraversal1(p.Left, q.Right)
	rsame := inOrderTraversal1(p.Right, q.Left)
	if lsame && rsame {
		return true
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
	// 使用中序遍历
	// 1. 左中右
	// 2. 右中左
	// 如果遍历后节点相同，则说明是对称
	if root == nil {
		return true
	}
	return inOrderTraversal1(root.Left, root.Right)
}

// func main() {
// 	proot := &TreeNode{Val: 1}
// 	pleft := &TreeNode{Val: 2}
// 	pright := &TreeNode{Val: 3}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	fmt.Println(isSymmetric(proot))

// 	proot = &TreeNode{Val: 1}
// 	pleft = &TreeNode{Val: 2}
// 	pright = &TreeNode{Val: 2}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	fmt.Println(isSymmetric(proot))

// 	proot = &TreeNode{Val: 1}
// 	pleft = &TreeNode{Val: 2}
// 	proot.Left = pleft
// 	fmt.Println(isSymmetric(proot))

// 	proot = &TreeNode{Val: 1}
// 	fmt.Println(isSymmetric(proot))
// }
