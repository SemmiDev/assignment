package main

import (
	"fmt"
	"unicode"
)

var vocal = map[string]string{
	"a": "a",
	"i": "i",
	"u": "u",
	"e": "e",
	"o": "o",
	"A": "a",
	"I": "i",
	"U": "u",
	"E": "e",
	"O": "o",
}

var nextVocal = map[string]string{
	"a": "e",
	"A": "E",
	"e": "i",
	"E": "I",
	"i": "o",
	"I": "O",
	"o": "u",
	"O": "U",
	"u": "a",
	"U": "A",
}

func Reverse(str string) string {
	var reversed string
	for _, v := range str {
		reversed = string(v) + reversed
	}
	return reversed
}

func ChangeToNextVocal(str string) string {
	var result string
	for _, v := range str {
		nextVocal, ok := nextVocal[string(v)]
		if ok {
			result += nextVocal
		} else {
			result += string(v)
		}
	}
	return result
}

func removeSpace(str string) string {
	var result string
	for _, v := range str {
		if !unicode.IsSpace(v) {
			result += string(v)
		}
	}
	return result
}

func changeLetterCase(str string) string {
	var result string
	for _, v := range str {
		// cek apakah dia ga huruf vokal
		if _, ok := vocal[string(v)]; !ok {
			// cek apakah hurufnya besar
			if unicode.IsUpper(v) {
				// ubah jadi huruf kecil
				result += string(unicode.ToLower(v))
			} else {
				result += string(unicode.ToUpper(v))
			}
		} else {
			if unicode.IsLower(v) {
				result += string(unicode.ToUpper(v))
			} else {
				result += string(v)
			}
		}
	}

	return result
}

func Generate(str string) string {
	reversedStr := Reverse(str)
	changedToNextVocal := ChangeToNextVocal(reversedStr)
	fmt.Println("----")
	fmt.Println(changedToNextVocal)
	fmt.Println("----")
	changedLetterCase := changeLetterCase(changedToNextVocal)
	removedSpace := removeSpace(changedLetterCase)
	return removedSpace
}

func isMinimumCheck(base int, data int) bool {
	return data >= base
}

func isContainSymbolCheck(str string) bool {
	for _, v := range str {
		if unicode.IsSymbol(v) {
			return true
		}
	}
	return false
}

func isContainPunctCheck(str string) bool {
	for _, v := range str {
		if unicode.IsPunct(v) {
			return true
		}
	}
	return false
}

func isContainsLetterCheck(str string) bool {
	for _, v := range str {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func mustNotContainsLetterCheck(str string) bool {
	for _, v := range str {
		if unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func mustNotContainsNumberCheck(str string) bool {
	for _, v := range str {
		if unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func isContainNumberCheck(str string) bool {
	for _, v := range str {
		if unicode.IsNumber(v) {
			return true
		}
	}
	return false
}

func isAllContainsNumberCheck(str string) bool {
	for _, v := range str {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func isAllContainsLetterCheck(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func isAllJustContainsNumberAndLetterCheck(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func isVeryWeakPasswordStrength(str string) bool {
	return isMinimumCheck(0, len(str))
}

func isWeakPasswordStrength(str string) bool {
	if !isMinimumCheck(7, len(str)) {
		return false
	}

	if isAllContainsNumberCheck(str) {
		return true
	}

	if isAllContainsLetterCheck(str) {
		return true
	}

	if isAllJustContainsNumberAndLetterCheck(str) {
		return true
	}

	return false
}

func isMidPasswordStrength(str string) bool {
	if !isMinimumCheck(7, len(str)) {
		return false
	}

	if !isContainSymbolCheck(str) && !isContainPunctCheck(str) {
		return false
	}

	if isContainNumberCheck(str) {
		return true
	}

	if isContainsLetterCheck(str) {
		return true
	}

	if isContainNumberCheck(str) && isContainsLetterCheck(str) {
		return true
	}

	return false
}

func isStrongPasswordStrength(str string) bool {
	if !isMinimumCheck(14, len(str)) {
		return false
	}
	return isMidPasswordStrength(str)
}

func CheckPassword(str string) string {
	if isWeakPasswordStrength(str) {
		return "lemah"
	} else if isVeryWeakPasswordStrength(str) {
		return "sangat lemah"
	} else if isStrongPasswordStrength(str) {
		return "kuat"
	} else if isMidPasswordStrength(str) {
		return "sedang"
	}
	return "tidak diketahui"
}

func PasswordGenerator(base string) (string, string) {
	generatedPassword := Generate(base)
	passwordStrength := CheckPassword(generatedPassword)
	return generatedPassword, passwordStrength
}

func main() {
	a, b := PasswordGenerator("admin")
	fmt.Println(a)
	fmt.Println(a == "NOMDe")
	fmt.Println(b == "lemah")
}
