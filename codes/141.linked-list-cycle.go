package codes

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 *
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (35.93%)
 * Total Accepted:    369.9K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, determine if it has a cycle in it.
 *
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [3,2,0,-4], pos = 1
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 *
 *
 * https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2], pos = 0
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 *
 *
 *https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png
 *
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1], pos = -1
 * Output: false
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 *
 *
 *
 *
 * Follow up:
 *
 * Can you solve it using O(1) (i.e. constant) memory?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func useSet(head *ListNode) bool {
	var nodeMap map[*ListNode]int
	nodeMap = make(map[*ListNode]int)
	for head != nil {
		_, ok := nodeMap[head]
		if ok {
			return true
		}
		nodeMap[head] = 1
		head = head.Next
	}
	return false
}

func useFastSlow(head *ListNode) bool {
	fastNode, slowNode := head, head
	first := 0
	for fastNode != nil && slowNode != nil {
		if first > 0 && fastNode == slowNode {
			return true
		}
		first = 1
		slowNode = slowNode.Next
		fastNode = fastNode.Next
		if fastNode == nil {
			return false
		}
		fastNode = fastNode.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	return useFastSlow(head)
}
