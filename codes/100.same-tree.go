package codes

// import "fmt"

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 *
 * https://leetcode.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (49.51%)
 * Total Accepted:    356.1K
 * Total Submissions: 718.7K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * Given two binary trees, write a function to check if they are the same or
 * not.
 *
 * Two binary trees are considered the same if they are structurally identical
 * and the nodes have the same value.
 *
 * Example 1:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:     1         1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * Output: false
 *
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

func inOrderTraversal(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	} else if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	lsame := inOrderTraversal(p.Left, q.Left)
	rsame := inOrderTraversal(p.Right, q.Right)
	if lsame && rsame {
		return true
	} else {
		return false
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return inOrderTraversal(p, q)
}

// func main() {
// 	proot := &TreeNode{Val: 1}
// 	pleft := &TreeNode{Val: 2}
// 	pright := &TreeNode{Val: 3}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	qroot := &TreeNode{Val: 1}
// 	qleft := &TreeNode{Val: 2}
// 	qright := &TreeNode{Val: 3}
// 	qroot.Left = qleft
// 	qroot.Right = qright
// 	fmt.Println(isSameTree(proot, qroot))

// 	proot = &TreeNode{Val: 1}
// 	pleft = &TreeNode{Val: 2}
// 	proot.Left = pleft
// 	qroot = &TreeNode{Val: 1}
// 	qright = &TreeNode{Val: 2}
// 	qroot.Right = qright
// 	fmt.Println(isSameTree(proot, qroot))

// 	proot = &TreeNode{Val: 1}
// 	pleft = &TreeNode{Val: 2}
// 	pright = &TreeNode{Val: 1}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	qroot = &TreeNode{Val: 1}
// 	qleft = &TreeNode{Val: 1}
// 	qright = &TreeNode{Val: 2}
// 	qroot.Left = qleft
// 	qroot.Right = qright
// 	fmt.Println(isSameTree(proot, qroot))

// 	proot = &TreeNode{Val: 1}
// 	qroot = &TreeNode{Val: 1}
// 	fmt.Println(isSameTree(proot, qroot))
// }
