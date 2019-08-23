// package main
// import "fmt"
/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// nlog(n) 归并排序和快速排序
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// func partition(A []int, p, r int) int{
// 	privot := A[r]
// 	i := p
// 	for j:=p; j<=r-1; j++{
// 		if A[j] < privot{
// 			A[i], A[j] = A[j], A[i]
// 			i++
// 		}
// 	}
// 	A[i], A[r] = A[r], A[i]
// 	return i
// }



// func quickSort(A []int, p, r int) []int {
// 	if p >= r {
// 		return A
// 	}
// 	q := partition(A, p, r) // 获取分区点
// 	quickSort(A, p, q-1)
// 	quickSort(A, q+1, r)
// 	return A
// }

// func mergeSort(A []int)[]int{
// 	if len(A) <=1{
// 		return A
// 	}
// 	mid := len(A) / 2
// 	left := mergeSort(A[:mid])
// 	right := mergeSort(A[mid:])
// 	result := merge(left, right)
// 	return result
// }

// func merge(left, right []int) []int{
// 	results := []int{}
//     i, j:=0, 0
// 	for i<len(left) && j<len(right) {
// 		if left[i] < right[j]{
// 			results = append(results, left[i])
// 			i++
// 		}else{
// 			results = append(results, right[j])
// 			j++
// 		}
// 	}
// 	results = append(results, left[i:]...)
// 	results = append(results, right[j:]...)
// 	return results
// }

func mergeList(left, right *ListNode) *ListNode{
	var curr, p, head *ListNode
	for left != nil && right != nil{
		p = curr
		if left.Val <= right.Val{
			curr = left
			left = left.Next
		}else{
			curr = right
			right = right.Next
		}
		if head == nil{
			head = curr
		}
		if p!=nil{
			p.Next = curr
		}
	}
	if left != nil{
		curr.Next = left
	}else if right != nil{
		curr.Next = right
	}
	return head
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	sp := &ListNode{}
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		// fmt.Println("fast: ", fast, "slow: ", slow)
		fast = fast.Next.Next
		sp = slow
		slow = slow.Next
	}
	mid := slow
	sp.Next = nil
	// fmt.Println("mid", mid, "sp: ", sp)
	left := sortList(head)
	right := sortList(mid)
	results := mergeList(left, right)
	return results
}

// func createList(a []int) *ListNode{
// 	head := &ListNode{Val: a[0]}
// 	curr := head
// 	for i:=1; i<len(a);i++{
// 		node := &ListNode{Val: a[i]}
// 		curr.Next = node
// 		curr = node
// 	}
// 	return head
// }

// func printList(head *ListNode){
// 	for head != nil{
// 		fmt.Println(head)
// 		head = head.Next
// 	}
// }

// func main(){
// 	a := []int{4, 3, 6, 5}
// 	head := createList(a)
// 	printList(head)
// 	sortList(head)
// 	fmt.Println(quickSort(a, 0, 3))
// 	fmt.Println(mergeSort(a))
// 	b := []int{4, 3, 6, 5, 10, 8, 7, 6}
// 	head = createList(b)
// 	shead := sortList(head)
// 	printList(shead)
// }

