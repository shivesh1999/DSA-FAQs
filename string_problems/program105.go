package main

import "fmt"

// pallindrome string
func main() {
	s := "acdcdca"
	result := isPalindrome(s)
	fmt.Println(result) // Output: true
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false // Mismatched characters
		}
		left++
		right--
	}
	return true // All characters matched
}
