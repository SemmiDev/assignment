package main

import (
	"fmt"
	"strings"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	
	// reversed string
	var result string

	// split per kata
	words := strings.Split(str, " ")

	for i := len(words) - 1; i >= 0; i-- {
	// [world] [hello]
	//    1       0
		
		// reverse per word, cuman smpe 1
		// len 5 = world
		for j := len(words[i]) - 1; j >= 1; j-- {
			// d_l_r_o_w
			result += string(words[i][j]) + "_"
		}

		result += string(words[i][0])
		// make sure setelah kata terakhir ga ada space
		if i != 0 {
			result += " "
		}
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
