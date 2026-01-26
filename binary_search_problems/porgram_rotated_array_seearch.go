package main

func main() {
	// Example usage
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	index := searchInRotatedArray(arr, target)
	println(index) // Output: 4
}

func searchInRotatedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		// Check if the left half is sorted
		if arr[left] <= arr[mid] {
			if target >= arr[left] && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right half is sorted
			if target > arr[mid] && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
