package codes

/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (35.34%)
 * Total Accepted:    209.8K
 * Total Submissions: 593.7K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
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

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeElements(head *ListNode, val int) *ListNode {
	var preNode *ListNode
	_head := head
	for head != nil {
		if head.Val == val {
			next := head.Next
			if preNode != nil {
				preNode.Next = next
			} else {
				if next == nil {
					return nil
				} else {
					// 如果第一个节点需要删除
					_head = next
				}
			}
		} else {
			preNode = head
		}
		head = head.Next
	}
	return _head
}

// func main() {
// 	var l1Head, tailNode *ListNode
// 	// l1 := [7]int{1, 2, 6, 3, 4, 5, 6}
// 	// l1 := [2]int{1, 2}
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
// 	_head := removeElements(l1Head, 1)
// 	for _head != nil {
// 		fmt.Println(_head.Val)
// 		_head = _head.Next
// 	}

// }
