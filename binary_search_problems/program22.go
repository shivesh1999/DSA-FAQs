package main

// binary search to find x in sorted array
func main() {
	arr := []int{3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	index := binarySearch(arr, target)
	if index != -1 {
		println("Element found at index", index)
	} else {
		println("Element not found")
	}
	low := 0
	high := len(arr) - 1
	index = recursionBinarySearch(arr, target, low, high)
	if index != -1 {
		println("Element found at index", index)
	} else {
		println("Element not found")
	}
}
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// time complexity of binary search is O(log2^n)
func recursionBinarySearch(arr []int, target int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return recursionBinarySearch(arr, target, mid+1, high)
	} else {
		return recursionBinarySearch(arr, target, low, mid-1)
	}
}

// incase of overflow
// mid := low + (high-low)/2
