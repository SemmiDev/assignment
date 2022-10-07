package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverse(s string) string {
	var result []string
	for _, char := range s {
		result = append([]string{string(char)}, result...)
	}
	return strings.Join(result, "")
}

func checkFirstAndLastChar(word string) string {
	if len(word) == 1 {
		if unicode.IsUpper(rune(word[0])) {
			return strings.ToUpper(word)
		} else {
			return strings.ToLower(word)
		}
	}
	lastChar := word[len(word)-1]
	firstChar := string(word[0])
	if unicode.IsUpper(rune(lastChar)) {
		firstChar = strings.ToUpper(string(firstChar))
		word = firstChar + word[1:]
		word = word[:len(word)-1] + strings.ToLower(string(lastChar))
	}
	word = word[:1] + strings.ToLower(word[1:len(word)-1]) + word[len(word)-1:]
	return word
}
func ReverseWord(str string) string {
	if str == "" {
		return str
	}
	words := strings.Split(str, " ")
	var result []string
	for _, word := range words {
		// allUpperCase := false
		// for _, char := range word {
		// 	if unicode.IsUpper(char) {
		// 		allUpperCase = true
		// 	} else {
		// 		allUpperCase = false
		// 		break
		// 	}
		// }

		// if allUpperCase {
		// 	result = append(result, strings.ToUpper(reverse(word)))
		// } else {
		result = append(result, checkFirstAndLastChar(reverse(word)))
		// }
	}
	return strings.Join(result, " ")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
	fmt.Println(ReverseWord("KITA SELALU BERSAMA"))
}
