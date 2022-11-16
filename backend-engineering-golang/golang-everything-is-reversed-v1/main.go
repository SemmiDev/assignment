package main

import "fmt"

func reverseItem(data int) int {
	var rev int
	for data > 0 {
		rev = rev*10 + data%10
		data /= 10
	}
	return rev
}

func ReverseData(arr [5]int) [5]int {
	var rev [5]int
	for i, v := range arr {
		/*
			rev[4] = hasil reverse angka pertama
			rev[3] = hasil reverse angka kedua
			rev[2] = hasil reverse angka ketiga
			rev[1] = hasil reverse angka keempat
			rev[0] = hasil reverse angka kelimat

			contoh:
			arr = [12,34,44,565,3]
			rev = []

			rev[4] = 21
			rev[3] = 43
			rev[2] = 44
			rev[1] = 565
			rev[0] = 3

			rev = [3,565,44,43,21]
		*/
		rev[len(arr)-1-i] = reverseItem(v) // reverse angka
	}
	return rev
}

func main() {
	fmt.Println(ReverseData([5]int{23, 456, 11, 1, 2}))
}
