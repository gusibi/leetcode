package codes

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (49.44%)
 * Total Accepted:    243.7K
 * Total Submissions: 491.8K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。.
 *
 * Example:
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 * 使用二分法，找中间节点，找中间节点的节点
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 1 {
		node := TreeNode{Val: nums[0]}
		return &node
	} else if length == 0 {
		return nil
	}
	mid := int(length / 2)
	node := TreeNode{Val: nums[mid]}
	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])
	node.Left = left
	node.Right = right
	return &node
}
