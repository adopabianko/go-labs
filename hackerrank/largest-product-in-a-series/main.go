package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// link : https://www.hackerrank.com/contests/guardianhero/challenges/euler008

// Example Input:
// 2
// 4 2
// 1234
// 5 3
// 12345

// Example Output:
// 12
// 60

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		num := readLine(reader)

		result := findMaxProduct(int(n), int(k), num)
		fmt.Println(result)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func findMaxProduct(n, k int, num string) int {
	maxProduct := 0

	for i := 0; i <= n-k; i++ {
		product := 1
		for j := 0; j < k; j++ {
			digit, err := strconv.Atoi(string(num[i+j]))
			if err != nil {
				return 0
			}

			product *= digit
		}

		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct
}
