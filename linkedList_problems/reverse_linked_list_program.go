package main

import "fmt"

type SLL struct {
	value int
	next  *SLL
}

func (l *SLL) add(value int) {
	newNode := &SLL{value: value, next: nil}

	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *SLL) reverse() *SLL {
	current := l
	var prev *SLL
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}

	return prev
}

func main() {
	LL := &SLL{value: 1, next: nil}
	LL.add(2)
	LL.add(3)
	LL.add(4)
	LL.add(5)

	head := LL
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}

	reversed := LL.reverse()

	head = reversed
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}
