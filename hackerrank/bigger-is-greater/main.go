package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// link : https://www.hackerrank.com/contests/guardianhero/challenges/bigger-is-greater

/** example 1 **/
/** input **/
// 5
// ab
// bb
// hefg
// dhck
// dkhc
/** output **/
// ba
// no answer
// hegf
// dhkc
// hcdk

/** example 2 **/
/** input **/
// lmno
// dcba
// dcbb
// abdc
// abcd
// fedcbabcd
/** output **/
// lmon
// no answer
// no answer
// acbd
// abdc
// fedcbabdc

func biggerIsGreater(w string) string {
	// Write your code here
	runes := []rune(w)
	n := len(runes)

	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}

	if i < 0 {
		return "no answer"
	}

	j := n - 1
	for runes[j] <= runes[i] {
		j--
	}

	runes[i], runes[j] = runes[j], runes[i]

	left, right := i+1, n-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
