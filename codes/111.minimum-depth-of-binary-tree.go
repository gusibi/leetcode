package main

import "fmt"

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (34.92%)
 * Total Accepted:    278.8K
 * Total Submissions: 797.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
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
 * return its minimum depth = 2.
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

func getDepth1(node *TreeNode, depth int) int {
	// depth 为node 节点的深度
	fmt.Println("depth: ", depth)
	if node.Left == nil && node.Right == nil {
		return depth
	}
	var ldepth, rdepth int
	depth++
	if node.Left != nil {
		ldepth = getDepth1(node.Left, depth)
	}
	if node.Right != nil {
		rdepth = getDepth1(node.Right, depth)
	}
	if ldepth == 0 { // 左子树为0 说明没有左子树
		return rdepth
	} else if rdepth == 0 { //右子树为0，说明没有右子树
		return ldepth
	} else if ldepth > rdepth { // 都不为0 取小
		return rdepth
	} else {
		return ldepth
	}
}

func getDepth2(node *TreeNode, depth int) int {
	// depth 为node 节点的深度
	fmt.Println("depth: ", depth)
	if node.Left == nil && node.Right == nil {
		return depth
	}
	depth++
	if node.Left == nil {
		return getDepth2(node.Right, depth)
	}
	if node.Right == nil {
		return getDepth2(node.Left, depth)
	}
	var ldepth, rdepth int
	ldepth = getDepth1(node.Left, depth)
	rdepth = getDepth1(node.Right, depth)
	if ldepth > rdepth { // 都不为0 取小
		return rdepth
	} else {
		return ldepth
	}
	return depth
}

func getChildren(depth int, nodes []*TreeNode) ([]*TreeNode, int) {
	var eNodes, cNodes []*TreeNode
	for index := range nodes {
		node := nodes[index]
		fmt.Println(node)
		if node.Left == nil && node.Right == nil {
			return eNodes, depth
		}
		if node.Left != nil {
			cNodes = append(cNodes, node.Left)
		}
		if node.Right != nil {
			cNodes = append(cNodes, node.Right)
		}
	}
	if len(cNodes) > 0 {
		depth++
	}
	return cNodes, depth
}

func minDepth(root *TreeNode) int {
	// 1. 按层遍历，找到第一个叶子节点
	// 2. 递归？ 还没找到方法
	if root == nil {
		return 0
	}
	return getDepth1(root, 1)
	// var nodes []*TreeNode
	// var depth int = 1
	// nodes = append(nodes, root)
	// for len(nodes) > 0 {
	// 	fmt.Println("nodes", nodes)
	// 	nodes, depth = getChildren(depth, nodes)
	// 	if len(nodes) == 0 {
	// 		return depth
	// 	}

	// }
	// return depth
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
// 	fmt.Println(minDepth(proot))
// 	pleft := &TreeNode{Val: 2}
// 	fmt.Println(minDepth(proot))
// 	pright := &TreeNode{Val: 3}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	fmt.Println(minDepth(proot))
// 	qleft := &TreeNode{Val: 5}
// 	qright := &TreeNode{Val: 6}
// 	pright.Left = qleft
// 	fmt.Println(minDepth(proot))
// 	qleft.Right = qright
// 	fmt.Println(minDepth(proot))
// 	fmt.Println(" 1 2 3 4 5")

// 	/*
// 				   1
// 				  /
// 				 2
// 				/
// 			   3
// 			  /
// 			 4
// 			/
// 		   5
// 	*/
// 	root := &TreeNode{Val: 0}
// 	left1 := &TreeNode{Val: 1}
// 	root.Left = left1
// 	left2 := &TreeNode{Val: 2}
// 	left1.Left = left2
// 	fmt.Println(minDepth(proot))
// 	left3 := &TreeNode{Val: 3}
// 	left2.Left = left3
// 	fmt.Println(minDepth(root))
// 	left4 := &TreeNode{Val: 4}
// 	left3.Left = left4
// 	fmt.Println(minDepth(root))
// }
