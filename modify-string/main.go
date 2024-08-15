package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(ModifyString("o1l2le334h"))
	fmt.Println(ModifyString("devreser"))
	fmt.Println(ModifyString("89yek23nom"))
}

func ModifyString(str string) string {
	strSplit := strings.Split(str, "")

	var (
		strs []string
		re   = regexp.MustCompile(`^[0-9]+$`)
	)

	for _, v := range strSplit {
		if !re.MatchString(v) {
			strs = append(strs, v)
		}
	}

	return reverseString(strings.Join(strs, ""))
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
