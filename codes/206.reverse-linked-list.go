/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	var preNode, nhead *ListNode
// 	for head != nil{
// 		next := head.Next
// 		nhead = head
// 		nhead.Next = preNode
// 		preNode = head
// 		head = next
// 	}
// 	return nhead
// }

func reverseList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

