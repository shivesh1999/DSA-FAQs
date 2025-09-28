package main

import "fmt"

// Program to count digit in a number
func main() {
	fmt.Println("Enter a number")
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("ERROR")
	}
	var count int = 0
	for number > 0 {
		number = number / 10
		count = count + 1
	}
	fmt.Println("Number of digits : ", count)
}
