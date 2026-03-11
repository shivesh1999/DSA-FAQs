package main

import "fmt"

func main() {

	arr := []int{1, 25, 3, 43654, 21, 5, 0}

	n := len(arr)
	nge := make([]int, n)

	stack := []int{}

	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			nge[i] = -1
		} else {
			nge[i] = stack[len(stack)-1]
		}

		stack = append(stack, arr[i])
	}

	for _, v := range nge {
		fmt.Println(v)
	}
}
