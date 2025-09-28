package main

// longest palindromic substring
func main() {
	s := "babad"
	result := longestPalindrome(s)
	println(result) // Output: "bab" or "aba"
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < n; i++ {
		len1 := expandAroundCenter(s, i, i)   // Odd length palindromes
		len2 := expandAroundCenter(s, i, i+1) // Even length palindromes
		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1 // Length of the palindrome
}
