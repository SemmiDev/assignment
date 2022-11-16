/*

country code = 628
tanpa country code = 08


// validasi

IF
// [08]123123213103
// first = 08 (tanpa country code)
// first = 62 (country code)

IF
// kalo country code
	// cek kalo panjang >= 10 untuk country code
// kalo without country code
	// cek kalo panjang >= 11 untuk country code


// detect provider nya (tel,xl,indosat,dll)
	08[11]123123781631782312 // tanpa country code
	11 sampai 15 = telkomsel

	// country code
	628[12]3132131231231
	11 sampai 15 = telkomsel

	return provider

*/


package main

import "fmt"

func main() {
	text := "081231312312"
	first := text[2:4] // 2-3

	fmt.Println(first)
}