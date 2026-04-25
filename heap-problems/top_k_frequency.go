package main

import (
	"container/heap"
)

type Pair struct {
	num  int
	freq int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: min heap
	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, Pair{num, freq})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Step 3: extract result
	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Pair).num)
	}

	return res
}
