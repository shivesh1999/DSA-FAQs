package main

// The two sum problem is to find two numbers in an array that add up to a specific target
func main() {
	arr := []int{2, 6, 5, 8, 11}
	target := 10
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				println("Pair Found: ", i, j)
			}
		}
	}
	println(Solution2(arr, target))
	println(Soelution3(arr, target))
}

// time complexity O(n^2) and space complexity O(n)
func Solution2(arr []int, target int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		first := arr[i]
		secound := target - first
		if _, ok := m[secound]; ok {
			return m[secound], i
		}
		m[first] = i
	}
	return -1, -1
}

// time complexity O(n) and space complexity O(1)
// This solution assumes that the array is sorted
// If the array is not sorted, you can sort it first and then use this solution
func Solution3(arr []int, target int) (int, int) {
	low := 0
	high := len(arr) - 1
	for low < high {
		if arr[low]+arr[high] == target {
			return low, high
		} else if arr[low]+arr[high] < target {
			low++
		} else {
			high--
		}
	}
	return -1, -1
}
