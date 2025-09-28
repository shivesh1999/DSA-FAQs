package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 7, 9}
	largest := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	fmt.Println("largest element : ", largest)
}
