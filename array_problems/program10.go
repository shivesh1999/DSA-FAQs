package main

import "fmt"

// Removes duplicates from the array and place it at the begining
func removeDuplicates(arr []int) []int {
	set := make(map[int]struct{})
	index := 0

	// Populate the set with unique values
	for _, value := range arr {
		if _, exists := set[value]; !exists {
			set[value] = struct{}{}
			arr[index] = value
			index++
		}
	}

	return arr
}

// Removes duplicates from the array and place it at the begining
func removeDuplicatesOptimized(arr []int) []int {
	//unique elements are at the start of the array
	i := 0
	for j := 1; j < len(arr); j++ {
		//if current element is not equal to the previous element
		//then we found a unique element
		//place it at the next position of the last found unique element
		//increment the index of the last found unique element
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return arr
}

// The main function to test the removeDuplicates and removeDuplicatesOptimized functions
func main() {
	arr := []int{1, 2, 1, 3, 2, 1, 4, 5}
	arrayWithNoDuplicate := removeDuplicates(arr)
	fmt.Println("array with removed duplicates at the start :", arrayWithNoDuplicate)
}
