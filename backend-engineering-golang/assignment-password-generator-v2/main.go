package main

import (
	"fmt"
	"strings"
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

/*
	konsonan kecil jadi besar
	konsonan besar jadi kecil

	kebalikannya adalah

	vokal kecil jadi besar
	vokal besar jadi kecil

*/

func changeLetterCase(str string) string {
	var result string
	for _, v := range str {
		if vocal, ok := vocal[string(v)]; !ok {
			if unicode.IsLower(v) {
				result += strings.ToUpper(string(v))
			} else {
				result += strings.ToLower(string(v))
			}
		} else {
			if unicode.IsLower(v) {
				result += strings.ToUpper(vocal)
			} else {
				result += strings.ToLower(vocal)
			}
		}

		//// cek apakah dia ga huruf vokal
		//if _, ok := vocal[string(v)]; !ok {
		//	// cek apakah hurufnya besar
		//	if unicode.IsUpper(v) {
		//		// ubah jadi huruf kecil
		//		result += string(unicode.ToLower(v))
		//	} else {
		//		result += string(unicode.ToUpper(v))
		//	}
		//} else {
		//	if unicode.IsLower(v) {
		//		result += string(unicode.ToUpper(v))
		//	} else {
		//		result += string(v)
		//	}
		//}
	}

	return result
}

func Generate(str string) string {
	// 1. reverse string
	reversedStr := Reverse(str)
	// 2. mengubah ke vokal berikutnya
	changedToNextVocal := ChangeToNextVocal(reversedStr)
	// 3. ubah huruf besar selain vokal jadi huruf kecil, dan sebaliknya
	changedLetterCase := changeLetterCase(changedToNextVocal)
	// 4. hapus spasinya
	return removeSpace(changedLetterCase)
}

func isContainSymbolCheck(str string) bool {
	for _, v := range str {
		if unicode.IsSymbol(v) {
			return true
		}
	}

	allsymbols := "!@#$%^&*()_+{}|:<>?`~-=[]\\;',./"
	for _, v := range allsymbols {
		if strings.ContainsRune(str, v) {
			return true
		}
	}
	return false
}

func isContainPunctCheck(str string) bool {
	for _, v := range str {
		if unicode.IsPunct(v) || unicode.IsMark(v) {
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
	total := len(str)
	totalLetter := 0
	totalNumber := 0

	for _, v := range str {
		if unicode.IsLetter(v) {
			totalLetter++
			continue
		}
		if unicode.IsNumber(v) {
			totalNumber++
			continue
		}
	}
	return total-(totalLetter+totalNumber) == 0
}

func isVeryWeakPasswordStrength(str string) bool {
	return len(str) < 7
}

func isWeakPasswordStrength(str string) bool {
	if len(str) < 7 {
		return false
	}

	if isAllContainsNumberCheck(str) ||
		isAllContainsLetterCheck(str) ||
		isAllJustContainsNumberAndLetterCheck(str) {
		return true
	}

	return false
}

func isMidPasswordStrength(str string) bool {
	if len(str) < 7 {
		return false
	}

	if !isContainSymbolCheck(str) || !isContainPunctCheck(str) {
		fmt.Println("1. tidak ada simbol")
		return false
	}

	if isContainNumberCheck(str) && (isContainSymbolCheck(str) || isContainPunctCheck(str)) {
		fmt.Println("2. number  dan simbol")
		return true
	}

	if isContainsLetterCheck(str) && (isContainSymbolCheck(str) || isContainPunctCheck(str)) {
		fmt.Println("3. huruf dan simbol")
		return true
	}

	if isContainNumberCheck(str) && isContainsLetterCheck(str) && (isContainSymbolCheck(str) || isContainPunctCheck(str)) {
		fmt.Println("4. number dan huruf dan simbol")
		return true
	}

	return false
}

func isStrongPasswordStrength(str string) bool {
	if len(str) < 14 {
		return false
	}
	return isMidPasswordStrength(str)
}

func CheckPassword(str string) string {
	if isVeryWeakPasswordStrength(str) {
		return "sangat lemah"
	}

	if isStrongPasswordStrength(str) {
		return "kuat"
	}

	if isMidPasswordStrength(str) {
		return "sedang"
	}
	
	if isWeakPasswordStrength(str) {
		return "lemah"
	}

	return "tidak diketahui"
}

func PasswordGenerator(base string) (string, string) {
	generatedPassword := Generate(base)
	passwordStrength := CheckPassword(generatedPassword)
	return generatedPassword, passwordStrength
}

func main() {
	a := Generate("Semangat Pagi 12!#")
	fmt.Println(a == "#!21OGEpTEGNEMIs")
}
