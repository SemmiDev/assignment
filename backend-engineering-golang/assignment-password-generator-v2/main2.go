package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Reverse(str string) string {
	var reverse string
	for _, v := range str {
		reverse = string(v) + reverse
	}
	return reverse
}

func Generate(str string) string {
	text := Reverse(str)
	var generate string
	var slice []string
	var vokal = map[string]string{
		"a": "E", "e": "I", "i": "O", "o": "U", "u": "A",
		"A": "e", "E": "i", "I": "o", "O": "u", "U": "a",
	}
	for _, v := range text {
		if string(v) == " " {
			continue
		}
		var inputVocal, outputVocal = vokal[string(v)]
		if outputVocal {
			slice = append(slice, inputVocal)
		} else {
			var change string
			if unicode.IsUpper(v) {
				change = strings.ToLower(string(v))
				slice = append(slice, change)
			} else {
				change = strings.ToUpper(string(v))
				slice = append(slice, change)
			}
		}
	}
	generate = strings.Join(slice, "")
	return generate
}

func checkKecilDari7(str string) bool {
	return len(str) < 7
}

func checkBesarDari7(str string) bool {
	return len(str) >= 7 && len(str) < 14
}

func checkBesarSamaDari14(str string) bool {
	return len(str) >= 14
}

func CheckPassword(str string) string {
	str1 := Generate(str)
	var isNumber, isChar, IsSymbol bool
	
	kecilDari7, besarDari7, besarSamaDari14 := checkKecilDari7(str1), checkBesarDari7(str1), checkBesarSamaDari14(str1)

	for _, v := range str1 {
		if !isChar && unicode.IsLetter(v) {
			isChar = true
		}
		if !isNumber && unicode.IsNumber(v) {
			isNumber = true
		}
		if !IsSymbol && (unicode.IsSymbol(v) || unicode.IsPunct(v)) {
			IsSymbol = true
		}
	}

	sangatLemah := kecilDari7
	lemah := besarDari7 && (isNumber || isChar) && !IsSymbol
	sedang := besarDari7 && (isNumber || isChar) && IsSymbol
	kuat := besarSamaDari14 && (isNumber || isChar) && IsSymbol

	return status(sangatLemah, lemah, sedang, kuat)
}

func status(sangatLemah, lemah, sedang, kuat bool) (result string) {
	if sangatLemah {
		result = "sangat lemah"
	} else if lemah {
		result = "lemah"
	} else if sedang {
		result = "sedang"
	} else if kuat {
		result = "kuat"
	}
	return
}

func CheckLetter(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func CheckNumber(str string) bool {
	for _, v := range str {
		if unicode.IsNumber(v) {
			return true
		}
	}
	return false
}

func CheckSymbol(str string) bool {
	var symbol = "!@#$%^&*()_+{}|:<>?`~-=[]\\;',./"
	for _, v := range symbol {
		if strings.ContainsRune(str, v) {
			return true
		}
	}
	return false
}

func PasswordGenerator(base string) (string, string) {
	generate := Generate(base)
	check := CheckPassword(base)
	return generate, check
}

func main() {
	data := "#$)321#!3312" // bisa digunakan untuk melakukan debug
	// data2 := "123213213123"  // bisa digunakan untuk melakukan debug
	// data3 := "adminadmin"    // bisa digunakan untuk melakukan debug
	// data4 := "adminadmin123" // bisa digunakan untuk melakukan debug
	res, check := PasswordGenerator(data)
	// fmt.Println(res, check)
	// res, check = PasswordGenerator(data2)
	// fmt.Println(res, check)
	// res, check = PasswordGenerator(data3)
	// fmt.Println(res, check)
	// res, check = PasswordGenerator(data4)
	// fmt.Println(res, check)
	fmt.Println(res)
	fmt.Println(check)
}