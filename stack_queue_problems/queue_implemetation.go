package main

import (
	"errors"
	"fmt"
)

type queue struct {
	items []int
	front int
	rear  int
	size  int
}

func newQueue(capacity int) *queue {
	return &queue{
		items: make([]int, capacity),
		front: 0,
		rear:  -1,
		size:  0,
	}
}

func (q *queue) enqueue(item int) error {
	if q.size == len(q.items) {
		return errors.New("queue overflow")
	}
	q.rear = (q.rear + 1) % len(q.items)
	q.items[q.rear] = item
	q.size++
	return nil
}

func (q *queue) dequeue() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue underflow")
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % len(q.items)
	q.size--
	return item, nil
}

func (q *queue) peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items[q.front], nil
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (q *queue) getSize() int {
	return q.size
}

func main() {
	queue := newQueue(3)
	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)

	item, err := queue.dequeue()
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Dequeued item:", item)
	}

	fmt.Println(queue.isEmpty())
	fmt.Println(queue.getSize())
}
