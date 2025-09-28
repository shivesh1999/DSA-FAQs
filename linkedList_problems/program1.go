package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) add(data int) {
	newNode := &Node{data: data, next: nil}
	if n == nil {
		n = newNode
		return
	}
	current := n
	// here we are traversing the linked list to find the last node
	// and then we are adding the new node at the end of the linked list
	for current.next != nil {
		current = current.next
	}
	// here we are adding the new node at the end of the linked list
	current.next = newNode
}

func main() {
	// linkedist is not stored in a contiguous block of memory
	// lnkedlist is dyamic in nature
	// start of linked list is called head
	// last node of linked list is called tail
	// linked list is used in stack, queue, graph, tree, etc.
	// linked list is used in browsers to store history
	head := &Node{data: 1, next: nil}
	head.add(2)
	head.add(3)
	head.add(4)
	head.add(5)
	fmt.Println("Linked List:")
	current := head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
	newHead := deletehead(head)
	traverseLL(newHead)
	arr := []int{1, 7, 11, 4, 9}
	ll := convertArrtoLL(arr)
	fmt.Println("Linked List from Array:", *ll)
	traverseLL(ll)
}

func deletehead(head *Node) *Node {
	if head == nil {
		return nil
	}
	// here we are deleting the head of the linked list
	// and returning the new head of the linked list
	newHead := head.next
	head = nil // optional, to help garbage collection
	return newHead
}

func convertArrtoLL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{data: arr[0], next: nil}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &Node{data: arr[i], next: nil}
		current.next = newNode
		current = newNode
	}
	return head
}

func traverseLL(head *Node) {
	if head == nil {
		fmt.Println("Linked List is empty")
		return
	}
	current := head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// never change the value of head in any LL Problem
// always make a new variable to traverse the linked list
