package main

import "fmt"

// Program to reverse a number
func main() {
	fmt.Println("Enter a number")
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("ERROR")
	}
	var reverse int
	var lastDigit int
	for number > 0 {
		lastDigit = number % 10
		reverse = reverse*10 + lastDigit
		number = number / 10
	}
	fmt.Println("Reverse of the number is : ", reverse)
}
