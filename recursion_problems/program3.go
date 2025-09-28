package main

import "fmt"

func main() {
	fmt.Println("Recursion example")
	couting(10)
}

func couting(number int) {
	// condition to stop the recursion from a stack overflow
	if number == 0 {
		return
	}
	couting(number - 1)
	println(number)
}
