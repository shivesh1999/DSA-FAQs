package main

import "fmt"

// maximmum consecutive ones
func main() {
	arr := []int{1, 1, 0, 1, 1, 1, 0, 1, 1}
	counter := 0
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			counter++
		} else {
			if counter > max {
				max = counter
			}
			counter = 0
		}
	}
	fmt.Println(max)
}
