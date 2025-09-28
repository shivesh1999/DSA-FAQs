package main

import "fmt"

func moveZerosToEnd(arr []int) []int {
	temp := make([]int, len(arr))
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp[index] = arr[i]
			index++
		}
	}
	for i := 0; i < len(arr); i++ {
		if i >= len(temp) {
			temp[i] = 0
		}
	}
	return temp
}

func optimizeMoveZerosToEnd(arr []int) []int {
	j := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}
	for i := j + 1; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	return arr
}

// remove zeros to the end to the array
func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println("Array before moving zeros to the end: ", arr)
	fmt.Println("Array after moving zeros to the end: ", moveZerosToEnd(arr))
	fmt.Println("Array before moving zeros to the end: ", optimizeMoveZerosToEnd(arr))
}
