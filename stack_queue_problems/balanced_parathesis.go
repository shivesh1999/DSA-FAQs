package main

import "fmt"

func main() {
	str := "[{}()]"
	fmt.Println(isBalanced(str))
}

func isBalanced(str string) bool {
	var stack []rune
	brakets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, char := range str {
		if open, exists := brakets[char]; exists {
			if len(stack) == 0 && stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
