package main

import "fmt"

var cadel = map[string]string{
	"S": "L",
	"R": "L",
	"Z": "L",
	"s": "l",
	"r": "l",
	"z": "l",
}

func SlurredTalk(words *string) {
	for i, v := range *words {
		if replacer, ok := cadel[string(v)]; ok {
			*words = (*words)[:i] + replacer + (*words)[i+1:]
		}
	}
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println()
}
