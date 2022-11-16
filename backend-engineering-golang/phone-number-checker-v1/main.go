package main

import (
	"strconv"
	"strings"
)

func isPhoneNumberWithCountryCodeValid(number *string) bool {
	if strings.HasPrefix(*number, "628") {
		if len(*number) >= 11 {
			*number = (*number)[3:]
			return true
		}
	}
	return false
}

func isPhoneNumberWithoutCountryCodeValid(number *string) bool {
	if strings.HasPrefix(*number, "08") {
		if len(*number) >= 11 {
			*number = (*number)[2:]
			return true
		}
	}
	return false
}

func validatePhoneNumber(number *string) bool {
	return isPhoneNumberWithCountryCodeValid(number) || isPhoneNumberWithoutCountryCodeValid(number)
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
	*result = "invalid"
	if validatePhoneNumber(&number) {
		*result = checkProvider(&number)
	}
}
