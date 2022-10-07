package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var result float64

	if gender == "laki-laki" {
		result = float64(height - 100)
		result = result - (result * 10 / 100)
		return result
	}
	result = float64(height - 100)
	result = result - (result * 15 / 100)
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
