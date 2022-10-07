package main

import "fmt"

func loop(count int, value int, result *[]int) {
	for i := 1; i <= count; i++ {
		*result = append(*result, value)
	}
}

func ExchangeCoin(amount int) []int {
	banknotes := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0, len(banknotes))

	for _, v := range banknotes {
		loop(amount/v, v, &result)
		amount = amount % v
	}
	return result
}

func main() {
	a := ExchangeCoin(1752)
	fmt.Println(a)
}
