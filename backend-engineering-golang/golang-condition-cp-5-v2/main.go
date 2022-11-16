package main

import "fmt"

func TicketPlayground(height, age int) int {
	switch {
	case age > 12:
		return 100_000
	case age == 12 || height > 160:
		return 60_000
	case (age >= 10 && age <= 11) || height > 150:
		return 40_000
	case (age >= 8 && age <= 9) || height > 135:
		return 25_000
	case (age >= 5 && age <= 7) || height > 120:
		return 15_000
	default:
		return -1
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
