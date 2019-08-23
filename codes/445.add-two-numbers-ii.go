/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// import "fmt"

func reverseList(head *ListNode) *ListNode {
	var preNode, nhead *ListNode
	for head != nil{
		next := head.Next
		nhead = head
		nhead.Next = preNode
		preNode = head
		head = next
	}
	return nhead
}

func divMod(m, n int) (int, int){
	return m/n, m%n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}
	nl1, nl2 := reverseList(l1), reverseList(l2)
	// fmt.Println(nl1, nl2)
	var m, val int // 进位和值
	var head *ListNode
	for nl1 != nil || nl2 != nil{
		if nl1 != nil && nl2 != nil{
			m, val = divMod((nl1.Val+nl2.Val+m), 10)
			nl1, nl2 = nl1.Next, nl2.Next
		}else if nl1 == nil && nl2 != nil{
			m, val = divMod((nl2.Val+m), 10)
			nl2 = nl2.Next
		}else if nl2 == nil && nl1 != nil{
			m, val = divMod((nl1.Val+m), 10)
			nl1 = nl1.Next
		}
		// fmt.Println("m: ", m, "val: ", val)
		node := &ListNode{Val: val}
		if head == nil{
			head = node
		}else{
			node.Next = head
			head = node
		}
	}
	if m>0{
		node := &ListNode{Val: m}
		node.Next = head
		head = node
	}
	return head
}

