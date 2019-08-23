// package main
// import (
// 	"fmt"
// 	"time"
// )
/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
//     Val int
//     Next *ListNode
// }
func insertNode(head, unhead *ListNode) *ListNode{
	newHead := head
	if head == unhead{
		head.Next = nil
	    return newHead
	}
	// fmt.Println("-----------sort....: ", head, unhead)
	pre := &ListNode{}
	curr, newHead := head, head
	for curr!= unhead{
		// fmt.Println(curr, unhead)
		if curr.Val <= unhead.Val{
		    if curr.Next == nil{
		    	curr.Next = unhead
		    	unhead.Next = nil
		    	break
		    }
			pre = curr
			curr = curr.Next
		}else{
			unhead.Next = curr
			if pre.Next != nil{
				pre.Next = unhead
			}else{
				newHead = unhead
			}
			break
		}
	}
	return newHead
}

func insertionSortList(head *ListNode) *ListNode {
	// 分两部分，前边为已排序，后边为未排序
	unhead := head
	// fmt.Println(head, unhead)
	for unhead != nil{
		// fmt.Println("unhead:", unhead)
		next := unhead.Next
		head = insertNode(head, unhead)
		unhead = next
	}
	return head
}

// func main(){
// 	// [-1,5,3,4,0]
// 	n1 := &ListNode{Val: -1}
// 	n2 := &ListNode{Val: 5}
// 	n3 := &ListNode{Val: 3}
// 	n4 := &ListNode{Val: 4}
// 	n5 := &ListNode{Val: 0}
// 	n1.Next = n2
// 	n2.Next = n3
// 	n3.Next = n4
// 	n4.Next = n5
// 	head := insertionSortList(n1)
// 	// fmt.Println(head, head.Next, head.Next.Next, head.Next.Next.Next)
// 	for head!= nil{
// 		fmt.Println(head)
// 		head = head.Next
// 		time.Sleep(time.Duration(1)*time.Second)
// 	}
// }

