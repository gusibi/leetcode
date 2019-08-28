package codes

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (35.42%)
 * Total Accepted:    236.2K
 * Total Submissions: 666.9K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
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

func isPalindrome(head *ListNode) bool {
	// 使用快慢指针定位中间节点
	first, middle := true, false
	fastNode, slowNode := head, head
	var backNode *ListNode
	for slowNode != nil {
		curNode := &ListNode{Val: slowNode.Val}
		if first {
			backNode = curNode
		}
		// fmt.Println("current node: ", slowNode)
		// slowNode = slowNode.Next
		if middle {
			if backNode.Val != slowNode.Val {
				return false
			}
			slowNode = slowNode.Next
			backNode = backNode.Next
		} else {
			if !first {
				curNode.Next = backNode
				backNode = curNode
			}
			fastNode = fastNode.Next
			if fastNode == nil {
				middle = true
				continue
			}
			fastNode = fastNode.Next
			if fastNode == nil {
				middle = true
			}
			slowNode = slowNode.Next
		}
		first = false
	}
	return true
}

// func main() {
// 	head := ListNode{Val: 1}
// 	// second := ListNode{Val: 1}
// 	// head.Next = &second
// 	// tail := ListNode{Val: 1}
// 	// second.Next = &tail
// 	fmt.Println(head, head.Next)
// 	fmt.Println(isPalindrome(&head))
// }
