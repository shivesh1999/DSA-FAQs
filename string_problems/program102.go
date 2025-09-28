package main

import "fmt"

// gorup anagram - sliding window
// Given an array of strings strs, group the anagrams together.
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	gorupedAnagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		gorupedAnagrams[sortedStr] = append(gorupedAnagrams[sortedStr], str)
	}
	for _, group := range gorupedAnagrams {
		fmt.Println(group)
	}
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
