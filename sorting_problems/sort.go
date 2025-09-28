package main

import "fmt"

// check minimum and move it to the start
func selectionSort(x []int) []int {
	for i := 0; i <= len(x)-2; i++ {
		min := i
		for j := i; j <= len(x)-1; j++ {
			if x[j] < x[min] {
				min = j
			}
		}
		x[min], x[i] = x[i], x[min]
		fmt.Println("Current Array", x)
	}
	return x
}

// check current element with next element and move it at the end
func bubbbleSort(x []int) []int {
	for i := len(x) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}

// check small arrays from the begining and sort them
func insertionSort(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
	}
	return arr
}

func merge(arr []int, low int, mid int, high int) {
	temp := []int{}
	left := low
	right := mid + 1
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

/*
slection sort, bubble sort and insertion sort takes O(n^2) to sort the array
merge sort - divide and merge
time complexity - log2n
space complexity - log(n)
*/
func mergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low
	j := high
	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}
		for arr[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

/*
quick sort -
pick a pivot - any element of the array
smaller on the left, greater on the right
divide and conquer
time complexity n log N
space complexity o(1)
*/
func quickSort(arr []int, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quickSort(arr, low, partitionIndex)
		quickSort(arr, partitionIndex+1, high)
	}
}

// sorting techniques
func main() {
	x := []int{5, 2, 4, 1, 3}
	mergeSort(x, 0, len(x)-1)
	fmt.Println("Sorted Array", x)
}
