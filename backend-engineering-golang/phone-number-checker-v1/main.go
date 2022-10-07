package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPhoneNumberValid(number *string) bool {
	numberWithCountryCode := strings.HasPrefix(*number, "628")
	numberWithoutCountryCode := strings.HasPrefix(*number, "08")

	if numberWithCountryCode {
		if len(*number) >= 11 {
			*number = (*number)[3:]
			return true
		}
	}

	if numberWithoutCountryCode {
		if len(*number) >= 10 {
			*number = (*number)[2:]
			return true
		}
	}

	return false
}

func checkProvider(number *string) string {
	twoFirstDigit, _ := strconv.Atoi((*number)[0:2])
	switch {
	case twoFirstDigit >= 11 && twoFirstDigit <= 15:
		return "Telkomsel"
	case twoFirstDigit >= 16 && twoFirstDigit <= 19:
		return "Indosat"
	case twoFirstDigit >= 21 && twoFirstDigit <= 23:
		return "XL"
	case twoFirstDigit >= 27 && twoFirstDigit <= 29:
		return "Tri"
	case twoFirstDigit >= 52 && twoFirstDigit <= 53:
		return "AS"
	case twoFirstDigit >= 81 && twoFirstDigit <= 88:
		return "Smartfren"
	default:
		return "invalid"
	}
}

func PhoneNumberChecker(number string, result *string) {
	if !isPhoneNumberValid(&number) {
		*result = "invalid"
		return
	}
	*result = checkProvider(&number)
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "0881111111111111111111111111111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
