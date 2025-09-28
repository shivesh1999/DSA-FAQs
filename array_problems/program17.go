package main

import "fmt"

// this function finds the longest subarray with a given sum k
func main() {
	arr := []int{1, 2, 3, 1, 1, 1, 1, 4, 2, 3}
	fmt.Println(longestSubarrayWithSumK(arr, 3))
	fmt.Println(betterLongestSubarrayWithSumK(arr, 5))
}

func betterLongestSubarrayWithSumK(arr []int, k int) int {
	maxLen := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == k {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}
	return maxLen
}

func longestSubarrayWithSumK(arr []int, target int) int {
	maxLen := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			if sum == target {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}
	return maxLen
}
