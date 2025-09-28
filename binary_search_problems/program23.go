package main

import "fmt"

// search in 2d array
func main() {
	//declare a 2darray
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	fmt.Println(search2D(arr, 7))                // Output: true;
	fmt.Println(search2DOptimized(arr, 7))       // Output: true;
	fmt.Println(search2DHighlyOptimized(arr, 7)) // Output: true;
}

func search2D(arr [][]int, target int) bool {
	//interate the number of rows and columns
	fmt.Println("number of rows:", len(arr))
	fmt.Println("number of columns:", len(arr[0]))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == target {
				return true //if found return true
			}
		}
	}
	return false //if not found return false
}

func search2DOptimized(arr [][]int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i][0] <= target && arr[i][len(arr[i])-1] >= target {
			return binarySearch1D(arr[i], target)
		}
	}
	return false //if not found return false
}

func binarySearch1D(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true //if found return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false //if not found return false
}

func search2DHighlyOptimized(arr [][]int, target int) bool {
	low := 0
	high := len(arr)*len(arr[0]) - 1
	for low <= high {
		mid := (low + high) / 2
		row := mid / len(arr[0]) // calculate row index
		col := mid % len(arr[0]) // calculate column index
		if arr[row][col] == target {
			return true //if found return true
		} else if arr[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false //if not found return false
}
