package main

import "errors"

type stack struct {
	items []int
	top   int
}

func newStack() *stack {
	return &stack{
		items: make([]int, 3),
		top:   -1,
	}
}

func (s *stack) push(item int) error {
	if s.top == len(s.items)-1 {
		return errors.New("stack overflow")
	}
	s.top++
	s.items[s.top] = item
	return nil
}

func (s *stack) pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("stack underflow")
	}
	item := s.items[s.top]
	s.top--
	return item, nil
}

func (s *stack) peek() (int, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	return s.items[s.top], nil
}

func isempty(s *stack) bool {
	return s.top == -1
}

func size(s *stack) int {
	return s.top + 1
}

func main() {
	stack := newStack()
	stack.push(10)
	stack.push(20)
	stack.push(30)

	item, err := stack.pop()
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Popped item:", item)
	}

	topItem, err := stack.peek()
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Top item:", topItem)
	}

	println("Is stack empty?", isempty(stack))
	println("Stack size:", size(stack))
}
