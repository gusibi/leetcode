package codes

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (33.23%)
 * Total Accepted:    357K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
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

func mergeLists(first, second *ListNode) *ListNode {
	// 使用归并排序
	var mergedList, tail, next *ListNode
	for first != nil || second != nil {
		if first == nil && second != nil {
			next = second
			second = second.Next
		} else if first != nil && second == nil {
			next = first
			first = first.Next
		} else {
			if first.Val >= second.Val {
				next = second
				second = second.Next
			} else {
				next = first
				first = first.Next
			}
		}
		// fmt.Println("next:", next)
		if mergedList == nil {
			mergedList = next
			tail = mergedList
		} else {
			tail.Next = next
			tail = tail.Next
		}
	}
	// fmt.Println(mergedList)
	return mergedList
}

func buildHeap(lists []*ListNode) []*ListNode {
	nodes := []*ListNode{nil}
	for index := range lists {
		node := lists[index]
		if node != nil {
			nodes = append(nodes, node)
		}
	}
	length := len(nodes)
	// fmt.Println(length, nodes)
	for i := length / 2; i > 0; i-- {
		nodes = heapfiy(nodes, length, i)
	}
	return nodes
}

func heapfiy(nodes []*ListNode, length, index int) []*ListNode {
	for {
		maxPos := index
		if index*2 < length && nodes[index].Val > nodes[index*2].Val {
			// 有左节点 且 父节点大于左节点
			maxPos = index * 2
		}
		if index*2+1 < length && nodes[maxPos].Val > nodes[index*2+1].Val {
			// 有右节点 且 父节点和左右节点中右节点最小
			maxPos = index*2 + 1
		}
		if maxPos == index {
			// 说明当前节点是最小节点
			break
		}
		// 交换节点
		nodes[index], nodes[maxPos] = nodes[maxPos], nodes[index]
		index = maxPos
	}
	return nodes
}

func mergeKListByHeap(lists []*ListNode) *ListNode {
	// 使用每个链表第一个元素创建一个最小堆
	nodes := buildHeap(lists)
	// fmt.Println(nodes)
	// 每次取堆顶元素，然后用堆顶元素的下一个节点补充，然后堆化
	var head, tail *ListNode
	length := len(nodes)
	for length > 1 {
		top := nodes[1]
		if head == nil {
			head = top
			tail = top
		} else {
			tail.Next = top
			tail = top
		}
		if top.Next != nil {
			// 如果有下一个节点，补上
			nodes[1] = top.Next
			heapfiy(nodes, length, 1)
		} else {
			// 如果没有下一个节点，把最后一个叶子节点放到堆顶，然后删除最后一个叶子节点，再堆化
			nodes[1] = nodes[length-1]
			length = length - 1
			nodes = heapfiy(nodes, length, 1)
		}
	}
	return head
}

func mergeKListsBySort(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	var mergedList *ListNode
	for i := 0; i < len(lists)-1; i++ {
		firstList := lists[i]
		secondList := lists[i+1]
		list := mergeLists(firstList, secondList)
		lists[i+1] = list
		mergedList = list
	}
	return mergedList
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListByHeap(lists)
}

// func main() {
// 	l1 := &ListNode{Val: 1}
// 	l11 := &ListNode{Val: 4}
// 	l1.Next = l11
// 	l12 := &ListNode{Val: 5}
// 	l11.Next = l12
// 	l2 := &ListNode{Val: 1}
// 	l21 := &ListNode{Val: 3}
// 	l2.Next = l21
// 	l22 := &ListNode{Val: 4}
// 	l21.Next = l22
// 	l3 := &ListNode{Val: 2}
// 	l31 := &ListNode{Val: 6}
// 	l3.Next = l31
// 	lists := []*ListNode{l1, l2, l3}
// 	mergedList := mergeKLists(lists)
// 	fmt.Println(mergedList)
// 	fmt.Println(mergedList.Next)
// 	fmt.Println(mergedList.Next.Next)
// 	fmt.Println(mergedList.Next.Next.Next)
// 	fmt.Println(mergedList.Next.Next.Next.Next)
// 	fmt.Println(mergedList.Next.Next.Next.Next.Next)
// 	fmt.Println(mergedList.Next.Next.Next.Next.Next.Next)
// 	fmt.Println(mergedList.Next.Next.Next.Next.Next.Next.Next)
// 	var l4 *ListNode
// 	lists1 := []*ListNode{l1, l4}
// 	mergedList1 := mergeKLists(lists1)
// 	fmt.Println(mergedList1)
// }
