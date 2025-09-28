package main

import "fmt"

// Sort an array of 0s, 1s, and 2s
// The Dutch National Flag problem is a classic problem in computer science and programming.
// It involves sorting an array of 0s, 1s, and 2s in a single pass.
func main() {
	arr := []int{0, 1, 2, 0, 1, 2, 1, 2, 0, 1, 2}
	count0 := 0
	count1 := 0
	count2 := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count0++
		}
		if arr[i] == 1 {
			count1++
		}
		if arr[i] == 2 {
			count2++
		}
	}
	for i := 0; i < count0; i++ {
		arr[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		arr[i] = 1
	}
	for i := count0 + count1; i < count0+count1+count2; i++ {
		arr[i] = 2
	}
	fmt.Println(arr)
	fmt.Println("Solution2: ", Solution2(arr))
}

// deatch national flag algorithm
func Solution2(arr []int) []int {
	low := 0
	mid := 0
	high := len(arr) - 1
	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else if arr[mid] == 2 {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}

// time	complexity O(n) and space complexity O(1) because we are not using any extra space
