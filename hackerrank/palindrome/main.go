package main

import "fmt"

func main() {
	str := "katak"
	fmt.Println("Golang program to check palindrome,\n Given Word =", str)

	result := isPalindrome(str)

	fmt.Printf("'%s' '%s'\n", result, str)
}

func isPalindrome(str string) string {
	var a []byte

	for i := 0; i < len(str); i++ {
		a = append(a, str[len(str)-1-i])
	}

	if string(a) == str {
		return "is palindrome"
	}

	return "is not paliondrome"
}
