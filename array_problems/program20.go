package main

import "fmt"

// The majority element is the element that appears more than n/2 times in the array
func main() {
	arr := []int{2, 2, 2, 3, 3, 1, 2, 2}
	fmt.Println(Solution(arr))
	fmt.Println(majorityElement(arr))
	fmt.Println(optimizedMajorityElement(arr))
}

// majority element
func Solution(arr []int) int {
	target := -1
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count > len(arr)/2 {
			target = arr[i]
			break
		}
	}
	return target
}

// better solution
func majorityElement(nums []int) int {
	mpp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mpp[nums[i]]; ok {
			mpp[nums[i]] = mpp[nums[i]] + 1
		} else {
			mpp[nums[i]] = 1
		}
	}
	for key, value := range mpp {
		if value > len(nums)/2 {
			return key
		}
	}
	return -1
}

// mores voting algorithm
func optimizedMajorityElement(nums []int) int {
	count := 0
	majorityElement := -1
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count = 1
			majorityElement = nums[i]
		}
		if nums[i] == majorityElement {
			count++
		} else {
			count--
		}
	}
	return majorityElement
}
