package main

import (
	"fmt"
)

// pallindrome string - ignores spaces and special characters
func main() {
	testCases := []string{
		"acdcdca",
		"A man, a plan, a canal: Panama",
		"race a car",
		"0P",
		".,",
		"a.",
		"12321",
		"12345",
		"a1@b@2@b@1a",   // alphanumeric mixed with special chars in middle
		"a#1#b#c#b#1#a", // another case with alphanumeric in middle
		"a-b-c-b-a",     // letters with hyphens in middle
		"1-2-3-2-1",     // numbers with hyphens
	}

	for _, s := range testCases {
		result := isPalindrome(s)
		fmt.Printf("isPalindrome(\"%s\") = %v\n", s, result)
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters from left
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// Skip non-alphanumeric characters from right
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// Compare lowercase versions of alphanumeric characters
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

// Helper function to check if a character is alphanumeric
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// Helper function to convert to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
