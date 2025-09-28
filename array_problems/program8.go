package main

import (
	"fmt"
	"math"
)

func secondLargest(arr []int) int {
	largest := arr[0]
	secondLargest := math.MinInt
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] < largest && arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}

func secondSmallest(arr []int) int {
	smallest := arr[0]
	secondSmallest := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			secondSmallest = smallest
			smallest = arr[i]
		} else if arr[i] > smallest && arr[i] < secondSmallest {
			secondSmallest = arr[i]
		}
	}
	return secondSmallest
}

// The main function to test the second largest and second smallest functionsW
func main() {
	arr := []int{3, 6, 1, 8, 3, 2, 7}
	fmt.Println("Second smallest number", secondSmallest(arr))
	fmt.Println("Second largest number", secondLargest(arr))
}
