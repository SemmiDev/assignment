package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string) map[string]int {
	// "samMMI sammi SaMMi" -> "sammi sammi sammi"
	sentence = strings.ToLower(sentence)
	// key = word, value = counter
	result := make(map[string]int)
	// "samMMI sammi SaMMi" = ["sammi","sammi","sammi"]
	words := strings.Split(sentence, " ")

	for _, v := range words {
		/*
			[sammi] = 1 = loop 1
			[sammi] = 2 = loop 2
			[sammi] = 3 = loop 3
		 */
		result[v]++
	}

	return result
}

func main() {
	fmt.Println(countWords("aku sUka coklat setiap hari minum susu Aku juga suka makan nasi goreng Goreng setiap"))
}
