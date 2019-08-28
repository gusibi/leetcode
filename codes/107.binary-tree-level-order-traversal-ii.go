package codes

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (45.80%)
 * Total Accepted:    213.1K
 * Total Submissions: 464.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 * 按层遍历
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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

func getChildren(nodes []*TreeNode) []*TreeNode {
	var cNodes []*TreeNode
	for index := range nodes {
		node := nodes[index]
		if node.Left != nil {
			cNodes = append(cNodes, node.Left)
		}
		if node.Right != nil {
			cNodes = append(cNodes, node.Right)
		}
	}
	return cNodes
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	var levelNodes [][]*TreeNode
	var nodes []*TreeNode
	nodes = append(nodes, root)
	levelNodes = append(levelNodes, nodes)
	for len(nodes) > 0 {
		nodes = getChildren(nodes)
		if len(nodes) > 0 {
			levelNodes = append(levelNodes, nodes)
		}

	}

	levelsCount := len(levelNodes)
	var results [][]int = make([][]int, levelsCount)
	for i := levelsCount - 1; i >= 0; i-- {
		children := levelNodes[i]
		values := make([]int, len(children))
		// 遍历子一层节点
		for index := range children {
			values[index] = children[index].Val
		}
		results[levelsCount-i-1] = values
	}
	return results
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
// 	var root *TreeNode
// 	fmt.Println(levelOrderBottom(root))
// 	proot := &TreeNode{Val: 1}
// 	fmt.Println(levelOrderBottom(proot))
// 	pleft := &TreeNode{Val: 2}
// 	fmt.Println(levelOrderBottom(proot))
// 	pright := &TreeNode{Val: 3}
// 	proot.Left = pleft
// 	proot.Right = pright
// 	fmt.Println(levelOrderBottom(proot))
// 	qleft := &TreeNode{Val: 5}
// 	qright := &TreeNode{Val: 6}
// 	pright.Left = qleft
// 	fmt.Println(levelOrderBottom(proot))
// 	qleft.Right = qright
// 	fmt.Println(levelOrderBottom(proot))
// }
