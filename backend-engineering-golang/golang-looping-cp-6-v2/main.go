package main

import (
	"fmt"
	"strconv"
)

func SplitNumber(number int) []int {
	var numbers []int
	for number > 0 {
		numbers = append(numbers, number%10)
		number /= 10
	}

	// reverse
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

/*
kalo sam flow nya gini kak:
eh kebanyakan sam nyelesai in nya pake materi2 berikutnya kak :)

misal number nya = 11223344
1. sam split dlu kak jadi array / slice
jadinya [1,1,2,2,3,3,4,4]

3. tu sam cek
kalo misal data nya cuman ada dua, let's say [1,2], langsung aja sam return angknya itu kak
return 12

4. kalo misal data nya lebih dari dua, sam looping kek biasa kak
krna data nya tadi beruba array
sebelum looping sam bkin variable biggestPairNUmber = 0, untuk nampung pair angka terbesar nantinya kak
next nya,	sam loop dari index 0 sampe index terakhir - 1
trus sam ambil pair nya = number[i] + number[i+1], sam tampung ke variale pairNumber
trus sam bandingin pairNumber nya dengan biggestPairNumber nya, klo lebih gede pair yg sekarang
sam replace nilai biggestPairNumber tadi jadi nilai pairNumber yg barusan kak
trus dh siap loop nya dpet deh hasilnya kak

flow code gini kak
numbers = 11223
split = [1,1,2,2,3]

biggestNumber = 0

	looping = {
		pair = 1 + 1 = 2
		pair > biggestNumber? if true, then biggestNumber = pair

		loop berikutnya =
			pair = 1 + 2 = 3
			pair > biggestNumber? if true, then biggestNumber = pair

		loop berikutnya =
			pair = 2 + 2 = 4
			pair > biggestNumber? if true, then biggestNumber = pair

		loop berikutnya =
			pair = 2 + 3 = 5
			pair > biggestNumber? if false, then biggestNumber = pair
	}

dpet deh biggestNumber nya kak
*/
func BiggestPairNumber(numbers int) int {
	numbersSplit := SplitNumber(numbers)
	var biggestPairNumber int
	// two digit number is the biggest pair number
	if len(numbersSplit) == 2 {
		return numbers
	}

	// fmt.Printf("%d%d\n", numbersSplit[0], numbersSplit[1])

	var candidateA, candidateB int
	for i := 0; i < len(numbersSplit)-1; i++ {
		pairNumber := numbersSplit[i] + numbersSplit[i+1]
		if pairNumber > biggestPairNumber {
			candidateA = numbersSplit[i]
			candidateB = numbersSplit[i+1]
			biggestPairNumber = pairNumber
		}
	}

	// split number to pair number
	s := fmt.Sprintf("%d%d", candidateA, candidateB)
	sInt, _ := strconv.Atoi(s)
	return sInt
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
	fmt.Println(BiggestPairNumber(89083278))
}
