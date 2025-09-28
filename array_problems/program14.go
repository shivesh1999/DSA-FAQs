package main

import "fmt"

// missing number problem
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	n := 9
	fmt.Println("missing number: ", missingNumber(arr, n))
	fmt.Println("missing number: ", missingNumberOptimized(arr, n))
	fmt.Println("missing number: ", missingNumberBestCase(arr, n))
}

// time complexity: O(n)
// space complexity: O(1)
func missingNumberBestCase(arr []int, n int) any {
	sum := n * (n + 1) / 2
	sum2 := 0
	for i := 0; i < len(arr); i++ {
		sum2 += arr[i]
	}
	return sum - sum2
}

// time complexity: O(2n)
// space complexity: O(n)
func missingNumberOptimized(arr []int, n int) int {
	temp := make([]int, n+1)
	for i := 0; i < len(arr); i++ {
		temp[arr[i]] = 1
	}
	for i := 1; i <= n; i++ {
		if temp[i] == 0 {
			return i
		}
	}
	return -1
}

// time complexity: O(n^2)
// space complexity: O(1)
func missingNumber(arr []int, n int) int {
	for i := 1; i <= n; i++ {
		flag := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == i {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return i
		}
	}
	return -1
}
