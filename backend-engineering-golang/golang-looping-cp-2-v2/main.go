package main

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	splitS := strings.Split(string(runes), " ")
	result := make([]string, 0, len(splitS))

	for _, v := range strings.Split(string(runes), " ") {
		splitV := strings.Join(strings.Split(v, ""), "_")
		result = append(result, splitV, " ")
	}

	return strings.TrimSuffix(strings.Join(result, ""), " ")
}

// gunakan untuk melakukan debug
func main() {
	a := ReverseString("Hello World")
	fmt.Println(a)
}
