package main

import "fmt"

func factorial(x int) int {
	var ans int = 1
	for i := 1; i < x; i++ {
		ans = ans * i
	}
	return ans
}

func factorial2(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial2(x-1)
}

func main() {
	var x int = 5
	result := factorial2(x)
	fmt.Println(result)
}
