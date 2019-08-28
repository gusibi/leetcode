package codes

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.03%)
 * Total Accepted:    362.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	getListLength := func(head *ListNode) int {
		length := 0
		for head != nil {
			length++
			head = head.Next
		}
		return length
	}
	// 1. 先判断长度，再找到节点删除
	// 2. 使用两个指针，第一个先前进 n 步, 第二个从头开始，第一个结束时，删掉第二个指针所在位置
	length := getListLength(head)
	if length == n {
		// 如果删第一个
		return head.Next
	}
	index := 1
	m := length - n
	_head := head
	for head != nil {
		if index == m {
			next := head.Next
			head.Next = next.Next
			return _head
		} else {
			head = head.Next
		}
		index++
	}
	return _head
}

// func main() {
// 	var l1Head, tailNode *ListNode
// 	// l1 := [7]int{1, 2, 6, 3, 4, 5, 6}
// 	l1 := [1]int{1}
// 	for i := range l1 {
// 		curNode := ListNode{Val: l1[i]}
// 		if i == 0 {
// 			l1Head = &curNode
// 			tailNode = &curNode
// 		} else {
// 			tailNode.Next = &curNode
// 			tailNode = &curNode
// 		}
// 	}
// 	_head := removeNthFromEnd(l1Head, 1)
// 	for _head != nil {
// 		fmt.Println(_head.Val)
// 		_head = _head.Next
// 	}
// }
