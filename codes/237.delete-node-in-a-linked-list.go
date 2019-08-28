package codes

/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 *
 * https://leetcode.com/problems/delete-node-in-a-linked-list/description/
 *
 * algorithms
 * Easy (52.22%)
 * Total Accepted:    268.3K
 * Total Submissions: 513.8K
 * Testcase Example:  '[4,5,1,9]\n5'
 *
 * 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。

 * 现有一个链表 -- head = [4,5,1,9]，它可以表示为:

 * https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/01/19/237_example.png
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [4,5,1,9], node = 5
 * Output: [4,1,9]
 * Explanation: You are given the second node with value 5, the linked list
 * should become 4 -> 1 -> 9 after calling your function.
 *
 *
 * Example 2:
 *
 *
 * Input: head = [4,5,1,9], node = 1
 * Output: [4,5,9]
 * Explanation: You are given the third node with value 1, the linked list
 * should become 4 -> 5 -> 9 after calling your function.
 *
 *
 *
 *
 * Note:
 *
 *
 * The linked list will have at least two elements.
 * All of the nodes' values will be unique.
 * The given node will not be the tail and it will always be a valid node of
 * the linked list.
 * Do not return anything from your function.
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	nextNext := next.Next
	if nextNext != nil {
		node.Next = nextNext
	} else {
		node.Next = nil
	}
}
