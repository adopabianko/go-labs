package main

import "fmt"

func main() {

	var input string
	fmt.Print("Masukkan simbol: ")
	fmt.Scan(&input)

	result := isValidString(input)

	fmt.Println(result)
}

func isValidString(input string) bool {
	if len(input) < 1 || len(input) > 4096 {
		return false
	}

	stack := []string{}

	matchingBracket := map[string]string{
		">": "<",
		"}": "{",
		"]": "[",
	}

	for _, char := range input {
		switch char {
		case '<', '{', '[':
			stack = append(stack, string(char))
		case '>', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matchingBracket[string(char)] {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}
