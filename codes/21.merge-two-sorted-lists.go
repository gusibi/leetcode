package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (46.08%)
 * Total Accepted:    524.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur, tail *ListNode
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 == nil {
			cur = l1
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			cur = l2
			l2 = l2.Next
		} else if l1.Val <= l2.Val {
			cur = l1
			l1 = l1.Next
		} else {
			cur = l2
			l2 = l2.Next
		}
		if head == nil {
			head = cur
			tail = cur
		} else {
			tail.Next = cur
			tail = cur
		}
		// fmt.Println("head: ", head, "cur", cur, "tail", tail)
	}
	return head
}

// func main() {
// 	var l1Head, l2Head, tailNode *ListNode
// 	l1 := [3]int{1, 2, 4}
// 	l2 := [3]int{1, 3, 4}
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
// 	for i := range l2 {
// 		curNode := ListNode{Val: l2[i]}
// 		if i == 0 {
// 			l2Head = &curNode
// 			tailNode = &curNode
// 		} else {
// 			tailNode.Next = &curNode
// 			tailNode = &curNode
// 		}
// 	}
// 	l := mergeTwoLists(l1Head, l2Head)
// 	fmt.Println(l, l.Next, l.Next.Next, l.Next.Next.Next, l.Next.Next.Next.Next)
// }
