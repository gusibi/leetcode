package codes

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	divMod := func(m, n int) (int, int) {
		return m / n, m % n
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	val, m := 0, 0 // 值和进位
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		// fmt.Println("l1: ", l1, l1==nil, "l2: ", l2, l2==nil)
		if l1 != nil && l2 != nil {
			m, val = divMod((l1.Val + l2.Val + m), 10)
			// fmt.Println("1: ", m, val)
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil && l2 != nil {
			m, val = divMod((l2.Val + m), 10)
			// fmt.Println("2: ", m, val)
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			m, val = divMod((l1.Val + m), 10)
			// fmt.Println("3: ", m, val)
			l1 = l1.Next
		}
		// fmt.Println("val", val)
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	if m > 0 {
		node := &ListNode{Val: m}
		tail.Next = node
	}
	return head
}
