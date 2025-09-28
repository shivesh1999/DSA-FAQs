package main

import "fmt"

// valid anagram - sliding window
// Given two strings s and t, return true if t is an anagram of s,and false otherwise.
func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	for _, char := range t {
		if counts[char] == 0 {
			return false
		}
		counts[char]--
	}
	return true
}
