package main

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 *
 * https://leetcode.com/problems/intersection-of-two-linked-lists/description/
 *
 * algorithms
 * Easy (32.54%)
 * Total Accepted:    278.5K
 * Total Submissions: 856K
 * Testcase Example:  '8\n[4,1,8,4,5]\n[5,0,1,8,4,5]\n2\n3'
 *
 * 编写一个程序，找到两个单链表相交的起始节点。

如下面的两个链表：

https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png

在节点 c1 开始相交。


示例 1：

https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png

    输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
    输出：Reference of the node with value = 8
    输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：

https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png

    输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
    输出：Reference of the node with value = 2
    输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例 3：

https://assets.leetcode.com/uploads/2018/12/13/160_example_3.png

    输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
    输出：null
    输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
    解释：这两个链表不相交，因此返回 null。

注意：

* 如果两个链表没有交点，返回 null.
* 在返回结果后，两个链表仍须保持原有的结构。
* 可假定整个链表结构中没有循环。
* 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
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

func solution1(headA, headB *ListNode) *ListNode {
	/*
		1. 先计算链表长度
		2. 找到两个链表的差值，长的先遍历，短的后遍历
		3. 开始遍历后，对比两个链表节点是否相同，如果相同则说明相交
	*/
	var lengthA int
	var lengthB int
	lengthA = getListLength(headA)
	lengthB = getListLength(headB)
	if lengthA > lengthB {
		diff := lengthA - lengthB
		return findIntersection(headA, headB, diff)
	} else {
		diff := lengthB - lengthA
		return findIntersection(headB, headA, diff)
	}
	return nil
}

func findIntersection(headA, headB *ListNode, diff int) *ListNode {
	index := 0
	for headA != nil {
		if index >= diff && headA == headB {
			// fmt.Println(headA, headB)
			return headA
		}
		headA = headA.Next
		index++
		if index > diff {
			headB = headB.Next
		}
	}
	return nil
}

func getListLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func solution2(headA, headB *ListNode) *ListNode {
	// 同时遍历 headA 和 headB，一个结束后从另一个的开头开始重新遍历
	// 直到两个都为nil
	// a = [4,1,8,4,5]
	// b = [5,0,1,8,4,5]
	// a 遍历顺序为 4 1 8 4 5 5 0 1 8 4 5
	// b 遍历顺序为 5 0 1 8 4 5 4 1 8 4 5
	// 如果有交点，总会相遇
	_headA, _headB := headA, headB
	if headA == nil || headB == nil {
		return nil
	}
	for headA != headB {
		if headA != nil {
			headA = headA.Next
		} else {
			headA = _headB
		}
		if headB != nil {
			headB = headB.Next
		} else {
			headB = _headA
		}
	}
	return headA
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return solution2(headA, headB)
}

// func main() {
// 	// a = [4,1,8,4,5]
// 	// b = [5,0,1,8,4,5]
// 	headA := &ListNode{Val: 4}
// 	a1 := &ListNode{Val: 1}
// 	headA.Next = a1
// 	a2 := &ListNode{Val: 8}
// 	a1.Next = a2
// 	a3 := &ListNode{Val: 4}
// 	a2.Next = a3
// 	a4 := &ListNode{Val: 5}
// 	a3.Next = a4
// 	headB := &ListNode{Val: 5}
// 	b1 := &ListNode{Val: 0}
// 	headB.Next = b1
// 	b1.Next = a1
// 	fmt.Println(getIntersectionNode(headA, headB))
// }
