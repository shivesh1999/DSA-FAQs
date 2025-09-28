package main

import "math"

// maximum subarray sum(kadanes algorihm)
func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	sum := 0
	max := math.MinInt
	ansstart := -1
	ansend := -1
	start := 0
	for i := 0; i < len(arr); i++ {
		if sum == 0 {
			start = i
		}
		sum = sum + arr[i]
		if sum > max {
			max = sum
			ansstart = start
			ansend = i
		}
		if sum < 0 {
			sum = 0
		}
	}
	if max < 0 {
		max = 0
	}
	println(max)
	print("subarray is: ")
	for i := ansstart; i <= ansend; i++ {
		print(arr[i], " ")
	}
}
