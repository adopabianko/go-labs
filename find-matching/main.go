package main

import "fmt"

func main() {
	var length int
	fmt.Print("Masukkan jumlah string: ")
	fmt.Scan(&length)

	input := make([]string, length)
	fmt.Println("Masukkan string:")
	for i := 0; i < length; i++ {
		fmt.Scan(&input[i])
	}

	result := findMatching(input, length)

	fmt.Println(result)
}

func findMatching(input []string, size int) interface{} {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			fmt.Println(input[i], input[j])
			if input[i] == input[j] {
				return fmt.Sprintf("%d %d", i+1, j+1)
			}
		}

	}

	return false
}
