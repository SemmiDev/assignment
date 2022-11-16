package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64 = float64(n)

	for i := 1; i < n; i++ {
		total += float64(i)*2 + 0.5
	}

	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
	fmt.Println(CountingNumber(5))
}
