package main

import (
	"fmt"
	"unicode"
)

func removeSpaces(str string) string {
	var res string
	for _, v := range str {
		if !unicode.IsSpace(v) {
			res += string(v)
		}
	}
	return res
}

type f func(s rune) bool

func filter(str string, fil f) string {
	var res string
	for _, v := range str {
		if fil(v) {
			res += string(v)
		}
	}
	return res
}

func CountVowelConsonant(str string) (int, int, bool) {
	vowel, conconant, status := 0, 0, false

	str = removeSpaces(str)
	str = filter(str, func(s rune) bool {
		return unicode.IsLetter(s)
	})

	for _, v := range str {
		switch v {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			vowel++
		default:
			conconant++
		}
	}

	if vowel == 0 || conconant == 0 {
		status = true
	}

	return vowel, conconant, status
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
