package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
// Left rotates the array by one
func leftRotateArrayByOne(arr []int) []int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
	return arr
}

// time complexity: O(n+d)
// space complexity: O(d)
func leftRotateArrayByN(arr []int, n int) []int {
	fmt.Println(arr)
	n = n % len(arr)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = arr[i]
	}

	for i := n; i < len(arr); i++ {
		arr[i-n] = arr[i]
		fmt.Println(arr)
	}
	for i := len(arr) - n; i < len(arr); i++ {
		arr[i] = temp[i-(len(arr)-n)]
		fmt.Println(arr)
	}
	return arr
}

func reverseArray(arr []int, startIndex int, endIndex int) []int {
	start := startIndex
	end := endIndex

	for start < end {
		// Swap elements
		arr[start], arr[end] = arr[end], arr[start]
		// Move pointers
		start++
		end--
	}

	return arr
}

func leftRotateArrayByNOptimized(arr []int, n int) []int {
	n = n % len(arr)
	reverseArray(arr, 0, n)
	reverseArray(arr, n+1, len(arr)-1)
	reverseArray(arr, 0, len(arr)-1)
	return arr
}

func rotate(nums []int, k int) []int {
	k = k % len(nums)
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[i]
	}
	for i := k; i < len(nums); i++ {
		nums[i-k] = nums[i]
	}
	for i := len(nums) - k; i < len(nums); i++ {
		nums[i] = temp[i-(len(nums)-k)]
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("left roatated array by 1: ", leftRotateArrayByOne(arr))
	// fmt.Println("left roatated array by n: ", leftRotateArrayByN(arr, 3))
	// fmt.Println("left roatated array by n: ", rotate(arr, 3))
	fmt.Println("left roatated array by n: ", leftRotateArrayByNOptimized(arr, 3))
}
