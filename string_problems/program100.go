package main

import "fmt"

// longest substring - sliding window
// Given a string s, find the length of the longest substring without repeating characters.
// Example: Input: s = "abcabcbb", Output: 3
func main() {
	s := "abcdacdefabc"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	mpp := make(map[rune]int)
	maxLength := 0
	start := 0
	for end, char := range s {
		if index, found := mpp[char]; found && index >= start {
			start = index + 1
		}
		mpp[char] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
