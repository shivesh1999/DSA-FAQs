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

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// an array is not heapify yet
	arrayForInit := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	arrayForPush := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}

	// apply Init mehtod
	minHeap := buildHeapByInit(arrayForInit)

	// apply Push mehtod
	buildHeapByPush(arrayForPush)
	fmt.Println()

	// apply Pop method: remove and return the minimum element
	elementToRemove := heap.Pop(minHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)
	fmt.Println()

	// apply Fix method
	fmt.Println("Current min heap: ", *minHeap)
	// change number from 15 t0 50 at the index 2
	(*minHeap)[2] = 50
	fmt.Println("Modify min heap value at index 2: ", *minHeap)
	heap.Fix(minHeap, 2)
	fmt.Println("Fix min heap: ", *minHeap)
	fmt.Println()

	// apply Remove method
	fmt.Println("Current min heap: ", *minHeap)
	// remove number at index 4
	heap.Remove(minHeap, 4)
	fmt.Println("Current min heap after remove element in index 4: ", *minHeap)
}

// Time: O(n)
func buildHeapByInit(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	*minHeap = array
	heap.Init(minHeap)
	fmt.Println("buildHeapByInit: ", *minHeap)
	return minHeap
}

// Time: O(nlogn)
func buildHeapByPush(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	for _, num := range array {
		heap.Push(minHeap, num)
	}
	fmt.Println("buildHeapByPush: ", *minHeap)
	return minHeap
}
