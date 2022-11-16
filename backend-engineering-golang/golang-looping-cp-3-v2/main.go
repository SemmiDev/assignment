package main

import "fmt"

var cadelListCharacters = map[rune]struct{}{
	'R': {},
	'r': {},
	'S': {},
	's': {},
	'T': {},
	't': {},
	'Z': {},
	'z': {},
}

func CountingLetter(text string) int {
	var counter int
	for _, v := range text {
		if _, ok := cadelListCharacters[v]; ok {
			counter++
		}
	}
	return counter
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
	fmt.Println(CountingLetter("Zebra Zig Zag"))
}
