package main

import "fmt"

func main() {
	// given a string, reverse it use only one loop and no extra space
	s := "hello world"
	runes := []rune(s) // convert string to rune slice for mutability

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // swap characters
	}

	reversed := string(runes) // convert back to string
	fmt.Println(reversed)     // Output: "dlrow olleh"
}
