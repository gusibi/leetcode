package codes

// import "fmt"
/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 */

type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

/** Initialize your data structure here. */
// func Constructor() MyLinkedList {
// 	return MyLinkedList{Size: 0}
// }

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index > this.Size-1 {
		return -1
	} else if index <= 0 {
		return this.Head.Val
	}
	head := this.Head
	var curr *Node
	for i := 0; i <= index; i++ {
		if i == 0 {
			curr = head
			continue
		}
		if curr != nil {
			curr = curr.Next
		} else {
			return -1
		}
	}
	if curr != nil {
		return curr.Val
	} else {
		return -1
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Head = &Node{Val: val, Next: this.Head}
	if this.Size == 0 {
		this.Tail = this.Head
	}
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newTail := &Node{Val: val}
	if this.Size == 0 {
		this.Head, this.Tail = newTail, newTail
	} else {
		this.Tail, this.Tail.Next = newTail, newTail
	}
	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 如果index 是负值，则 index = size + index
	// 比如：size=5，index=-1，则真正的index=4
	switch {
	case index > this.Size:
		return
	case index == this.Size:
		this.AddAtTail(val)
		return
	case index <= 0:
		this.AddAtHead(val)
		return
	}
	curr := this.Head
	for i := 1; i < index; i++ {
		if curr != nil {
			curr = curr.Next
		}
	}
	newNode := &Node{Val: val}
	newNode.Next = curr.Next
	curr.Next = newNode
	this.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	switch {
	case index >= this.Size || index < 0:
		return
	case this.Size == 1:
		this.Head, this.Tail = nil, nil
		return
	case index == 0:
		this.Head.Next = nil
		this.Head = this.Head.Next
	default:
		curr := this.Head
		for i := 1; i < index; i++ {
			if curr != nil {
				curr = curr.Next
			}
		}
		curr.Next = curr.Next.Next
	}
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// func main(){
// 	obj := Constructor()
// 	param_1 := obj.Get(1)
// 	fmt.Println("param_1: -1==", param_1)
// 	// obj.AddAtHead(10)
// 	// param_2 := obj.Get(0)
// 	// fmt.Println("param_2: 10==", param_2)
// 	// obj.AddAtTail(11)
// 	// param_3 := obj.Get(1)
// 	// fmt.Println("param_3: 11==", param_3)
// 	obj.AddAtIndex(-1,0)
// 	param_4 := obj.Get(-1)
// 	fmt.Println("param_4: 12==", param_4)
// 	fmt.Println(obj.Size)
// 	obj.DeleteAtIndex(0)
// 	param_5 := obj.Get(1)
// 	fmt.Println("param_5: 11==", param_5)
// }
