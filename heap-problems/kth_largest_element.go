package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	element := findkthLargestElement(arr, 2)
	fmt.Println(element)
}

func findkthLargestElement(arr []int, k int) int {
	min := &MinHeap{}
	heap.Init(min)
	for _, value := range arr {
		heap.Push(min, value)
		if min.Len() > k {
			heap.Pop(min)
		}
	}
	return (*min)[0]
}
