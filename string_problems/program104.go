package main

// valid parentheses
func main() {
	s := "([]{})"
	result := isValid(s)
	println(result) // Output: true
}

func isValid(s string) bool {
	stack := []rune{}
	mpp := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, exists := mpp[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false // Mismatched or unbalanced parentheses
			}
			stack = stack[:len(stack)-1] // Pop the last element
		} else {
			stack = append(stack, char) // Push the opening parenthesis onto the stack
		}
	}

	return len(stack) == 0
}
