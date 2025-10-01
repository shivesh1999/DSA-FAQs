package main

import "fmt"

// linkedlist
type LL struct {
	data int
	next *LL
}

func deleteHead(n *LL) *LL {
	if n == nil {
		return nil
	}
	n = n.next
	return n
}

func deleteTail(n *LL) *LL {
	if n == nil || n.next == nil {
		return nil
	}
	temp := n
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return n
}

func convertArraytoLL(arr []int) *LL {
	if len(arr) == 0 {
		return nil
	}
	head := &LL{data: arr[0], next: nil}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &LL{data: arr[i], next: nil}
		current.next = newNode
		current = newNode
	}
	return head
}

func traversetheLL(head *LL) {
	if head == nil {
		fmt.Println("Linked List is empty")
		return
	}
	// never change the value of head in any LL Problem
	// always make a new variable to traverse the linked list
	current := head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// delete a node from linked list and insert a node in linked list
func main() {
	arr := []int{1, 7, 11, 4, 9}
	ll := convertArraytoLL(arr)
	fmt.Println("Original list:")
	traversetheLL(ll)

	fmt.Println("After deleting tail:")
	ll = deleteTail(ll)
	traversetheLL(ll)

}
