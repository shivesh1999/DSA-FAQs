package main

import "fmt"

func checkSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
		} else {
			return false
		}
	}
	return true
}

func checkRotatedSortedArray(arr []int) bool {
	i := 1
	for ; i < len(arr) && arr[i-1] <= arr[i]; i++ {
	}
	if i == len(arr) {
		return true
	}
	for i++; i < len(arr) && arr[i-1] <= arr[i]; i++ {
	}

	return i == len(arr) && arr[len(arr)-1] <= arr[0]
}

// The main function to test the checkSorted and checkRotatedSortedArray functions
func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println("The array is sorted :", checkSorted(arr))
	fmt.Println("The array is a rotated sorted array :", checkRotatedSortedArray(arr))
}
