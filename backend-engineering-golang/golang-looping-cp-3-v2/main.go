package main

import "fmt"

func CountingLetter(text string) int {
	counter := 0
	for _, v := range text {
		switch v {
		case 'R', 'S', 'T', 'Z', 'r', 's', 't', 'z':
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
