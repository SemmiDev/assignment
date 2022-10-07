package main

import "fmt"

func reverseItem(data int) int {
	var rev int
	for data > 0 {
		// 123 % 10 = 3
		// 12 % 10 = 2
		// 1 % 10 = 1
		rev = rev*10 + data%10
		// 123 / 10 = 12
		// 12 / 10 = 1
		data /= 10
	}
	return rev
}

func ReverseData(arr [5]int) [5]int {
	var rev [5]int
	for i, v := range arr {
		rev[len(arr)-1-i] = reverseItem(v)
	}
	return rev
}

func main() {
	fmt.Println(ReverseData([5]int{23, 456, 11, 1, 2}))
}
