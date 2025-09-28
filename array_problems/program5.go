package main

import "fmt"

func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		// Swap elements
		arr[start], arr[end] = arr[end], arr[start]
		// Move pointers
		start++
		end--
	}

	return arr
}

func main() {
	var arr = []int{5, 4, 3, 2, 1}
	arr2 := reverseArray(arr)
	fmt.Println("Reversed Array: ", arr2)
}
