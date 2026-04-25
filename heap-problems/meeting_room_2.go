package main

import (
	"container/heap"
	"fmt"
	"sort"
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
	meetingTimes := [][]int{{0, 10}, {5, 10}, {15, 20}}
	sort.Slice(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][0] < meetingTimes[j][0]
	})
	MinHeap := &MinHeap{}
	heap.Init(MinHeap)
	heap.Push(MinHeap, meetingTimes[0][1])
	for i := 1; i < len(meetingTimes); i++ {
		if meetingTimes[i][0] >= (*MinHeap)[0] {
			heap.Pop(MinHeap)
		}
		heap.Push(MinHeap, meetingTimes[i][1])
	}
	fmt.Println(MinHeap.Len())
}
